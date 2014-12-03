// Copyright (c) 2014 Akeda Bagus <admin@gedex.web.id>
// Licensed under MIT license.
//
// Utilitas CLI untuk menterjemahkan teks (menggunakan Yandex API)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
)

const (
	TranslateURL  = "https://translate.yandex.net/api/v1.5/tr.json/translate"
	DefaultAPIKey = "trnsl.1.1.20141203T051248Z.447b0f450a99eeab.923b88210e4887f07e143074a15631f39a313156"
)

type Response struct {
	Code int
	Lang string
	Text []string
}

var (
	from    = flag.String("from", "en", "Terjemahin dari bahasa XX, default en.")
	to      = flag.String("to", "id", "Terjemahin ke bahasa YY, default id.")
	api_key = flag.String("api_key", DefaultAPIKey, fmt.Sprintf("Yandex API key, default %s", DefaultAPIKey))

	responseErrors map[int]string
)

func init() {
	responseErrors = map[int]string{
		401: "Invalid API key",
		402: "API key sudah diblok",
		403: "Permintaan sudah mencapai limit harian",
		404: "Permintaan sudah mencapai limit harian untuk jumlah teks yang diterjemahkan",
		413: "Panjang teks melebihi batas",
		422: "Teks tidak dapat diterjemahkan",
		501: fmt.Sprintf("Tidak mendukung penterjemahan dari %s ke %s", *from, *to),
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: terjemahin [flags] teks\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	terjemahin(flag.Args())
}

func terjemahin(text []string) {
	v := url.Values{}
	v.Set("key", *api_key)
	v.Set("lang", fmt.Sprintf("%s-%s", *from, *to))
	for _, t := range text {
		v.Add("text", t)
	}

	resp, err := http.PostForm(TranslateURL, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when requesting: %s\n", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	if errMsg, foundErr := responseErrors[resp.StatusCode]; foundErr {
		fmt.Fprintf(os.Stderr, "Error: %s\n", errMsg)
		os.Exit(2)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error status code: %d\n", resp.StatusCode)
		os.Exit(2)
	}

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding JSON: %s\n", err)
		os.Exit(2)
	}
	fmt.Println(strings.Join(r.Text, " "))
}
