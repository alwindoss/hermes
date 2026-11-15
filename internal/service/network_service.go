package service

import (
	"context"
	"hermes/internal/discovery"
	"log/slog"
	"net"
	"sync"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func NewNetworkService(app *application.App, mcGroup *net.UDPAddr) *NetworkService {
	svc := new(NetworkService)
	svc.multicastGroup = mcGroup
	svc.app = app
	return svc
}

type NetworkService struct {
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
	multicastGroup *net.UDPAddr
	app            *application.App
}

// ServiceStartup is called by Wails before the event loop starts.
func (s *NetworkService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	s.ctx, s.cancel = context.WithCancel(ctx)

	// if err := s.startHTTP(); err != nil {
	// 	return err
	// }
	// if err := s.startTCP(); err != nil {
	// 	return err
	// }
	if err := s.startDiscoveryServer(); err != nil {
		return err
	}
	return nil
}

// ServiceShutdown is automatically invoked by Wails during application exit.
func (s *NetworkService) ServiceShutdown() error {
	slog.Info("Shutting down network services...")

	// 1. Signal background loops to stop
	if s.cancel != nil {
		s.cancel()
	}

	s.wg.Wait()
	slog.Info("All network services stopped cleanly.")
	return nil
}

func (s *NetworkService) startDiscoveryServer() error {
	s.wg.Add(1)
	go func(wg *sync.WaitGroup) {
		cfg := discovery.DiscoverConfig{
			Group: s.multicastGroup,
			App:   s.app,
		}
		discovery.RunDiscoveryServer(&cfg)
		wg.Done()
	}(&s.wg)
	return nil
}

func (s *NetworkService) startHTTP() error {
	return nil
}

func (s *NetworkService) startTCP() error {
	return nil
}
