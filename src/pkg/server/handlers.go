package server

import (
	"fmt"
	"net/http"
	"config-as-a-service/m/v2/pkg/log"
)

const (
	uriBase   			= "/"
	uriConfigTemplate 	= "/config_template"
	uriPing				= "/ping"
)

func setHandlers() {
	http.HandleFunc(uriPing, handlePing)
	http.HandleFunc(uriConfigTemplate, handleConfigTemplate)
	//http.Handle(requestBase, http.FileServer(http.Dir(indexPage)))
}

func handleError(w http.ResponseWriter, err string, code int) {
	log.Logf(err)
	http.Error(w, err, code)
}

func printToResponseBody(writer http.ResponseWriter, format string, args ...interface{}) {
	log.Logf(format, args...)
	_, err := fmt.Fprintf(writer, format, args...)
	if err != nil {
		handleError(writer, err.Error(), http.StatusInternalServerError)
	}
}

func closeResponseBody(writer http.ResponseWriter, request *http.Request) {
	if err := recover(); err != nil {
		log.Logf("Recovered from panic: %s.", err)
		handleError(writer, "Internal error.", http.StatusInternalServerError)
	}

	if err := request.Body.Close(); err != nil {
		log.Logf(err.Error())
	}
}

func handlePing(writer http.ResponseWriter, request *http.Request) {
	printToResponseBody(writer, "OK")
	closeResponseBody(writer, request)
}

func handleConfigTemplate(writer http.ResponseWriter, request *http.Request) {
	defer closeResponseBody(writer, request)

	log.Logf("%s %s", request.Method, request.RequestURI)

	switch request.Method {
	//case http.MethodGet:
		// TODO handle get
	//case http.MethodPost:
		// TODO handle post
	//case http.MethodPut:
		// TODO handle put
	default:
		handleError(writer, "Unsupported request type", http.StatusBadRequest)
		return
	}
}
