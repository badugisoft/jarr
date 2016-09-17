package main

import (
	"fmt"
	"time"

	"github.com/badugisoft/jarr"
)

type MapIntTestItem struct {
	Name string
	Func func(data []int) []int
}

var testInt = []MapIntTestItem{
	MapIntTestItem{
		Name: "jarr.Map",
		Func: func(data []int) []int {
			return jarr.Map(data, func(e interface{}) interface{} {
				return e.(int) * e.(int)
			}).([]int)
		},
	},
	MapIntTestItem{
		Name: "jarr.MapInt",
		Func: func(data []int) []int {
			return jarr.MapInt(data, func(e int) int {
				return e * e
			})
		},
	},
	MapIntTestItem{
		Name: "jarr.MapToInt",
		Func: func(data []int) []int {
			return jarr.MapToInt(data, func(e interface{}) int {
				return e.(int) * e.(int)
			})
		},
	},
	MapIntTestItem{
		Name: "for (append)",
		Func: func(data []int) []int {
			res := []int{}
			for _, e := range data {
				res = append(res, e*e)
			}
			return res
		},
	},
	MapIntTestItem{
		Name: "for (allocated)",
		Func: func(data []int) []int {
			res := make([]int, len(data))
			for i, e := range data {
				res[i] = e * e
			}
			return res
		},
	},
}

func main() {
	dataSize := 10
	testTimes := 100000
	data := []int{}

	for i := 0; i < dataSize; i++ {
		data = append(data, i)
	}

	for _, test := range testInt {
		start := time.Now()
		for i := 0; i < testTimes; i++ {
			test.Func(data)
		}
		fmt.Printf("%14s %v\n", test.Name, time.Now().Sub(start))
	}
}
