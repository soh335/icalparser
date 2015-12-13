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

## SEE ALSO

* https://tools.ietf.org/html/rfc5545

## LICENSE

* MIT
