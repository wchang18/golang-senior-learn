package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.Now().AddDate(0, 0, -3))
}
