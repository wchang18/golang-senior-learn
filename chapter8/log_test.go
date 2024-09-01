package chapter8

import (
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	log.Printf("Info:%s\n", "status:200")
	log.Printf("Error:%s\n", "status:500")
}

func TestLogFile(t *testing.T) {
	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.Printf("Info:%s\n", "status:200")
	log.Printf("Error:%s\n", "status:501")
}
