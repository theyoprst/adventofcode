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

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/anaskhan96/soup"

	"github.com/theyoprst/adventofcode/aoc/htmlparser"
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

func (yd YearDay) ProblemURL() string {
	return fmt.Sprintf("https://adventofcode.com/%d/day/%d", yd.Year, yd.Day)
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
	var downloadInput []YearDay
	var downloadProblem []YearDay
	err = filepath.Walk(curDir, filepath.WalkFunc(func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			yd := parseYearDay(path)
			if yd.IsValid() {
				// Check that file input.txt exists in this path:
				if _, err := os.Stat(filepath.Join(yd.Path, "input.txt")); os.IsNotExist(err) {
					downloadInput = append(downloadInput, yd)
				}
				if _, err := os.Stat(filepath.Join(yd.Path, "part2.md")); os.IsNotExist(err) {
					downloadProblem = append(downloadProblem, yd)
				}
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}

	log.Printf("Downloading inputs for %d days: \n%s", len(downloadInput), humanReadableDays(downloadInput))
	for _, yd := range downloadInput {
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
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
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

	log.Printf("Downloading problems for %d days: \n%s", len(downloadProblem), humanReadableDays(downloadProblem))
	for _, yd := range downloadProblem {
		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, yd.ProblemURL(), nil)
		if err != nil {
			return err
		}
		req.Header.Add("User-Agent", "aoc-input/1.0 Go-http-client/1.1")
		req.AddCookie(&http.Cookie{Name: "session", Value: config.SessionCookie})
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		html, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		doc := soup.HTMLParse(string(html))
		articles := doc.FindAll("article")
		var paragraphs []soup.Root
		for i, article := range articles {
			articleHTMLPath := filepath.Join(yd.Path, fmt.Sprintf("part%d.html", i+1))
			html := innerHTML(article)
			if err := os.WriteFile(articleHTMLPath, []byte(html), 0o600); err != nil {
				return err
			}
			log.Printf("Successfully wrote to %q: %d bytes", articleHTMLPath, len(html))

			articleMDPath := filepath.Join(yd.Path, fmt.Sprintf("part%d.md", i+1))
			markdown, err := htmltomarkdown.ConvertString(article.HTML())
			if err != nil {
				return err
			}
			if err := os.WriteFile(articleMDPath, []byte(markdown), 0o600); err != nil {
				return err
			}
			log.Printf("Successfully wrote to %q: %d bytes", articleMDPath, len(markdown))

			paragraphs = append(paragraphs, article.Children()...)
		}
		examples, err := htmlparser.ExtractExamples(paragraphs)
		if err != nil {
			return fmt.Errorf("extract examples: %w", err)
		}
		if len(examples) == 0 {
			return fmt.Errorf("no examples found")
		}
		for i, example := range examples {
			examplePath := filepath.Join(yd.Path, fmt.Sprintf("input_ex%d.txt", i+1))
			if err := os.WriteFile(examplePath, []byte(example), 0o600); err != nil {
				return err
			}
			log.Printf("Successfully wrote to %q: %d bytes", examplePath, len(example))
		}
	}

	return nil
}

func innerHTML(node soup.Root) string {
	var sb strings.Builder
	for _, c := range node.Children() {
		sb.WriteString(c.HTML())
	}
	return sb.String()
}

func main() {
	err := do()
	if err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
