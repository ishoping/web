# web

# web开发框架
    这个小程序使用martini框架，按道理实现这个小程序用go的net/http就可以，但为了试用一下，  
    所以用了一下martini，

# 运行命令

    ljx@ljx-X550JX:~/project/go/web/src$ go run main.go -p9090  
    [martini] listening on :9090 (development)  
    [martini] Started GET /testuser for 127.0.0.1:36118  
    [martini] Completed 200 OK in 93.484µs

# 使用 curl 测试

    $ curl -v http://localhost:9090/testuser  
    *   Trying 127.0.0.1...
    * Connected to localhost (127.0.0.1) port 9090 (#0)
    > GET /testuser HTTP/1.1
    > Host: localhost:9090
    > User-Agent: curl/7.47.0
    > Accept: */*
    >
    < HTTP/1.1 200 OK
    < Date: Sat, 04 Nov 2017 09:18:39 GMT
    < Content-Length: 15
    < Content-Type: text/plain; charset=utf-8
    <
    Hello testuser
    * Connection #0 to host localhost left intact
    
# 对 web 执行压力测试

    $ ab -n 1000 -c 100 http://localhost:9090/you
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 100 requests
    Completed 200 requests
    Completed 300 requests
    Completed 400 requests
    Completed 500 requests
    Completed 600 requests
    Completed 700 requests
    Completed 800 requests
    Completed 900 requests
    Completed 1000 requests
    Finished 1000 requests


    Server Software:        
    Server Hostname:        localhost
    Server Port:            9090

    Document Path:          /you
    Document Length:        10 bytes

    Concurrency Level:      100
    Time taken for tests:   0.099 seconds
    Complete requests:      1000
    Failed requests:        0
    Total transferred:      127000 bytes
    HTML transferred:       10000 bytes
    Requests per second:    10146.62 [#/sec] (mean)
    Time per request:       9.856 [ms] (mean)
    Time per request:       0.099 [ms] (mean, across all concurrent requests)
    Transfer rate:          1258.42 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    0   0.4      0       2
    Processing:     0    9   6.7      9      37
    Waiting:        0    9   6.8      8      37
    Total:          0   10   7.0      9      37

    Percentage of the requests served within a certain time (ms)
    50%      9
    66%      9
    75%     10
    80%     11
    90%     20
    95%     26
    98%     33
    99%     33
    100%     37 (longest request)

