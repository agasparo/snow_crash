package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	file := os.Args[1:]
	dat, _ := ioutil.ReadFile(file[0])
	decrypt(string(dat))
}

func decrypt(crypted string) {

	new_str := ""

	for i := 0; i < len(crypted); i++ {
		new_str += string(crypted[i] - byte(i))
	}

	fmt.Println(new_str)
}