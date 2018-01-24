# read from file and post to st2 (stackstorm) webhook

```
$ gost2netpro -h
Copyright 2018 @IrekRomaniuk. All rights reversed.
Usage of gost2netpro:
  -a string
        file to read targets from (default "./targets/pinglist.txt")
  -api string
        st2 api key
  -c string
        cmd to run on targets (default "uptime")
  -p string
        target password (default "pass")
  -u string
        target username (default "user")
  -url string
        st2 webhook (default "https://ubuntu-st2/api/v1/webhooks/netpro")
  -v    Prints current version
```

## Under Development

```
$ go run main.go -api='' -u=user -p='passowrd'

me@ubuntu-st2:~$ st2 execution get 5a6238a750140059dc113abf
id: 5a6238a750140059dc113abf
status: running (26s elapsed)
parameters: 
  cmd: uptime
  hosts: 1.1.1.1,2.2.2.2
  password: '********'
  username: user
result: None
```

```
gost2netpro -api='' -u=user -p='password' -a="./list.txt"
```