package main

import (
	"os"
	"strings"
)

func main() {

	println("HeloSo")
	lisDir := strings.Split("auth comment restaurants food user session", " ")
	for _, r := range lisDir {
		os.MkdirAll("proto/"+r, 0777)

	}
}
