package remote

import (
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
	userCognito := &cognito.SignUpInput{
		ClientId: aws.String(c.appClientID),
		Username: aws.String(user.CPF),
		Password: aws.String(user.CPF),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
		},
	}

	_, err := c.cognitoClient.SignUp(userCognito)

	if err != nil {
		return err
	}

	return nil
}
