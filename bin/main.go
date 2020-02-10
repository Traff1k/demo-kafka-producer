package main

import (
	"fmt"
	"github.com/Traff1k/demo-kafka-producer/dummydata"
)

func main() {
	fmt.Println(dummydata.SimpleGet(dummydata.GetRequester, "https://httpbin.org/ip"))
}
