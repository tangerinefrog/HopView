package network

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"tangerinefrog/HopView/internal/api"
	"tangerinefrog/HopView/internal/commands"
	"tangerinefrog/HopView/internal/models"
)

const tracerouteRegex string = `^\s*\d+[\s\*]+([a-zA-Z-\.\d]*) \((((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4})\)\s+(\d+\.*\d*) ms`

var regex *regexp.Regexp

func TraceRoute(url string) ([]models.Node, error) {
	out, err := commands.Execute("traceroute", url)
	if err != nil {
		return nil, err
	}

	regex = regexp.MustCompile(tracerouteRegex)
	nodes := make([]models.Node, 0)

	results := strings.Split(out, "\n")
	for _, r := range results {
		node := extractNode(r)
		if node != nil {
			resp, err := api.GetIpLocation(node.IP)
			fmt.Printf("LOCATION: %v", err)
			if err != nil {
				continue
			}

			location := resp.Location
			lat, _ := strconv.ParseFloat(location.Latitude, 64)
			lon, _ := strconv.ParseFloat(location.Longitude, 64)

			if lat != 0 && lon != 0 {
				node.Latitude = lat
				node.Longitude = lon

				nodes = append(nodes, *node)
			}
		}
	}

	return nodes, nil
}

func extractNode(hopLine string) *models.Node {
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
