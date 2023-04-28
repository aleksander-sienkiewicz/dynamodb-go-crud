package health

import (
	"errors"
	"net/http"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/handlers"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/repository/adapter"
	HttpStatus "github.com/aleksander-sienkiewicz/dynamodb-go-crud/utils/http"
)

type Handler struct { //handler of type strcut
	handlers.Interface                   //has handlers and repository
	Repository         adapter.Interface //this means we gotta import our two pkgs adapter and handlers
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

// get func
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

// post
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

// put
func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

// del.
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

// options
func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
