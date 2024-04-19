package main

import (
	"os"
	"testing"
)

// ==> HOW TO BENCHMARK **AMD** PERFORMANCE DEGRADATION (https://github.com/golang/go/issues/66866)
//
//  1. Make sure your PC have an AMD CPU.
//
//  2. Find the right value for your CPU on line 37 (arround 1 minute).
//     Use command: go build && ./perfect28 --loop=N
//
//  3. Run the following command (it will take a long time):
//
//     printf "\n==## SPLIT (preheating the cpu) ##==\n" \
//     && go test -bench=BenchmarkMain -count=2 > old.txt \
//     && sed -i -e 's/fmt.Print()/\/\/fmt.Print()/g' main.go \
//     && printf "\n==## SPLIT (git diff should show change ==> //fmt.Print()) ##==\n\n" \
//     && git diff --unified=0 \
//     && go clean \
//     && printf "\n==## SPLIT (execution time of bench old.txt will show next) ##==\n" \
//     && time go test -bench=BenchmarkMain -count=8 > old.txt \
//     && sed -i -e 's/\/\/fmt.Print()/fmt.Print()/g' main.go \
//     && printf "\n==## SPLIT (git diff should show nothing) ##==\n" \
//     && git diff --unified=0 \
//     && go clean \
//     && printf "\n==## SPLIT (execution time of bench new.txt will show next) ##==\n" \
//     && time go test -bench=BenchmarkMain -count=8 > new.txt \
//     && printf "\n"
//
//  4. Open old.txt and new.txt and remove the benchmark output (keep benchmark stats only).
//
//  5. Run following command: benchstat old.txt new.txt
func BenchmarkMain(b *testing.B) {
	// Please adjust loop count to match you CPU performance.
	// It is suggested to find a number that makes the benchmark last 1 minute.
	os.Args = []string{"TEST", "--loop", "1000000000000"}
	main()
}
