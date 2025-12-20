package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/edipermadi/sadako/pkg/core/grid"
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

	_ = grid.Permutation(func(g grid.Grid) bool {
		v := g.Values()
		if _, err := fmt.Fprintf(file, "INSERT INTO grids (id, c1, c2, c3, c4, c5, c6, c7, c8, c9) VALUES (%d, %d, %d, %d, %d, %d, %d, %d, %d, %d);\n", g, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8]); err != nil {
			log.Warn().Err(err).Str("filename", filename).Msg("failed to write grid")
		}
		return false
	})
}
