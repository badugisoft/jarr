package main

import (
	"fmt"
	"time"

	"github.com/badugisoft/jarr"
	"github.com/badugisoft/tuple"
)

var eachTests = []tuple.StringPair{
	tuple.NewStringPair("jarr.Each", func(data []int) interface{} {
		return jarr.Each(data, func(e interface{}) {
			p := e.(*int)
			*p = (*p) * (*p)
		})
	}),
	tuple.NewStringPair("jarr.EachInt", func(data []int) interface{} {
		return jarr.EachInt(data, func(e *int) {
			*e = (*e) * (*e)
		})
	}),
	tuple.NewStringPair("for", func(data []int) interface{} {
		for i, e := range data {
			data[i] = e * e
		}
		return data
	}),
}

var mapTests = []tuple.StringPair{
	tuple.NewStringPair("jarr.Map", func(data []int) {
		return jarr.Map(data, func(e interface{}) interface{} {
			return e.(int) * e.(int)
		})
	}),
	tuple.NewStringPair("jarr.MapInt", func(data []int) interface{} {
		return jarr.MapInt(data, func(e int) int {
			return e * e
		})
	}),
	tuple.NewStringPair("jarr.MapToInt", func(data []int) interface{} {
		return jarr.MapToInt(data, func(e interface{}) int {
			return e.(int) * e.(int)
		})
	}),
	tuple.NewStringPair("for (append)", func(data []int) interface{} {
		res := []int{}
		for _, e := range data {
			res = append(res, e*e)
		}
		return res
	}),
	tuple.NewStringPair("for (allocated)", func(data []int) interface{} {
		res := make([]int, len(data))
		for i, e := range data {
			res[i] = e * e
		}
		return res
	}),
}

var reduceTests = []tuple.StringPair{
	tuple.NewStringPair("jarr.Reduce", func(data []int) interface{} {
		return jarr.Reduce(data, func(p, e interface{}) interface{} {
			return p.(int) + e.(int)
		}, 0)
	}),
	tuple.NewStringPair("jarr.ReduceInt", func(data []int) interface{} {
		return jarr.ReduceInt(data, func(p, e int) int {
			return p + e
		}, 0)
	}),
	tuple.NewStringPair("jarr.ReduceToInt", func(data []int) interface{} {
		return jarr.ReduceToInt(data, func(p int, e interface{}) int {
			return p + e.(int)
		}, 0)
	}),
	tuple.NewStringPair("for", func(data []int) interface{} {
		sum := 0
		for _, e := range data {
			sum += e
		}
		return sum
	}),
}

func makeIntSlice(size int) []int {
	return jarr.EachIntIndex(make([]int, size), func(i int, e *int) { *e = i })
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
