package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
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
		ls = append(ls, val)
	}
	fmt.Println(ls)
	bubble := BubbleSort(ls)
	checkOrder(bubble)
	heap := HeapSort(ls)
	checkOrder(heap)
	insert := InsertionSort(ls)
	checkOrder(insert)
}

func BubbleSort(array []int) []int {
	flag := false
	counter := 0
	for i := 0; i < len(array)-1; i++ {
		if flag == true {
			break
		}
		counter++
		flag = true
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				flag = false
				array = swap(array, j, j+1)
			}
			counter++
		}
	}
	return array
}

func HeapSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	lower := array[0 : len(array)/2]
	upper := array[(len(array) / 2):len(array)]
	lowerSorted := HeapSort(lower)
	upperSorted := HeapSort(upper)

	i := 0
	j := 0
	var sorted []int

	for index := 0; index < len(array); index++ {
		if lowerSorted[i] > upperSorted[j] {
			sorted = append(sorted, upperSorted[j])
			j++
		} else {
			sorted = append(sorted, lowerSorted[i])
			i++
		}

		if i == len(lowerSorted) {
			sorted = append(sorted, upperSorted[j:len(upperSorted)]...)
			break
		} else if j == len(upperSorted) {
			sorted = append(sorted, lowerSorted[i:len(lowerSorted)]...)
			break
		}
	}
	return sorted
}

func InsertionSort(array []int) []int {
	for index := 0; index < len(array); index++ {
		for key := index; key > 0; key-- {
			if array[key-1] > array[key] {
				array = swap(array, key-1, key)
			} else {
				break
			}
		}
	}
	return array
}

func swap(array []int, i int, j int) []int {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
	return array
}
func checkOrder(array []int) {
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			fmt.Println("WRONG")
			return
		}
	}
	fmt.Println("CORRECT")
}
