package utils

import (
	"fmt"
	"os"
)

// Check error
func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
