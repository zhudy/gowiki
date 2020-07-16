//https://golang.google.cn/doc/effective_go.html#web_server
package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))  //这个也魔幻，有这个必要吗？

func main() {
    flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	println(*addr)  //不知道闹啥，就是:1718啊，搞这么玄乎干嘛？
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}


func QR(w http.ResponseWriter, req *http.Request) {
	println(req.FormValue("s"))
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
