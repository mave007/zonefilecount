# zonefilecount
DNS zonefile count in CSV alike format

It will print number of the following DNS RR: 

A, AAAA, CNAME, NS,  MX


## Requirements
* dns: DNS library in Go - https://github.com/miekg/dns

## Build
```
go get github.com/mave007/zonefilecount
go build github.com/mave007/zonefilecount
```

## Usage
```
zonefilecount <zonefile>
```

## Acknowledgements
Based on https://github.com/fcambus/statzone

