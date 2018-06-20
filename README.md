# fanin-out
This program includes N generators and M readers. It uses worker pools, channels, threads (concurrency)

Program output
```go
iteractions := 5
worker_generators := 3
readers := 1
```
  


    >go run fanin-out.go
    Generator 1 has wrote 1. Counter by now 1
    Generator 1 has wrote 2. Counter by now 4
    Generator 1 has wrote 3. Counter by now 5
    Generator 1 has wrote 4. Counter by now 6
    Generator 1 has wrote 5. Counter by now 7
    Generator 2 has wrote 1. Counter by now 3
    read value:  1
    Have read by now 1
    read value:  1
    Have read by now 2
    read value:  1
    Have read by now 3
    read value:  2
    Have read by now 4
    read value:  3
    Have read by now 5
    read value:  4
    Have read by now 6
    read value:  5
    Have read by now 7
    read value:  2
    Have read by now 8
    Generator 2 has wrote 2. Counter by now 8
    Generator 2 has wrote 3. Counter by now 9
    Generator 2 has wrote 4. Counter by now 10
    Generator 2 has wrote 5. Counter by now 11
    Generator 3 has wrote 1. Counter by now 2
    Generator 3 has wrote 2. Counter by now 12
    Generator 3 has wrote 3. Counter by now 13
    Generator 3 has wrote 4. Counter by now 14
    Generator 3 has wrote 5. Counter by now 15
    Generators finished: 15
    read value:  3
    Have read by now 9
    read value:  4
    Have read by now 10
    read value:  5
    Have read by now 11
    read value:  2
    Have read by now 12
    read value:  3
    Have read by now 13
    read value:  4
    Have read by now 14
    read value:  5
    Have read by now 15
    We got 15 from 15 iteractionsLenght of CHANNEL 0 
    total time taken
    0.1036066 seconds
    
