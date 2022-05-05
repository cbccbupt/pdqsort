# 字节跳动-青训营-算法与数据结构代码附件

**不要在生产环境使用本仓库，仅供参考**

## Benchmark

运行 benchmark(性能测试)

```
go test -bench=. -cpu=1 -timeout=1h
```

运行 benchmark(性能测试) 快速版本，可能出现数据不准确

```
go test -bench=. -cpu=1 -benchtime=1000x -timeout=1h
```

