# Go tests
Examples to demonstrate use of go tests with/without parallel flags
1. Sequential - 5 Tests each having 1S sleep without `t.Parallel()` statements.
2. Parallel - 5 Tests each having 1S sleep with `t.Parallel()` statements.

## Running tests without parallel flag (Default parallel flag value is GOMAXPROCS, equals runtime.NumCPU)
GOMAXPROCS - The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously

1. Sequential tests
```bash
go test -v .\sequential_test.go
```
Results 

```bash
=== RUN   TestSerial
=== RUN   TestSerial/First
***** Running First test =====
===== First test complete *****
=== RUN   TestSerial/Second
***** Running Second test =====
===== Second test complete *****
=== RUN   TestSerial/Third
***** Running Third test =====
===== Third test complete *****
=== RUN   TestSerial/Fourth
***** Running Fourth test =====
===== Fourth test complete *****
=== RUN   TestSerial/Fifth
***** Running Fifth test =====
===== Fifth test complete *****
--- PASS: TestSerial (5.03s)
    --- PASS: TestSerial/First (1.00s)
    --- PASS: TestSerial/Second (1.01s)
    --- PASS: TestSerial/Third (1.01s)
    --- PASS: TestSerial/Fourth (1.01s)
    --- PASS: TestSerial/Fifth (1.01s)
PASS
ok      command-line-arguments  7.226s

```
2. Parallel tests
```bash
go test -v .\parallel_test.go
```

Results
```bash
=== RUN   TestParallel
=== RUN   TestParallel/First
=== PAUSE TestParallel/First
=== RUN   TestParallel/Second
=== PAUSE TestParallel/Second
=== RUN   TestParallel/Third
=== PAUSE TestParallel/Third
=== RUN   TestParallel/Fourth
=== PAUSE TestParallel/Fourth
=== RUN   TestParallel/Fifth
=== PAUSE TestParallel/Fifth
=== CONT  TestParallel/First
***** Running First test =====
=== CONT  TestParallel/Second
***** Running Second test =====
=== CONT  TestParallel/Fifth
***** Running Fifth test =====
=== CONT  TestParallel/Fourth
***** Running Fourth test =====
=== CONT  TestParallel/Third
***** Running Third test =====
===== Second test complete *****
===== Third test complete *****
===== Fourth test complete *****
===== First test complete *****
===== Fifth test complete *****
--- PASS: TestParallel (0.00s)
    --- PASS: TestParallel/Third (1.01s)
    --- PASS: TestParallel/Fourth (1.01s)
    --- PASS: TestParallel/Fifth (1.01s)
    --- PASS: TestParallel/First (1.01s)
    --- PASS: TestParallel/Second (1.01s)
PASS
ok      command-line-arguments  1.131s

```


## Running tests with parallel flag = 1

1. Sequential tests
```bash
go test -v .\sequential_test.go -parallel 1
```
Results
```bash
=== RUN   TestSerial
=== RUN   TestSerial/First
***** Running First test =====
===== First test complete *****
=== RUN   TestSerial/Second
***** Running Second test =====
===== Second test complete *****
=== RUN   TestSerial/Third
***** Running Third test =====
===== Third test complete *****
=== RUN   TestSerial/Fourth
***** Running Fourth test =====
===== Fourth test complete *****
=== RUN   TestSerial/Fifth
***** Running Fifth test =====
===== Fifth test complete *****
--- PASS: TestSerial (5.04s)
    --- PASS: TestSerial/First (1.00s)
    --- PASS: TestSerial/Second (1.01s)
    --- PASS: TestSerial/Third (1.01s)
    --- PASS: TestSerial/Fourth (1.01s)
    --- PASS: TestSerial/Fifth (1.01s)
PASS
ok      command-line-arguments  5.178s

```

2. Parallel tests
```bash
go test -v .\parallel_test.go -parallel 1
```
Results
```bash
=== RUN   TestParallel
=== RUN   TestParallel/First
=== PAUSE TestParallel/First
=== RUN   TestParallel/Second
=== PAUSE TestParallel/Second
=== RUN   TestParallel/Third
=== PAUSE TestParallel/Third
=== RUN   TestParallel/Fourth
=== PAUSE TestParallel/Fourth
=== RUN   TestParallel/Fifth
=== PAUSE TestParallel/Fifth
=== CONT  TestParallel/First
***** Running First test =====
===== First test complete *****
=== CONT  TestParallel/Fourth
***** Running Fourth test =====
===== Fourth test complete *****
=== CONT  TestParallel/Fifth
***** Running Fifth test =====
===== Fifth test complete *****
=== CONT  TestParallel/Third
***** Running Third test =====
===== Third test complete *****
=== CONT  TestParallel/Second
***** Running Second test =====
===== Second test complete *****
--- PASS: TestParallel (0.00s)
    --- PASS: TestParallel/First (1.00s)
    --- PASS: TestParallel/Fourth (1.01s)
    --- PASS: TestParallel/Fifth (1.01s)
    --- PASS: TestParallel/Third (1.01s)
    --- PASS: TestParallel/Second (1.01s)
PASS
ok      command-line-arguments  5.160s

```

## Running tests with parallel flag = 8
1. Sequential tests
```bash
go test -v .\sequential_test.go -parallel 8
```
Results
```bash
=== RUN   TestSerial
=== RUN   TestSerial/First
***** Running First test =====
===== First test complete *****
=== RUN   TestSerial/Second
***** Running Second test =====
===== Second test complete *****
=== RUN   TestSerial/Third
***** Running Third test =====
===== Third test complete *****
=== RUN   TestSerial/Fourth
***** Running Fourth test =====
===== Fourth test complete *****
=== RUN   TestSerial/Fifth
***** Running Fifth test =====
===== Fifth test complete *****
--- PASS: TestSerial (5.04s)
    --- PASS: TestSerial/First (1.01s)
    --- PASS: TestSerial/Second (1.01s)
    --- PASS: TestSerial/Third (1.01s)
    --- PASS: TestSerial/Fourth (1.01s)
    --- PASS: TestSerial/Fifth (1.01s)
PASS
ok      command-line-arguments  5.165s

```

2. Parallel tests
```bash
go test -v .\parallel_test.go -parallel 8
```
Results 
```bash
=== RUN   TestParallel
=== RUN   TestParallel/First
=== PAUSE TestParallel/First
=== RUN   TestParallel/Second
=== PAUSE TestParallel/Second
=== RUN   TestParallel/Third
=== PAUSE TestParallel/Third
=== RUN   TestParallel/Fourth
=== PAUSE TestParallel/Fourth
=== RUN   TestParallel/Fifth
=== PAUSE TestParallel/Fifth
=== CONT  TestParallel/First
***** Running First test =====
=== CONT  TestParallel/Fifth
***** Running Fifth test =====
=== CONT  TestParallel/Fourth
***** Running Fourth test =====
=== CONT  TestParallel/Third
***** Running Third test =====
=== CONT  TestParallel/Second
***** Running Second test =====
===== Fifth test complete *****
===== Second test complete *****
===== First test complete *****
===== Third test complete *****
===== Fourth test complete *****
--- PASS: TestParallel (0.00s)
    --- PASS: TestParallel/Fifth (1.01s)
    --- PASS: TestParallel/First (1.01s)
    --- PASS: TestParallel/Second (1.01s)
    --- PASS: TestParallel/Fourth (1.01s)
    --- PASS: TestParallel/Third (1.01s)
PASS
ok      command-line-arguments  1.135s

```