package net

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/go-ping/ping"
)

var (
	ntw string
)

func ScanNetwork(networkCIDR string, timeout time.Duration) {
	ips, err := getIPAddresses(networkCIDR)
	if err != nil {
		log.Fatalf("Failed to get IP addresses: %v", err)
	}

	var wg sync.WaitGroup
	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if isReachable(ip, timeout) {
				hostname, err := net.LookupAddr(ip)
				if err != nil || len(hostname) == 0 {
					fmt.Printf("IP: %s, Hostname: unknown\n", ip)
				} else {
					fmt.Printf("IP: %s, Hostname: %s\n", ip, hostname[0])
				}
			}
		}(ip)
	}

	wg.Wait()
}

func isReachable(ip string, timeout time.Duration) bool {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Printf("Failed to create pinger for %s: %v", ip, err)
		return false
	}

	pinger.Count = 1
	pinger.Timeout = timeout
	pinger.SetPrivileged(true) // Required for Windows; optional for Unix/Linux
	err = pinger.Run()
	if err != nil {
		log.Printf("Failed to ping %s: %v", ip, err)
		return false
	}

	stats := pinger.Statistics()
	return stats.PacketsRecv > 0
}

func getIPAddresses(networkCIDR string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(networkCIDR)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %v", err)
	}

	var ips []string
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}
	
	return nil, fmt.Errorf("network CIDR too small")
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] != 0 {
			break
		}
	}
}

// scanNetCmd represents the scanNet command
var scanNetCmd = &cobra.Command{
	Use:   "scan-network",
	Short: "This command scans a network.",
	Long: `This command scans a network. It is used to check for active hosts on a network.`,
	Run: func(cmd *cobra.Command, args []string) {
		timeout := 1 * time.Second
		fmt.Printf("Scanning network %s...\n", ntw)
		ScanNetwork(ntw, timeout)
	},
}

func init() {
	scanNetCmd.Flags().StringVarP(&ntw, "network", "n", "192.168.1.0/24", "Network CIDR to scan")

	NetCmd.AddCommand(scanNetCmd)
}
