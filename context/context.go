package context

import (
	"fmt"
	"math"
)

type perfect28Mode int

const (
	mode_undefined perfect28Mode = iota
	mode_custom
	mode_pn8
	mode_m24
	mode_s24
)

type Batch struct {
	BatchId      int
	IsCheckpoint bool
	Percentage   int
	Start        uint64
	Stop         uint64
}

type Context struct {
	mode    perfect28Mode
	version string

	BatchCount          int
	BatchCountPerThread int
	Batches             []Batch
	Checkpoint          int
	LoopCount           uint64
	LoopCountPerBatch   uint64
	ModeLabel           string
	RepeatCount         uint64
	ThreadCount         int
}

func NewContext(version string) Context {
	ctx := newContextFromFlags(version)
	ctx.setModeLabel()

	batchCount := math.MaxUint16 - (math.MaxUint16 % ctx.ThreadCount)
	checkpoint := batchCount / 20

	ctx.BatchCount = batchCount
	ctx.BatchCountPerThread = batchCount / ctx.ThreadCount
	ctx.Checkpoint = checkpoint
	ctx.LoopCountPerBatch = ctx.LoopCount / uint64(batchCount)
	ctx.setBatches()
	return ctx
}

func (c *Context) PrintVersion() {
	fmt.Printf("perfect28 v%s\n", c.version)
}

func (c *Context) setBatches() {
	c.Batches = make([]Batch, c.BatchCount)
	batchSize := c.LoopCount / uint64(c.BatchCount)
	diffCount := c.LoopCount - batchSize*uint64(c.BatchCount)

	start := uint64(1)
	for i := 0; i < c.BatchCount; i++ {
		stop := start + batchSize
		if diffCount > 0 {
			stop++
			diffCount--
		}

		b := Batch{
			BatchId:      (i + 1),
			IsCheckpoint: (i+1)%c.Checkpoint == 0,
			Start:        start,
			Stop:         stop,
		}
		if b.IsCheckpoint {
			b.Percentage = b.BatchId / c.Checkpoint * 5
		}
		c.Batches[i] = b

		start = stop
	}
}

func (c *Context) setModeLabel() {
	switch c.mode {
	case mode_m24:
		c.ModeLabel = "Benchmark: Multi thread 2024 (m24)"
	case mode_s24:
		c.ModeLabel = "Benchmark: Single thread 2024 (s24)"
	case mode_pn8:
		c.ModeLabel = "Benchmark: Perfect number 8 (pn8) :: Good Luck, Have Fun! :)"
	case mode_custom:
		c.ModeLabel = "Custom loop enabled"
	default:
		panic("UNKNOWN MODE")
	}
}
