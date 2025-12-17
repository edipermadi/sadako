package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/edipermadi/sadako/pkg/core/permutation"
	"github.com/rs/zerolog"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: seed <filename>")
		os.Exit(1)
		return
	}

	filename := filepath.Clean(args[0])
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal().Err(err).Str("filename", filename).Msg("failed to open file")
		os.Exit(1)
		return
	}
	defer func() { _ = file.Close() }()

	ctx := permutation.Context{}
	ctx.Permutation(grid.Grid{})
	for _, g := range ctx.Grids {
		if _, err := fmt.Fprintf(file, "INSERT INTO grids (c1, c2, c3, c4, c5, c6, c7, c8, c9) VALUES (%d, %d, %d, %d, %d, %d, %d, %d, %d);\n", g[0], g[1], g[2], g[3], g[4], g[5], g[6], g[7], g[8]); err != nil {
			log.Fatal().Err(err).Str("filename", filename).Msg("failed to write grid")
			os.Exit(1)
			return
		}
	}
}
