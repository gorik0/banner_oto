
package main

import (
	"os"
	"strings"
)

func main() {
	lisDir := strings.Split("alias constants props recorder myerrors regex sanitizer", " ")
	for _, v := range lisDir {

		os.MkdirAll("internal/utils/"+v, 0777)
	}
}
