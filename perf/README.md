## Each Test
```
data size = 100000, test times = 100
------------------------------------------
jarr.Each        3.416759ms
jarr.EachInt     299.014µs
for              109.796µs
```
```
data size = 100, test times = 100000
------------------------------------------
jarr.Each        3.523µs
jarr.EachInt     294ns
for              109ns
```
```
data size = 10, test times = 1000000
------------------------------------------
jarr.Each        447ns
jarr.EachInt     30ns
for              13ns
```
## Map Test
```
data size = 100000, test times = 100
------------------------------------------
jarr.Map         14.743159ms
jarr.MapInt      446.134µs
jarr.MapToInt    7.170509ms
for (append)     997.02µs
for (allocated)  229.257µs
```
```
data size = 100, test times = 100000
------------------------------------------
jarr.Map         14.88µs
jarr.MapInt      592ns
jarr.MapToInt    7.19µs
for (append)     1.128µs
for (allocated)  299ns
```
```
data size = 10, test times = 1000000
------------------------------------------
jarr.Map         1.662µs
jarr.MapInt      108ns
jarr.MapToInt    845ns
for (append)     374ns
for (allocated)  80ns
```
## Reduce Test
```
data size = 100000, test times = 100
------------------------------------------
jarr.Reduce      9.937255ms
jarr.ReduceInt   432.825µs
jarr.ReduceToInt 6.389297ms
for              61.716µs
```
```
data size = 100, test times = 100000
------------------------------------------
jarr.Reduce      9.996µs
jarr.ReduceInt   417ns
jarr.ReduceToInt 6.872µs
for              68ns
```
```
data size = 10, test times = 1000000
------------------------------------------
jarr.Reduce      1.105µs
jarr.ReduceInt   38ns
jarr.ReduceToInt 735ns
for              8ns
```
