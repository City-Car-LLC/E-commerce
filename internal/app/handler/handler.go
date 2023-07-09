package handler

import (
	"e-commerce/internal/app/service"
	"encoding/json"
	"net/http"

	"github.com/go-playground/form"
	"github.com/shopspring/decimal"
)

type Handler struct {
	Service service.Service
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
