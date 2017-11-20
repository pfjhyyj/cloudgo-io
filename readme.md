# 服务计算作业

学号：15331366



# 支持静态文件服务

```
curl -v http://localhost:8080/assets/css/main.css > 1.txt
StatusCode        : 200
StatusDescription : OK
Content           : body {
                      background-color: #eff7de;
                    }
                    
RawContent        : HTTP/1.1 200 OK
                    Accept-Ranges: bytes
                    Content-Length: 38
                    Content-Type: text/css; charset=utf-8
                    Date: Mon, 20 Nov 2017 17:46:28 GMT
                    Last-Modified: Wed, 15 Nov 2017 03:58:14 GMT
                    
                    body {
                      backgrou...
Forms             : {}
Headers           : {[Accept-Ranges, bytes], [Content-Length, 38], [Content-Type, text/css; charset=utf-8], [Date, Mon,
                     20 Nov 2017 17:46:28 GMT]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 38
```

# 支持简单js访问

```
curl -v http://localhost:8080/api/test > 2.txt
StatusCode        : 200
StatusDescription : OK
Content           : {
                      "message": "Hello"
                    }
                    
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 25
                    Content-Type: application/json; charset=UTF-8
                    Date: Mon, 20 Nov 2017 17:55:37 GMT
                    
                    {
                      "message": "Hello"
                    }
                    
Forms             : {}
Headers           : {[Content-Length, 25], [Content-Type, application/json; charset=UTF-8], [Date, Mon, 20 Nov 2017 17:
                    55:37 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : System.__ComObject
RawContentLength  : 25
```

# 支持表单提交

···
curl -v http://localhost:8080/login -d "username=123" > 3.txt
> POST /login HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Length: 12
> Content-Type: application/x-www-form-urlencoded
>
} [12 bytes data]
* upload completely sent off: 12 out of 12 bytes
< HTTP/1.1 200 OK
< Content-Type: text/html; charset=UTF-8
< Date: Mon, 20 Nov 2017 17:56:39 GMT
< Content-Length: 412
<
{ [412 bytes data]
···

# 对/unknown处理

```
curl -v http://localhost:8080/login2 > 4.txt
> GET /login2 HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.47.0
> Accept: */*
>
< HTTP/1.1 501 Not Implemented
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Mon, 20 Nov 2017 17:59:42 GMT
< Content-Length: 17
<
{ [17 bytes data]
```