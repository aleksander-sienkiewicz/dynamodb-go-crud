package env

//how to get the environment
import "os"

func GetEnv(env, defaultValue string) string { //get all environment variables, reutn a string
	environment := os.Getenv(env) //proccess that var here
	if environment == "" {        //if input is empty string
		return defaultValue //use default environment
	}

	return environment //else use the environment in input
}
