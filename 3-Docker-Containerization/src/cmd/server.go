package main

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MESSAGE string `env:"EXAMPLE_MSG" env-required:"true"`
}

func main() {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		slog.Error("Something wrong on config load", "error", err)
	}

	fmt.Println("Golang Server")
}
