package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
)

type Response struct {
	Status  string              `json:"status"`
	Message map[string][]string `json:"message"`
}

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://dog.ceo/api/breeds/list/all", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	a := responseObject.Message
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintln(w, "Breed\tSub-breed count\tSub-breeds\t")
	for b, c := range a {
		fmt.Fprintln(w, b, "\t", len(c), "\t", strings.Join(c, ","))
	}
	w.Flush()
}
