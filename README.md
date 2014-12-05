terjemahin
=========

Utilitas CLI untuk menterjemahkan teks (menggunakan Yandex API)

*Program ini dibuat untuk screencast ["Menggunakan grunt init go main Untuk Membuat Aplikasi terjemahin"](https://www.youtube.com/watch?v=6AwlPkBpjQI) dibuat*.

## Install

```
go get github.com/golang-id/terjemahin
```

## Usage

```
$ terjemahin "This works like a charm"
Ini bekerja seperti mantra

$ terjemahin -from=id -to=fr "selamat makan"
bon appétit
```

## License

MIT license.
