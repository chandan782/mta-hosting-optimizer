package service

import (
	"github.com/chandan782/mta-hosting-optimizer.git/model"
)

type MtaService struct {
}

func NewMtaService() *MtaService {
	return &MtaService{}
}

func (s *MtaService) GetHostnamesWithLessOrEqualActiveIPs(threshold int) []string {
	ipConfigData := getMockIpConfigData()
	hostnamesCount := make(map[string]int)

	// Count the number of active IPs per hostname
	for _, ipConfig := range ipConfigData {
		if ipConfig.Active {
			hostnamesCount[ipConfig.Hostname]++
		} else if !ipConfig.Active {
			count, ok := hostnamesCount[ipConfig.Hostname]
			if !ok {
				hostnamesCount[ipConfig.Hostname] = 0
			}
			hostnamesCount[ipConfig.Hostname] = count
		}
	}

	// Filter hostnames based on the threshold
	filteredHostnames := make([]string, 0)
	for hostname, count := range hostnamesCount {
		if count <= threshold {
			filteredHostnames = append(filteredHostnames, hostname)
		}
	}

	return filteredHostnames
}

func getMockIpConfigData() []model.IpConfig {
	// Mock data for testing
	return []model.IpConfig{
		{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
		{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
		{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
		{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
	}
}
