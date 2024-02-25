# Concurrent-synchronous I/O

The program demonstrates how an I/O operation can be done much faster with the use of buffers.

It makes use of buffers for reading large chunks of data and writing them to a file, reducing system calls. The size of the buffer used can vary based on the size of the file and the underlying system's memory.

The program also uses goroutines for concurrent programming and go channels for synchronization. For smaller files, direct I/O is recommended.

## buffer-size to speed relationship
To test the relationship between buffer-size and I/O speed, an operation was done, involving reading from a very large file and writing the data to a new file.

The following results show the performance of the I/O operation with varying buffer sizes:

                      |   Execution time (seconds)       |                  
| Buffer-size (bytes) | test1 | test2 | test3  | Average |
|---------------------|-------|-------|--------|---------|
|          30         |   90  |   90  |   107  |    95   |
| 128                 | 26    | 26    | 27     | 26      |
| 256                 | 14    | 14    | 14     | 14      |
| 512                 | 7     | 7     | 7      | 7       |
| 1024                | 4     | 3     | 3      | 3       |
| 2048                | 2     | 2     | 2      | 2       |
| 4096                | 1     | 1     | 1      | 1       |
