<p align="center">
<img src="https://github.com/rocketlaunchr/react/raw/master/logo.png" alt="react" />
</p>

# Go with React [![GoDoc](http://godoc.org/github.com/rocketlaunchr/react?status.svg)](http://godoc.org/github.com/rocketlaunchr/react) [![Go Report Card](https://goreportcard.com/badge/github.com/rocketlaunchr/react)](https://goreportcard.com/report/github.com/rocketlaunchr/react)

Facebook's React is one of the most dominant libraries for front-end development around. Google's Go programming language is one of the most elegantly crafted languages for server development. Why not combine the two?

This package is an extremely thin wrapper over the native react.js API. The objective was to make it light-weight, developer-friendly and intuitive. You shouldn’t have to scour the documentation to get going — a few peeks should be adequate. If you know your way around the React API and you know a bit of Go, then you should be able to make prototypes and production-worthy applications in no time.

This package is best suited for making cross-platform Desktop applications using these technologies:

-   [webview](https://github.com/zserge/webview)
-   [lorca](https://github.com/zserge/lorca)
-   [go-astilectron](https://github.com/asticode/go-astilectron)
-   [gotron](https://github.com/Equanox/gotron)
-   **[electron.js](https://electronjs.org/)**

The package is **production ready**. An optional (but highly convenient) `elements` sub-package is also included.

See [Tutorial here](https://medium.com/@rocketlaunchr.cloud/go-with-react-de5ee4f01df9).

⭐ **the project to show your appreciation.**

## Dependencies

-   [React 18.2.0](https://www.npmjs.com/package/react) (it will probably work with lower)
-   [Gopherjs](https://github.com/gopherjs/gopherjs) (Go to Javascript transpiler)
-   [create-react-class](https://www.npmjs.com/package/create-react-class)
-   [react-addons-pure-render-mixin](https://www.npmjs.com/package/react-addons-pure-render-mixin) (optional: For creating a `PureComponent`)

## Installation

```
go get -u github.com/rocketlaunchr/react
```

## Examples

The examples can be [found here](https://github.com/rocketlaunchr/react/tree/master/examples):

### Uptime Timer

-   How to create React class components
-   How to pass props from parent to child
-   How to use **UnmarshalProps()** and **UnmarshalState()**
-   How to use **state()** and **setState()**
-   How to create strongly-typed structured props and states

### Event Handling

-   How to create React functional components
-   How to handle events (and pass extra arguments)
-   How to create a Ref and interact with dom object directly

### [Desktop Application](https://github.com/rocketlaunchr/desktop-application)

-   100% written in Go
-   Cross-platform (macOS, Win, Linux)
-   [Electron.js](https://electronjs.org/) based
-   How to bundle javascript dependencies using [rollup.js](https://rollupjs.org)

## Performance Tips

-   Use `-m` command line flag to instruct gopher.js to minify code. Then bundle+minify further with [rollup.js](https://rollupjs.org) xor [Webpack/UglifyJS](https://github.com/gopherjs/gopherjs/issues/136). A Webpack tutorial can be [found here](https://medium.com/ag-grid/webpack-tutorial-understanding-how-it-works-f73dfa164f01).
-   Apply [gzip compression](https://en.wikipedia.org/wiki/HTTP_compression)
-   Use int instead of (u)int8/16/32/64
-   Use float64 instead of float32
-   Avoid importing `fmt` at all costs (including indirectly). Use [fmtless](https://github.com/rocketlaunchr/react/forks/fmtless) instead.
-   Avoid importing `net/http` for http requests (including indirectly). Use [gopherjs-xhr](https://github.com/rocketlaunchr/gopherjs-xhr) instead.
-   Until GopherJS supports Go1.13+, avoid using the standard libraries `context` package because it uses `fmt`. Instead use [`context`](https://github.com/rocketlaunchr/react/forks/context) from `forks` sub-package.
-   Avoid importing `honnef.co/go/js/dom` if possible.
-   Use **react.JSFn()** and use native javascript functions as much as possible.
-   https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas
-   See if [jsgo](https://github.com/dave/jsgo) is appropriate for your web-based project.
-   To reduce file size, copy only what's required from `elements` sub-package.
-   For json unmarshaling, try [slim-decoder](https://github.com/gopherjs/gopherjs/wiki/Using-native-JSON-parsing-to-realize-a-slim-JSON-decoder) or use [json](https://github.com/rocketlaunchr/react/forks/encoding/json)

## Future Work

-   WebAssembly version ![Help Required](https://img.shields.io/badge/help-required-blueviolet)

## Other useful packages

- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Statistics and data manipulation
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go
- [electron-alert](https://github.com/rocketlaunchr/electron-alert) - SweetAlert2 for Electron Applications
- [igo](https://github.com/rocketlaunchr/igo) - A Go transpiler with cool new syntax such as fordefer (defer for for-loops)
- [mysql-go](https://github.com/rocketlaunchr/mysql-go) - Properly cancel slow MySQL queries
- [remember-go](https://github.com/rocketlaunchr/remember-go) - Cache slow database queries

#

### Legal Information

The license is a modified MIT license. Refer to `LICENSE` file for more details.

**© 2018-20 PJ Engineering and Business Solutions Pty. Ltd.**

### Final Notes

Feel free to enhance features by issuing pull-requests.
