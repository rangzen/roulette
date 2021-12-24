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

const jobsByBatch = 10000

type Job struct {
	Id       int
	Conf     *SimulationConf
	Strategy *Strategy
}

type Simulation struct {
	conf SimulationConf
}

func NewSimulation(conf SimulationConf) Simulation {
	return Simulation{
		conf: conf,
	}
}

func worker(id int, jobs <-chan Job, resultsChan chan<- []RunResult) {
	for j := range jobs {
		//fmt.Println("worker", id, "started job", j.Id)
		results := make(Results, 0, jobsByBatch)
		for i := 0; i < jobsByBatch; i++ {
			payroll := j.Conf.StartAmount
			result := make(RunResult, 0)
			for payroll >= j.Strategy.MinimalBet() {
				result = append(result, payroll)
				spin := j.Conf.EntropyEngine.Spin()
				payroll = payroll - j.Strategy.MinimalBet() + j.Conf.Roulette.PayoutWith(spin, *j.Strategy)
				if len(result) >= j.Conf.MaxSpins {
					break
				}
			}
			results = append(results, result)
		}
		resultsChan <- results
	}
}

func (s *Simulation) RunWith(strategy Strategy) Results {
	numJobs := s.conf.NbRun / jobsByBatch
	jobsChan := make(chan Job, numJobs)
	resultsChan := make(chan []RunResult, numJobs)

	for w := 1; w <= 4; w++ {
		go worker(w, jobsChan, resultsChan)
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- Job{
			Id:       j,
			Conf:     &s.conf,
			Strategy: &strategy,
		}
	}
	close(jobsChan)

	var results Results
	results = make([]RunResult, 0, s.conf.NbRun)
	for a := 1; a <= numJobs; a++ {
		rs := <-resultsChan
		results = append(results, rs...)
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
