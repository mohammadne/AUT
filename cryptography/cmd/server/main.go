package server

import (
	"os"

	"github.com/mohammadne/university/cryptography/internal/config"
	"github.com/spf13/cobra"
)

const (
	use   = "call"
	short = "run heliograph call server"
)

func Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { main(cfg, trap) }
	cmd := &cobra.Command{Use: use, Short: short, Run: run}
	return cmd
}

func main(cfg *config.Config, trap chan os.Signal) {}
