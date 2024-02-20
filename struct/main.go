package main

import (
	"fmt"
	//"bufio"
	//"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//https://gobyexample.com/reading-files
	//seems deprecated
	//
	//dat, err := os.ReadFile("/Users/kelley/14feb/build-go-bin/scenario/sinfo/test1/outputs/d41d8cd98f00b204e9800998ecf8427e.txt")
	//check(err)
	//fmt.Println(string(dat))
	bin := os.Getenv("BIN")
	fmt.Println(bin)
	scenario := os.Getenv("SCENARIO")
	fmt.Println(scenario)
	path := fmt.Sprintf("/Users/kelley/14feb/build-go-bin/scenario/%s/%s/outputs/d41d8cd98f00b204e9800998ecf8427e.txt", bin, scenario)
	fmt.Println(path)
	f, err := os.ReadFile("/Users/kelley/14feb/build-go-bin/scenario/sinfo/test1/outputs/d41d8cd98f00b204e9800998ecf8427e.txt")
	check(err)

	//fmt.Println(os.ReadFile("/Users/kelley/14feb/build-go-bin/scenario/sinfo/test1/outputs/d41d8cd98f00b204e9800998ecf8427e.txt"))
	//fmt.Println(io.ReadAll(f))
	fmt.Println(f)
	//b1 := make([]byte, 5)
	//n1, err
}
