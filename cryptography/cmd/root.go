package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammadne/university/cryptography/cmd/server"
	"github.com/mohammadne/university/cryptography/internal/config"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	short = "short description"
	long  = `long description`
)

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)

	root := &cobra.Command{Short: short, Long: long}
	root.AddCommand(
		server.Command(config, channel),
	)

	if err := root.Execute(); err != nil {
		fmt.Println(err.Error())
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}
}
