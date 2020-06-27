package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {

    // Compile the expression once, usually at init time.
    // Use raw strings to avoid having to quote the backslashes
    var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
    fmt.Println(validID.MatchString("adam[23]"))    // true
    fmt.Println(validID.MatchString("eve[7]"))      // true
    fmt.Println(validID.MatchString("Job[48]"))     // false
    fmt.Println(validID.MatchString("snakey"))      // false

    // func Match(pattern string, b []byte) (matched bool, err error)
    matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
    fmt.Println(matched, err)                                   // true <nil>
    matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
    fmt.Println(matched, err)                                   // false <nil>
    matched, err = regexp.Match(`a(b`, []byte(`seafood`))
    fmt.Println(matched, err)                                   // false error parsing regexp: missing closing ): `a(b`
    matched, err = regexp.Match(regexp.QuoteMeta(`a(b`), []byte(`a(bc`))
    fmt.Println(matched, err)                                   // true <nil>

    // func MatchString(pattern string, s string) (matched bool, err error)
    matched, err = regexp.MatchString(`foo.*`, "seafood")
    fmt.Println(matched, err)                                   // true <nil>
    matched, err = regexp.MatchString(`bar.*`, "seafood")
    fmt.Println(matched, err)                                   // false <nil>
    matched, err = regexp.MatchString(`a(b`, "seafood")
    fmt.Println(matched, err)                                   // false error parsing regexp: missing closing ): `a(b`

    // func QuoteMeta(s string) string
    fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))   // Escaping symbols like: \.\+\?\(\)\|\[\]\{\}\^\$

    // func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
    content := []byte(`
        # comment line
        option1: value1
        option2: value2

        # another comment line
        option3: value3
    `)
    // Regex pattern captures "key: value" pair from the content
    pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
    // Template to convert "key: value" to "key=value" by referencing the values captured by the regex pattern.
    template := []byte("$key=$value\n")
    result := []byte{}
    // For each match of the regex in the content
    for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
        // Apply the captured submatches to the template and append the output to the result.
        result = pattern.Expand(result, template, content, submatches)
    }
    fmt.Println(string(result))

    // func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
    contentStr := `
        # comment line
        option1: value1
        option2: value2

        # another comment
        option3: value3
    `
    // Regex pattern captures "key: value" pair from the content.
    // pattern = regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
    // Template to convert "key: value" to "key=value" by referencing the values captured by the regex pattern.
    templateStr := "$key=$value\n"
    result = []byte{}
    // For each match of the regex in the content.
    for _, submatches := range pattern.FindAllStringSubmatchIndex(contentStr, -1) {
        // Apply the captured submatches to the template and append the output to the result.
        result = pattern.ExpandString(result, templateStr, contentStr, submatches)
    }
    fmt.Println(string(result))

    // func (re *Regexp) Find(b []byte) []byte
    re := regexp.MustCompile(`foo.?`)
    fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))  // "food"

    // func (re *Regexp) FindAll(b []byte, n int) [][]byte
    fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))  // ["food" "fool"]
    fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), 0))   // []
    fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), 1))   // ["food"]
    fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), 2))   // ["food" "fool"]
    fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), 3))   // ["food" "fool"]

    // func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
    content = []byte("London")
    re = regexp.MustCompile(`o.`)
    fmt.Println(re.FindAllIndex(content, 1))        // [[1 3]]
    fmt.Println(re.FindAllIndex(content, -1))       // [[1 3] [4 6]]

    // func (re *Regexp) FindAllString(s string, n int) []string
    re = regexp.MustCompile(`a.`)
    fmt.Println(re.FindAllString("paranormal", -1)) // [ar an al]
    fmt.Println(re.FindAllString("paranormal", 2))  // [ar an]
    fmt.Println(re.FindAllString("graal", -1))      // [aa]
    fmt.Println(re.FindAllString("none", -1))       // []

    // func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
    re = regexp.MustCompile(`a(x*)b`)
    fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))        // [["ab" ""]]
    fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxxb-", -1))     // [["axxb" "xx"]]
    fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))    // [["ab" ""] ["axb" "x"]]
    fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))   // [["axxb" "xx"] ["ab" ""]]

    // func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
    re = regexp.MustCompile(`a(x*)b`)
    fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))          // [[1 3 2 2]]
    fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))        // [[1 5 2 4]]
    fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))      // [[1 3 2 2] [4 7 5 6]]
    fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))     // [[1 5 2 4] [6 8 7 7]]
    fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))         // []

    // func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
    re = regexp.MustCompile(`foo(.?)`)
    fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1))  // [["food" "d"] ["fool" "l"]]

    // func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
    content = []byte(`
        # comment line
        option1: value1
        option2: value2
    `)
    // Regex pattern captures "key: value" pair from the content.
    pattern = regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
    allIndexes := pattern.FindAllSubmatchIndex(content, -1)     // [[...] [...]]
    for _, loc := range allIndexes {
        fmt.Println(loc)                                        // [...]
        fmt.Println(string(content[loc[0]:loc[1]]))             // option1: value1  // option2: value2
        fmt.Println(string(content[loc[2]:loc[3]]))             // option1          // option2
        fmt.Println(string(content[loc[4]:loc[5]]))             // value1           // value2
    }

    // func (re *Regexp) FindIndex(b []byte) (loc []int)
    content = []byte(`
        # comment line
        option1: value1
        option2: value2
    `)
    // Regex pattern captures "key: value" pair from the content.
    pattern = regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
    loc := pattern.FindIndex(content)
    fmt.Println(loc)
    fmt.Println(string(content[loc[0]:loc[1]]))             // option1: value1

    // func (re *Regexp) FindString(s string) string
    re = regexp.MustCompile(`foo.?`)
    fmt.Printf("%q\n", re.FindString("seafood fool"))       // "food"
    fmt.Printf("%q\n", re.FindString("meat"))               // ""

    // func (re *Regexp) FindStringIndex(s string) (loc []int)
    re = regexp.MustCompile(`ab?`)
    fmt.Println(re.FindStringIndex("tablett"))          // [1 3]
    fmt.Println(re.FindStringIndex("foo") == nil)       // true

    // func (re *Regexp) FindStringSubmatch(s string) []string
    re = regexp.MustCompile(`a(x*)b(y|z)c`)
    fmt.Println("%q\n", re.FindStringSubmatch("-axxxbyc-")) // ["axxxbyc" "xxx", "y"]
    fmt.Println("%q\n", re.FindStringSubmatch("-abzc-"))    // ["abzc" "" "z"]

    // func (re *Regexp) FindSubmatch(b []byte) [][]byte
    re = regexp.MustCompile(`foo(.?)`)
    fmt.Printf("%q\n", re.FindSubmatch([]byte(`seafood fool`))) // ["food" "d"]

    // func (re *Regexp) FindSubmatchIndex(b []byte) []int
    re = regexp.MustCompile(`a(x*)b`)
    fmt.Println(re.FindSubmatchIndex([]byte("-ab-")))           // [1 3 2 2]
    fmt.Println(re.FindSubmatchIndex([]byte("-axxb-")))         // [1 5 2 4]
    fmt.Println(re.FindSubmatchIndex([]byte("-ab-axb-")))       // [1 3 2 2]
    fmt.Println(re.FindSubmatchIndex([]byte("-axxb-ab-")))      // [1 5 2 4]
    fmt.Println(re.FindSubmatchIndex([]byte("-foo-")))          // []

    // func (re *Regexp) Longest()
    re = regexp.MustCompile(`a(|b)`)
    fmt.Println(re.FindString("ab"))    // a
    re.Longest()
    fmt.Println(re.FindString("ab"))    // ab

    // func (re *Regexp) Match(b []byte) bool
    re = regexp.MustCompile(`foo.?`)
    fmt.Println(re.Match([]byte(`seafood fool`)))       // true
    fmt.Println(re.Match([]byte(`something else`)))     // false

    // func (re *Regexp) MatchString(s string) bool
    re = regexp.MustCompile(`(gopher){2}`)
    fmt.Println(re.MatchString("gopher"))               // false
    fmt.Println(re.MatchString("gophergopher"))         // true
    fmt.Println(re.MatchString("gophergophergopher"))   // true

    // func (re *Regexp) NumSubexp() int
    re = regexp.MustCompile(`a.`)
    fmt.Printf("%d\n", re.NumSubexp())                  // 0
    re = regexp.MustCompile(`(.*)((a)b)(.*)a`)
    fmt.Println(re.NumSubexp())                         // 4

    // func (re *Regexp) ReplaceAll(src, repl []byte) []byte
    re = regexp.MustCompile(`a(x*)b`)
    fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))    // -T-T-
    fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))   // --xx-
    fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))  // ---
    fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))// -W-xxW-

    // func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
    re = regexp.MustCompile(`a(x*)b`)
    fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))           // -T-T-
    fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))          // -$1-$1-
    fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))       // -${1}-${1}-

    // func (re *Regexp) ReplaceAllString(src, repl string) string
    re = regexp.MustCompile(`a(x*)b`)
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))      // -T-T-
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))     // --xx-
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))    // ---
    fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))  // -W-xxW-

    // func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
    re = regexp.MustCompile(`[^aeiou]`)
    fmt.Println(re.ReplaceAllStringFunc("seafood fool", strings.ToUpper))   // SeaFooD FooL

    // func (re *Regexp) Split(s string, n int) []string
    s := regexp.MustCompile(`a*`).Split("abaabaccadaaae", 5)
    fmt.Printf("%q\n", s)                   // ["" "b" "b" "c" "cadaaae"]
    a := regexp.MustCompile(`a`)
    fmt.Println(a.Split("banana", -1))      // [b n n]
    fmt.Println(a.Split("banana", 0))       // []
    fmt.Println(a.Split("banana", 1))       // [banana]
    fmt.Println(a.Split("banana", 2))       // [b nana]
    zp := regexp.MustCompile(`z+`)
    fmt.Println(zp.Split("pizza", -1))      // [pi a]
    fmt.Println(zp.Split("pizza", 0))       // []
    fmt.Println(zp.Split("pizza", 1))       // [pizza]
    fmt.Println(zp.Split("pizza", 2))       // [pi za]

    // func (re *Regexp) SubexpNames() []string
    re = regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
    fmt.Println(re.MatchString("Alan Turing"))                  // true
    fmt.Printf("%q\n", re.SubexpNames())                        // ["" "first" "last"]
    reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
    fmt.Println(reversed)                                       // ${last} ${first}
    fmt.Println(re.ReplaceAllString("Alan Turing", reversed))   // Turing Alan

}
