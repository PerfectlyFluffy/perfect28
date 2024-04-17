package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/PerfectlyFluffy/perfect28/context"
)

var version string = "-DEBUG-BUILD"

func main() {
	fmt.Print() // Temporary fix for performance degradation.
	// Command to check if bug is present: go build && time ./perfect28 --loop=1000000000000

	fmt.Print() // => 2024-04-16 21:30:06 ::   5% done in 4.046703732s => bad
	fmt.Print() // => 2024-04-16 21:30:35 ::   5% done in 3.016144871s => good
	fmt.Print() // => 2024-04-16 21:31:06 ::   5% done in 4.045381827s => bad
	fmt.Print() // => 2024-04-16 21:31:38 ::   5% done in 2.99919815s => good

	var w sync.WaitGroup
	w.Add(1) // => 2024-04-16 21:32:18 ::   5% done in 2.999105851s => good
	w.Add(1) // => 2024-04-16 21:32:48 ::   5% done in 4.052358235s => bad
	w.Add(1) // => 2024-04-16 21:33:34 ::   5% done in 2.988433397s => good
	w.Add(1) // => 2024-04-16 21:33:55 ::   5% done in 3.001821616s => good
	w.Add(1) // => 2024-04-16 21:34:33 ::   5% done in 4.05221464s => bad

	ctx := context.NewContext(version)
	runtime.GOMAXPROCS(ctx.ThreadCount)
	printHeader(ctx)

	buffer := make(chan int, ctx.ThreadCount)
	var wg sync.WaitGroup

	runCount := uint64(0)
	for {
		runCount++
		printRepeatHeader(ctx, runCount)
		printCheckpointZero()

		wg.Add(ctx.BatchCount)
		timeStarted := time.Now()
		for _, b := range ctx.Batches {
			buffer <- b.BatchId
			go func(b context.Batch) {
				defer wg.Done()
				findPerfectNumbers(b)
				printCheckpoint(b, timeStarted)
				<-buffer
			}(b)
		}
		wg.Wait()

		if runCount == ctx.RepeatCount {
			break
		}
	}
	printDashedLine()
	close(buffer)
}

func findPerfectNumbers(b context.Batch) {
	for n := b.Start; n < b.Stop; n++ {
		if n%2 == 1 {
			continue
		}

		smallDivisor := uint64(1)
		aliquotSum := uint64(1)
		bigDivisor := n / 2
		for smallDivisor < bigDivisor && bigDivisor%2 == 0 {
			smallDivisor *= 2
			aliquotSum += smallDivisor + bigDivisor
			bigDivisor /= 2
		}

		if bigDivisor%2 == 1 {
			aliquotSum += smallDivisor*2 + bigDivisor
			if aliquotSum == n && isPerfect(n, smallDivisor) {
				fmt.Println("     Perfect number found ==>", n)
			}
		}
	}
}

func isPerfect(n uint64, smallDivisor uint64) bool {
	for odd := uint64(3); odd < smallDivisor; odd += 2 {
		if n%odd == 0 {
			return false
		}
	}
	return true
}
