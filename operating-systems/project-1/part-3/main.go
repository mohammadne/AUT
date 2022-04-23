package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go server()
	time.Sleep(time.Second)
	go client()

	blocker := make(chan struct{})
	<-blocker
}

const listeningPort = 5678

func server() {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		time := time.Now().Format(time.ANSIC)
		fmt.Fprintf(w, time+"\n")
	}

	http.HandleFunc("/time", handler)
	url := fmt.Sprintf(":%d", listeningPort)
	http.ListenAndServe(url, nil)

}

func client() {
	worker := func() {
		url := fmt.Sprintf("http://localhost:%d/time", listeningPort)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

	ticker := time.NewTicker(1 * time.Second)
	for {
		<-ticker.C
		worker()
	}
}
