package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

var( 
	KB = uint64(1024)
)

var diskUsageCmd = &cobra.Command{
	Use:   "disk-usage",
	Short: "Disk usage information",
	Long: `This command provides information about the disk usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage(".")
		
		fmt.Println("Free:", usage.Free())
		fmt.Println("Available:", usage.Available())
		fmt.Println("Size:", usage.Size())
		fmt.Println("Used:", usage.Used())
		fmt.Println("Usage:", usage.Usage(), "%")
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
}
