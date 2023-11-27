package middleware

import(
    "fmt"
    "net/http"
    helper "golang-jwt-artica/helpers"
    "github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientToken := c.Request.Header.Get("token")
        if clientToken == "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint("No Authorization header provided")})
            c.Abort()
            return
        }

        claims, msg := helper.ValidateToken(clientToken)
        if msg != "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error while authenticating: %s", msg)})
            c.Abort()
            return
        }

        c.Set("email", claims.Email)
        c.Set("first_name", claims.FirstName)
        c.Set("last_name", claims.LastName)
        c.Set("user_type", claims.UserType)
        c.Set("uid", claims.Uid)
        c.Next()

    }
}