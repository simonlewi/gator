package main

import (
	"fmt"
	"time"
)

func aggCommand(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	// Parse the duration string
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	// Start the ticker and run the scrape feeds function in a loop
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeedHandler(s, cmd)
		if err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
			// Loop will continue despite errors
		}
	}

	// This line will never be reached due to infinite loop
	return nil
}
