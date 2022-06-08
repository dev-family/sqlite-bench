# How to run

```bash
$ docker-compose up
```

# Result

```
test_1  | goos: linux
test_1  | goarch: amd64
test_1  | pkg: benching
test_1  | cpu: Intel(R) Core(TM) i5-2500K CPU @ 3.30GHz
test_1  | BenchmarkInsertsSqlite10000-4               50           1237933 ns/op
test_1  | BenchmarkInsertsPostgres10000-4             50          12592869 ns/op
test_1  | BenchmarkUpdateSqlite-4                     50           1218852 ns/op
test_1  | BenchmarkUpdatePostgres-4                   50           1137043 ns/op
test_1  | BenchmarkSelectSqlite-4                     50             20605 ns/op
test_1  | BenchmarkSelectPostgres-4                   50            597675 ns/op
test_1  | PASS
test_1  | ok    benching        4.410s

```