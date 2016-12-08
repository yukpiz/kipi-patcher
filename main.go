package main

import (
	"github.com/yukpiz/kipi-patcher/patcher"
)

func main() {
	if err := patcher.Execute(); err != nil {
		panic(err)
	}
	return
}
