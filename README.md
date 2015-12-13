[![wercker status](https://app.wercker.com/status/5d3bc7d53a19bfc2a9100e5f4eb2d4d9/s/master "wercker status")](https://app.wercker.com/project/bykey/5d3bc7d53a19bfc2a9100e5f4eb2d4d9)

# icalparser

lexical ical parser for golang

## USAGE

```go
f, err := os.Open("/path/to/ical")
if err != nil {
        log.Fatal(err)
}
obj, err := icalparser.NewParser(f).Parse()
if err != nil {
        log.Fatal(err)
}

var b bytes.Buffer
icalparser.NewPrinter(obj).WriteTo(&b)
```

## TODO

* many, many refactoring
* writing document
* folding with rune

## SEE ALSO

* https://tools.ietf.org/html/rfc5545

## LICENSE

* MIT
