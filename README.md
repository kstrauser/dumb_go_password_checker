**To run the server:**

`$ go run server.go`

**Example success:**

```HTTP
$ curl -v -d'{"password":"foofoo1!"}' http://localhost:8080/
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> POST / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
> Content-Length: 23
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 23 out of 23 bytes
< HTTP/1.1 200 OK
< Date: Fri, 03 Jun 2016 05:10:04 GMT
< Content-Length: 17
< Content-Type: text/plain; charset=utf-8
<
{"status": "ok"}
```

**Example failure:**

```HTTP
$ curl -v -d'{"password":"ಠ_ಠ"}' http://localhost:8080/
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> POST / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
> Content-Length: 22
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 22 out of 22 bytes
< HTTP/1.1 400 Bad Request
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Fri, 03 Jun 2016 05:11:19 GMT
< Content-Length: 91
<
{"Errors":["The password contains invalid characters.","The password isn't long enough."]}
```
