package entity

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SimulationConf describes the configuration of the Simulation
type SimulationConf struct {
	Writer      *os.File
	Entropy     Entropy
	Roulette    Roulette
	NbRun       int
	MaxSpins    int
	StartAmount int
}

// Simulation is the base type for any simulations
type Simulation struct {
	conf SimulationConf
}

// NewSimulation creates a new simulation with a particular configuration
func NewSimulation(conf SimulationConf) Simulation {
	fmt.Println("Simulation configuration:")
	fmt.Println("  Number of runs:", conf.NbRun)
	fmt.Println("  Max spins:", conf.MaxSpins)
	fmt.Println("  Start amount:", conf.StartAmount)
	fmt.Println("  Roulette:", conf.Roulette.Name())	
	return Simulation{
		conf: conf,
	}
}

// RunWith runs a simulation with a particular Strategy, then send back Results
func (s *Simulation) RunWith(strategy Strategy) Results {
	results := make(Results, 0, s.conf.NbRun)
	ttlResultSize := 0
	avgResultSize := 0
	for i := 0; i < s.conf.NbRun; i++ {
		bankroll := s.conf.StartAmount
		brHist := make(BankrollHistory, 0, avgResultSize)
		for bankroll >= strategy.MinimalBet() {
			brHist = append(brHist, bankroll)
			spin := s.conf.Entropy.Spin()
			bankroll = bankroll - strategy.MinimalBet() + s.conf.Roulette.PayoutWith(spin, strategy)
			if len(brHist) >= s.conf.MaxSpins {
				break
			}
		}
		results = append(results, brHist)
		ttlResultSize += len(brHist)
		avgResultSize = ttlResultSize / (i + 1)
	}
	results.Print(s.conf, strategy)

	return results
}

// Results represents results from every runs
type Results []BankrollHistory

// BankrollHistory represents the history of every bankroll before a turn during a run
type BankrollHistory []int

// Print send to a writer some statistics about the simulation's results
func (r *Results) Print(simConf SimulationConf, s Strategy) {
	// Bankroll for each run
	/*	for i, rr := range *r {
			fmt.Fprintln(simConf.Writer, fmt.Sprintf("run %d,%s", i+1, IntToString2(rr)))
		}
	*/

	// Nb of spins for each run
	/*	fmt.Fprint(simConf.Writer, "spins")
		for _, rr := range *r {
			fmt.Fprint(simConf.Writer, ","+strconv.Itoa(len(rr)))
		}
		fmt.Fprintln(simConf.Writer)
	*/

	// Average of number of spins
	fmt.Fprint(simConf.Writer, s.Name, ",average spins before broke,")
	var avg int
	for _, rr := range *r {
		avg += len(rr)
	}
	fmt.Fprintf(simConf.Writer, "%.2f\n", float32(avg)/float32(len(*r)))

	// Percentage of games running to the max spins limit
	fmt.Fprint(simConf.Writer, s.Name, ",percentage of games running to the max spins limit,")
	var maxSpinsCount int
	for _, rr := range *r {
		if len(rr) == simConf.MaxSpins {
			maxSpinsCount++
		}
	}
	fmt.Fprintf(simConf.Writer, "%.2f\n", float32(maxSpinsCount*100)/float32(simConf.NbRun))

	// Average bankroll when surviving the max number of spins
	fmt.Fprint(simConf.Writer, s.Name, ",average surviving bankroll,")
	var avgBrl int
	var avgBrlCount int
	for _, rr := range *r {
		if len(rr) == simConf.MaxSpins {
			avgBrl += rr[simConf.MaxSpins-1]
			avgBrlCount++
		}
	}
	fmt.Fprintf(simConf.Writer, "%.2f\n", float32(avgBrl)/float32(avgBrlCount))
}

// IntToString2 is an helper to transform a slice of int into a string
func IntToString2(a []int) string {
	// https://stackoverflow.com/a/42159097/337726
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}

// Entropy represents how the random spin is done
type Entropy interface {
	Spin() int
}
