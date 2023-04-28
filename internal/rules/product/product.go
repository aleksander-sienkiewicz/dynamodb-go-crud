package product

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/entities"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/entities/product"

	//this is all from AWS terminology library
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Rules struct{}

// Create rules struct
func NewRules() *Rules {
	return &Rules{}
}

/*
struct method

takes data, and model. returns interface and error
*/

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil { //if we have no data
		return nil, errors.New("body is invalid") //body invalid, no data recieved
	}
	return model, json.NewDecoder(data).Decode(model) //get io data, decode into a model we can use.
}

/*
struct method
just need connection to dynamoDB so will be a quick func.
*/
func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	return r.createTable(connection) //return r.createTable with connection passed.
}

/*
create table struct method
this is mosty from DynamoDB Documentations

so they say to stay using the tables directly u have to ensure the tables  exist so you have to
initialize them youself; unlike lets say mongoDB that does it for u
but ya when u make the connection with mltp

so ya heres createtable func <3  USE THIS THING WHEN USING DYNAMODB
*/
func (r *Rules) createTable(connection *dynamodb.DynamoDB) error {
	table := &product.Product{}

	input := &dynamodb.CreateTableInput{ //CreateTableInput provided by dynamoDB
		//capture that in variable called input
		AttributeDefinitions: []*dynamodb.AttributeDefinition{ //from library
			{
				AttributeName: aws.String("_id"), //from aws package
				AttributeType: aws.String("S"),   //from aws package Type string
			}, //just telling dynamo that u will have name for table input called ID
		},
		KeySchema: []*dynamodb.KeySchemaElement{ //from library, reference same slice
			{
				AttributeName: aws.String("_id"), //attribute name is ID
				KeyType:       aws.String("HASH"),
			},
		}, //now u have to define something called ProvisionedThroughput
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10), //read capacity units - int 10
			WriteCapacityUnits: aws.Int64(10), //write capacity units - int 10
		},
		TableName: aws.String(table.TableName()), //use table to get access to TableName

	}
	response, err := connection.CreateTable(input) //capture in response, and err
	if err != nil && strings.Contains(err.Error(), "Table already exists") {
		return nil //return nil for error if not same
	}
	if response != nil && strings.Contains(response.GoString(), "TableStatus: \"CREATING\"") {
		time.Sleep(3 * time.Second)     //sleep program when creating
		err = r.createTable(connection) //create table
		if err != nil {
			return err //error status if encountered
		}
	}
	return err
}
func (r *Rules) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{ //created base file in entities/base.go
			ID:        uuid.New(), //create new id
			CreatedAt: time.Now(), //update now
			UpdatedAt: time.Now(), //update now
		},
		Name: uuid.New().String(), //assign name
	}
}

/*
struct meth
take model of type interface from iotostruct function and pass it through validate.
hopefully err returns nil
*/
func (r *Rules) Validate(model interface{}) error {
	productModel, err := product.InterfaceToModel(model) //pass thru func
	if err != nil {                                      //error handle
		return err
	}

	return Validation.ValidateStruct(productModel,
		Validation.Field(&productModel.ID, Validation.Required, is.UUIDv4),
		//take id, check for validation, it is required, check on uuid version 4
		Validation.Field(&productModel.Name, Validation.Required, Validation.Length(3, 50)),
		//check name, check validation, length is 3,50 - light
	)
}
