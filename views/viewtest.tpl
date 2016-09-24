{{range .Pages}}
{{.Num}} of {{$.Total}}
{{end}}

<br>

{{with $value := "My name is %s"}}
{{printf . "slene"}}
{{end}}

<br>

{{define "loop"}}
<li>{{.Name}}</li>
{{end}}

<ul>
    {{range .Items}}
    {{template "loop" .}}
    {{end}}
</ul>


{{printf "nums is %s %d" (printf "%d %d" 1 2) 3}}

<br>
{{index .Maaaap "name"}}

<br>
{{urlquery "http://beego.me"}}

<br>

{{range $key, $val := .s}}
{{$key}}
{{$val}}
{{end}}

<br>
the username is {{.m.name}}
the username is {{.m.nickname}}

<br>
{{range $key, $val := .m}}
{{$key}}
{{$val}}
{{end}}

<br>
<form action="" method="post">
    {{.Form | renderform}}
</form>

<br>
{{dateformat .Time "2006-01-02"}}