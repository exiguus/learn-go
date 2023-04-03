# Go introduction

## Install

<https://go.dev/doc/install>

```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf ~/Downloads/go1.20.2.linux-amd64.tar.gz
```

## Run

```bash
go run . 
```

## Build

```bash
go build -o build/helloworld .
# run build
./build/helloworld
```

## Cross Compile

```bash
GOOS=windows GOARCH=amd64 go build -o build/helloworld.exe .
```

```bash
GOOS=js GOARCH=wasm go build -o build/helloworld.wasm .
```

## Resources

- <https://go.dev/doc/>
- <https://go.dev/doc/tutorial/getting-started>
- <https://www.youtube.com/watch?v=eqSjKOPt7dg> (in german)
- <https://www.golang-book.com/books/intro>
