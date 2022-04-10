package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/callbackend1", backend1Handler)
	http.HandleFunc("/callbackend2", backend2Handler)
	http.HandleFunc("/sumbackends", sumHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func backend1Handler(w http.ResponseWriter, r *http.Request) {
	be1 := os.Getenv("BACKEND1")
	resp, _ := http.Get("http://" + be1 + "/number")
	body, _ := ioutil.ReadAll(resp.Body)
	sb := string(body)
	fmt.Fprintf(w, sb)
}

func backend2Handler(w http.ResponseWriter, r *http.Request) {
	be2 := os.Getenv("BACKEND2")
	resp, _ := http.Get("http://" + be2 + "/number")
	body, _ := ioutil.ReadAll(resp.Body)
	sb := string(body)
	fmt.Fprintf(w, sb)
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	be1 := os.Getenv("BACKEND1")
	resp1, _ := http.Get("http://" + be1 + "/number")
	body1, _ := ioutil.ReadAll(resp1.Body)
	sb1 := string(body1)
	log.Println("be1: " + sb1)

	be2 := os.Getenv("BACKEND2")
	resp2, _ := http.Get("http://" + be2 + "/number")
	body2, _ := ioutil.ReadAll(resp2.Body)
	sb2 := string(body2)
	log.Println("be2: " + sb2)

	n1, _ := strconv.Atoi(sb1)
	n2, _ := strconv.Atoi(sb2)

	dat, _ := os.ReadFile("/operation.txt")
	log.Println("op: " + string(dat))

	if strings.TrimSpace(string(dat)) == "add" {
		n3 := n1 + n2
		sb3 := strconv.Itoa(n3)
		fmt.Fprintf(w, sb3)
		log.Println("output: " + sb3)
	}

}

