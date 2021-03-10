package test

import (
	"encoding/json"
	"log"
	"testing"
)

type nvrProperty struct {
	SizeReduce int64 `json:"path-size-reduce"`
}

func JsonDemo() {
	sParams0 := nvrProperty{SizeReduce: 999}
	body, err := json.Marshal(sParams0)
	if err != nil {
		log.Println("Marshal fail")
		return
	}

	sParams := nvrProperty{}
	err = json.Unmarshal(body, &sParams)
	if err == nil {
		log.Printf("%+v", sParams)
		return
	}
}

func TestJson(t *testing.T) {
	JsonDemo()
}
