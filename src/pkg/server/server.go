package server

import (
	"fmt"
	"sync"
	"net/http"
	"config-as-a-service/m/v2/pkg/log"
)

const (
	defaultAddress = ""
	defaultPort = 80

	codeErr = -1
	codeOk  = 0
)

func getServerAddressString(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}

func listen(address string, wg sync.WaitGroup) {
	defer wg.Done()

	err := http.ListenAndServe(address, nil)
	if err != nil && err != http.ErrServerClosed {
		log.Logf(err.Error())
	} else {
		log.Logf("Server closed.")
	}
}

func RunServer(wg *sync.WaitGroup) int {
	log.Logf("Configuring server...")
	setHandlers()

	address := getServerAddressString(defaultAddress, defaultPort)

	waitForClose := sync.WaitGroup{}
	waitForClose.Add(1)

	go listen(address, waitForClose)

	log.Logf("Server is listening on %s", address)
	waitForClose.Wait()

	return codeOk
}
