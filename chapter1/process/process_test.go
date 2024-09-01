package process

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestProcess1(t *testing.T) {
	for {
		log.Println("run")
		time.Sleep(time.Second)
	}
}

func TestProcess2(t *testing.T) {
	go func() {
		for {
			log.Println("run")
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("hello")
	select {}
}

func TestProcess3(t *testing.T) {
	go func() {
		for {
			log.Println("run")
			time.Sleep(time.Second)
		}
	}()

	ch := make(chan int)
	<-ch
}
