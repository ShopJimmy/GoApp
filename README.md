# GoApp

This repository contains a basic cross-platform GUI application written in Go using the [Fyne](https://fyne.io/) toolkit. The application opens a square white window with a simple left-side navigation containing **Home** and **Settings**.

## Building

You will need Go installed. Fetch the dependencies and build the application:

```bash
go mod tidy
go build
```

To cross-compile for Windows or macOS:

```bash
# Windows 64bit
env GOOS=windows GOARCH=amd64 go build -o GoApp.exe

# macOS 64bit
env GOOS=darwin GOARCH=amd64 go build -o GoApp
```

Running these commands requires network access to download modules the first time.
