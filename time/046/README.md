# Exercise 46: Time Basics and Formatting

## Working with Time in Go

Go's `time` package provides comprehensive functionality for working with dates and times. Understanding time formatting is essential for displaying times in human-readable formats.

## Getting Current Time

```go
now := time.Now()        // Current local time
utc := time.Now().UTC()  // Current time in UTC
```

## Go's Unique Time Formatting

Go uses a unique approach to time formatting. Instead of using format codes like `%Y %m %d`, Go uses a **reference time** to show the desired format.

## The Reference Time

Go's reference time is: **`Mon Jan 2 15:04:05 MST 2006`**

This represents:
- `Mon` - Monday (weekday)
- `Jan` - January (month name)  
- `2` - 2nd (day)
- `15:04:05` - 3:04:05 PM (time)
- `MST` - Mountain Standard Time (timezone)
- `2006` - Year

## Memory Aid: 1 2 3 4 5 6 -7

To remember the reference time, use the sequence:
- `1` = Month (January)
- `2` = Day
- `3` = Hour (15 = 3 PM)
- `4` = Minute  
- `5` = Second
- `6` = Year (2006)
- `-7` = Timezone offset

## Common Format Patterns

```go
time.Format("2006-01-02")           // 2023-12-07
time.Format("2006 Jan 02")          // 2023 Dec 07  
time.Format("15:04:05")             // 15:30:45
time.Format("3:04 PM")              // 3:30 PM
time.Format("January 2, 2006")      // December 7, 2023
```

## Your Task

Complete the exercise by:
1. Importing the `time` package
2. Creating a `current` variable with the current UTC time
3. Filling in the missing Format() calls to display the time in the requested formats

## Expected Output

You should see the current date and time displayed in the specified formats.

## Key Concept

Remember: Show the reference time (`Jan 2 15:04:05 2006 MST`) in the format you want your actual time to appear.

```go
// Exercise: Time

// import the "time" package
// create a variable named "current" that has the current time in UTC
// Print the Date and time with the format [YYYY Mo dd] and [HH:MM:SS]

package main

import (
    "fmt"
  
)

func main() {
  
  fmt.Println("Date: " + current.Format("2006 Jan 02")) //(*)
  fmt.Println("Time: " + current.Format()) // (*)
  fmt.Println("Date: " + current.Format()) //(*)
}

// (*)
    // The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
```

<details>
<summary> Solution: </summary>

```go
// Exercise: Time

// import the "time" package
// create a variable named "current" that has the current time in UTC
// Print the Date and time with the format [YYYY Mo dd], then with [YYYY/Mo/dd] and finally [HH:MM:SS]

package main

import (
    "fmt"
    "time"
)

func main() {
  current := time.Now().UTC()
  fmt.Println("Date: " + current.Format("2006 Jan 02")) //(*)
  fmt.Println("Date: " + current.Format("2006/Jan/02")) //(*)
  fmt.Println("Time: " + current.Format("03:04:05")) // (*)
}

// (*)
    // The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
```

</details>