package adapter

//this file is all abt talking to our database
//This contains all funcs that talk to dynamodb, this is the basic crud.
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Database struct {
	connection *dynamodb.DynamoDB //this is our defines session
	logMode    bool               //true or false, on or off boi!
}

type Interface interface { //create interface
	Health() bool
	//deinfes funcs
	FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
}

/*
returns interface
*/
func NewAdapter(con *dynamodb.DynamoDB) Interface { //accepts connection of type *dynamodb.DynamoDB
	return &Database{ //return database
		connection: con, //con == connectoion
		logMode:    false,
	}
}

// struct method
func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{}) //if tables are getting listed
	return err == nil                                               //everything is fine, return nil for error
}

/*
this & FindOne are so important!!!! same w/ createorupdate/delete these directly talk to our dynamodb and get info that we need
so from controllers we will call these two funcs mostly - these r our database lvels functions -
*/
//expression from import
func (db *Database) FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{ //create input, type scaninput

		//in condition we get a few things, lets assign em to variables
		ExpressionAttributeNames:  condition.Names(), //these will be used by dynamodb to proccess the info and find data from the database
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	} //so ya things on left collum is what dynamoDB needs, right side is what got passed into the function
	//this from dynamo documentations
	return db.connection.Scan(input)
}

/*this is sorta just dynamodb syntax... just get used to it kid (the table name and key part and input := &dyn...)

 */
//pass condition of type map, and its string & interface? unreal.. and  we pass tablename too for dynamodb func
func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) { //return response! of type *dynamodb.GetItemOutput, or we can also return error here
	//condition just means what values do u want to match to what values
	//ie. keyvalues to keyvalues as ints, or as name = to (string here)
	//condition can also be empty brackets which means grab the whole data set, no condition passed
	conditionParsed, err := dynamodbattribute.MarshalMap(condition) //
	if err != nil {                                                 //check for error, if there is an error
		return nil, err //return no response, and error
	}

	input := &dynamodb.GetItemInput{ /*so ya what ever we return from the dynamodb func, almost like mongodb func
		that is below we capture our input var, which we return at the end of the func.
		*/
		TableName: aws.String(tableName), //in mongodb u say collection name, here u say tablename w/ dynamodb
		//also pass table namehere
		Key: conditionParsed, //requires key
	}
	return db.connection.GetItem(input) //return input.
}

// super important
func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) { //return puritemoutput or error
	entityParsed, err := dynamodbattribute.MarshalMap(entity) //entity is grabbed and passed thru func here
	if err != nil {                                           //same error
		return nil, err //return error and no response if there is error
	}
	input := &dynamodb.PutItemInput{ //dynamodb code part again
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input) //input is what we get back from database func
} //also this is so little work compared to mysql or mongodb, in terms or syntax. very little code here
//love dynamodb for that <3 create or update are the same func! not two dif ones how awesome
//also documentation for dynamodb is very easy to understand

// super importatnt
// delete is very similar to find one, cuz first we find it, then delet eit. so will naturally be similar to findone
func (db *Database) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) { //different cuz returns deleteitemoutput, error option is always there
	conditionParsed, err := dynamodbattribute.MarshalMap(condition) //same as findone func
	if err != nil {
		return nil, err
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,       //yes yes
		TableName: aws.String(tableName), //yes yes, so swag
	}
	return db.connection.DeleteItem(input) //return deleteitem input
}
