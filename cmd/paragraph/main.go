package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nicluo/paragraph"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(paragraph.GenerateSentence(10))
}
