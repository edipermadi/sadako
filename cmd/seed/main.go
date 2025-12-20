package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/edipermadi/sadako/pkg/core/board"
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

	_ = board.Permutation(func(g board.Board) (bool, bool) {
		if _, err := fmt.Fprintf(file, "INSERT INTO boards (g1, g2, g3, g4, g5, g6, g7, g8, g9) VALUES (%d, %d, %d, %d, %d, %d, %d, %d, %d);\n", g[0], g[1], g[2], g[3], g[4], g[5], g[6], g[7], g[8]); err != nil {
			log.Warn().Err(err).Str("filename", filename).Msg("failed to write board")
		}
		return false, false
	})
}
