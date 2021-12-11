go-input [![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs] [![Travis](https://img.shields.io/travis/qiuzhanghua/go-input.svg?style=flat-square)][travis] [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
====

[godocs]: http://godoc.org/github.com/qiuzhanghua/go-input
[travis]: https://travis-ci.org/qiuzhanghua/go-input
[license]: /LICENSE

`go-input` is a Go package for reading user input in console.

Here is the some good points compared with other/similar packages. It can handle `SIGINT` (`Ctrl+C`) while reading input and returns error. It allows to change IO interface as `io.Writer` and `io.Reader` so it's easy to test of your go program with this package (This package is also well-tested!). It also supports raw mode input (reading input without prompting) for multiple platform (Darwin, Linux and Windows). Not only this it allows prompting complex input via Option struct. 

The documentation is on [GoDoc][godocs].

## Install

Use `go get` to install this package:

```bash
$ go get github.com/qiuzhanghua/go-input
```

## Usage

The following is the simple example,

```golang
ui := &input.UI{
    Writer: os.Stdout,
    Reader: os.Stdin,
}

query := "What is your name?"
name, err := ui.Ask(query, &input.Options{
    Default: "Bob",
    Required: true,
    Loop:     true,
})

```

## I10N
```bash
go get github.com/qiuzhanghua/i10n
```
add local resource to your app, see 

https://github.com/qiuzhanghua/SimpleInput

```golang
i10n.SetDefaultLang("zh-CN")

// ...

input.T=i10n.T
```

You can check other examples in [here](/_example).

## Author

[Taichi Nakashima](https://github.com/tcnksm)

[邱张华](https://github.com/qiuzhanghua)