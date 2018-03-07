# meminfo

This is based off of Guillermo Álvarez Fernández's go.procmeminfo (https://github.com/guillermo/go.procmeminfo)

[![GoDoc](http://godoc.org/github.com/guillermo/go.procmeminfo?status.png)](http://godoc.org/github.com/guillermo/go.procmeminfo)

Package procmeminfo provides an interface for /proc/meminfo

```golang
    import "github.com/guillermo/go.procmeminfo"
    meminfo := &procmeminfo.MemInfo{}
    meminfo.Update()

    (*meminfo)['Cached'] // Get cached memory
    (*meminfo)['Buffers'] // Get buffers size
    (*meminfo)['...'] // Any field in /proc/meminfo

    meminfo.Total() // Total memory size in bytes
    meminfo.TotalAvailable() // Free Memory (Free + Cached + Buffers)
    meminfo.Available() // Free Memory (Free)
    meminfo.Used() // Total - Used (Total - Available)
    meminfo.TotalUsed() // (Total - TotalAvailable)
    meminfo.Cached() // Portion of RAM cached by Kernel, but available to programs if needed
    meminfo.Buffers() // -um- Buffers.
```


## Docs

Visit: http://godoc.org/github.com/guillermo/go.procmeminfo

## LICENSE

BSD
