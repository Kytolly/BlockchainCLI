package cli

import(
	"os"
)

func(cli *CLI) validateArgs(){
	// TODO: 验证是否给出命令。
	if len(os.Args) < 2 {
		panic("use command `help` to check out usage")
	}
}