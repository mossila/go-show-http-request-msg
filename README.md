# go-show-http-request-msg
Show raw data that post to mock service 

#Get
```go
go get github.com/mossila/go-show-http-request-msg
```

#run 
```bash
go run go-show-http-request-msg 8080
```

#Result 
เมื่อมีการ ทำ http request เข้ามาที่ mock server ก็แสดงผลออกมาที่หน้าจอ terminal 
```html
POST / HTTP/1.1
Host: localhost:1234
Connection: keep-alive
Content-Length: 139
Cache-Control: no-cache
Origin: chrome-extension://fhbjgbiflinjbdggehcddcbncdddomop
Content-Type: application/json
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36
Postman-Token: 31e2a755-eea4-4641-755f-abcac415f0e0
Accept: */*
Accept-Encoding: gzip, deflate
Accept-Language: en-US,en;q=0.8

------WebKitFormBoundaryBXhnVSHJX694UlPO
Content-Disposition: form-data; name="key"

value
------WebKitFormBoundaryBXhnVSHJX694UlPO--
```

