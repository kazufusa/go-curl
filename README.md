# go-curl

```sh
$ go run . -h
Usage of /tmp/go-build3191817544/b001/exe/go-curl:
  -H value
        request header
  -d string
        request body
  -v    show request
  -X string
        method (default "GET")

$ go run . -x GET -v -d "{}" -H aaa=bbb -H ccc=ddd:eee example.com
GET / HTTP/1.1
Host: example.com
Aaa: bbb
Ccc: ddd:eee

{}
<?xml version="1.0" encoding="iso-8859-1"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
         "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
        <head>
                <title>400 - Bad Request</title>
        </head>
        <body>
                <h1>400 - Bad Request</h1>
        </body>
</html>
```
