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
