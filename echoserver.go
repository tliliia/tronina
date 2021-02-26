package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//comment
const postURLPath string = "/echo"
const contentTypeHeaderJSON string = "application/json"
const contentTypeHeaderURLEncoded string = "application/x-www-form-urlencoded"

func main() {
	fmt.Println("Start server")
	http.HandleFunc(postURLPath, PostHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//PostHandler is a func that handling http post request
func PostHandler(w http.ResponseWriter, r *http.Request) {

	var contentType = r.Header.Get("Content-Type")
	if contentType == contentTypeHeaderURLEncoded {
		r.ParseForm()
		var params strings.Builder
		for key, value := range r.Form {
			fmt.Fprintf(&params, "{%s:%s} ", key, value)
		}
		io.WriteString(w, params.String())
		return
	}

	if contentType == contentTypeHeaderJSON {
		var m map[string]string
		b, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(b, &m)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		params := echoStringFromParams(m)
		io.WriteString(w, params)
		return
	}
}

func echoStringFromParams(m map[string]string) string {
	var params strings.Builder
	for key, value := range m {
		fmt.Fprintf(&params, "{%s:%s} ", key, value)
	}
	return params.String()
}
