package moderator

import (
	"os"

	"github.com/mohammadne/university/cryptography/internal/config"
	"github.com/mohammadne/university/cryptography/internal/network"
	"github.com/mohammadne/university/cryptography/pkg/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	use   = "moderator"
	short = "run moderator server, it's the insecure channel"
)

func Command(cfg *config.Config, trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) { main(cfg, trap) }
	cmd := &cobra.Command{Use: use, Short: short, Run: run}
	return cmd
}

func main(cfg *config.Config, trap chan os.Signal) {
	log := logger.NewZap(cfg.Log)

	alice := network.NewClient(cfg.Alice, log)
	bob := network.NewClient(cfg.Bob, log)
	_, _ = alice, bob

	field := zap.String("signal trap", (<-trap).String())
	log.Info("exiting by recieving a unix signal", field)
}
