package Util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileTester(fileName string) []int {
	b, err := ioutil.ReadFile("array.txt")
	if err != nil {
		panic(err)
	}

	lsString := strings.Split(string(b), ",")
	var ls []int
	for i := 0; i < len(lsString); i++ {
		val, err := strconv.Atoi(lsString[i])
		if err != nil {
			panic(err)
		}
		_ = append(ls, val)
	}
	return ls
}
