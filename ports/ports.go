package ports

import (
	pb "../portsgrpc"
	cfg "../types"
	"../utils"
	"./types"
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//GetPorts get info from ports.json file endpoint
func GetPorts(w http.ResponseWriter, r *http.Request) {
	w, err := utils.CheckMethod(w, r, "GET")
	w, err = utils.CheckContentType(w, r)
	w.Header().Set("Content-Type", "application/json")
	config := cfg.SetupConfig()

	if err == nil {
		skip := 0
		if r.URL.Query().Get("skip") != "" {
			skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
			utils.Check(w, "ports.GetPorts-1", err)
		}

		limit := skip + 100
		if r.URL.Query().Get("limit") != "" {
			limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
			utils.Check(w, "ports.GetPorts-2", err)

			if err == nil && limit <= skip {
				err = errors.New("The parameter 'Limit' must be greater than parameter 'skip'")
				utils.Check(w, "ports.GetPorts-3", err)
			}

			//check whether difference between limit and skip is bigger
			//than 1000 records
			if (limit - skip) > config.RecordLimit {
				err = errors.New(fmt.Sprintf("Difference between parameter 'Limit' and parameter 'skip' must be lower than %d records", config.RecordLimit))
				utils.Check(w, "ports.GetPorts-4", err)
			}
		}

		if err == nil {
			Unmarshal(w, r, skip, limit)
		}
	}
}

//Unmarshal Read []byte from ports.json and Unmarshal as Ports type
//slicing the object and counting to delivery only the number
//of records informed from the skip to limit
func Unmarshal(w http.ResponseWriter, r *http.Request, skip int, limit int) ([]byte, error) {
	// open ports.json file
	f, err := os.Open("/tmp/ports.json")
	if utils.Check(w, "ports.Unmarshal-1", err); err != nil {
		return nil, err
	}
	defer f.Close()

	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.
	rf := bufio.NewReader(f)
	var portsArr []types.Ports
	recordCount := 0
	var recordStart = regexp.MustCompile(`("[A-Z]{5}"\: \{)`) // head
	var recordEnd = regexp.MustCompile(`(  \}(\,|)\n)`)       // tail
	record := ""
	writingRecord := false
	pbPorts := pb.Ports{}

	for {
		fb, errf := rf.ReadBytes('\n')
		str := string(fb)

		if recordStart.MatchString(str) {
			writingRecord = true
		}

		if writingRecord {
			if recordCount >= skip {
				record += str
			}
		}

		if recordEnd.MatchString(str) {
			writingRecord = false

			if recordCount >= skip {
				log.Printf("[if recordCount(%v) >= skip(%v)]: record(%v)\n", recordCount, skip, record)
				if record != "" {
					if record[len(record)-2:len(record)-1] == "," {
						record = record[:len(record)-2]
					}

					record = "{\n" + record + "\n}"
					fmt.Println(record)

					var ports types.Ports
					err = json.Unmarshal([]byte(record), &ports)
					if utils.Check(w, "ports.Unmarshal-2", err); err != nil {
						return nil, err
					}
					portsArr = append(portsArr, ports)

					//gRPC array append
					portId := reflect.ValueOf(ports).MapKeys()[0].String()
					pbPorts.PortsBody = append(pbPorts.PortsBody, &pb.PortsBody{
						PortId: portId,
						Name: ports[portId].Name,
						City: ports[portId].City,
						Country: ports[portId].Country,
						Alias: ports[portId].Alias,
						Regions: ports[portId].Regions,
						Coordinates: ports[portId].Coordinates,
						Province: ports[portId].Province,
						Timezone: ports[portId].Timezone,
						Unlocs: ports[portId].Unlocs,
						Code: ports[portId].Code,
					})
					record = ""
				}
			}

			recordCount++
			fmt.Println(recordCount)
		}

		if recordCount == limit || errf == io.EOF {
			fmt.Println(errf)
			break
		}

	}

	//send parallelized gRPC message
	go func(in *pb.Ports) {
		config := cfg.SetupConfig()

		/***** Initialize manual resolver and Dial *****/
		r := manual.NewBuilderWithScheme("whatever")
		// Set up a connection to the server.
		conn, err := grpc.Dial(r.Scheme()+":///test.server",
			grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
		if err != nil {
			log.Printf("did not connect: %v", err)
		}
		defer conn.Close()

		// Manually provide resolved addresses for the target.
		grpcAddresses := strings.Split(config.GrpcAddress, ",")
		address := []resolver.Address{}
		for _, add := range grpcAddresses {
			address = append(address, resolver.Address{Addr: add})
		}
		r.UpdateState(resolver.State{Addresses: address})

		c := pb.NewPortsDbClient(conn)

		// Setting a 150ms timeout on the RPC.
		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()

		resp, err := c.Upsert(ctx, in)
		if err != nil {
			log.Printf("could not reach out gRPC server: %v\n", err)
		}
		if resp != nil {
			log.Printf("Code: %s - Message: %s", resp.Code, resp.Message)
		}

		/***** Wait for user exiting the program *****/
		// Unless you exit the program (e.g. CTRL+C), channelz data will be available for querying.
		// Users can take time to examine and learn about the info provided by channelz.
		select {}
	}(&pbPorts)

	ret, err := json.Marshal(portsArr)
	if utils.Check(w, "ports.Unmarshal-3", err); err != nil {
		return nil, err
	}
	fmt.Fprintf(w, string(ret))

	return nil, nil
}
