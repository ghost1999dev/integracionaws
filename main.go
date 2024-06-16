package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/integracionaws/awsgo"
	"github.com/integracionaws/db"
	"github.com/integracionaws/models"
)

func main(){
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error)  {
	awsgo.StartAWS()
	if !ValidateParams() {
		fmt.Println("You have a error in this params. Must send 'SecretManager' ")
		err:= errors.New("Error in the parameters you have send params")
		return event, err
	}

	var data models.SignUp
	for row, att := range event.Request.UserAttributes{
		switch row{
		case "email":
			data.UserEmail = att
			fmt.Println("Email = "+ data.UserEmail)
		case "sub":
			data.UserUUID =att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil{
		fmt.Println("Error at read the secret "+err.Error())
		return event,err
	}
	err = db.SigunUp(data)
	return event, err
}

func ValidateParams() bool {
	var getParams bool
	_, getParams = os.LookupEnv("SecretName")
	return getParams
}