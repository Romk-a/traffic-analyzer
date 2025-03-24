package packet

import (
	"fmt"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// PacketInfo хранит информацию о пакете
type PacketInfo struct {
	Timestamp       time.Time
	SourceIP        net.IP
	DestinationIP   net.IP
	SourcePort      int
	DestinationPort int
	Protocol        string // "TCP", "UDP", "ICMP", etc.
	Payload         []byte // Полезная нагрузка.
	// ... добавьте другие нужные поля
}

// ParsePacket разбирает пакет и возвращает PacketInfo.
func ParsePacket(packet gopacket.Packet) (*PacketInfo, error) {
	info := &PacketInfo{
		Timestamp: packet.Metadata().Timestamp,
	}

	// Разбираем Ethernet-заголовок
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		// Можем что-то делать с MAC адресами, но пока не будем.
	}

	// Разбираем IPv4-заголовок.
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		info.SourceIP = ip.SrcIP
		info.DestinationIP = ip.DstIP
	} else {
		// Обработка IPv6, если надо.
		ipLayer := packet.Layer(layers.LayerTypeIPv6)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv6)
			info.SourceIP = ip.SrcIP
			info.DestinationIP = ip.DstIP
		}
	}

	// Разбираем TCP или UDP заголовок.
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		info.SourcePort = int(tcp.SrcPort)
		info.DestinationPort = int(tcp.DstPort)
		info.Protocol = "TCP"
		info.Payload = tcp.Payload
		return info, nil // Возвращаем инфу о пакете.
	}

	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		info.SourcePort = int(udp.SrcPort)
		info.DestinationPort = int(udp.DstPort)
		info.Protocol = "UDP"
		info.Payload = udp.Payload
		return info, nil
	}

	// ICMP
	icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
	if icmpLayer != nil {
		info.Protocol = "ICMP"
		return info, nil
	}

	// Если ничего не нашли.
	return info, fmt.Errorf("unknown protocol")
}
