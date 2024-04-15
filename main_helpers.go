package main

import (
	"fmt"
	"time"

	"github.com/PerfectlyFluffy/perfect28/context"
)

func printCheckpoitZero() {
	now := time.Now()
	fmt.Printf("%s :: %3d%% done in %s\n", now.Format("2006-01-02 15:04:05"), 0, "0s")
}

func printCheckpoint(b context.Batch, timeStarted time.Time) {
	if b.IsCheckpoint {
		now := time.Now()
		fmt.Printf("%s :: %3d%% done in %s\n", now.Format("2006-01-02 15:04:05"), b.Percentage, now.Sub(timeStarted))
	}
}

func printDashedLine() {
	fmt.Println("--------------------------------------------------")
}

func printHeader(ctx context.Context) {
	repeatLabel := ""
	switch ctx.RepeatCount {
	case 0:
		repeatLabel = "Forever"
	default:
		repeatLabel = fmt.Sprintf("%d", ctx.RepeatCount)
	}

	printDashedLine()
	ctx.PrintVersion()
	printDashedLine()
	fmt.Println(ctx.ModeLabel)
	printDashedLine()
	fmt.Println("Loop count:", ctx.LoopCount)
	fmt.Println("Thread count:", ctx.ThreadCount)
	fmt.Println("Repeat count:", repeatLabel)
	fmt.Println("Batch count:", ctx.BatchCount)
	fmt.Println("Batch count per thread:", ctx.BatchCountPerThread)
	fmt.Println("Loop count per batch:", ctx.LoopCountPerBatch)
}

func printRepeatHeader(ctx context.Context, runCount uint64) {
	printDashedLine()
	if ctx.RepeatCount != 1 {
		fmt.Printf("  ==> REPEAT MODE ENABLED ==> RUN #%d\n", runCount)
		printDashedLine()
	}
}
