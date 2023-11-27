package helper

import(
    "context"
    "fmt"
    "log"
    "os"
    "time"
    "golang-jwt-artica/database"
    jwt "github.com/dgrijalva/jwt-go"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
    Email string
    FirstName string
    LastName string
    Uid string
    UserType string
    jwt.StandardClaims
}

var userCollection *mongoCollection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstname string, lastname string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
    claims := &SignedDetails{
        Email: email,
        FirstName: firstname,
        LastName: lastname,
        Uid: uid,
        UserType: userType,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Local().Add(time.Minutes * time.Duration(5)).Unix(),
        },
    }

    refreshClaims := &SignedDetails{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(10)).Unix(),
        }
    }

    token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
    refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

    if err != nil {
        log.Panic(err)
        return
    }

    return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

    var updateObj primitive.D

    updateObj = append(updateObj, bson.E{"token", signedToken})
    updateObj = append(updateObj, bson.E{"refreshToken", signedRefreshToken})

    UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
    updateObj = append(updateObj, bson.E{"updated_at", UpdatedAt})

    upsert := true
    filter := bson.M{"user_id": userId}
    opt := options.UpdateOptions{
        Upsert : &upsert,
    }

    _, err := userCollection.UpdateOne(
        ctx,
        filter,
        bson.D{
            {"$set", updateObj},
        },
        &opt,
    )

    defer cancel()

    if err != nil {
        log.Panic(err)
        return
    }
    return
}