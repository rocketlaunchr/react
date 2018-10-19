Go with React [![GoDoc](http://godoc.org/github.com/rocketlaunchr/react?status.svg)](http://godoc.org/github.com/rocketlaunchr/react) [![Go Report Card](https://goreportcard.com/badge/github.com/rocketlaunchr/react)](https://goreportcard.com/report/github.com/rocketlaunchr/react)
===============

Facebook's React is one of the most dominant libraries for front-end development around. Google's Go programming language is one of the most elegantly crafted languages for server development. Why not combine the two?

This package is an extremely thin wrapper over the native react.js API. The objective was to make it light-weight, developer-friendly and intuitive. You shouldn’t have to scour the documentation to get going — a few peeks should be adequate. If you know your way around the React API and you know a bit of Go, then you should be able to make prototypes and production-worthy applications in no time. 

The package is **production ready** but it should be noted that the `elements` sub-package is not stable yet. Once stability is reached, version `1.0.0` will be tagged. The `elements` sub-package is optional but highly convenient.
It is recommended your package manager locks to a commit id instead of the master branch directly.

See [Tutorial here](https://medium.com/@rocketlaunchr.cloud/go-with-react-de5ee4f01df9).

## Dependencies

* [React 16.5.2](https://www.npmjs.com/package/react) (it will probably work with lower)
* [Gopherjs](https://github.com/gopherjs/gopherjs)
* [create-react-class](https://www.npmjs.com/package/create-react-class)

## Installation

```
go get -u github.com/rocketlaunchr/react
```

## Examples

The examples can be [found here](https://github.com/rocketlaunchr/react/tree/master/examples):
### Uptime Timer

* How to create React class components
* How to pass props from parent to child
* How to use **HydrateProps()** and **HydrateState()**
* How to use **state()** and **setState()**
* How to create strongly-typed structured props and states

### Event Handling

* How to create React functional components
* How to handle events (and pass extra arguments)
* How to create a Ref and interact with dom object directly

## Performance Tips

* Use `-m` command line flag to instruct gopher.js to minify code. [Then minify further with webpack](https://github.com/gopherjs/gopherjs/issues/136).
* Apply [gzip compression](https://en.wikipedia.org/wiki/HTTP_compression)
* Use int instead of (u)int8/16/32/64
* Use float64 instead of float32
* Try to avoid importing `fmt` package (either directly or indirectly)
* Use **react.JSFn()** and use native javascript functions as much as possible.
* https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas

#

### Legal Information

The license is a modified MIT license. Refer to `LICENSE` file for more details.

**© 2018 PJ Engineering and Business Solutions Pty. Ltd.**

### Final Notes

Feel free to enhance features by issuing pull-requests.

**Star** the project to show your appreciation.
