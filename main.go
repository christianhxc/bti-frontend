package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title"`
	Info  ItemInfo `json:"info"`
}

type Payload struct {
	Movies []Item `json:"movies"`
	Count  int    `json:"count"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(os.Getenv("BTI_BACKEND_ENDPOINT"))
	if err != nil {
		log.Fatal("Got error getting the list of movies:", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error! Look at the logs for more details.")
	}

	defer resp.Body.Close()

	var payload Payload
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &payload)

	t, _ := template.ParseFiles("home.html")
	t.Execute(w, payload)
	log.Print("Done getting the list of movies ...")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
