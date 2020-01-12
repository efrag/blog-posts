package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"strings"

	"github.com/efrag/blog-posts/go_routines/domain"

	"github.com/efrag/blog-posts/go_routines/logger"

	"github.com/efrag/blog-posts/go_routines/app"

	"github.com/efrag/blog-posts/go_routines/utils"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	ctx := context.Background()
	RunSupermarket(ctx)
}

func RunSupermarket(ctx context.Context) {
	options := []string{"sequential", "one-routine", "one-routine-wait", "n-routines", "n-routines-lock"}
	if len(os.Args) < 2 {
		fmt.Printf("Need to provide one argument. Choose from: %v\n", options)
		os.Exit(1)
	}

	arg := os.Args[1]
	if !utils.AllowedOption(options, arg) {
		fmt.Printf("Not allowed argument. Choose from: %v\n", options)
		os.Exit(1)
	}

	if len(os.Args) > 2 && utils.FlagIsPresent(os.Args, "debug") {
		logger.Debug = true
	}

	if len(os.Args) > 2 && utils.FlagIsPresent(os.Args, "verbose") {
		logger.Verbose = true
	}

	numCounters := 10
	numPeople := 50

	// Initialize the queue with the people and also the counters
	q := utils.InitializeQueue(ctx, numPeople)
	cs := utils.InitializeCounters(ctx, numCounters)

	switch os.Args[1] {
	case "sequential":
		app.RunSupermarketSequential(ctx, q, cs)
		output("Run Supermarket", cs)
	case "one-routine":
		app.RunSupermarketGoRoutine(ctx, q, cs)
		output("Run Supermarket with: 1 Go Routine", cs)
	case "one-routine-wait":
		app.RunSupermarketGoRoutineWait(ctx, q, cs)
		output("Run Supermarket with: 1 Go Routine, WaitGroup", cs)
	case "n-routines":
		app.RunSupermarketNGoRoutines(ctx, q, cs)
		output(fmt.Sprintf("Run Supermarket with: %d Go routines, WaitGroup", numCounters), cs)
	case "n-routines-lock":
		app.RunSuperMarketGoRLock(ctx, q, cs)
		output(fmt.Sprintf("Run Supermarket with: %d Go routines, WaitGroup and Mutex", numCounters), cs)
	}
}

func output(name string, cs []*domain.Counter) {
	dupe := utils.GetDuplicatePeople(cs)

	if len(dupe) > 0 || logger.Verbose {
		fmt.Println(strings.Repeat("-", len(name)))
		fmt.Println(name)
		fmt.Println(strings.Repeat("-", len(name)))

		for _, c := range cs {
			fmt.Printf("Counter %d: %v\n", c.Id, c.PeopleIDs())
		}

		fmt.Println("\nPeople processed more than once:", utils.GetDuplicatePeople(cs))
	}
}
