package main

import "testing"

func BenchmarkInsertsSqlite10000(b *testing.B) {
	benchmarkInsertsSqlite(10_000)
}

func BenchmarkInsertsPostgres10000(b *testing.B) {
	benchmarkInsertsPostgres(10_000)
}

func BenchmarkUpdateSqlite(b *testing.B) {
	benchmarkPrepareSqlite(10_000)

	b.ResetTimer()

	benchmarkUpdatesSqlite()
}

func BenchmarkUpdatePostgres(b *testing.B) {
	benchmarkPreparePostgres(10_000)

	b.ResetTimer()

	benchmarkUpdatesPostgres()
}

func BenchmarkSelectSqlite(b *testing.B) {
	benchmarkPrepareSqlite(10_000)

	b.ResetTimer()

	benchmarkSelectSqlite()
}

func BenchmarkSelectPostgres(b *testing.B) {
	benchmarkPreparePostgres(10_000)

	b.ResetTimer()

	benchmarkSelectPostgres()
}
