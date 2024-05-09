/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/ifuryst/lol"
	"io"
	"log"
	"net/http"
)

func isHTTPPacket(payload []byte) bool {
	return bytes.HasPrefix(payload, []byte(http.MethodGet+" ")) || bytes.HasPrefix(payload, []byte(http.MethodPost+" ")) ||
		bytes.HasPrefix(payload, []byte(http.MethodPut+" ")) || bytes.HasPrefix(payload, []byte(http.MethodDelete+" ")) ||
		bytes.HasPrefix(payload, []byte(http.MethodHead+" ")) || bytes.HasPrefix(payload, []byte(http.MethodPatch+" ")) ||
		bytes.HasPrefix(payload, []byte(http.MethodConnect+" ")) || bytes.HasPrefix(payload, []byte(http.MethodOptions+" ")) ||
		bytes.HasPrefix(payload, []byte(http.MethodTrace+" "))
}

func isHTTPHeadersMatch(headers http.Header, args Argument) bool {
	if len(args.Headers) == 0 {
		return true
	}
	for key, values := range headers {
		for _, value := range values {
			for _, header := range args.Headers {
				switch header.Type {
				case HeaderTypeKey:
					if header.Key == key {
						return true
					}
				case HeaderTypeValue:
					if header.Value == value {
						return true
					}
				case HeaderTypeBoth:
					if header.Key == key && header.Value == value {
						return true
					}
				}
			}
		}
	}
	return false
}

func Capture(args Argument) {
	handle, err := pcap.OpenLive(args.Nic, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Println("Failed to open device:", err)
		return
	}
	defer handle.Close()

	portFilter := ""
	for _, port := range args.Ports {
		if portFilter == "" {
			portFilter += " and "
		} else {
			portFilter += " or "
		}
		portFilter += fmt.Sprintf("port %d", port)
	}
	filter := "tcp" + portFilter
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Println("Failed to set BPF filter:", err)
		return
	}
	fmt.Printf("Capturing HTTP traffic on interface %s...\nApply BPF filter: %s\n", args.Nic, filter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		if appLayer := packet.ApplicationLayer(); appLayer != nil {
			payload := appLayer.Payload()
			err = processPayload(args, payload)
			if err != nil {
				log.Println("Failed to process payload:", err)
			}
		}
	}
}

func processPayload(args Argument, payload []byte) error {
	if !isHTTPPacket(payload) {
		return nil
	}

	req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(payload)))
	if err != nil {
		return err
	}

	if !lol.Include(args.HttpMethods, req.Method) {
		return nil
	}

	if !lol.Include(args.Uris, req.URL.Path) {
		return nil
	}

	if !isHTTPHeadersMatch(req.Header, args) {
		return nil
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	_ = req.Body.Close()

	Replay(args, req.Method, bodyBytes, req.Header, *req.URL)
	return nil
}
