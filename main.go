package main

import(
    "os"
    "log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    routes "golang-jwt-artica/routes"
)

func main(){

    port := os.Getenv("PORT")
    if port==""{
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())

    routes.AuthRoutes(router)
    routes.UserRoutes(router)

    router.GET("/api-1", func(c *gin.Context) {
        c.JSON(200, gin.H{"success" : "Access Granted for API-1"})
    })

    router.GET("/api-2", func(c *gin.Context) {
        c.JSON(200, gin.H{"success" : "Access Granted for API-2"})
    })

    router.Run(":" + port)
}