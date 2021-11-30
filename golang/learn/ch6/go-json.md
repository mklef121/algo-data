## Working with JSON

The Go standard library includes `encoding/json` module which is for working with JSON data.
Go allows you to **add support** for JSON fields in Go structures using **tags**.


### Using `Marshal()` and `Unmarshal()`

- **Marshaling** is the process of converting a Go structure into a JSON record
- **Unmarshaling** is the process of converting a JSON record given as a byte slice into a Go structure.

Take a look at this

```go
type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}
```

This metadata tells us is that the **Name** field of the UseAll structure is translated to **username** in the JSON record, and vice versa, the **Surname** field is translated to **surname**, and vice versa, and the **Year** structure field is translated to **created** in the JSON record, and vice versa.

```go

    func main() {
        useall := UseAll{Name: "Mike", Surname: "Tsoukalos", Year: 2021}

        // Encoding a structure as a string
        t, err := json.Marshal(useall)

        if err != nil {
            fmt.Println("Error doing Json encoding", err)
        } else {
            fmt.Printf("Value %s\n", t)
        }

        //Decoding a string into a structure

        // Decoding JSON data given as a string
        str := `{"username": "M.", "surname": "Ts", "created":2020}`

        newVal := UseAll{}
        err = json.Unmarshal([]byte(str), &newVal)

        if err != nil {
            fmt.Println("\n There was an error decoding Json", err)
        } else {
            fmt.Println("\nstruct Value ", newVal)
        }
    }
```