package main

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Struct struct {
    Name string
}

func insert(collection *mongo.Collection) error {
    data := &Struct { Name: "WatanabeJunna" }

    _, err := collection.InsertOne(context.Background(), data)
    return err
}

func find(collection *mongo.Collection) error {
    cur, err := collection.Find(context.Background(), bson.D{})

    if err != nil { 
        return err
    }

    defer cur.Close(context.Background())

    for cur.Next(context.Background()) {
        var result bson.M

        err := cur.Decode(&result)

        if err != nil { 
            return err
        }

        log.Println(result)
    }

    if err := cur.Err(); err != nil {
        return err
    }

    return nil
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

    col := client.Database("test").Collection("user")

    if err = insert(col); err != nil {
        log.Fatalln(err)
    }

    if err = find(col); err != nil {
        log.Fatalln(err)
    }
}