package utils

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {
	f, err := os.OpenFile("lazynix.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	Log = log.New(f, "", log.LstdFlags)
}
