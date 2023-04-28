package rules

import (
	"io"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// create interface for rules
type Interface interface {
	ConvertIoReaderToStruct(data io.Reader, model interface{}) (body interface{}, err error)
	//ioreader type to struct type, *abstraction*<3
	GetMock() interface{}                        //gets nmock values
	Migrate(connection *dynamodb.DynamoDB) error //returns error, or nil!
	Validate(model interface{}) error            //returns error, or nil!
}
