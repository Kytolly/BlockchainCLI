package cli

import(
	"fmt"
)

func(cli *CLI) addBlock(data string){
	// TODO：命令行接管添加区块的工作
	cli.BC.AddBlock(data)
	fmt.Println("Block added successfully:", data)
}