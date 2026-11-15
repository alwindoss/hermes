package discovery_test

import (
	"fmt"
	"hermes/internal/discovery"
	"net"
	"sync"
	"testing"
	"time"
)

func TestDiscover(t *testing.T) {
	cfg := discovery.DiscoverConfig{
		Group: &net.UDPAddr{
			IP:   net.ParseIP("224.0.0.167"),
			Port: 53317,
		},
	}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		err := discovery.RunDiscoveryServer(&cfg)
		if err != nil {
			t.Fatalf("Expected err to be nil but was %v", err)
		}
		wg.Done()
	}()

	numberOfMessages := 10
	time.Sleep(2 * time.Second)
	for i := range numberOfMessages {
		fmt.Println("Message Number", i+1)
		sender(t, cfg.Group)
		time.Sleep(3 * time.Second)
	}

	wg.Wait()
	t.FailNow()
}

func sender(t *testing.T, group *net.UDPAddr) {
	t.Helper()

	conn, err := net.DialUDP("udp4", nil, group)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	message := []byte(fmt.Sprintf("Hello from my Go application @ %s", time.Now().String()))

	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}
}
