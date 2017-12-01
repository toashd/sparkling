# Sparkling
Render sparklines on the command line.

## Install
Build and install the sparkling binary with

```bash
$ go build -o sparkling
```

## Usage
After building and installing the binary just run sparkling and pass it a space-delimited list of numbers. It takes both, integer and float values. You can also pass an optional title flag to name the series.

```bash
  $ sparkling '0 30 55 80 33 150'
    ▁▂▃▄▂█

  $ echo 0 30 55 80 33 150 | sparkling
    ▁▂▃▄▂█

  $ sparkling -t=Awesome '23 45 23 5 1 67 8 5'
    Awesome ▃▅▃▁▁█▁▁
```
Invoke help and usage info with `sparkling -help`

## Other usages

For more awesome usages use the [sparkling package](https://github.com/toashd/sparkling). No matter how awesome you sparkle, keep me posted!

## License

Sparkling is available under the MIT license. See the LICENSE file for more info.
