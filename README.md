# speed
Simple speed test server and client written in Go

[![Donate](https://dl.ugjka.net/Donate-PayPal-green.svg)](https://www.paypal.me/ugjka)

## output

### download

```
[ugjka@archee ~]$ speedc 
Download: 44.637149 Mbit/s | Avg: 38.120646 Mbit/s | 1024KB piece ^C
speedc terminated 
```

### upload

```
[ugjka@archee ~]$ speedc -u
Upload: 9.224652 Mbit/s | Avg: 10.425750 Mbit/s | 512KB piece ^C
speedc terminated 
```

### server

```
Dec 27 14:07:06 ugjka speedd[20861]: speedd 2018/12/27 14:07:06 client connected: 90.139.36.111:47186
Dec 27 14:07:06 ugjka speedd[20861]: speedd 2018/12/27 14:07:06 client sent the down command: 90.139.36.111:47186
Dec 27 14:07:40 ugjka speedd[20861]: speedd 2018/12/27 14:07:40 client disconnected: 90.139.36.111:47186
Dec 27 14:08:09 ugjka speedd[20861]: speedd 2018/12/27 14:08:09 client connected: 90.139.36.111:47194
Dec 27 14:08:09 ugjka speedd[20861]: speedd 2018/12/27 14:08:09 client sent the up command: 90.139.36.111:47194
Dec 27 14:08:19 ugjka speedd[20861]: speedd 2018/12/27 14:08:19 client disconnected: 90.139.36.111:47194
Dec 27 14:08:22 ugjka speedd[20861]: speedd 2018/12/27 14:08:22 client connected: 90.139.36.111:47198
Dec 27 14:08:22 ugjka speedd[20861]: speedd 2018/12/27 14:08:22 client sent the down command: 90.139.36.111:47198
Dec 27 14:08:30 ugjka speedd[20861]: speedd 2018/12/27 14:08:30 client disconnected: 90.139.36.111:47198
```

### usage

```
[ugjka@archee speedc]$ speedc -help
Usage of speedc:
  -d    download test (default true)
  -p uint
        piece size in KB
  -s string
        server adress (default ":44444")
  -u    upload test
```

you can set `$SPEEDCSRV` variable in your profile containing your server adress

```
[ugjka@archee speedd]$ speedd -help
Usage of speedd:
  -p uint
        listener port (default 44444)
```
