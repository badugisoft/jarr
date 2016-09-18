package main

import (
	"fmt"
	"time"

	"github.com/badugisoft/jarr"
	"github.com/badugisoft/tuple"
)

var eachTests = []tuple.StringPair{
	tuple.StringPair{
		First: "jarr.Each",
		Second: func(data []int) {
			jarr.Each(data, func(e interface{}) {
				p := e.(*int)
				*p = (*p) * (*p)
			})
		},
	},
	tuple.StringPair{
		First: "jarr.EachInt",
		Second: func(data []int) {
			jarr.EachInt(data, func(e *int) {
				*e = (*e) * (*e)
			})
		},
	},
	tuple.StringPair{
		First: "for",
		Second: func(data []int) {
			for i, e := range data {
				data[i] = e * e
			}
		},
	},
}

var mapTests = []tuple.StringPair{
	tuple.StringPair{
		First: "jarr.Map",
		Second: func(data []int) {
			jarr.Map(data, func(e interface{}) interface{} {
				return e.(int) * e.(int)
			})
		},
	},
	tuple.StringPair{
		First: "jarr.MapInt",
		Second: func(data []int) {
			jarr.MapInt(data, func(e int) int {
				return e * e
			})
		},
	},
	tuple.StringPair{
		First: "jarr.MapToInt",
		Second: func(data []int) {
			jarr.MapToInt(data, func(e interface{}) int {
				return e.(int) * e.(int)
			})
		},
	},
	tuple.StringPair{
		First: "for (append)",
		Second: func(data []int) {
			res := []int{}
			for _, e := range data {
				res = append(res, e*e)
			}
		},
	},
	tuple.StringPair{
		First: "for (allocated)",
		Second: func(data []int) {
			res := make([]int, len(data))
			for i, e := range data {
				res[i] = e * e
			}
		},
	},
}

var reduceTests = []tuple.StringPair{
	tuple.StringPair{
		First: "jarr.Reduce",
		Second: func(data []int) {
			jarr.Reduce(data, func(p, e interface{}) interface{} {
				return p.(int) + e.(int)
			}, 0)
		},
	},
	tuple.StringPair{
		First: "jarr.ReduceInt",
		Second: func(data []int) {
			jarr.ReduceInt(data, func(p, e int) int {
				return p + e
			}, 0)
		},
	},
	tuple.StringPair{
		First: "jarr.ReduceToInt",
		Second: func(data []int) {
			jarr.ReduceToInt(data, func(p int, e interface{}) int {
				return p + e.(int)
			}, 0)
		},
	},
	tuple.StringPair{
		First: "for",
		Second: func(data []int) {
			sum := 0
			for _, e := range data {
				sum += e
			}
		},
	},
}

func makeIntSlice(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func main() {
	tests := []tuple.StringPair{
		tuple.NewStringPair("Each", eachTests),
		tuple.NewStringPair("Map", mapTests),
		tuple.NewStringPair("Reduce", reduceTests),
	}

	factors := []tuple.IntsPair{
		tuple.NewIntsPair(100000, 100),
		tuple.NewIntsPair(100, 100000),
		tuple.NewIntsPair(10, 1000000),
	}

	for _, k := range tests {
		fmt.Printf("## %v Test\n", k.First)
		for _, f := range factors {
			fmt.Println("```")
			fmt.Printf("data size = %v, test times = %v\n", f.First, f.Second)
			fmt.Println("------------------------------------------")
			data := makeIntSlice(f.First)

			for _, t := range k.Second.([]tuple.StringPair) {
				start := time.Now()
				for i := 0; i < f.Second; i++ {
					t.Second.(func(data []int))(data)
				}
				fmt.Printf("%-16s %v\n", t.First, time.Now().Sub(start)/time.Duration(f.Second))
			}
			fmt.Println("```")
		}
	}
}
