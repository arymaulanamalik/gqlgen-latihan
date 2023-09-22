package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIModel struct {
	Data  interface{} `json:"data"`
	Error *Failure    `json:"error"`
}

func Success(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	model := APIModel{Data: data, Error: nil}
	js, err := json.Marshal(model)
	if err != nil {
		Error(w, WithMessage(CodeInternal, "Unknown error"), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(js)
}

func Error(w http.ResponseWriter, errModel Failure, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	model := APIModel{Data: nil, Error: &errModel}
	js, err := json.Marshal(model)
	if err != nil {
		Error(w, WithMessage(CodeInternal, "Unknown error"), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(js)
}

func Unmarshal(b []byte, target interface{}) error {
	var out APIModel
	err := json.Unmarshal(b, &out)
	if err != nil {
		return err
	}

	if out.Error != nil {
		byt, _ := json.Marshal(out.Error)
		return errors.New(string(byt))
	}

	dataByt, _ := json.Marshal(out.Data)

	err = json.Unmarshal(dataByt, &target)
	return err
}
