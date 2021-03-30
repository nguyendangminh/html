# HTML

This package is mainly copied from [gofiber framework](https://github.com/gofiber/template/blob/master/html/README.md), added some minor modifications for my own use cases.

HTML is the official Go template engine [html/template](https://golang.org/pkg/html/template/), to see the original syntax documentation please [click here](https://pkg.go.dev/html/template) or read this [cheatsheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet#actions)

### Basic Example

_**./views/index.html**_
```html
{{template "partials/header" .}}

<h1>{{.Title}}</h1>

{{template "partials/footer" .}}
```
_**./views/partials/header.html**_
```html
<h2>Header</h2>
```
_**./views/partials/footer.html**_
```html
<h2>Footer</h2>
```
_**./views/layouts/main.html**_
```html
<!DOCTYPE html>
<html>

<head>
  <title>Main</title>
</head>

<body>
  {{embed}}
</body>

</html>
```

```go
package main

import (
	"log"
	"net/http"

	"github.com/nguyendangminh/html"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")
	if err := html.SetDefaultEngine(engine); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/layout", func(w http.ResponseWriter, r *http.Request) {
		html.Render(w, "index", map[string]string{"Title": "Hello, world!"}, "layouts/main")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html.Render(w, "index", map[string]string{"Title": "Hello, world!"})
	})

	log.Println("listen and serve at http://localhost:1203")
	log.Fatal(http.ListenAndServe(":1203", nil))
}

```

### Example with embed.FS

Note:
- Need Go 1.16 or later
- Template names must be prefixed with template directory `views`
```go
html.Render(w, "views/index", map[string]string{"Title": "Hello, world!"}, "views/layouts/main")
```

```
{{template "views/partials/footer" .}}
```


```go
package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/nguyendangminh/html"
)

//go:embed views/*
var viewsfs embed.FS

func main() {
	// Create a new engine
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")
	if err := html.SetDefaultEngine(engine); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/layout", func(w http.ResponseWriter, r *http.Request) {
		html.Render(w, "views/index", map[string]string{"Title": "Hello, world!"}, "layouts/main")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html.Render(w, "views/index", map[string]string{"Title": "Hello, world!"})
	})

	log.Println("listen and serve at http://localhost:1203")
	log.Fatal(http.ListenAndServe(":1203", nil))
}
```