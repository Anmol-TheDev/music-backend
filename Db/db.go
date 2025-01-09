package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() {
	db_password := os.Getenv("MONGO_PASSWORD")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://anmol:%s@anmol.9253c.mongodb.net/?retryWrites=true&w=majority&appName=Anmol",db_password)).SetServerAPIOptions(serverAPI)
	
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
