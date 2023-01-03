package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/libremfg/go-tools/cmd/sortSchema/graphql"
)

const version = "1.0.0"

var endpoint string
var help bool
var log bool

func init() {
	flag.StringVar(&endpoint, "endpoint", "http://localhost:8080", "the graphql server to proxy the request to")
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&log, "log", false, "log payloads before and after sorting")
}

func main() {

	flag.Parse()

	if help {
		fmt.Printf(`
sortSchema: http proxy a graphql schema request and returns a sorted schema

Useful for committing schema files under version control. 

Version: %s

Options: 
  --endpoint <host>  the endpoint to proxy the request to. Path is forwarded
	                   onto the host.

	--help:            shows this message

	--log:             logs the payload to the current directory before and
	                   after sorting.
`, version)
		os.Exit(0)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		contentType := r.Header.Get("Content-Type")
		resp, err := http.Post(endpoint+r.URL.Path, contentType, r.Body)
		if err != nil {
			msg := fmt.Sprintf("Failed to POST: %s", err)
			panic(msg)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			msg := fmt.Sprintf("Failed to convert Body: %s", err)
			panic(msg)
		}

		var respBody graphql.Response

		if err := json.Unmarshal(body, &respBody); err != nil {
			msg := fmt.Sprintf("Failed to Unmarshal: %s", err)
			panic(msg)
		}

		respBody.Sort()

		ordered, err := json.Marshal(respBody)
		if err != nil {
			msg := fmt.Sprintf("Failed to Marshal: %s", err)
			panic(msg)
		}

		w.WriteHeader(http.StatusOK)
		count, err := w.Write(ordered)

		if err != nil {
			msg := fmt.Sprintf("Failed to Write Response: %s", err)
			panic(msg)
		}

		fmt.Printf("%s: wrote %d bytes\n", time.Now().Format(time.RFC3339), count)

		if log {
			fileNameFragment := fmt.Sprintf("result-%d", time.Now().Unix())
			os.WriteFile(fmt.Sprintf("./%s.json", fileNameFragment), ordered, 0777)
			os.WriteFile(fmt.Sprintf("./%s-raw.json", fileNameFragment), body, 0777)
		}
	})

	addr := "0.0.0.0:8081"
	fmt.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err)
	}
}
