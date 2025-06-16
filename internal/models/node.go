package models

type Node struct {
	IP         string  `json:"ip"`
	DomainName string  `json:"domainName"`
	LatencyMs  int     `json:"latencyMs"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
