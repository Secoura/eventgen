package processor

import (
	"strings"

	"github.com/icrowley/fake"
)

type HostNameProcessor struct{}

var hostNameKey = "%HostName%"

func (p HostNameProcessor) process(template string) string {
	if strings.Contains(template, hostNameKey) {
		return strings.Replace(template, hostNameKey, fake.IPv4(), -1)
	}
	// TODO
	// ideas for when we only want a specific set of IPs, not random everytime
	// random number to get how many IPs we want in a set
	// generate a set of IPs to use
	// randomly assign the IP from the set to the template
	return template
}
