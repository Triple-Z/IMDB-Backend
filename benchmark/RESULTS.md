# Benchmark

## Http load

```bash
$ http_load -p 10 -s 30 benchmark/urls
```

```bash
87600 fetches, 10 max parallel, 6.17512e+08 bytes, in 30 seconds
7049.22 mean bytes/connection
2920 fetches/sec, 2.05837e+07 bytes/sec
msecs/connect: 0.137123 mean, 6.395 max, 0.028 min
msecs/first-response: 3.08429 mean, 32.685 max, 0.032 min
20701 bad byte counts
HTTP response codes:
  code 200 -- 79548
  code 400 -- 8052
```

## MySQL status

```
Uptime: 6087  Threads: 11  Questions: 766302  Slow queries: 0  Opens: 304  Flush tables: 1  Open tables: 298  Queries per second avg: 125.891
```
