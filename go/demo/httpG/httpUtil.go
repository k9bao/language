package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	flagDebug = flag.Bool("debug", true, "Print message type")
)

func httpGet(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Get error:", err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ReadAll error:", err)
		return err
	}
	if v, ok := resp.Header["Date"]; ok {
		log.Println("read Data:", v[0])
		t1, _ := time.Parse(time.RFC1123, v[0])
		fmt.Println(t1)
		fmt.Println(t1.Format(time.StampMilli))
	}

	fmt.Println(string(body))
	return nil
}

func main() {
	log.Println("in")
	flag.Parse()
	httpGet("http://127.0.0.1:8061/version")

	log.Println("end")
}
