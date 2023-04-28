package config //this name must be exact same as the folder name, and package name, not even a capital difference boi

import (
	"strconv" //string convert

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/utils/env" //import environment
)

type Config struct {
	Port    int //always need a port for servers to interact with
	Timeout int //if it doesnt work for x time then we should cut it off, otherwise always pending
	//always do timeout for api or back end server so front end doesnt wait for back end 4EVER
	//this also need time func when we pass an actual value into it
	Dialect     string //like will it be sql, all that
	DatabaseURI string //where to go to locate the database, which poort .
}

func GetConfig() Config { //returns config structure
	return Config{ //return config
		//config has:
		Port: parseEnvToInt("PORT", "8080"), //port, if doesnt work then use a dif port or make sure
		//8080 isnt in use or what ever we choose. java programs will also often use same port so if it doesnt work
		//could be a java proj running in the background using the port
		//if WE DO GET A PORT ISSUE - u kill the ports - cuz u gave up on a project over that bozo.
		Timeout:     parseEnvToInt("TIMEOUT", "30"),   //timeout
		Dialect:     env.GetEnv("DIALECT", "sqlite3"), //dialect
		DatabaseURI: env.GetEnv("DATABASE_URI", ":memory:"),
		//dynamodb uri's, youll see a link that stars with ":memory:" or something
	}
}

// takes envName, default value type string, returns an integer
func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(env.GetEnv(envName, defaultValue))
	//.atoi is in the strconv documentation. or golangtour.com

	if err != nil { //if theres an error
		return 0 //return 0
	}

	return num //if no error, return our int
}
