package discovery

import (
	"fmt"
	"net"
	"sync"
	"time"

	"wtg_desktop/internal/config"
	"wtg_desktop/internal/logger"
)

type DiscoveryService struct {
	stopCh chan struct{}
	wg     sync.WaitGroup
}

func NewDiscoveryService() *DiscoveryService {
	return &DiscoveryService{
		stopCh: make(chan struct{}),
	}
}

func (ds *DiscoveryService) Start() {
	if !config.IsUDPDiscoveryEnabled() {
		logger.Info("UDP Discovery Service is disabled in configuration.")
		return
	}

	ds.wg.Add(1)
	go ds.sendDiscoveryMessages()
}

func (ds *DiscoveryService) Stop() {
	if !config.IsUDPDiscoveryEnabled() {
		return
	}
	close(ds.stopCh)
	ds.wg.Wait()
	logger.Info("UDP Discovery Service stopped.")
}

func (ds *DiscoveryService) getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", fmt.Errorf("failed to dial UDP for local IP: %w", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func (ds *DiscoveryService) sendDiscoveryMessages() {
	defer ds.wg.Done()

	localIP, err := ds.getLocalIP()
	if err != nil {
		logger.Error("Error getting local IP for UDP discovery: %v", err)
		return
	}

	listenAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", config.GetUDPDiscoveryPort()))
	if err != nil {
		logger.Error("Error resolving UDP listen address: %v", err)
		return
	}

	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		logger.Error("Error listening on UDP for discovery: %v", err)
		return
	}
	defer conn.Close()

	broadcastAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%d", config.GetUDPDiscoveryBroadcastIP(), config.GetUDPDiscoveryPort()))
	if err != nil {
		logger.Error("Error resolving UDP broadcast address: %v", err)
		return
	}

	message := []byte(fmt.Sprintf("WTG_SERVER_INFO:%v:%v", localIP, config.GetServerPort()))
	logger.Info("Starting UDP discovery service. Sending message: %v to %v every %v", string(message), broadcastAddr.String(), config.GetUDPDiscoveryInterval())

	ticker := time.NewTicker(config.GetUDPDiscoveryInterval())
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, err := conn.WriteToUDP(message, broadcastAddr)
			if err != nil {
				logger.Error("Error sending UDP discovery message: %v", err)
			} else {
				logger.Debug("Sent UDP discovery message to %v", broadcastAddr.String())
			}
		case <-ds.stopCh:
			return
		}
	}
}
