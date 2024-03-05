package main

import (
	"bufio"
	"os"

	"github.com/haegtaw/trnsmttr/ex00/api"
	"github.com/haegtaw/trnsmttr/ex00/config"
	"github.com/haegtaw/trnsmttr/ex00/logging"
	"github.com/haegtaw/trnsmttr/ex00/net"
	"github.com/haegtaw/trnsmttr/ex00/net/gateway"
	"github.com/haegtaw/trnsmttr/ex00/storage"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup utilities things
	config.SetupConfig()
	logging.SetupLog()

	log.Info("Starting the Transmitter.Server ...")

	// Initialize storage
	storage.InitDB()
	net.LoadChannelsFromDatabase()

	// Init the gateway
	gateway.InitGateway()

	// Init the api
	api.InitAPI()

	// Waiting for input
	for {
		reader := bufio.NewReader(os.Stdin)
		cmd, _ := reader.ReadString('\n')
		//TODO: Handle the input command
		_ = cmd
	}
}
