package types

import (
	"os"
	"strconv"
)

//Config all configurations variable used to connect to services
type Config struct {
	GrpcAddress string
	APIAddress  string
	RecordLimit int
}

//SetupConfig Using environment variables to setup all the configuration's variables
func SetupConfig() Config {
	config := Config{}
	if os.Getenv("GRPC_ADDRESS") == "" {
		panic("No gRPC server address informed.")
	} else {
		config.GrpcAddress = os.Getenv("GRPC_ADDRESS")
	}

	if os.Getenv("API_ADDRESS") == "" {
		config.APIAddress = ":10000"
	} else {
		config.APIAddress = os.Getenv("API_ADDRESS")
	}

	if os.Getenv("RECORD_LIMIT") == "" {
		config.RecordLimit = 5000
	} else {
		config.RecordLimit, _ = strconv.Atoi(os.Getenv("RECORD_LIMIT"))
	}

	return config
}
