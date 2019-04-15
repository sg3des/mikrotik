package mikrotik

import (
	"time"
)

// ====================================
//
// Entities
//
// ====================================

type Ping struct {
	mikrotik *Mikrotik
	Address  string

	Count        int
	Interface    string
	Interval     int
	RoutingTable string
	Size         int
	SrcAddress   string
	TTL          int `mikrotik:"ttl"`
}

type PingResponse struct {
	Host string

	Status string

	Sent       int
	Received   int
	PacketLoss int

	TTL  int `mikrotik:"ttl"`
	Seq  int
	Size int

	Time   time.Duration
	MinRTT time.Duration `mikrotik:"min-rtt"`
	AvgRTT time.Duration `mikrotik:"avg-rtt"`
	MaxRTT time.Duration `mikrotik:"max-rtt"`
}
