package main

import (
	"fmt"
)

const dummyResponse = "{123:22}"

func GetDummyData() string {
	return dummyResponse
}

func main() {
	fmt.Println(GetDummyData())
}
