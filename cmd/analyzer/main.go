package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"TRAFFIC-ANALYZER/internal/capture"
	"traffic-analyzer/internal/analysis"
	"traffic-analyzer/internal/packet"
)

func main() {
	// Флаги командной строки.
	interfaceName := flag.String("i", "eth0", "Network interface to capture from")
	bpfFilter := flag.String("f", "", "BPF filter")
	// Добавьте другие флаги, если нужно (например, для вывода в файл, для БД и т.д.)
	flag.Parse()

	// Загрузка конфигурации
	// conf, err := config.LoadConfig("config.yaml") // пример с YAML файлом

	// Создание конфигурации для захвата
	captureConfig := capture.CaptureConfig{
		Interface:   *interfaceName,
		BPFFilter:   *bpfFilter,
		SnapshotLen: 65535,            // Максимальный размер пакета
		Promiscuous: true,             // Включаем promiscuous mode
		Timeout:     30 * time.Second, // Время ожидания
	}

	// Запуск захвата трафика.
	packetChan, err := capture.StartCapture(captureConfig)
	if err != nil {
		log.Fatal(err) // Завершаем с фатальной ошибкой.
	}
	defer close(packetChan)

	// Создаем анализатор.
	analyzer := analysis.NewAnalyzer()

	// Канал для обработки сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // CTRL+C

	// Запускаем горутину, которая будет слушать сигналы завершения программы
	go func() {
		<-sigChan //ждем сигнала
		fmt.Println("\nReceived interrupt signal, shutting down...")
		// Здесь можно добавить логику для корректного завершения (например, закрытия файлов, соединений).
		os.Exit(0)
	}()

	// Обрабатываем пакеты в цикле.
	fmt.Println("Capturing packets.  Press Ctrl+C to stop.")
	for pkt := range packetChan {
		p, err := packet.ParsePacket(pkt)
		if err != nil {
			//fmt.Println("Error parsing packet", err) // ошибки при парсинге.
			continue // Пропускаем пакет, если не удалось разобрать.
		}
		analyzer.AnalyzePacket(p) // Передаем пакет анализатору.
	}
	fmt.Println("Packet channel closed.")
}
