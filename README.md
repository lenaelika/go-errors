# Go Errors

Primitives for cascading error handling.  
In-place replacement for `errors.New()`, analogue of `fmt.Errorf()`.  
Based on [pkg/errors](https://github.com/pkg/errors) without a stack trace for simple logging.

## Usage

```go
// make a simple error
if token == "" {
    return errors.New("token is not provided")
}

// make an error with a formatted message
if _, ok := dict[key]; !ok {
    return errors.Errorf("key %q is not set", key)
}

// add context to an error
content, err := ioutil.ReadFile("data.txt")
if err != nil {
    return errors.Wrap(err, "failed to get records")
}

// add a formatted annotation to an error
db, err := gorm.Open("postgres", url)
if err != nil {
    return errors.Wrapf(err, "failed to connect to %s", url)
}

// retrieve the cause of an error
switch errors.Cause(err).(type) {
case *os.PathError:
    // handle specifically
default:
    // general error
}
```

## Installation

```
$ go get github.com/lenaelika/go-errors
```

Please feel free to submit issues and send pull requests. 🇷🇺🇬🇧
