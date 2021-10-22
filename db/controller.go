package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/yakumo-saki/ofuroNotifyGo/config"
)

var connection *dynamo.DB

var awsConfig aws.Config

func Init(config *config.ConfigStruct) {
	awsConfig.Region = aws.String(config.Region)
	awsConfig.Endpoint = aws.String(config.Endpoint)
	awsConfig.DisableSSL = aws.Bool(config.DisableSSL)
}

func getConnection() *dynamo.DB {

	if connection != nil {
		return connection
	}

	sess := session.Must(session.NewSession())

	connection := dynamo.New(sess, &awsConfig)

	return connection
}
