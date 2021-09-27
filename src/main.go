package main

import (
	"config-as-a-service/m/v2/pkg/log"
	"config-as-a-service/m/v2/pkg/server"
	"os"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	returnCode := 0
	go func() {
		returnCode = server.RunServer(wg)
	}()
	wg.Wait()

	log.Logf("exit with code %d", returnCode)
	os.Exit(returnCode)
}
