# Go-based KVS

Learn Go by using it to build a KVS.

## Plan

The end goal is to build a distributed fault-tolerant KVS.
I want to make it high-performance, reliable.
This means I want to add replication, recovery, and consensus protocols etc.

Also, I will build in a disaggregated way: the management part will be separated from the data parts.

Project timeline: roughly two weeks.
Nov 14, 2021 to Nov 30, 2021. 

1. Planning the architecture.
2. data node KVS core
3. networking stack
4. APIs exposed
5. replication, consensus

## Related Work

- TiDB
- MangoDB
- https://github.com/gostor/awesome-go-storage
