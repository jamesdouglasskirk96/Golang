package main

import (
	"fmt"
	"github.com/jamesdouglasskirk96/Golang/Project4Testing/variables"
	// "net/http"
	// "io/ioutil"
)

func sum (x int) int {
	return 10 + x
}

func main() {
	fmt.Println(variables.MyCharCount)
	fmt.Println(sum(24))

	var x int = 1
	increment := func () int {
		x += 1
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	// res, _ := http.Get("http://www.geekwiseacademy.com/")
	// page, _ := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// fmt.Printf("%s", page)
}
