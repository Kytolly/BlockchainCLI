package cli

import(
	"fmt"
)
func (cli *CLI) printUsage(){
	fmt.Println("Usage:")
    fmt.Println("  blockchain addblock -data <data>  // Add a new block to the blockchain")
    fmt.Println("  blockchain printchain             // Print the entire blockchain")
}