package instance

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//this files starts a session with data base and conects to dynamodb session

//thanks dynamodb session library! <3

func GetConnection() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{ //start new session, cmnd from the
		//dynamodb documentation
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	return dynamodb.New(sess) //return new sess(ion)
}
