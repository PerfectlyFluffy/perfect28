package context

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
)

const MAX_LOOP_COUNT uint64 = math.MaxUint64 - 200
const MIN_LOOP_COUNT uint64 = 50000000000

const BENCH_PN8_LOOP uint64 = 2305843008139952128
const BENCH_M24_LOOP uint64 = 10000000000000
const BENCH_S24_LOOP uint64 = 500000000000

var MAX_THREAD_COUNT int = runtime.NumCPU()

var HELP_MESSAGE = `A benchmark who use brute force to find perfect numbers

BENCHMARKS:
  --pn8
      perfect 8: Loops until perfect number 8 is reached.
      It will take a very long time since perfect number 8 is ` + strconv.FormatUint(BENCH_PN8_LOOP, 10) + `

  --m24
      Multi thread 2024: Loops using all ` + strconv.Itoa(MAX_THREAD_COUNT) + ` threads until ` + strconv.FormatUint(BENCH_M24_LOOP, 10) + ` loops are done

  --s24
      Single thread 2024: Loops using a single thread until ` + strconv.FormatUint(BENCH_S24_LOOP, 10) + ` loops are done

CUSTOM LOOP:
  --loop=<COUNT>
      Min:` + strconv.FormatUint(MIN_LOOP_COUNT, 10) + `, Max:` + strconv.FormatUint(MAX_LOOP_COUNT, 10) + `

  --repeat=<COUNT>
      If COUNT=0, restarts the workload every time it completes until you kill it.
      If COUNT>1, restarts until COUNT is reached

  --thread=<COUNT>
      Number of thread to use. Default/Max: ` + strconv.Itoa(MAX_THREAD_COUNT) + `

--help
    Show this message

--version
    Show version`

// Flag variables
var _version bool
var help bool
var pn8 bool
var m24 bool
var s24 bool
var loop uint64
var repeat uint64
var thread uint

func newContextFromFlags(version string) Context {
	// Setting flags
	if flag.Lookup("version") == nil {
		flag.BoolVar(&_version, "version", false, "")
	}
	if flag.Lookup("help") == nil {
		flag.BoolVar(&help, "help", false, "")
	}
	if flag.Lookup("pn8") == nil {
		flag.BoolVar(&pn8, "pn8", false, "")
	}
	if flag.Lookup("m24") == nil {
		flag.BoolVar(&m24, "m24", false, "")
	}
	if flag.Lookup("s24") == nil {
		flag.BoolVar(&s24, "s24", false, "")
	}
	if flag.Lookup("loop") == nil {
		flag.Uint64Var(&loop, "loop", 0, "")
	}
	if flag.Lookup("repeat") == nil {
		flag.Uint64Var(&repeat, "repeat", math.MaxUint64, "")
	}
	if flag.Lookup("thread") == nil {
		flag.UintVar(&thread, "thread", 0, "")
	}
	flag.Usage = func() {}
	flag.Parse()

	// Setting context
	ctx := Context{
		LoopCount:   BENCH_PN8_LOOP,
		mode:        mode_undefined,
		ThreadCount: MAX_THREAD_COUNT,
		RepeatCount: 1,
		version:     version,
	}

	if pn8 {
		ctx.LoopCount = BENCH_PN8_LOOP
		ctx.mode = mode_pn8
	} else if m24 {
		ctx.LoopCount = BENCH_M24_LOOP
		ctx.mode = mode_m24
	} else if s24 {
		ctx.LoopCount = BENCH_S24_LOOP
		ctx.mode = mode_s24
		ctx.ThreadCount = 1
	}

	if loop != 0 {
		ctx.mode = mode_custom
		ctx.LoopCount = loop

		if ctx.LoopCount < MIN_LOOP_COUNT {
			ctx.LoopCount = MIN_LOOP_COUNT
		} else if ctx.LoopCount > MAX_LOOP_COUNT {
			ctx.LoopCount = MAX_LOOP_COUNT
		}
	}

	if repeat != math.MaxUint64 {
		ctx.mode = mode_custom
		ctx.RepeatCount = repeat
	}

	if thread > 0 && int(thread) < MAX_THREAD_COUNT {
		ctx.mode = mode_custom
		ctx.ThreadCount = int(thread)
	}

	if _version {
		ctx.PrintVersion()
		os.Exit(0)
	}

	if ctx.mode == mode_undefined || help {
		fmt.Println(HELP_MESSAGE)
		os.Exit(0)
	}

	return ctx
}
