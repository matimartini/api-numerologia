package database

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func GetDesciption(number int, nameDocument string) *firestore.DocumentSnapshot {

	ctx := context.Background()
	sa := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, sa)

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	numberTxt := strconv.Itoa(number)
	document, err := client.Collection(nameDocument).Doc(numberTxt).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	}

	if !document.Exists() {
		fmt.Println("Error not description name document: ", nameDocument)
	}
	return document
}
