**核心API:**

`Handler -- `

```go

/*

*/
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```

`ServeMux -- `

```go

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}

```

`muxEntry -- `

```go

type muxEntry struct {
	h       Handler
	pattern string
}

```

`HandlerFunc -- `

```go

type HandlerFunc func(ResponseWriter, *Request)

```

**其他API:**
