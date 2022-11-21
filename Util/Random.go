package Util

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func GenerateArray() {
	min := 1
	max := 100
	var b strings.Builder
	for i := 0; i < 50; i++ {
		fmt.Fprintf(&b, "%d,", rand.Intn(max-min)+min)
	}
	b.WriteString(strconv.Itoa(rand.Intn(max-min) + min))

	file, err := os.Create("array.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := file.Write([]byte(b.String())); err != nil {
		panic(err)
	}
}
