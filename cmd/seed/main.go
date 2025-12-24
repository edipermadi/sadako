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

	grids := grid.Permutation(func(g grid.Grid) bool {
		v := g.Values()
		if _, err := fmt.Fprintf(file, "INSERT INTO grids (id, c1, c2, c3, c4, c5, c6, c7, c8, c9) VALUES (%d, %d, %d, %d, %d, %d, %d, %d, %d, %d);\n", g, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8]); err != nil {
			log.Warn().Err(err).Str("filename", filename).Msg("failed to write grid")
		}
		return true
	})

	formatFn := func(g grid.Grid) string {
		if g == 0 {
			return "NULL"
		}
		return fmt.Sprintf("%d", g)
	}

	prefixed := append([]grid.Grid{0}, grids...)
	for _, h1 := range prefixed {
		for _, h2 := range prefixed {
			for _, v1 := range prefixed {
				for _, v2 := range prefixed {
					for _, g := range grid.GenerateGrids([]grid.Grid{h1, h2}, []grid.Grid{v1, v2}, grids) {
						if _, err := fmt.Fprintf(file, "INSERT INTO positions (g, h1, h2, v1, v2) VALUES (%d, %s, %s, %s, %s);\n", g, formatFn(h1), formatFn(h2), formatFn(v1), formatFn(v2)); err != nil {
							log.Warn().Err(err).Str("filename", filename).Msg("failed to write grid")
						}
					}
				}
			}
		}
	}
}
