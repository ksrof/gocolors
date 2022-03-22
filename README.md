# gocolors
Color library for `Golang` which works along the standard library `fmt` and `log` packages. It can render text in a wide variety of colors giving your programs a much more outstanding output.

![Gocolors output](https://github.com/ksrof/gocolors/blob/main/images/gocolors_output.png)
## Roadmap
- [x] 8 bit color rendering
- [x] 16 bit color rendering
- [x] Implement tests
- [x] Documentation
- [x] CI workflow

The following ideas should be included in further releases, but the intentions are to have all this features added in the first stable version of the **gocolors color library**.

**Important note** This project is still in its early stages so expect breaking changes in the API, experiment, play and contribute but don't use it in production code or something that is serious enough.

### Gocolors Version 1.0.0
- [ ] 256 bit, RGB and True colors support
- [ ] File separation between color types
- [ ] Custom built color
- [ ] Option to plug in with the standard library `fmt` or `log`

## Examples
You can start using the color library just by printing something:
```go
package main

import (
	color "github.com/ksrof/gocolors"
)

func main() {
	// The Color method takes three parameters, the first is the
	// content that is going to be printed, the second is the color
	// and the third is the decoration.
	fmt.Println(color.Color("This should be underlined and green", "green", "underline"))
}
```
As you see it is pretty easy to use, but still there's a lot of room to improve!

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/gocolors/blob/main/LICENSE) for more details.
