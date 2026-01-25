# Exercise 87: Fuzz Testing

## What is Fuzz Testing?

Fuzz testing (or fuzzing) is an automated testing technique that provides random, unexpected, or invalid inputs to your functions to discover bugs, crashes, and security vulnerabilities that traditional tests might miss.

## Go's Built-in Fuzzing

Go 1.18+ includes native fuzz testing support:
- Generates random inputs automatically
- Finds edge cases you didn't think of
- Reproduces failures with minimal test cases
- Integrates with the standard `go test` command

## Fuzz Function Structure

Fuzz functions must:
- Start with `Fuzz` prefix
- Accept `*testing.F` parameter
- Call `f.Add()` to provide seed inputs
- Call `f.Fuzz()` with the target function

```go
func FuzzExample(f *testing.F) {
    // Seed the fuzzer with known inputs
    f.Add("hello")
    f.Add("")
    f.Add("123")
    
    // Define the fuzz target
    f.Fuzz(func(t *testing.T, input string) {
        result := YourFunction(input)
        // Add assertions here
        if len(result) < 0 {
            t.Error("Invalid result length")
        }
    })
}
```

## Seed Corpus

- `f.Add()` provides initial "interesting" inputs
- Fuzzer uses these as starting points
- Mutates seeds to generate new test cases
- Good seeds improve fuzzing effectiveness

## Running Fuzz Tests

```bash
go test -fuzz .           # Run all fuzz tests
go test -fuzz FuzzHex     # Run specific fuzz test
go test -fuzz . -fuzztime 30s  # Run for specific duration
```

## Your Task

Look at the main_test.go file and complete the exercise by:
1. Understanding hex encoding/decoding operations
2. Adding seed values to the fuzzer with f.Add()
3. Implementing hex encode/decode operations in the fuzz target
4. Learning how fuzzing finds edge cases automatically

This exercise demonstrates how fuzz testing can automatically discover bugs in data processing functions like encoding/decoding operations.

```go

```

<details>
<summary> Solution: </summary>

```go
package main

// import the "testing" library
import (
    "testing"
    "encoding/hex"
    "bytes"
)

// Again, like the Test or Benchmark signature, go also accepts Fuzzing (https://en.wikipedia.org/wiki/Fuzzing)
// The syntax of the Fuzzer function  name will be FuzzXxxxxx and the signature (f *testing.F)
// Create a fuzzer function named FuzzHex
func FuzzHex(f *testing.F) {
    // now we will seed the fuzzer with the function *f.Add() and the seed as the parameter
    for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
      //here!
      f.Add(seed)
    }

    // now we will call the *f.Fuzz() function, this is considered the target. 
    // documentation: A fuzz target must accept a *T parameter, followed by one or more parameters for random inputs. 
    f.Fuzz(func(t *testing.T, in []byte) {
      // here we will use enc (hex.EncodeToString(in))
      enc := hex.EncodeToString(in)
      // and here out := hex.DecodeString(enc)
      out, err := hex.DecodeString(enc)
      if err != nil {
        t.Fatalf("%v: decode: %v", in, err)
      }
      if !bytes.Equal(in, out) {
        t.Fatalf("%v: not equal after round trip: %v", in, out)
      }
    })
}
```

</details>