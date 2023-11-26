package database

import(
"fmt"
"log"
"time"
"os"
"github.com/joho/godotenv"
"go.mongodb.org/mongodb-driver/mongo"
"go/mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("error Loading env file")
    }

    MongoDb := os.Getenv("MONGODB_URL")
    client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDb")
    return client
}