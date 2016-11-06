package main

import (
	"C"
	"fmt"
	"math"
	"sort"
)

type Percentile struct {
	Values    []float64
	Length    int
	Percent   float64
	Rank      float64
	Truncated float64
	Result    float64
}

//export percentile_go
func percentile_go(array []float64, percent float64) float64 {
	sort.Float64s(array)

	p := Percentile{Values: array, Length: len(array), Percent: percent}
	p.Rank = p.calculate_rank()

	if p.Length == 0 {
		return 0
	}

	if p.truncated_rank() == 0 {
		p.Result = p.truncated_for_zero()
	}

	if p.truncated_rank() == float64(p.Length) {
		p.Result = p.truncated_for_values()
	}

	if p.Result == 0 {
		if p.fractional_part() > 0 && p.truncated_rank() < float64(p.Length) {
			p.Result = p.calculate_percentile()
		} else {
			p.Result = p.Values[int(p.Rank)-1]
		}
	}

	return p.Result
}

func (p *Percentile) fractional_part() float64 {
	return math.Abs(p.Rank - p.truncated_rank())
}

func (p *Percentile) calculate_rank() float64 {
	return (p.Percent / 100) * float64(p.Length-1)
}

func (p *Percentile) calculate_percentile() float64 {
	truncated := int(p.truncated_rank())

	part_one := p.Values[truncated-1]
	part_two := p.Values[truncated]

	return (p.fractional_part() * (part_two - part_one)) + part_one
}

func (p *Percentile) truncated_rank() float64 {
	return math.Trunc(p.Rank)
}

func (p *Percentile) truncated_for_zero() float64 {
	return p.Values[0]
}

func (p *Percentile) truncated_for_values() float64 {
	return p.Values[p.Length-1]
}

func main() {
	a := []float64{1, 2, 3, 4, 5}
	fmt.Println(percentile_go(a, 25))
}
