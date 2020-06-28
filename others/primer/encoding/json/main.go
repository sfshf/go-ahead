package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func main() {

    /*
        type Decoder struct {
            // contains filtered or unexported fields
        }
    */
    var jsonStream = `
        {"Name": "Ed", "Text": "Knock knock."}
        {"Name": "Sam", "Text": "Who's there?"}
        {"Name": "Ed", "Text": "Go fmt."}
        {"Name": "Sam", "Text": "Go fmt who?"}
        {"Name": "Ed", "Text": "Go fmt yourself!"}
        {"Name": "Aaron Swartz", "Text": "One creator of markdown!"}
    `
    type Message struct {
        Name, Text string
    }

    dec := json.NewDecoder(strings.NewReader(jsonStream))

    for {
        var m Message
        if err := dec.Decode(&m); err == io.EOF { break } else if err != nil { log.Fatal(err) }
        fmt.Printf("%s: %s -- inputOffset: %d\n", m.Name, m.Text, dec.InputOffset())
    }

    // func (dec *Decoder) Decode(v interface{}) error
    jsonStream = `
        [
            {"Name": "Ed", "Text": "Knock knock."},
            {"Name": "Sam", "Text": "Who's there?"},
            {"Name": "Ed", "Text": "Go fmt."},
            {"Name": "Sam", "Text": "Go fmt who?"},
            {"Name": "Ed", "Text": "Go fmt yourself!"},
            {"Name": "Aaron Swartz", "Text": "One creator of markdown!"}
        ]
    `
    dec = json.NewDecoder(strings.NewReader(jsonStream))
    // read open bracket
    t, err := dec.Token()
    if err != nil { log.Fatal(err) }
    fmt.Printf("%T: %v\n", t, t)
    // while the array contains values
    for dec.More() {
        var m Message
        // decode an array value (Message)
        err := dec.Decode(&m)
        if err != nil { log.Fatal(err) }
        fmt.Printf("%v: %v\n", m.Name, m.Text)
    }
    // read closing bracket
    t, err = dec.Token()
    if err != nil { log.Fatal(err) }
    fmt.Printf("%T: %v\n", t, t)

    // func (dec *Decoder) Token() (Token, error)
    jsonStream = `
        {"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
    `
    dec = json.NewDecoder(strings.NewReader(jsonStream))
    for {
        t, err := dec.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%T: %v", t, t)
        if dec.More() {
            fmt.Printf(" (more)")
        }
        fmt.Printf("\n")
    }

    // func HTMLEscape(dst *bytes.Buffer, src []byte)
    var out bytes.Buffer
    json.HTMLEscape(&out, []byte(`{"Name":"<b>HTML content</b>"}`))
    out.WriteTo(os.Stdout)
    println()

    // func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
    type Road struct {
        Name string
        Number int
    }
    roads := []Road{
        {"Diamond Fork", 29},
        {"Sheep Creek", 51},
    }
    b, err := json.Marshal(roads)
    if err != nil { log.Fatal(err) }
    // var out bytes.Buffer
    json.Indent(&out, b, "=", "\t")
    out.WriteTo(os.Stdout)
    println()

    // func Marshal(v interface{}) ([]byte, error)
    type ColorGroup struct {
        ID int
        Name string
        Colors []string
    }
    group := ColorGroup{
        ID: 1,
        Name: "Reds",
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    }
    b, err = json.Marshal(group)
    if err != nil { log.Fatal(err) }
    os.Stdout.Write(b)
    os.Stdout.Write([]byte("\n"))

    // func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
    data := map[string]int{
        "a": 1,
        "b": 2,
    }
    jsonStr, err := json.MarshalIndent(data, "<prefix>", "<indent>")
    if err != nil { log.Fatal(err) }
    os.Stdout.Write(jsonStr)
    os.Stdout.Write([]byte("\n"))

    // type RawMessage []byte -- Marshal
    h := json.RawMessage(`{"precomputed": true}`)
    c := struct {
        Header *json.RawMessage `json:"header"`
        Body string `json:"body"`
    }{
        Header: &h,
        Body: "Hello Gophers!",
    }
    b, err = json.MarshalIndent(&c, "", "\t")
    if err != nil { log.Fatal(err) }
    os.Stdout.Write(b)
    os.Stdout.Write([]byte("\n"))
    // type RawMessage []byte -- Unmarshal
    type Color struct {
        Space string
        Point json.RawMessage   // delay parsing until we know the color space
    }
    type RGB struct {
        R uint8
        G uint8
        B uint8
    }
    type YCbCr struct {
        Y uint8
        Cb int8
        Cr int8
    }
    var j = []byte(`[
        {"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
        {"Space": "RGB", "Point": {"R": 98, "G": 218, "B": 255}}
    ]`)
    var colors []Color
    err = json.Unmarshal(j, &colors)
    if err != nil { log.Fatal(err) }
    for _, c := range colors {
        var dst interface{}
        switch c.Space {
        case "RGB":
            dst = new(RGB)
        case "YCbCr":
            dst = new(YCbCr)
        }
        err := json.Unmarshal(c.Point, dst)
        if err != nil { log.Fatal(err) }
        fmt.Println(c.Space, dst)
    }

    // func Unmarshal(data []byte, v interface{}) error
    var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll", "Order": "Dasyuromorphia"}
    ]`)
    type Animal struct {
        Name string
        Order string
    }
    var animals []Animal
    err = json.Unmarshal(jsonBlob, &animals)
    if err != nil { log.Fatal(err) }
    fmt.Printf("%+v\n", animals)

    // func Valid(data []byte) bool
    goodJSON := `{"example": 1}`
    badJSON := `{"example":2:]}}`
    fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))

}
