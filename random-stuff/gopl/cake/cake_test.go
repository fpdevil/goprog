package cake_test

import (
	"testing"
	"time"

	"github.com/fpdevil/goprog/random-stuff/gopl/cake"
)

var defaults = cake.Shop{
	Verbose:      testing.Verbose(),
	Cakes:        20,
	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber
	// Each steps takes exactly 10ms, No buffers
	cs := defaults
	cs.Work(b.N)
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	// increases total time due to channel delays.
	cs := defaults
	cs.BakeStdDev = cs.BakeTime / 4
	cs.IceStdDev = cs.IceTime / 4
	cs.InscribeStdDev = cs.InscribeTime / 4
	cs.Work(b.N)
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	// delays resulting from variability.
	cs := defaults
	cs.BakeStdDev = cs.BakeTime / 4
	cs.IceStdDev = cs.IceTime / 4
	cs.InscribeStdDev = cs.InscribeTime / 4
	cs.BakeBuf = 10
	cs.IceBuf = 10
	cs.Work(b.N)
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	// adds directly to the critical path.
	cs := defaults
	cs.IceTime = 50 * time.Millisecond
	cs.Work(b.N)
}

func BenchmarkSlowIcingMayIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	// to its sequential component, following Amdahl's Law.
	cs := defaults
	cs.IceTime = 50 * time.Millisecond
	cs.NumIcers = 5
	cs.Work(b.N)
}
