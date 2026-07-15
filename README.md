# rand-cty-funcs

A Go module providing random number and list-sampling functions for use in [go-cty](https://github.com/zclconf/go-cty) / HCL2 evaluation contexts.

## Installation

```
go get github.com/tsarna/rand-cty-funcs
```

## Usage

```go
import (
    randcty "github.com/tsarna/rand-cty-funcs"
    "github.com/zclconf/go-cty/cty/function"
)

// Register all functions in an HCL eval context
funcs := randcty.GetRandomFunctions()
// funcs is map[string]function.Function — merge into your eval context
```

The functions are namespaced under `rand::` (e.g. `rand::int`, `rand::choice`). HCL
parses `a::b(x)` natively as a single flat map key, so no special handling is needed —
the keys of the returned map simply contain `::`.

## Functions

### Scalar Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `rand::float` | `rand::float() number` | Returns a random float in `[0.0, 1.0)` |
| `rand::int` | `rand::int(a, b number) number` | Returns a random integer N such that `a <= N <= b` (inclusive) |
| `rand::uniform` | `rand::uniform(a, b number) number` | Returns a random float N such that `a <= N <= b` |
| `rand::gauss` | `rand::gauss(mu, sigma number) number` | Returns a random float from a Gaussian distribution with mean `mu` and standard deviation `sigma` |

#### `rand::float()`

Returns a uniformly distributed random float in `[0.0, 1.0)`.

```hcl
x = rand::float()  # e.g. 0.7342819...
```

#### `rand::int(a, b)`

Returns a uniformly distributed random integer N such that `a <= N <= b`. Both endpoints are inclusive. Returns an error if `b < a`.

```hcl
die = rand::int(1, 6)   # 1, 2, 3, 4, 5, or 6
```

#### `rand::uniform(a, b)`

Returns a uniformly distributed random float N such that `a <= N <= b`. Returns an error if `b < a`.

```hcl
temp = rand::uniform(36.0, 38.0)
```

#### `rand::gauss(mu, sigma)`

Returns a random float sampled from a Gaussian (normal) distribution with the given mean `mu` and standard deviation `sigma`.

```hcl
noise = rand::gauss(0.0, 1.0)   # standard normal
value = rand::gauss(100.0, 15.0) # IQ-like distribution
```

---

### List Functions

| Function | Signature | Description |
|----------|-----------|-------------|
| `rand::choice` | `rand::choice(list) element` | Returns a random element from a list |
| `rand::sample` | `rand::sample(list, k number) list` | Returns k unique random elements (without replacement) |
| `rand::shuffle` | `rand::shuffle(list) list` | Returns a shuffled copy of the list |

#### `rand::choice(list)`

Returns a single randomly selected element from `list`. The return type matches the element type of the list. Returns an error if the list is empty.

```hcl
color = rand::choice(["red", "green", "blue"])
```

#### `rand::sample(list, k)`

Returns a new list of `k` unique elements drawn from `list` without replacement, in random order. Returns an error if `k < 0` or `k > len(list)`. Returns an empty list if `k == 0`.

```hcl
winners = rand::sample(["alice", "bob", "carol", "dave"], 2)
```

#### `rand::shuffle(list)`

Returns a shuffled copy of `list`. The original list is not modified. Returns an empty list if the input is empty.

```hcl
deck = rand::shuffle(["A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"])
```

---

## Notes

- All randomness uses Go's `math/rand` package (automatically seeded in Go 1.20+).
- `rand::int` and `rand::uniform` truncate float arguments to `int64` and `float64` respectively.
- List functions preserve the element type of the input list.
