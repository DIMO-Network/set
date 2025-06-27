A basic set type, backed by a struct with an unexported map field. It is not safe for concurrent use, but it is safe to use without it having been initialized:

```go
var s Set[int]

s.Add(1)
s.Add(-3)
s.Add(1)
s.Len() // 2
s.Contains(1) // true

for x := range s.All() {
    fmt.Println(x)
}

sl := slices.Collect(s.All()) // Some ordering of [1 -3].
```

The code rips off [Ian Lance Taylor's snippet](https://github.com/golang/go/issues/69230#issuecomment-2327176386) from a recent Go project issue ([#69230](https://github.com/golang/go/issues/69230)) about adding a set type to the standard library. It is unclear whether that addition will happen.
