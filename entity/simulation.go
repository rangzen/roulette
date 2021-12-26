package entity

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SimulationConf struct {
	Writer      *os.File
	Entropy     Entropy
	Roulette    Roulette
	NbRun       int
	MaxSpins    int
	StartAmount int
}

type Simulation struct {
	conf SimulationConf
}

func NewSimulation(conf SimulationConf) Simulation {
	return Simulation{
		conf: conf,
	}
}

func (s *Simulation) RunWith(strategy Strategy) Results {
	results := make(Results, 0, s.conf.NbRun)
	ttlResultSize := 0
	avgResultSize := 0
	for i := 0; i < s.conf.NbRun; i++ {
		payroll := s.conf.StartAmount
		result := make(RunResult, 0, avgResultSize)
		for payroll >= strategy.MinimalBet() {
			result = append(result, payroll)
			spin := s.conf.Entropy.Spin()
			payroll = payroll - strategy.MinimalBet() + s.conf.Roulette.PayoutWith(spin, strategy)
			if len(result) >= s.conf.MaxSpins {
				break
			}
		}
		results = append(results, result)
		ttlResultSize += len(result)
		avgResultSize = ttlResultSize / (i + 1)
	}
	results.Print(s.conf, strategy)

	return results
}

type Results []RunResult

type RunResult []int

func (r *Results) Print(simConf SimulationConf, s Strategy) {
	// Payroll for each run
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

	// Average payroll when surviving the max number of spins
	fmt.Fprint(simConf.Writer, s.Name, ",average surviving payroll,")
	var avgPrl int
	var avgPrlCount int
	for _, rr := range *r {
		if len(rr) == simConf.MaxSpins {
			avgPrl += rr[simConf.MaxSpins-1]
			avgPrlCount++
		}
	}
	fmt.Fprintf(simConf.Writer, "%.2f\n", float32(avgPrl)/float32(avgPrlCount))
}

func IntToString2(a []int) string {
	// https://stackoverflow.com/a/42159097/337726
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}

type Entropy interface {
	Spin() int
}
