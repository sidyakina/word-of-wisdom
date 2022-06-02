package quotesrepo

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Repo struct {
	quotes []string
}

func New(filePath string) (*Repo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}()

	var quotes []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	log.Printf("loaded %v quotes", len(quotes))

	if len(quotes) == 0 {
		return nil, errors.New("empty quotes list")
	}

	return &Repo{quotes: quotes}, nil
}

func (r *Repo) GetQuote() (string, error) {
	n := rand.Intn(len(r.quotes))

	quote := r.quotes[n]

	log.Printf("quotes[%v] = %v", n, quote)

	return quote, nil
}
