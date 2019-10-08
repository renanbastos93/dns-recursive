# dns-recursive
This is a POC to test DNSR lib


### Getting starter

```bash
$ git clone https://github.com/renanbastos93/dns-recursive.git
$ cd dns-recursive
$ go get go get github.com/domainr/dnsr
$ go run main.go
```


## In the future
We plan to create a resolver recursive DNS, it was think some slices to implement it.

#### Todo
 - [ ] Dynamic cache
    - [ ] Expire
    - [ ] Mutex
    - [ ] Clean data
 - [ ] Recursive Resolver
 
 
 ### Suggestions
  Here we go comment some suggestion to implement this lib
  
  - wrap to libunbound or others
  - using net library with lookup methods
  - using context library with the library miek/dns
