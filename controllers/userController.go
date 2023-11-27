package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongodb-driver/bson"
	"go.mongodb.org/mongodb-driver/bson/primitive"
	"go.mongodb.org/mongodb-driver/mongo"
	"golang-jwt-artica/database"
	helper "golang-jwt-artica/helpers"
	"golang-jwt-artica/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client)
var validate = validator.New()

func HashPassword(password string) string {
    bcrypt.GenerateFromPassword([]byte(password, 14)
    if err != nil {
        log.Panic(Err)
        return err
    }
    return string(bytes), err
}


func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
    err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
    check := true
    msg := ""

    if err != nil {
        msg := fmt.Sprintf("password is incorrect: %s". err.Error())
        check = false
    }

    return check, msg
}

func GetUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := helper.CheckUserType(c, "ADMIN"); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second))
        recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
        if err != nil || recordPerPage < 1 {
            recordPerPage = 10
        }
        page, err1 := strconv.Atoi(c.Query("Page"))
        if err1 != nil || page < 1 {
            page = 1
        }
        startIndex := (page - 1) * recordPerPage
        startIndex, err = strconv.Atoi(c.Query("startIndex"))

        matchStage := bson.D{{"$match", bson.D{{}}}}
        groupStage := bson.D{{"$group", bson.D{
            {"_id", bson.D{{"_id", "null"}}},
            {"total_count", bson.D{{"$sum", 1}}},
            {"data", bson.D{{"$push", "$$ROOT"}}}
        }}}
        projectStage := bson.D{
            {"$project", bson.D{
                {"_id", 0},
                {"total_count", 1},
                {"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage} }}},
            }}
        }
        userCollection.Aggregate(ctx, mongo.Pipeline{
            matchStage, groupStage, projectStage
        })
        defer cancel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while querying to mongo"})
        }
        var allUsers []bson.M
        if err = result.All(ctx, &allUsers); err != nil {
            log.Fatal(err)
        }
        c.JSON(http.StatusOk, allUsers[0])
    }
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return err
		}
		c.JSON(http.StatusOk, user)
	}
}

func Login() gin.HandlerFunc {
    return func(c *gin.Context) {
       var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
       var user models.User
       var foundUser models.User

       if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
       }

       userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
       defer cancel()
       if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "email does not exist"})
            return
       }
       passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
       defer cancel()
       if passwordIsValid != true {
            c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
            return
       }

       if foundUser.Email == nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
       }
       token, refreshToken, _ := helper.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName)
       helper.UpdateAllTokens(token, refreshToken, foundUser.UserId)
       err = userCollection.FindOne(ctx, bson.M{"user_id": foundUser.UserId}).Decode(&foundUser)

       if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
       }
       c.JSON(http.StatusOk, foundUser)
    }
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while checking existing emails"})
		}

        password := HashPassword(*user.Password)
        user.Password = &password

		count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error while checking existing phones"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userId with email or phone already exists"})
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.UserId, _ = user.ID.Hex()
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UserType, *&user.UserId)
		user.Token = &token
		user.RefreshToken = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("user item was not created : %s", insertErr.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOk, resultInsertionNumber)
	}
}
