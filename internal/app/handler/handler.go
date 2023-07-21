package handler

import (
	"e-commerce/internal/app/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/form"
	"github.com/shopspring/decimal"
)

type Handler struct {
	Service service.Service
}

func H(handle func(r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handle(r)
		if err == nil {
			reply(w, data)
			return
		}
		if e, ok := service.IsValidationError(err); ok {
			validationError(w, e)
		} else if _, ok := service.IsUnauthenticatedError(err); ok {
			unauthenticatedError(w, err)
		} else {
			internalError(w, err)
		}
	}
}

func unauthenticatedError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	log.Println(err)
	log.Printf("%+v\n", err)
}

func internalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Println(err)
}

func validationError(w http.ResponseWriter, e service.ValidationError) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(e.Error()))
}

func reply(w http.ResponseWriter, data interface{}) {
	if data == nil {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
func unmarshal(r *http.Request, ptr interface{}) error {
	switch r.Method {
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(ptr)
		if err != nil {
			return service.NewValidationError(err)
		}
	case http.MethodGet:
		err := formDecoder.Decode(ptr, r.URL.Query())
		if err != nil {
			return service.NewValidationError(err)
		}
	}

	return nil
}

var formDecoder = newFormDecoder()

func newFormDecoder() *form.Decoder {
	formDecoder := form.NewDecoder()
	formDecoder.SetTagName("json")
	formDecoder.SetMode(form.ModeImplicit)
	formDecoder.RegisterCustomTypeFunc(func(ss []string) (interface{}, error) {
		v := ss[0]
		return decimal.NewFromString(v)
	}, decimal.Decimal{})
	return formDecoder
}
