package cake

// Cake Simulation
// Imagine three cooks in a cake shop.package cake
// one baking,
// one icing and
// one inscribing  each cake before  passing it on to  the next
// cook in the  assembly line. In a kitchen  with little space,
// each cook  that has finished a  cake must wait for  the next
// cook  to  become ready  to  accept  it; this  rendezvous  is
// analogous to communic at ion over an unbuf fered channel.
//
// To run the benchmarks
// $ go test -bench= . cake

import (
	"fmt"
	"math/rand"
	"time"
)

// Shop struct describes a typical cake shop parameters
type Shop struct {
	Verbose        bool          // handler for verbose display
	Cakes          int           // number of cakes to bake
	BakeTime       time.Duration // time to bake a cake
	BakeStdDev     time.Duration // standard deviation of baking time
	BakeBuf        int           // buffer slots between baking and icing
	NumIcers       int           // number of cooks doing icing
	IceTime        time.Duration // time to ice one cake
	IceStdDev      time.Duration // standard deviation of icing time
	IceBuf         int           // buffer slots between icing and inscribing
	InscribeTime   time.Duration // time to inscribe one cake
	InscribeStdDev time.Duration // standard deviation of inscribing time
}

type cake int

// work function blocks the calling go routine for a period of time
// that is normally distributed around d with a standard deviation
// of stddev.
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

func (s *Shop) baker(baked chan<- cake) {
	for i := 0; i < s.Cakes; i++ {
		c := cake(i)
		if s.Verbose {
			fmt.Printf("baking cake %v\n", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(baked <-chan cake, iced chan<- cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Printf("icing cake %v\n", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscribe(iced <-chan cake) {
	for i := 0; i < s.Cakes; i++ {
		c := <-iced
		if s.Verbose {
			fmt.Printf("inscribing cake %v\n", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Printf("finished cake %v\n", c)
		}
	}
}

// Work runs the simulatio `runs` times.
func (s *Shop) Work(runs int) {
	for run := 0; run < runs; run++ {
		baked := make(chan cake, s.BakeBuf)
		iced := make(chan cake, s.IceBuf)
		go s.baker(baked)
		for i := 0; i < s.NumIcers; i++ {
			go s.icer(baked, iced)
		}
		s.inscribe(iced)
	}
}
