# byteformat
Human readable and machine parsable byte formatter

# Features

- Converting long numbers to one of two human readable formats: 3294201 -> 3.14M or 3,294,201
- Customisable precision: 3294201 can be converted to 3M, 3.1M, 3.14M, 3.142M, 3.1416M. 3.14159M
- Working with sizes up 16 exabytes or 18.446.744.073.709.551.615 bytes
- Converting from human readable format to bytes: 10.5G -> 11274289152

# Examples
```go
byteformat.SetPrecision(2)
byteformat.HumanizeSize(3294201) // 3.14M
bytefotmat.HumanizeBytes(3294201) // 3,294,201
byteformat.FromString("10.5G") // 11274289152
```

# Functions

## func FromString
```go
func FromString(input string) (uint64, error)
```
Parses human readable number like 1.5M and returns ```uint64``` number. Can return ```UnknownUnit``` error.

## func HumanizeBytes
```go
func HumanizeBytes(size uint64) string
```
Breaks given number to 3-digit chunks splitted by a comma.

## func HumanizeSize
```go
func HumanizeSize(size uint64) string
```
Converts ```size``` to numan readable format.

## func SetPrecision
```go
func SetPrecision(precision uint8) error
```
Sets precision for future ```HumanizeSize()``` conversions. Default precision is 2. Can return ```TooPrecise``` error.

# Errors
## UnknownUnit
```go
var UnknownUnit = errors.New("Unknown unit")
```
Error is returned by ```FromString()``` when letter after the float(or integer) number is not one of b, k, m, g, t, p, e (case insensitive)

## TooPrecise
```go
var TooPrecise = errors.New("Precision must not exceed 5")
```
Error can be returned by ```SetPrecision()``` if given precision is greater than 5.
