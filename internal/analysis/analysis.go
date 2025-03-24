package analysis

import (
	"fmt"
	"go-network-analyzer/internal/packet" // Импортируем пакет packet
)

type Analyzer struct {
	// Какие-то счетчики, статистика, и т.д.
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

// AnalyzePacket анализирует один пакет.
func (a *Analyzer) AnalyzePacket(packetInfo *packet.PacketInfo) {
	// Реализуйте логику анализа здесь.
	// Например, подсчет TCP пакетов:
	if packetInfo.Protocol == "TCP" {
		// a.TC ব্যাপারটাCount++
		fmt.Printf("TCP Packet: %s:%d -> %s:%d\n", packetInfo.SourceIP, packetInfo.SourcePort, packetInfo.DestinationIP, packetInfo.DestinationPort)
	}
}
