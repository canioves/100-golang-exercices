# Exercise 69: JSON Marshaling (Encoding)

## What is JSON Marshaling?

**Marshaling** is the process of converting Go data structures into JSON format. It's the opposite of unmarshaling - instead of JSON → Go, it's Go → JSON.

## json.Marshal() Function

```go
data, err := json.Marshal(goStruct)
if err != nil {
    // Handle error
}
jsonString := string(data) // Convert []byte to string
```

## Marshaling Examples

**Simple struct:**
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

person := Person{Name: "Alice", Age: 30}
jsonData, _ := json.Marshal(person)
// Result: {"name":"Alice","age":30}
```

**Array of structs:**
```go
people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
}
jsonData, _ := json.Marshal(people)
// Result: [{"name":"Alice","age":30},{"name":"Bob","age":25}]
```

## Key Points

- `json.Marshal()` returns `[]byte`, not string
- Use `string(data)` to convert to readable format
- Struct fields must be exported (capitalized) to be marshaled
- Use struct tags to control JSON field names

## Your Task

1. Create a `user` struct with Username and Password fields
2. Create an array of 3 users
3. Marshal the array to JSON using `json.Marshal()`
4. Print the resulting JSON string

This demonstrates how to convert Go data to JSON format for APIs, file storage, or data transmission.

```go
// Exercise: JSON!! - Marshal (encode)

// We will create a json object.
// First, create a struct called 'user' which contains 2 attributes of type string called 'Username' and 'Password'
// Then, create an array of 3 users 'John', 'Anna' and 'Mike' with any random password
// Use the json.Marshal() function to encode the go array into a JSON object.

package main

import (
    "fmt"
    "encoding/json"
)

type user struct {
  
}

func main() {
  

  // print encoded json data
  fmt.Println(string(out))
}
```

<details>
<summary> Solution: </summary>

```go
// Exercise: JSON!! - Marshal (encode)

// We will create a json object.
// First, create a struct called 'user' which contains 2 attributes of type string called 'Username' and 'Password'
// Then, create an array of 3 users 'John', 'Anna' and 'Mike' with any random password
// Use the json.Marshal() function to encode the go array into a JSON object.

package main

import (
    "fmt"
    "encoding/json"
)

type user struct {
  Username string
  Password string
}

func main() {
  users := []user{
    {"John","123456"},
    {"Anna","42"},
    {"Mike","33"},
  }

  // use marshal func to convert to json
  out, err := json.Marshal(users)
  if err != nil {
      fmt.Println(err)
      return
  }

  // print encoded json data
  fmt.Println(string(out))
}
```

</details>