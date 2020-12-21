// Types responsible to keep all the configurations
// variables used in the whole service
package types

import (
	"os"
	"strconv"
)

//Config keep all configurations properties to be used in the whole service
type Config struct {
	// GrpcAddress address and port used to
	// connect to the gRPC server
	GrpcAddress string
	// APIAddress address used to the Rest API listener
	APIAddress  string
	// Max number of records sent to response
	RecordLimit int
}

//SetupConfig Get the environment variables values and set accordingly
//all the related Config properties
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
