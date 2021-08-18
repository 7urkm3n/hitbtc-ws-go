package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"orange-currency/src/domain/currency"

	"github.com/go-chi/chi/v5"
)

// AccessTokenHandler interface
type CurrencyHandler interface {
	GetBySymbol(http.ResponseWriter, *http.Request)
	All(http.ResponseWriter, *http.Request)
}

type currencyHandler struct {
	service currency.Service
}

// NewHandler returns AccessTokenHandler
func NewHandler(service currency.Service) CurrencyHandler {
	return &currencyHandler{
		service: service,
	}
}

func (handler *currencyHandler) GetBySymbol(w http.ResponseWriter, req *http.Request) {
	sybmol := strings.TrimSpace(chi.URLParam(req, "symbol"))
	data, err := handler.service.GetBySymbol(sybmol)
	if err != nil {
		// w.Write([]byte(fmt.Sprintf("%v, %v", err.Status, err)))
		fmt.Fprintf(w, "%v", err)
		return
	}

	res, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Fprintf(w, "%v", err1)
		return
	}
	fmt.Fprintf(w, "%s", res)
}

func (handler *currencyHandler) All(w http.ResponseWriter, req *http.Request) {
	data, err := handler.service.All()
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	res, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Fprintf(w, "%v", err1)
		return
	}
	fmt.Fprintf(w, "%s", res)
}
