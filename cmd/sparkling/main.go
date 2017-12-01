package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/toashd/sparkling"
)

var (
	title = flag.String("t", "", "Sparkline Title")
)

var usage = `Usage: sparkling [options...] <values>

Options:
  -t  Title of the sparkline.

Examples:
  $ sparkling 0 30 55 80 33 150
    ▁▂▃▄▂█

  $ sparkling -t=Awesome 23 45 23 5 1 67 8 5
    Awesome ▃▅▃▁▁█▁▁
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	var input string

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// data being piped to stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			input = s.Text()
		}
	} else {
		// stdin is from a terminal
		flag.Parse()
		if flag.NArg() < 1 {
			usageAndExit("")
		}
		if flag.Args()[0] == "version" {
			fmt.Println(sparkling.Version)
			os.Exit(0)
		}
		input = strings.Join(flag.Args(), "")
	}

	data, err := parseInputData(input)
	if err != nil {
		usageAndExit(err.Error())
	}

	sp := sparkling.New(os.Stdout)
	sp.AddSeries(data, *title)
	sp.Render()
}

// usageAndExit prints usage.
func usageAndExit(message string) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

// parseInputData parses and transforms the input data.
func parseInputData(in string) ([]float64, error) {
	var input []string
	if strings.Contains(in, ",") {
		input = strings.Split(in, ",")
	} else {
		input = strings.Split(in, " ")
	}
	var numbers []float64
	for _, arg := range input {
		if n, err := strconv.ParseFloat(strings.Trim(arg, " "), 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	if len(numbers) < 1 {
		return nil, errors.New("Could not parse provided input")
	}
	return numbers, nil
}
