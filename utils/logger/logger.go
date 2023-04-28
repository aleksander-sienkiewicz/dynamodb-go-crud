package logger

//so modulare everything in its own folder
import (
	"log"
)

//set up logging here
//either print out panic error or information, so those two funcs here

func PANIC(message string, err error) {
	if err != nil { //if error exists
		log.Panic(message, err) //print message and error
	}
}

func INFO(message string, data interface{}) {
	log.Print(message, data) //print message and data.
}
