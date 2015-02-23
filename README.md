# strsearch
Use [Boyer Moore Horspool](http://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm) algorithm 
to search for a pattern in a text

Golang already has a `Contains` method on the `strings` package. However, that uses the Rabin-Karp algorithm.  
This library uses the Boyer Moore Horspool algorithm.

To use this library, import as:

```golang
import "github.com/callmekatootie/strsearch"
```

## Examples

```golang
strsearch.Find(<text>, <pattern>)
strsearch.Find("Hello World!", "orl") // 7
```
Returns the position of the pattern in the text - position is 0 based  


```golang
strsearch.Find("What sorcery is this?", "nothing") // -1
```
Returns -1 if the pattern can't be found in the text

## License
MIT
