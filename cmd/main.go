package main

import (
	"context"

	"github.com/emaforlin/mAVS/pkg/cli"
	"github.com/emaforlin/mAVS/pkg/node"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	config, err := cli.ParseFlags()
	if err != nil {
		log.Fatal().Err(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	n := node.New(ctx)

	if err := n.Start(0, uint(config.MiningDifficulty)); err != nil {
		log.Fatal().Err(err)
	}
}
