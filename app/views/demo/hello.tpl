<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>layout</title>
  </head>
  <body>
    <h3>userslist:</h3>
    <p>userlist:</p>{{ .}}
    <p>users:</p>
    {{range $i,$e := .}}
        {{$i}}:{{$e}}
    {{end}}
    {{range .}}
        {{ .}}
    {{end}}
  </body>
</html>
