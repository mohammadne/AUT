package config

import (
	"github.com/mohammadne/university/cryptography/internal/network"
	"github.com/mohammadne/university/cryptography/pkg/logger"
)

type Config struct {
	Log *logger.Config `koanf:"log"`

	Alice *network.Config
	Bob   *network.Config
}
