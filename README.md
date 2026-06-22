# Golang Session

A beginner-friendly practice repository for learning the basics of **Go (Golang)** through small example files.

## Repository

GitHub repo: https://github.com/aliraza910/golang-session

---

# What is in this repo?

This repository contains practice files for learning Go basics, including topics like:

* Hello World
* Variables
* Constants
* Arrays
* For loops
* If/Else
* Switch statements

Current practice files include:

* `hello.go`
* `varibales.go`
* `Constants.go`
* `arrays.go`
* `forLoop.go`
* `ifelse.go`
* `switch.go`

---

# 1) Install Go (Golang)

## Windows

1. Go to the official Go download page:

    * https://go.dev/dl/
2. Download the **Windows installer** (`.msi`)
3. Run the installer and complete the setup
4. Open **Command Prompt** or **PowerShell**
5. Verify installation:

```bash
go version
```

You should see output like:

```bash
go version go1.xx.x windows/amd64
```

---

## Linux

### Ubuntu / Debian

```bash
sudo apt update
sudo apt install golang-go -y
go version
```

> Note: The apt version may be older. For the latest version, use the official tarball from https://go.dev/dl/

### Official tarball method

```bash
wget https://go.dev/dl/go1.24.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

---

## macOS

### Using installer

1. Visit: https://go.dev/dl/
2. Download the macOS package
3. Install it
4. Verify:

```bash
go version
```

### Using Homebrew

```bash
brew install go
go version
```

---

# 2) Learn Go (Best beginner resources)

## Official Go Tour

Best for absolute beginners:

* https://go.dev/tour/

## Official Go Documentation

* https://go.dev/doc/

## Go by Example

Very useful for short, practical examples:

* https://gobyexample.com/

## Learn X in Y Minutes - Go

Quick syntax overview:

* https://learnxinyminutes.com/docs/go/

## Effective Go

Best once you know the basics:

* https://go.dev/doc/effective_go

---

# 3) Clone this repository

Run:

```bash
git clone https://github.com/aliraza910/golang-session.git
cd golang-session
```

---

# 4) How to run the practice files

Go files can be run directly with:

```bash
go run filename.go
```

## Examples

### Run Hello World

```bash
go run hello.go
```

### Run variables practice

```bash
go run varibales.go
```

### Run constants practice

```bash
go run Constants.go
```

### Run arrays practice

```bash
go run arrays.go
```

### Run for loop practice

```bash
go run forLoop.go
```

### Run if/else practice

```bash
go run ifelse.go
```

### Run switch practice

```bash
go run switch.go
```

---

# 5) Suggested learning order for this repo

If you are new to Go, follow this order:

1. `hello.go`
2. `varibales.go`
3. `Constants.go`
4. `arrays.go`
5. `ifelse.go`
6. `switch.go`
7. `forLoop.go`

This gives you a smoother path from basic syntax to control flow.

---

# 6) Recommended practice workflow

For each file:

1. Open the file
2. Read the code carefully
3. Run it with `go run <filename.go>`
4. Change some values and run it again
5. Add your own examples inside the same file
6. Try to write the same logic from scratch in a new file

Example:

```bash
go run hello.go
```

Then edit `hello.go` and change the printed message.

---

# 7) Create your own Go practice file

You can create a new file like `functions.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("Learning Go functions!")
}
```

Run it with:

```bash
go run functions.go
```

---

# 8) Topics to add next in this repo

If you want to expand this practice repository, good next topics are:

* Functions
* Structs
* Slices
* Maps
* Pointers
* Packages
* Error handling
* Goroutines
* Channels
* File handling
* JSON handling
* HTTP server in Go

Suggested future file names:

* `functions.go`
* `slices.go`
* `maps.go`
* `structs.go`
* `pointers.go`
* `errors.go`
* `goroutines.go`
* `channels.go`
* `json.go`
* `httpserver.go`

---

# 9) Useful Go commands

## Run a file

```bash
go run hello.go
```

## Build a file

```bash
go build hello.go
```

## Format Go code

```bash
go fmt
```

## Check Go environment

```bash
go env
```

## Check installed Go version

```bash
go version
```

---

# 10) Troubleshooting

## `go` command not found

This means Go is not added to your system `PATH`.

### Fix:

* Reinstall Go from the official website
* Restart terminal after installation
* Run:

```bash
go version
```

---

## `git` command not found

Install Git first:

* https://git-scm.com/downloads

Then clone the repo again.

---

## File does not run

Make sure:

* You are inside the repo folder
* The file name is correct
* Go is installed properly

Example:

```bash
cd golang-session
go run hello.go
```

---

# 11) Goal of this repo

This repository is meant to help beginners practice Go fundamentals through small standalone examples.
The best way to use it is:

* run each file
* understand the output
* modify the code
* write your own variations

---

# Author

**Ali Raza**
GitHub: https://github.com/aliraza910

---
