# read from file and post to st2 (stackstorm) webhook

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