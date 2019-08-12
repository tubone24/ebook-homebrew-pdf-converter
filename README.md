# ebook-homebrew-rust-client

> Make pdf file which you use e-books by take in some image files such as jpeg, png and gif.

[![CircleCI](https://circleci.com/gh/tubone24/ebook-homebrew-rust-client.svg?style=svg)](https://circleci.com/gh/tubone24/ebook-homebrew-rust-client)

## What is this app?

This app is a Demo client CLI Client for using [ebook-homebrew](https://github.com/tubone24/ebook_homebrew).

This app is created `Rust`, so you have to build app for `cargo`.

## Project build

```
cargo run
```

Build success, create execute file on `target/debug/ebook-homebrew-rust-client`

## Usage

Show help, set opt `-h`

```
$ cd target/debug
$ ./ebook-homebrew-rust-client -h

ebook_homebrew_rust_client 0.1.0
tubone24 <tubo.yyyuuu@gmail.com>
Ebook-homebrew Command Line Tools

USAGE:
    ebook-homebrew-rust-client.exe [SUBCOMMAND]

FLAGS:
    -v, --version    CLI version
    -h, --help       Prints help information

SUBCOMMANDS:
    status      check server status
    upload      upload image files
    convert     convert image files to PDF
    download    download converted PDF file
    help        Prints this message or the help of the given subcommand(s)
```

### Example

```
$ cd target/debug
$ ./ebook-homebrew-rust-client -h
$ ebook-homebrew-rust-client upload image_directory png

============================
Upload image file
============================
SKIP: aaaa\66623244_2156448614404555_1652487473135091712_n.jpg
ADD: aaaa\standalone-choose-img.png
ADD: aaaa\standalone-top.png
ADD: aaaa\test_001.png
{"release_date": "2019-08-02T14:39:56.001293", "upload_id": "/tmp/tmpz79kzx2l"}

$ ebook-homebrew-rust-client convert "/tmp/tmpz79kzx2l" png

============================
Convert PDF
============================
{"release_date": "2019-08-02T14:40:51.017644", "upload_id": "/tmp/tmpz79kzx2l"}

$ ebook-homebrew-rust-client download "/tmp/tmpz79kzx2l" test.pdf

============================
Download PDF
============================
Response { url: "https://ebook-homebrew.herokuapp.com/convert/pdf/download", status: 200, headers: {"connection": "keep-alive", "date": "Fri, 02 Aug 2019 14:42:03 GMT", "server": "uvicorn", "content-type": "application/pdf", "vary": "Accept-Encoding", "via": "1.1 vegur"} }
```

## Completions

This CLI has prepared shell completions include shell type below.

* Bash
* Zsh
* Fish
* PowerShell

Use Those!

## Test

Use cargo test.
```
cargo test
```
