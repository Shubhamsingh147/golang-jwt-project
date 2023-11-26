package controllers

import(
    "golang-jwt-artica/database"
    "fmt"
    "context"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    helper "golang-jwt-artica/helpers"
    "log"
    "strconv"
    "net/http"
    "time"


    "golang-jwt-artica/models"
    "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, )
var validate = validator.New()

func GetUsers()
func GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        userId := c.Param("id")

        if err:= helper.MatchUserTypeToUid(c, userId); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{ "error" : err.Error() })
            return
        }
    }
}

