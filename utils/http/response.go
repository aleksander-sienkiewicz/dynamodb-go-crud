package http

//responses defined here
import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct { //check documentation to see what ur response will be
	Status int         `json:"status"` //like 500, 400, type of response from server
	Result interface{} `json:"result"`
}

/*
called any time we make a response,
takes data
takes status (could be error or statusOK) ir. 200,300,400,500
*/
func newResponse(data interface{}, status int) *response { //grab response from above
	return &response{ //return our response struct that we defined
		Status: status, //status is the status
		Result: data,   //result is an interface, the data we got
	}
}

// struct method
func (resp *response) bytes() []byte { //return slice of bytes
	data, _ := json.Marshal(resp) //marshal the response, get in a variable called data
	return data                   //return data
}

func (resp *response) string() string { //return strng
	return string(resp.bytes()) //return what we recieve, convert it to a str.
}

// response & response writer
// sends our response
func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Status)   //respond w/ status
	_, _ = w.Write(resp.bytes()) //send resp.bytes. takes two values but we dont need em so just send it out
	log.Println(resp.string())   //PRINT THTTTTTT
	//whenever we wanna use a struct method, so we have struct response right,
	//and struct methods []byte and string
	//when ever we wanna use a struct method call is like resp.Status, or resp.string. (i knew that alr bro idk why i wrote that like it was some real heat)
} //sike i know its cuz its used ALLTHEDAMNTIMEEEEE

/*standardize types of response, very good practice for prod. code
this is good cuz u can acc know what going on in ur program when u get errors (woah)
so down below we go define all our lil errors
*/

// 200
func StatusOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	newResponse(data, http.StatusOK).sendResponse(w, r) //when everything works all good
}

// 204
func StatusNoContent(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusNoContent).sendResponse(w, r) //when we have NO CONTENT
}

// 400
func StatusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()} //map stri interface, define error
	newResponse(data, http.StatusBadRequest).sendResponse(w, r)
}

// 404
func StatusNotFound(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusNotFound).sendResponse(w, r)
}

// 405
func StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	newResponse(nil, http.StatusMethodNotAllowed).sendResponse(w, r)
} //no data has to be passed cuz not allowed

// 409
func StatusConflict(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusConflict).sendResponse(w, r)
}

// 500
func StatusInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	data := map[string]interface{}{"error": err.Error()}
	newResponse(data, http.StatusInternalServerError).sendResponse(w, r)
}
