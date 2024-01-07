package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type Config struct {
	SessionCookie string `json:"sessionCookie"`
}

type YearDay struct {
	Year, Day int
	Path      string
}

func (yd YearDay) IsValid() bool {
	return yd.Year > 0 && yd.Day > 0
}

func (yd YearDay) URL() string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", yd.Year, yd.Day)
}

func parseYearDay(path string) YearDay {
	dayStr := filepath.Base(path)
	yearStr := filepath.Base(filepath.Dir(path))
	day, _ := strconv.Atoi(dayStr)
	year, _ := strconv.Atoi(yearStr)
	return YearDay{Year: year, Day: day, Path: path}
}

func ReadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cfgData, err := os.ReadFile(homeDir + "/.aoc-input.json")
	if err != nil {
		return nil, err
	}
	var config Config
	if err = json.Unmarshal(cfgData, &config); err != nil {
		return nil, fmt.Errorf("cannot parse config: %v", err)
	}
	if config.SessionCookie == "" {
		return nil, fmt.Errorf("no session cookie in the config")
	}
	return &config, nil
}

func humanReadableDays(yearDays []YearDay) string {
	slices.SortFunc(yearDays, func(first, second YearDay) int {
		if first.Year != second.Year {
			return first.Year - second.Year
		}
		return first.Day - second.Day
	})
	years := map[int][]int{}
	for _, yd := range yearDays {
		years[yd.Year] = append(years[yd.Year], yd.Day)
	}
	strs := make([]string, 0, len(years))
	for year, days := range years {
		var daysS []string
		for _, d := range days {
			daysS = append(daysS, strconv.Itoa(d))
		}
		strs = append(strs, fmt.Sprintf("%d: %v", year, strings.Join(daysS, ", ")))
	}
	return strings.Join(strs, "\n")
}

func do() error {
	config, err := ReadConfig()
	if err != nil {
		return err
	}
	_ = config

	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	var yearDays []YearDay
	err = filepath.Walk(curDir, filepath.WalkFunc(func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			yd := parseYearDay(path)
			if yd.IsValid() {
				yearDays = append(yearDays, yd)
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}
	log.Printf("Downloading inputs for %d days: \n%s", len(yearDays), humanReadableDays(yearDays))
	for _, yd := range yearDays {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, yd.URL(), nil)
		if err != nil {
			return err
		}
		req.Header.Add("User-Agent", "aoc-input/1.0 Go-http-client/1.1")
		req.AddCookie(&http.Cookie{Name: "session", Value: config.SessionCookie})
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		inputData, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(yd.Path, "input.txt")
		if err := os.WriteFile(targetPath, inputData, 0o600); err != nil {
			return err
		}
		log.Printf("Successfully wrote to %q: %d bytes", targetPath, len(inputData))
	}

	return nil
}

func main() {
	err := do()
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
