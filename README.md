# Mikrotik/RouterOS API Client 

## WIP

Common methods fully workable, but not all API elements released.


### install:

```go
go get -u github.com/icadsistemi/lib-mikrotik
```

### usage:

```go
router, err := mikrotik.Dial(addr, user, pass)
// OR
router, err := mikrotik.DialTimeout(addr, user, pass, timeout)
```

API methods are presented how a tree, similar to the CLI commands RouterOS.

```go
//get router name
router.System.Identity.Name()

//add ip address
ip := mikrotik.IPAddress{Address: ..., Interface: ...}
router.IP.Address.Add(&ip)

//get all network interfaces
var intfs []mikrotik.Interface
router.Interface.List(&intfs)
```

Most methods accepts a pointer on the appropriate structure, example: `mikrotik.IPAddress` , `mikrotik.NATRule` etc... Structure field names can by founded by tag `mikrotik`. If tag not specified, go field name auto convert to RouterOS like format, example: `FieldName` converted to `field-name`, and back.
