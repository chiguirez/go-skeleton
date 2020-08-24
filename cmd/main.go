package main

import (
	"context"

	"github.com/chiguirez/go-skeleton/internal/bootstrap"
	"github.com/chiguirez/snout"
)

func main() {
	kernel := snout.Kernel{
		RunE: bootstrap.Run,
	}

	kernelBootstrap := kernel.Bootstrap(
		"go-skeleton",
		&bootstrap.Config{},
	)

	if err := kernelBootstrap.Initialize(); err != nil {
		if err != context.Canceled {
			panic(err)
		}
	}
}
