// Copyright (c) 2023 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package device

import (
	"context"
	"net"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/joshuar/go-hass-agent/internal/hass"
	"github.com/rs/zerolog/log"
)

var ipLookupHosts = map[string]map[string]string{
	"icanhazip": {"IPv4": "https://4.icanhazip.com", "IPv6": "https://6.icanhazip.com"},
	"ipify":     {"IPv4": "https://api.ipify.org", "IPv6": "https://api6.ipify.org"},
}

type address struct {
	addr net.IP
}

func (a *address) Name() string {
	switch {
	case a.addr.To4() != nil:
		return "External IPv4 Address"
	case a.addr.To16() != nil:
		return "External IPv6 Address"
	default:
		return "External IP Address"
	}
}

func (a *address) ID() string {
	switch {
	case a.addr.To4() != nil:
		return "external_ipv4_address"
	case a.addr.To16() != nil:
		return "external_ipv6_address"
	default:
		return "external_ip_address"
	}
}

func (a *address) Icon() string {
	switch {
	case a.addr.To4() != nil:
		return "mdi:numeric-4-box-outline"
	case a.addr.To16() != nil:
		return "mdi:numeric-6-box-outline"
	default:
		return "mdi:ip"
	}
}

func (a *address) SensorType() hass.SensorType {
	return hass.TypeSensor
}

func (a *address) DeviceClass() hass.SensorDeviceClass {
	return 0
}

func (a *address) StateClass() hass.SensorStateClass {
	return 0
}

func (a *address) State() interface{} {
	return a.addr.String()
}

func (a *address) Units() string {
	return ""
}

func (a *address) Category() string {
	return "diagnostic"
}

func (a *address) Attributes() interface{} {
	return nil
}

func lookupExternalIPs(ctx context.Context) []*address {

	ip4 := &address{}
	ip6 := &address{}

	for host, addr := range ipLookupHosts {
		log.Debug().Caller().
			Msgf("Trying to find external IP addresses with %s", host)
		for ver, url := range addr {
			var s string
			err := requests.
				URL(url).
				ToString(&s).
				Fetch(ctx)
			log.Debug().Caller().
				Msgf("Fetching %s address from %s", ver, url)
			if err != nil {
				log.Debug().Caller().Err(err).
					Msgf("Unable to retrieve external %s address", ver)
			} else {
				s = strings.TrimSpace(s)
				switch ver {
				case "IPv4":
					ip4.addr = net.ParseIP(s)
				case "IPv6":
					ip6.addr = net.ParseIP(s)
				}
			}
		}
		return []*address{ip4, ip6}
	}
	// At this point, we've gone through all IP checkers and not found an
	// external address
	log.Debug().Caller().Msg("Couldn't retrieve any external IP address.")
	return nil
}

func updateExternalIPSensors(ctx context.Context, status chan interface{}) {
	ips := lookupExternalIPs(ctx)
	for _, ip := range ips {
		if ip.addr != nil {
			status <- ip
		}
	}
}

func ExternalIPUpdater(ctx context.Context, status chan interface{}) {
	// Set up a ticker with the interval specified to check if the external IPs
	// have changed.
	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				log.Debug().Caller().Msg("Checking for external IP update...")
				updateExternalIPSensors(ctx, status)
			}
		}
	}()
}
