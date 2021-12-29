package network

import (
	"fmt"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	net := New(t, DefaultConfig())
	for i := 0; i < 10; i++ {
		fmt.Println(net.LatestHeight())
		time.Sleep(time.Second * 3)
	}
	validator := net.Validators[0]
	// validator.
}
