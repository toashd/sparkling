# Sparkling
Go sparklines for the command lines.

[![CI Status](http://img.shields.io/travis/toashd/sparkling.svg?style=flat)](https://travis-ci.org/toashd/sparkling)

## Installation

Easily generate sparkling on the command line. Installation is as simple as

```bash
$ go get github.com/toashd/sparkling
```

## Usage
Use sparkling as a [binary](https://github.com/toashd/sparkling/tree/master/cmd) or just build your own sparkles with

```go
import "github.com/toashd/sparkling"
```

Create a new sparkling container, add a few different series of data and render to the specified `io.Writer` interface.

```go
sp := sparkling.New(os.Stdout)
sp.AddSeries([]float64{0, 30, 55, 80, 33, 150}, "1st series")
sp.AddSeries([]float64{23, 45, 23, 5, 1, 67, 8, 5}, "2nd series")
sp.Render()
```

This will give you

```bash
1st series  ▁▂▃▄▂█
2nd series  ▃▅▃▁▁█▁▁
```

Isn't it sparkling? Enjoy!

## Todo

There are a lot of things that can be improved. For example, but not limited to,

* add support for different colors,
* add more meta information, like labels etc,
* output different series/lines in parallel,
* add duration parameter,
* allow setting writers on series or render.

## Contribution

Please feel free to suggest any kind of improvements, refactorings, or just file an
issue, fork and submit a pull requests.

## License

Sparkling is available under the MIT license. See the LICENSE file for more info.
