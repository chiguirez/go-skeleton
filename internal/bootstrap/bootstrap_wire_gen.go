// +build !wireinject

package bootstrap

import (
	"context"

	"github.com/chiguirez/kfk/v2"
	"golang.org/x/sync/errgroup"
)

type Config struct {
}

func Run(ctx context.Context, cfg Config) error {
	g, ctx := errgroup.WithContext(ctx)

	// example with kfk consumer
	consumer, err := InitializeKafkaSubscriber(ctx, cfg)
	if err != nil {
		return err
	}

	g.Go(func() error {
		return consumer.Start(ctx)
	})

	return g.Wait()
}

func InitializeKafkaSubscriber(ctx context.Context, cfg Config) (kfk.KafkaConsumer, error) {
	return kfk.KafkaConsumer{}, nil
}
