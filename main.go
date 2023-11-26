package main

import(
    "os"
    "github.com/gin-gonic/gin"
    routes "golang-jwt-artica/routes"
)

func main(){

    port := os.Getenv("PORT")
    if port==""{
        port = "8000"
    }

    router := gin.New()

}