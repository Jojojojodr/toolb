package net

import (
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/spf13/cobra"
)

var (
	wg sync.WaitGroup
	ip string
	startPort int
	endPort int
	openPorts []int
)

func scanPort(ip string, port int) {
	address := ip + ":" + strconv.Itoa(port)

	_, err := net.Dial("tcp", address)
	if err != nil {
		return
	} else {
		fmt.Printf("Port %d is open\n", port)
		openPorts = append(openPorts, port)
	}
}

var scanPortCmd = &cobra.Command{
	Use:   "scan-port",
	Short: "This command scans a range of ports.",
	Long: `This command scans a range of ports. It is used to check for open ports on a system`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Scanning ports %d to %d on %s\n\n", startPort, endPort, ip)
		for port := startPort; port <= endPort; port++ {
			wg.Add(1)
			go func(ip string, port int, wg *sync.WaitGroup) {
				defer wg.Done()
				scanPort(ip, port)
			}(ip, port, &wg)
		}
		wg.Wait()

		if len(openPorts) == 0 {
			fmt.Printf("No open ports found on %s\n", ip)
		} else {
			fmt.Printf("\nOpen ports on %s: %v\n", ip, len(openPorts))
		}
	},
}

func init() {
	scanPortCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "IP address to scan")
	scanPortCmd.Flags().IntVarP(&startPort, "port", "p", 1, "Start port to scan")
	scanPortCmd.Flags().IntVarP(&endPort, "end-port", "e", 10000, "End port to scan")

	NetCmd.AddCommand(scanPortCmd)
}
