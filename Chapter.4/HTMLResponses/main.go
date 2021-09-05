package main

import (
	"html/template"
	"net/http"
	"os"
)

var x = `
<html>
	<body>
		Hello {{.}}
	</body>
</html>
`

func main() {

	t, err := template.New("hello").Parse(x)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/h", func(rw http.ResponseWriter, r *http.Request) {
		t.Execute(os.Stdout, "<script>alert('world')</script>")
	})

	http.ListenAndServe(":8000", nil)
}
