package network

import (
	"context"
	"regexp"
	"strconv"
	"tangerinefrog/HopView/internal/api"
	"tangerinefrog/HopView/internal/commands"
	"tangerinefrog/HopView/internal/models"
)

const tracerouteRegex string = `^\s*\d+[\s\*]+([a-zA-Z-\.\d]*) \((((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4})\)\s+(\d+\.*\d*) ms`

var regex *regexp.Regexp

func TraceRoute(ctx context.Context, url string, out chan<- *models.Node) error {
	lineChan := make(chan string)

	regex = regexp.MustCompile(tracerouteRegex)

	args := []string{"-4", "-m 50", "-f 2", "-N 8", url}

	go commands.StreamCommandOutput(ctx, "traceroute", lineChan, args...)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case line, ok := <-lineChan:
			if !ok {
				close(out)
				return nil
			}
			node := parseNode(line)
			if node != nil {
				resp, err := api.GetIpLocation(node.IP)
				if err != nil {
					continue
				}

				location := resp.Location
				lat, _ := strconv.ParseFloat(location.Latitude, 64)
				lon, _ := strconv.ParseFloat(location.Longitude, 64)

				if lat != 0 && lon != 0 {
					node.Latitude = lat
					node.Longitude = lon
					select {
					case <-ctx.Done():
						return ctx.Err()
					case out <- node:
					}
				}
			}
		}
	}
}

func parseNode(hopLine string) *models.Node {
	match := regex.FindStringSubmatch(hopLine)
	if len(match) < 6 {
		return nil
	}

	latency, err := strconv.ParseFloat(match[6], 32)

	if err != nil {
		return nil
	}

	node := models.Node{
		DomainName: match[1],
		IP:         match[2],
		LatencyMs:  int(latency),
	}

	return &node
}
