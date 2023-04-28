// create a very basic outline for main.go first
package main

import (
	"fmt"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/config"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/repository/adapter"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/repository/instance"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/routes"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/rules"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	//"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/rules/product"
	"log"
	"net/http"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/utils/logger"
	//aws-sdk-go has rlly great documentation for so many dif things, dianamodb, lambda, etc.
) //db just a small part of the huge library aws gives

func main() {
	configs := config.GetConfig() //takes config, and gets attributes

	connection := instance.GetConnection()       //use getconection func
	repository := adapter.NewAdapter(connection) //use newadapter func, pass connection from above.
	//newadapter returns the whole database to u with the connection and logNode, etc.

	logger.INFO("Waiting service starting.... ", nil) //uses logger package.

	errors := Migrate(connection) //basic database migration.
	if len(errors) > 0 {          //handle multiple errors, if more than 0
		for _, err := range errors { //range over errors
			logger.PANIC("Error on migrate: ", err) //panic logger,print error on migration
		}
	}
	logger.PANIC("", checkTables(connection)) //check tables, if no accesss error.

	//standard procedure when ever creating project with database to use migrating seeds and scripts aand all that and make
	//sure the data tables exist before u start workin too

	//Sprintf helps u format a string
	port := fmt.Sprintf(":%v", configs.Port)            //print out the port , format it too
	router := routes.NewRouter().SetRouters(repository) //
	logger.INFO("Service running on port ", port)       //print to log what port u on

	server := http.ListenAndServe(port, router) //listen&Serve, /http package gives us it, pass port and router
	//and u get a server
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error { //gives us connection to our DB
	var errors []error //return slice of errors

	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})

	return errors
}

// takes errors, connection,
func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection) //pass into err
	if err != nil {                 //if theres error
		*errors = append(*errors, err) //append errors
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{}) //instance.getconection is connection
	//return list of all our tables
	if response != nil { //if theres a response
		if len(response.TableNames) == 0 { //if len of list is 0
			logger.INFO("Tables not found: ", nil) //no tables
		}
		for _, tableName := range response.TableNames { //other wise,  look thru the list index'
			logger.INFO("Table found: ", *tableName) //log table found, log table name
		}
	}
	return err //otherwise return err
}
