pprofのおためし  
https://christina04.hatenablog.com/entry/golang-pprof-basic を参考にした

```bazaar
❯ go tool pprof main cpu.pprof 
File: main
Type: cpu
Time: Aug 27, 2021 at 12:29am (JST)
Duration: 5.21s, Total samples = 4.41s (84.61%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 4380ms, 99.32% of 4410ms total
Dropped 6 nodes (cum <= 22.05ms)
      flat  flat%   sum%        cum   cum%
    4090ms 92.74% 92.74%     4390ms 99.55%  main.fib
     290ms  6.58% 99.32%      290ms  6.58%  runtime.newstack
         0     0% 99.32%     4390ms 99.55%  main.main
         0     0% 99.32%     4390ms 99.55%  runtime.main
```