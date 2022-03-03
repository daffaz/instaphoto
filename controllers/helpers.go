package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// ParseForm parse the incoming request and return error if the parsing failed
// after that, using gorilla/schema decoder, we parse the incoming request
// to dst struct
func ParseForm(r *http.Request, dst interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(dst, r.PostForm)

	if err != nil {
		return err
	}

	return nil
}
