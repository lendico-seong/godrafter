package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/lendico-seong/godrafter"
	"io"
	"io/ioutil"
	"os"
)

var in string

func main() {
	flag.StringVar(&in, "i", "", "input file (*.apib)")
	flag.Parse()
	if in == "" {
		fmt.Println("no input file provided. usage: -i input.apib")
		os.Exit(1)
	}
	f, err := os.Open(in)
	if err != nil {
		fmt.Printf("error while opening file %s: %v\n", in, err)
		os.Exit(1)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	input, err := ioutil.ReadAll(rd)
	if err != nil {
		fmt.Printf("error while reading file %s: %v\n", in, err)
		os.Exit(1)
	}
	errorCode := 0
	b, err := godrafter.DrafterParse(input, 0)
	if err != nil {
		fmt.Printf("error while parsing: %v\n", err)
		errorCode = 1
	}
	if b != nil {
		rd = bufio.NewReader(bytes.NewReader(b))
		_, err = io.Copy(os.Stdout, rd)
		if err != nil {
			fmt.Printf("io error: %v\n", err)
			errorCode = 2
		}
	}
	os.Exit(errorCode)
}
