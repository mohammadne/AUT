package server

import (
	"os"

	"github.com/mohammadne/university/cryptography/internal/config"
	"github.com/mohammadne/university/cryptography/internal/network"
	"github.com/mohammadne/university/cryptography/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	use   = "server"
	short = "run server"
)

func Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	return &cobra.Command{
		Use: use, Short: short,
		Run:       func(_ *cobra.Command, args []string) { main(cfg, args, trap) },
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"alice", "bob"},
	}
}

func main(cfg *config.Config, args []string, trap chan os.Signal) {
	log := logger.NewZap(cfg.Log)

	if len(args) != 1 {
		log.Fatal("invalid arguments given", zap.Any("args", args))
	}

	if args[0] == "alice" {
		network.NewServer(cfg.Alice, log)
	} else {
		network.NewServer(cfg.Bob, log)
	}

	field := zap.String("signal trap", (<-trap).String())
	log.Info("exiting by recieving a unix signal", field)
}
