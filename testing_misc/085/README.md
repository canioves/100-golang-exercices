# Exercise 85: Multiple Test Cases

## Testing with Multiple Scenarios

Real-world functions need testing with various inputs and edge cases. Writing multiple test cases helps ensure your function works correctly across different scenarios and handles edge conditions properly.

## Test Case Strategy

When testing functions, consider:
- **Happy Path**: Normal, expected inputs
- **Edge Cases**: Boundary values, empty inputs
- **Error Cases**: Invalid inputs, error conditions
- **Special Values**: Zero, negative numbers, very large values

## Table-Driven Tests

A common Go pattern for multiple test cases:

```go
func TestReverse(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {"hello", "olleh"},
        {"", ""},
        {"a", "a"},
        {"12345", "54321"},
    }
    
    for _, test := range tests {
        result := Reverse(test.input)
        if result != test.expected {
            t.Errorf("Reverse(%q) = %q, want %q", 
                test.input, result, test.expected)
        }
    }
}
```

## Testing String Manipulation

The Reverse function demonstrates common string manipulation challenges:
- Converting strings to byte slices
- Index manipulation and bounds checking
- Character-by-character operations
- Handling empty strings and single characters

## Your Task

Look at the main.go file and complete the exercise by:
1. Understanding the provided Reverse function
2. Creating comprehensive test cases for the function
3. Testing various input scenarios including edge cases

This exercise teaches you to write thorough tests that validate function behavior across multiple scenarios.

```go

```

<details>
<summary> Solution: </summary>

```go
// TESTING!
package main

// In this exercise, we are going to create 3 tests for the reverse function, that reverses a string
// if it gets "Hello" it will return "olleH"

import "fmt"

func Reverse(s string) string {
  b := []byte(s)
  for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
      b[i], b[j] = b[j], b[i]
  }
  return string(b)
}

func main() {
  input := "The quick brown fox jumped over the lazy dog"
  rev := Reverse(input)
  doubleRev := Reverse(rev)
  fmt.Printf("original: %q\n", input)
  fmt.Printf("reversed: %q\n", rev)
  fmt.Printf("reversed again: %q\n", doubleRev)
}
```

</details>