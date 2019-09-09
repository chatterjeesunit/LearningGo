package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}


func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/Users/in-sunit.chatterjee/test.logger", os.O_CREATE|os.O_APPEND|os.O_RDWR,	0744)
	log.SetOutput(logFileLocation)
}


func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}