package service

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostnamesWithLessOrEqualActiveIPs(t *testing.T) {
	mtaService := NewMtaService()

	// Test case 1: Threshold is 0
	hostnames := mtaService.GetHostnamesWithLessOrEqualActiveIPs(0)
	sort.Strings(hostnames)
	assert.ElementsMatch(t, []string{"mta-prod-3"}, hostnames)

	// Test case 2: Threshold is 1
	hostnames = mtaService.GetHostnamesWithLessOrEqualActiveIPs(1)
	sort.Strings(hostnames)
	assert.ElementsMatch(t, []string{"mta-prod-1", "mta-prod-3"}, hostnames)

	// Test case 3: Threshold is 2
	hostnames = mtaService.GetHostnamesWithLessOrEqualActiveIPs(2)
	sort.Strings(hostnames)
	assert.ElementsMatch(t, []string{"mta-prod-1", "mta-prod-2", "mta-prod-3"}, hostnames)
}
