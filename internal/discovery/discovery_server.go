package discovery

import (
	"hermes/internal/message"
	"log"
	"log/slog"
	"net"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DiscoverConfig struct {
	Group *net.UDPAddr
	App   *application.App
}

func RunDiscoveryServer(cfg *DiscoverConfig) error {
	conn, err := net.ListenMulticastUDP("udp4", nil, cfg.Group)
	if err != nil {
		return err
	}
	defer conn.Close()
	slog.Info("Listening for multicast messages...", "Group", cfg.Group)
	buffer := make([]byte, 2048)
	for {
		n, source, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("Read error:", err)
			continue
		}
		// fmt.Printf("Received from %s: %s\n", source, string(buffer[:n]))
		msg := string(buffer[:n])
		cfg.App.Event.Emit(message.DeviceMessageReceived, msg)
		slog.Info("Received from", "Source", source, "Data", string(buffer[:n]))
	}
}
