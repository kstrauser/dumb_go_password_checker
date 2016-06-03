package main

import (
	"encoding/json"
	"fmt"
	"github.com/kstrauser/dumb_go_password_checker/validators"
	"io/ioutil"
	"net/http"
)

type Errors struct {
	Errors []string
}

type RequestBody struct {
	Password string
}

func handleError(writer http.ResponseWriter, err error) {
	errors := Errors{[]string{err.Error()}}
	body, err := json.Marshal(errors)
	if err != nil {
		http.Error(writer, "Total fail", 500)
	}
	http.Error(writer, string(body), 400)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		handleError(writer, err)
		return
	}

	var message RequestBody
	err = json.Unmarshal(body, &message)
	if err != nil {
		handleError(writer, err)
		return
	}

	rules := validators.NewRulesWithDefaults()
	errorList := rules.Validate(message.Password)
	if len(errorList) == 0 {
		fmt.Fprintf(writer, "{\"status\": \"ok\"}\n")
		return
	}

	errors := Errors{errorList}
	body, err = json.Marshal(errors)
	if err != nil {
		http.Error(writer, "Total fail", 500)
		return
	}
	http.Error(writer, string(body), 400)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
