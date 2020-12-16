package ports

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"../utils"
	"./types"
)

//GetPorts get info from ports.json file endpoint
func GetPorts(w http.ResponseWriter, r *http.Request) {
	w, err := utils.CheckMethod(w, r, "GET")
	w, err = utils.CheckContentType(w, r)
	w.Header().Set("Content-Type", "application/json")

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
				err = errors.New("The parameter 'Limit' must to be greater than parameter 'skip'")
				utils.Check(w, "ports.GetPorts-3", err)
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
				record = ""
			}

			recordCount++
			fmt.Println(recordCount)
		}

		if recordCount == limit || errf == io.EOF {
			fmt.Println(errf)
			break
		}

	}

	ret, err := json.Marshal(portsArr)
	if utils.Check(w, "ports.Unmarshal-3", err); err != nil {
		return nil, err
	}
	fmt.Fprintf(w, string(ret))

	return nil, nil
}
