package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/sbinet/go-gnuplot"
)

func collatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func countCollatzLength(from int, to int, maxKey *int, maxVal *int) map[int]int {
	var n int
	var results map[int]int = make(map[int]int)
	var numOfElements int = 0

	for i := from; i <= to; i++ {
		n = i
		numOfElements = 0
		for n != 1 {
			n = collatz(n)
			numOfElements++
		}

		if numOfElements > *maxVal {
			*maxVal = numOfElements
			*maxKey = i
		}

		results[i] = numOfElements
	}

	return results
}

func printLengths(results map[int]int) {
	var keys []int = make([]int, 0, len(results))

	for k := range results {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("Number %d has %d elements\n", k, results[k])
	}
}

func prepareLengthPlot(x []float64, y []float64) {
	fname := ""
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()

	p.SetStyle("dots")
	p.PlotXY(y, x, "some data")
	p.SetXLabel("number")
	p.SetYLabel("number of elements in a collatz sequence")
	p.CheckedCmd("set terminal pdf")
	p.CheckedCmd("set output 'plot002.pdf'")
	p.CheckedCmd("replot")

	p.CheckedCmd("q")
}

func convertMapToSlices(m map[int]int) ([]float64, []float64) {
	var x []float64 = make([]float64, 0, len(m))
	var y []float64 = make([]float64, 0, len(m))

	for k, v := range m {
		x = append(x, float64(k))
		y = append(y, float64(v))
	}

	return x, y
}

func main() {
	from := flag.Int("from", 10, "From which number to start")
	to := flag.Int("to", 100, "To which number to end")
	flag.Parse()

	if *to < *from || *from < 1 || *to < 1 {
		fmt.Println("Wrong arguments")
		return
	}

	var maxKey int = 0
	var maxVal int = 0

	fmt.Printf("Number %d has the longest sequence with %d elements\n", maxKey, maxVal)

	results := countCollatzLength(*from, *to, &maxKey, &maxVal)

	// printLengths(results)

	x, y := convertMapToSlices(results)
	prepareLengthPlot(x, y)

	fmt.Println()
}
