package remote

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/thiagoluis88git/tech1/internal/core/data/model"
)

type CognitoRemoteDataSource interface {
	SignUp(user *model.Customer) error
}

type CognitoRemoteDataSourceImpl struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientID   string
}

func NewCognitoRemoteDataSource(appClientId string, region string) CognitoRemoteDataSource {
	config := &aws.Config{Region: aws.String(region)}
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	client := cognito.New(sess)

	client.AdminUpdateUserAttributes(&cognito.AdminUpdateUserAttributesInput{
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("true"),
			},
		},
	})

	// teste, testee := client.CreateUserPoolRequest()
	return &CognitoRemoteDataSourceImpl{
		cognitoClient: client,
		appClientID:   appClientId,
	}
}

func (c *CognitoRemoteDataSourceImpl) SignUp(user *model.Customer) error {
	messageAction := "SUPPRESS"

	pass := fmt.Sprintf("%v%v", user.Email, "12!@Az")

	userCognito := &cognito.AdminCreateUserInput{
		UserPoolId:        aws.String(c.appClientID),
		Username:          aws.String(user.Email),
		MessageAction:     &messageAction,
		TemporaryPassword: &pass,
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
			{
				Name:  aws.String("email_verified"),
				Value: aws.String("True"),
			},
		},
	}

	_, err := c.cognitoClient.AdminCreateUser(userCognito)

	if err != nil {
		return err
	}

	password := fmt.Sprintf("%v%v", user.Email, "1234&$sWa")
	permanent := true

	setPasswordInput := &cognito.AdminSetUserPasswordInput{
		Password:   &password,
		UserPoolId: aws.String(c.appClientID),
		Username:   aws.String(user.Email),
		Permanent:  &permanent,
	}

	_, errPasswd := c.cognitoClient.AdminSetUserPassword(setPasswordInput)

	if errPasswd != nil {
		return errPasswd
	}

	return nil
}
