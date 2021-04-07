package main

import (
	"html"
	"io/ioutil"
	"net/http"
	"strings"
)

var head = `
<!DOCTYPE html>
<html>
<head>
	<title>简单markdown展示器</title>
	<script type="text/javascript" src="/public/dist/showdown.min.js"></script>
</head>

<body>

<div style="display:none"id="ctt">
`
var end = `
</div>

<div  id="show">
</div>
</body>

<script type="text/javascript">
	 var text = document.getElementById("ctt").innerHTML;
      var converter = new showdown.Converter();
      var html = converter.makeHtml(text);
      document.getElementById("show").innerHTML = html;


</script>


</html>
	
`

func handlerMd(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".md") {
			filePath := "./" + r.URL.Path
			file, err := ioutil.ReadFile(filePath)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte(head))
				s := string(file)
				e := html.EscapeString(s)
				w.Write([]byte(e))
				w.Write([]byte(end))
			}
		} else {
			h.ServeHTTP(w, r)
		}

	})
}
