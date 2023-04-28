package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/controllers/product"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/entities/product"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/handlers"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/repository/adapter"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/rules/product"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/utils/http"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// define handler struct
type Handler struct {
	handlers.Interface

	Controller product.Interface
	Rules      Rules.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Controller: product.NewController(repository),
		Rules:      RulesProduct.NewRules(),
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if chi.URLParam(r, "ID") != "" {
		h.getOne(w, r)
	} else {
		h.getAll(w, r)
	}
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil { //bad request
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}
	//internalservererror
	response, err := h.Controller.ListOne(ID)
	if err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	response, err := h.Controller.ListAll()
	if err != nil { //internal server error
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, response)
}

//post func

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	//take request body, validate it, then we get our productBoody
	productBody, err := h.getBodyAndValidate(r, uuid.Nil) //access handler, pass r and nil for uuid
	if err != nil {                                       //if there was an error
		HttpStatus.StatusBadRequest(w, r, err) //badrequest
		return                                 //return
	}

	ID, err := h.Controller.Create(productBody)
	if err != nil { //internal server error
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}

	HttpStatus.StatusOK(w, r, map[string]interface{}{"id": ID.String()})
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID")) //id will be passed to rest of funcs
	if err != nil {                              //bad request
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return
	}

	productBody, err := h.getBodyAndValidate(r, ID)
	if err != nil { //bad req
		HttpStatus.StatusBadRequest(w, r, err)
		return
	}

	if err := h.Controller.Update(ID, productBody); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return //int err
	}

	HttpStatus.StatusNoContent(w, r)
}

// struct func
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ID, err := uuid.Parse(chi.URLParam(r, "ID")) //id to find exact record w/ data
	//uuid lets us parse, CHI router gives us accesss to parametres, the id
	//capture it in ID
	if err != nil { //if theres an error
		HttpStatus.StatusBadRequest(w, r, errors.New("ID is not uuid valid"))
		return //send err
	}

	if err := h.Controller.Remove(ID); err != nil {
		HttpStatus.StatusInternalServerError(w, r, err)
		return
	}
	//empty data
	HttpStatus.StatusNoContent(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}

// when ever u use put ur getting some body from the request. when ur posting u need all data put putting or updating
// u get a lil bit, the id and data for that id. thats whaty we wana update, needs to be validated
func (h *Handler) getBodyAndValidate(r *http.Request, ID uuid.UUID) (*EntityProduct.Product, error) {
	productBody := &EntityProduct.Product{}
	body, err := h.Rules.ConvertIoReaderToStruct(r.Body, productBody) //defined in a file called rules
	if err != nil {
		return &EntityProduct.Product{}, errors.New("body is required")
	}

	productParsed, err := EntityProduct.InterfaceToModel(body) //interfacetomodel in entityprod,
	//interface -> model function.
	if err != nil {
		return &EntityProduct.Product{}, errors.New("error on convert body to model")
	}

	setDefaultValues(productParsed, ID)

	return productParsed, h.Rules.Validate(productParsed)
}

func setDefaultValues(product *EntityProduct.Product, ID uuid.UUID) {
	product.UpdatedAt = time.Now()
	if ID == uuid.Nil {
		product.ID = uuid.New()
		product.CreatedAt = time.Now()
	} else {
		product.ID = ID
	}
}
