# HTTP Toolbox

## Commands

| Command        | Returns             |
| -------------- | ------------------- |
| /dump          | request dump        |
| /ip            | remote address      |
| /status/[code] | status code         |
| /useragent     | header "User-agent" |

## Examples

Get your requesting IP address
```sh
$ curl htbx.00101010.org/ip
1.2.3.4
```

Get a response with requested status code
```sh
$ curl -v htbx.00101010.org/status/418
* Hostname was NOT found in DNS cache
*   Trying 172.217.23.147...
* Connected to htbx.00101010.org (172.217.23.147) port 80 (#0)
> GET /status/418 HTTP/1.1
> User-Agent: curl/7.35.0
> Host: htbx.00101010.org
> Accept: */*
> 
< HTTP/1.1 418 I'm a Teapot
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< X-Cloud-Trace-Context: ab65efab06faba461f5fb1f4eaed5e8a;o=1
< Date: Sun, 12 Mar 2017 14:37:24 GMT
* Server Google Frontend is not blacklisted
< Server: Google Frontend
< Content-Length: 13
< 
I'm a teapot
* Connection #0 to host htbx.00101010.org left intact
```
