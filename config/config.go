package config

import (
	"log"
	"os"
)

var (
	sigKey string
)

func init(){
	sigKey = os.Getenv("SIG_KEY")
	if sigKey == "" {
		log.Println("SIG_KEY environment variable is empty")
	}
}

func GetSigKey() string {
	return sigKey
}