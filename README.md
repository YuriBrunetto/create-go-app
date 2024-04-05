# create-go-app 🪤

> Minor CLI tool to create a simple Go project folder structure with `Makefile`

## Why?

Simply because it got boring to always have to create by hand a new folder with the necessary files and commands inside `Makefile`. And of course, for study reasons 🤓

# Usage

```bash
$ create-go-app your-app
```

## The structure

```
your-app
├── Makefile
└── main.go
```

And the `Makefile` will look with something like this:

```Makefile
build:
	@go build -o bin/your-app

run: build
	@./bin/your-app
```
