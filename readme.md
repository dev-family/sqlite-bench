<a href="https://dev.family/?utm_source=github&utm_medium=sqlite-bench&utm_campaign=readme"><img width="auto" center src="https://github.com/dev-family/sqlite-bench/blob/main/sqlite-bench.png?raw=true" /></a>

# PostgreSQL & SQLite Speed Test

We tried to compare the speed of SQLite and PostgreSQL during the simplest operations of inserting, updating and extracting data from a spreadsheet.  
You should know that SQLite is a single file, while PostgreSQL is a whole server with many additional functions.  
As we can see, the speed of SQLite due to its simplicity significantly exceeds PostgreSQL. For example:
- The insertion speed is 10 times higher.
- The sampling rate turned out to be 30 times higher.
- The update rate is almost identical to PostgreSQL.

That’s why we can conclude that SQLite will be an excellent choice if you need a database to store a small amount of data and do not need to do complex manipulations.

### How to run

```bash
$ docker-compose up
```

### Result

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