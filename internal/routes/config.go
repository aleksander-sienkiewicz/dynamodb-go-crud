package routes

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
)

type Config struct { //struct called config, all it has is time.duration (from import)
	timeout time.Duration
}

func NewConfig() *Config { //returns config
	return &Config{}
}

// SET UP OUR CORS
// c *config funcs are struct methods
// pass handler here
// in here we say whats allowed, if it is allowed then go on.
func (c *Config) Cors(next http.Handler) http.Handler { //handle Cors
	return cors.New(cors.Options{ //in these options we define where we want origins to be
		AllowedOrigins:   []string{"*"}, //allow all origins
		AllowedMethods:   []string{"*"}, //allow methods
		AllowedHeaders:   []string{"*"}, //everything allowed! thats why *
		ExposedHeaders:   []string{"*"}, //
		AllowCredentials: true,
		MaxAge:           5,
	}).Handler(next) //returns handler, next sends it to next func which is http handler
}

func (c *Config) SetTimeout(timeInSeconds int) *Config { //set timeout
	c.timeout = time.Duration(timeInSeconds) * time.Second //time specified by duration of a second defined in time file
	return c
}

func (c *Config) GetTimeout() time.Duration { //get timeout, no params passed
	return c.timeout
}
