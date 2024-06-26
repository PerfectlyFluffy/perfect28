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
	// fmt.Print() // COMMENT/UNCOMMENT TO FIX AMD PERFORMANCE DEGRADATION

	ctx := context.NewContext(version)
	runtime.GOMAXPROCS(ctx.ThreadCount)
	printHeader(ctx)

	buffer := make(chan int, ctx.ThreadCount)
	var wg sync.WaitGroup

	for runCount := uint64(1); runCount <= ctx.RepeatCount || ctx.RepeatCount == 0; runCount++ {
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
