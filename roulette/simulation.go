package roulette

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type SimulationConf struct {
	Writer        *os.File
	EntropyEngine EntropyEngine
	Roulette      Roulette
	NbRun         int
	MaxSpins      int
	StartAmount   int
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
	rc := make(chan RunResult)
	for i := 0; i < s.conf.NbRun; i++ {
		go func() {
			payroll := s.conf.StartAmount
			result := make(RunResult, 0)
			for payroll >= strategy.MinimalBet() {
				result = append(result, payroll)
				spin := s.conf.EntropyEngine.Spin()
				payroll = payroll - strategy.MinimalBet() + s.conf.Roulette.PayoutWith(spin, strategy)
				if len(result) >= s.conf.MaxSpins {
					break
				}
			}
			rc <- result
		}()
	}

	results := make(Results, 0, s.conf.NbRun)
	for i := 0; i < s.conf.NbRun; i++ {
		result := <-rc
		results = append(results, result)
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
	fmt.Fprint(simConf.Writer, s.name, ",average spins before broke,")
	var avg int
	for _, rr := range *r {
		avg += len(rr)
	}
	fmt.Fprintln(simConf.Writer, fmt.Sprintf("%.2f", float32(avg)/float32(len(*r))))

	// Percentage of games running to the max spins limit
	fmt.Fprint(simConf.Writer, s.name, ",percentage of games running to the max spins limit,")
	var maxSpinsCount int
	for _, rr := range *r {
		if len(rr) == simConf.MaxSpins {
			maxSpinsCount++
		}
	}
	fmt.Fprintln(simConf.Writer, fmt.Sprintf("%.2f", float32(maxSpinsCount*100)/float32(simConf.NbRun)))

	// Average payroll when surviving the max number of spins
	fmt.Fprint(simConf.Writer, s.name, ",average surviving payroll,")
	var avgPrl int
	var avgPrlCount int
	for _, rr := range *r {
		if len(rr) == simConf.MaxSpins {
			avgPrl += rr[simConf.MaxSpins-1]
			avgPrlCount++
		}
	}
	fmt.Fprintln(simConf.Writer, fmt.Sprintf("%.2f", float32(avgPrl)/float32(avgPrlCount)))
}

func IntToString2(a []int) string {
	// https://stackoverflow.com/a/42159097/337726
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}

type EntropyEngine interface {
	Spin() int
}

type RandomEngine struct {
	interval int
}

func NewRandomEngine(nb int) EntropyEngine {
	return RandomEngine{
		interval: nb,
	}
}

func (r RandomEngine) Spin() int {
	return int(rand.Int31n(int32(r.interval)))
}

type ControlledEngine struct {
	target int
}

func NewControlledEngine(target int) EntropyEngine {
	return ControlledEngine{
		target: target,
	}
}

func (c ControlledEngine) Spin() int {
	return c.target
}
