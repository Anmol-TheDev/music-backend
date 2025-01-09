package db

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func Db()(*firestore.Client,error) {
	opt := option.WithCredentialsFile("/home/anmol/Documents/firebase.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	database,err:=app.Firestore(ctx)

	if err != nil {
		return nil,err
	}

	return database,nil

}
