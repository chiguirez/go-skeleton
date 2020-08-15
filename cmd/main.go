package main

import (
	"context"

	"github.com/chiguirez/go-sekeleton/internal/bootstrap"
	"github.com/chiguirez/snout"
)

func main() {
	kernel := snout.Kernel{
		RunE: bootstrap.Run,
	}

	kernelBootstrap := kernel.Bootstrap(
		"telemetry-consumer",
		&bootstrap.Config{},
	)

	if err := kernelBootstrap.Initialize(); err != nil {
		if err != context.Canceled {
			panic(err)
		}
	}
}
