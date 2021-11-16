package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

// type KvHTTPRequest struct {
// 	Key   string `json:"Key"`
// 	Value string `json:"Value"`
// }

// Example usages
// PUT: http://localhost:9999/?op=put&key=123&value=hello
// GET: http://localhost:9999/?op=get&key=123
// DUMP: http://localhost:9999/?op=
func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
	dump, err := httputil.DumpRequest(r, false)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%q\n\n", dump)

	key := r.URL.Query().Get("key")
	op := r.URL.Query().Get("op")
	op = strings.ToLower(op)
	if op == "get" {
		kv := Get(key)
		if kv != nil {
			fmt.Fprintf(w, "GET: %+v\n\n", kv)
		} else {
			fmt.Fprintf(w, "Key %s does not exist\n", key)
		}
	} else if op == "put" {
		value := r.URL.Query().Get("value")
		if value == "" {
			fmt.Fprintf(w, "No value specified")
			return
		}

		kv := Put(key, value)
		fmt.Fprintf(w, "PUT successful. %+v\n\n", kv)
	} else {
		DumpHashTable(w)
	}
}

// Start a http server @port
func kvDataNodeServer(port int) error {
	http.HandleFunc("/", httpHandler)

	portString := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal(err)
		return errors.New("Fail to start http server")
	}
	return nil
}
