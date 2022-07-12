# int2words

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Integer to word converter

## Features
- Zero dependencies
- Easy to integrate.

## Installation
```shell
go get github.com/neotoolkit/int2words
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/neotoolkit/int2words"
)

func main() {
	fmt.Println(int2words.Convert(12345))
}
```
```console
twelve thousand three hundred forty five
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/neotoolkit/int2words/workflows/build/badge.svg
[build-url]: https://github.com/neotoolkit/int2words/actions
[pkg-img]: https://pkg.go.dev/badge/neotoolkit/int2words
[pkg-url]: https://pkg.go.dev/github.com/neotoolkit/int2words
[reportcard-img]: https://goreportcard.com/badge/neotoolkit/int2words
[reportcard-url]: https://goreportcard.com/report/neotoolkit/int2words
[coverage-img]: https://codecov.io/gh/neotoolkit/int2words/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/neotoolkit/int2words
[version-img]: https://img.shields.io/github/v/release/neotoolkit/int2words
[version-url]: https://github.com/neotoolkit/int2words/releases

## Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=neotoolkit">
    <img src="https://raw.githubusercontent.com/neotoolkit/.github/main/assets/sponsored_by_evrone.svg"
      alt="Sponsored by Evrone">
  </a>
</p>
