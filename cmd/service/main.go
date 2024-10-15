package main

import (
	"blockchain/internal/cli"
	"blockchain/pkg/mid"
	"log/slog"
)
func main() {
	slog.SetDefault(mid.L)

	cli := cli.CLI{BC: nil}
	cli.Run()
}