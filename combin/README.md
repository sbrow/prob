

# combin
`import "github.com/sbrow/prob/combin"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package combin contains various tools for computing combinatorics.

### License
This Program is licensed under the GNU General Public License V3.0.
A copy of this license is provided with the program, and can be found
here:


	<a href="https://github.com/sbrow/combin/blob/master/LICENSE">https://github.com/sbrow/combin/blob/master/LICENSE</a>




## <a name="pkg-index">Index</a>
* [func CombineSets(a, b Set) [][]interface{}](#CombineSets)
* [func Fact(n int) int](#Fact)
* [func NCR(rep bool, n int, r ...int) int](#NCR)
* [func NPR(rep bool, n int, r ...int) int](#NPR)
* [func PermuteR(set string, min, max int) []string](#PermuteR)
* [func Product(i, n int) (prod int)](#Product)
* [func ProductFunc(k, n int, f func(int, ...int) int, params ...int) int](#ProductFunc)
* [func SumFunc(k, n int, f func(int, ...int) int, params ...int) int](#SumFunc)
* [type Set](#Set)
  * [func NewSet(v ...interface{}) Set](#NewSet)
  * [func (s *Set) Add(v interface{})](#Set.Add)
  * [func (s Set) Combine() [][2]interface{}](#Set.Combine)
  * [func (s *Set) Size() int](#Set.Size)

#### <a name="pkg-examples">Examples</a>
* [CombineSets](#example_CombineSets)
* [NCR (Repitition)](#example_NCR_repitition)
* [NPR (NoRepitition)](#example_NPR_noRepitition)
* [NPR (Repitition)](#example_NPR_repitition)
* [PermuteR](#example_PermuteR)
* [SumFunc](#example_SumFunc)

#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/sbrow/prob/combin/doc.go) [funcs.go](/src/github.com/sbrow/prob/combin/funcs.go) [permute.go](/src/github.com/sbrow/prob/combin/permute.go) [set.go](/src/github.com/sbrow/prob/combin/set.go) 





## <a name="CombineSets">func</a> [CombineSets](/src/target/set.go?s=1180:1222#L50)
``` go
func CombineSets(a, b Set) [][]interface{}
```
CombineSets returns all combinations (without replacement)
of the items in set a with the items in set b.



## <a name="Fact">func</a> [Fact](/src/target/funcs.go?s=2640:2660#L113)
``` go
func Fact(n int) int
```
Fact returns the factorial of n. Specifically, it returns Product(1, n).



## <a name="NCR">func</a> [NCR](/src/target/funcs.go?s=428:467#L18)
``` go
func NCR(rep bool, n int, r ...int) int
```
NCR returns the number of combinations when choosing r objects from n. rep determines
whether or not to count with repetition.

If passed more than one r, NCR will return the sum of each combination of n and r.

This means that calling:


	NCR(false, 10, 1, 2, 3)

Will return the same result as:


	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NCR(false, n, i)
	}



## <a name="NPR">func</a> [NPR](/src/target/funcs.go?s=1353:1392#L60)
``` go
func NPR(rep bool, n int, r ...int) int
```
NPR returns the number of permutations when choosing r objects from n. rep determines
whether or not to count with repetition.

If passed more than one r, NCR will return the sum of each permutation of n and r.

This means that calling:


	NPR(false, 10, 1, 2, 3)

Will return the same result as:


	sum := 0
	for i := 1; i <= 3; i++ {
		sum += NPR(false, n, i)
	}



## <a name="PermuteR">func</a> [PermuteR](/src/target/permute.go?s=242:290#L9)
``` go
func PermuteR(set string, min, max int) []string
```
PermuteR returns all permutations, (with repetition)
of the characters in the set with length in the interval (min, max).

Output from PermuteR will always match the Regexp "/[\1]+/" where "\1" = set



## <a name="Product">func</a> [Product](/src/target/funcs.go?s=1932:1965#L88)
``` go
func Product(i, n int) (prod int)
```
Product returns the product of all numbers on the interval [i, n].



## <a name="ProductFunc">func</a> [ProductFunc](/src/target/funcs.go?s=2152:2222#L94)
``` go
func ProductFunc(k, n int, f func(int, ...int) int, params ...int) int
```
ProductFunc returns the product of all numbers on the interval [i, n],
performing function f on each number.



## <a name="SumFunc">func</a> [SumFunc](/src/target/funcs.go?s=2415:2481#L104)
``` go
func SumFunc(k, n int, f func(int, ...int) int, params ...int) int
```
SumFunc returns the sum of all numbers on the interval [i, n],
performing function f on each number.




## <a name="Set">type</a> [Set](/src/target/set.go?s=89:132#L6)
``` go
type Set struct {
    // contains filtered or unexported fields
}

```
Set represents a set of items that can be combined.







### <a name="NewSet">func</a> [NewSet](/src/target/set.go?s=184:217#L11)
``` go
func NewSet(v ...interface{}) Set
```
NewSet creates a new set from the given items.





### <a name="Set.Add">func</a> (\*Set) [Add](/src/target/set.go?s=585:617#L27)
``` go
func (s *Set) Add(v interface{})
```
Add ads a new object to the set.




### <a name="Set.Combine">func</a> (Set) [Combine](/src/target/set.go?s=725:764#L32)
``` go
func (s Set) Combine() [][2]interface{}
```
Combine returns all combinations of the given set with itself.




### <a name="Set.Size">func</a> (\*Set) [Size](/src/target/set.go?s=1014:1038#L44)
``` go
func (s *Set) Size() int
```
Size returns the number of elements in the set.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
