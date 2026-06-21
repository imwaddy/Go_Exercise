package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func fetchURL(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Fetched %s with status %d\n", url, resp.StatusCode)
	return nil
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://gophercon.com",
	}

	// Create a new error group and a context
	g, ctx := errgroup.WithContext(context.Background())

	// Iterate over the URLs and launch a goroutine for each one
	for _, url := range urls {
		url := url // Capture the current value of url
		g.Go(func() error {
			return fetchURL(ctx, url)
		})
	}

	// Wait for all HTTP fetches to complete
	if err := g.Wait(); err != nil {
		fmt.Printf("Failed to fetch URLs: %v\n", err)
	} else {
		fmt.Println("Successfully fetched all URLs")
	}
}
