package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Name   string  `json:"name"`
	Preamp float64 `json:"preamp"`
	Bands  []Band
}

type Band struct {
	Type      int     `json:"type"`
	Frequency float64 `json:"frequency"`
	Q         float64 `json:"q"`
	Gain      float64 `json:"gain"`
}

func main() {
	var cfg Config

	const dirName = "poweramp"

	dir, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		sourceBytes, err := os.ReadFile(filepath.Join(dirName, entry.Name()))
		if err != nil {
			panic(err)
		}

		if err = json.Unmarshal(sourceBytes, &cfg); err != nil {
			panic(err)
		}

		var builder bytes.Buffer

		builder.WriteString(fmt.Sprintf("Preamp: %f db\n", cfg.Preamp))
		for i, band := range cfg.Bands {
			builder.WriteString(fmt.Sprintf("Filter %d: ON PK Fc %f Hz Gain %f dB Q %f\n", i+1, band.Frequency, band.Gain, band.Q))
		}

		os.WriteFile(filepath.Join("apo", cfg.Name+".txt"), builder.Bytes(), 0644)

	}
}
