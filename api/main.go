package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Struct struct {
    Name string
}

func insert(col *mongo.Collection) error {
    data := &Struct { Name: "WatanabeJunna" }

    _, err := col.InsertOne(context.Background(), data)
    return err
}

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://db:27017"))

    if err != nil {
        log.Fatalln(err)
    }

    if err = client.Connect(context.Background()); err != nil {
        log.Fatalln(err)
    }

    defer client.Disconnect(context.Background())

    col := client.Database("test").Collection("col")

    if err = insert(col); err != nil {
        log.Fatalln(err)
    }
}