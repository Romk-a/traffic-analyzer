package capture

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// CaptureConfig хранит конфигурацию для захвата.
type CaptureConfig struct {
	Interface   string
	BPFFilter   string // BPF фильтр (Berkeley Packet Filter)
	SnapshotLen int32
	Promiscuous bool
	Timeout     time.Duration
}

// StartCapture начинает захват трафика.
func StartCapture(config CaptureConfig) (<-chan gopacket.Packet, error) {
	// Пытаемся открыть интерфейс.
	handle, err := pcap.OpenLive(config.Interface, config.SnapshotLen, config.Promiscuous, config.Timeout)
	if err != nil {
		return nil, fmt.Errorf("pcap.OpenLive failed: %w", err)
	}

	// Установка фильтра.
	if config.BPFFilter != "" {
		if err := handle.SetBPFFilter(config.BPFFilter); err != nil {
			handle.Close() // Не забываем закрывать handle в случае ошибки.
			return nil, fmt.Errorf("SetBPFFilter failed: %w", err)
		}
	}

	// Создаем канал для пакетов
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetChan := packetSource.Packets() // Получаем канал из которого будем читать

	// Возвращаем канал.  Вызывающая функция будет читать из него пакеты.
	return packetChan, nil
}
