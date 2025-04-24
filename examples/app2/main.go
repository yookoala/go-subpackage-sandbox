package main

import (
	"fmt"

	parent "github.com/yookoala/go-subpackage-sandbox"
	"github.com/yookoala/go-subpackage-sandbox/subpackage"
)

func main() {
	fmt.Println(parent.Hello())
	fmt.Println(subpackage.Hello())
}
