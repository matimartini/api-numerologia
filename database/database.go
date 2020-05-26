package database

import (
	"context"
	"log"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

/*
func Init() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, nil, sa)

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	return client

	/* 	log.Println(client)
	   	iter := client.Collection("test-1").Documents(ctx)

	   	for {
	   		doc, err := iter.Next()

	   		if err == iterator.Done {
	   			break
	   		}
	   		if err != nil {
	   			log.Fatalf("Failed to iterate: %v", err)
	   		}
	   		fmt.Println(doc.Data())
	   	}
} */

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

	client.Close()
	return document
}

/* func GetDescriptionSoulAmbition(number int, nameDocuent string) {
	document := GetDesciption(number, nameDocuent)
	var soulAmbition service.SoulAmbition
	document.DataTo(&soulAmbition)
	fmt.Println(soulAmbition)
} */
