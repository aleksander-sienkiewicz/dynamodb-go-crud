package handlers

//handlers called by controllers, that call repository functions, whcih are our adapters
import "net/http"

type Interface interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Options(w http.ResponseWriter, r *http.Request)
}
