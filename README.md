# go-shell - a minimal shell in Go

<div align="left">
<img style="vertical-align: middle; height: 20px; width: 59px;" src="https://img.shields.io/badge/-Linux-grey?logo=linux" />
<img style="vertical-align: middle; height: 20px; width: 59px;" src="https://img.shields.io/badge/-macOS-black?logo=apple" />
<img style="vertical-align: middle; height: 20px; width: 59px;" src="https://img.shields.io/badge/-Windows-red" /><br><br>
<img style="vertical-align: middle; height: 40px; width: 40px;" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" />

## Description

A simple shell interpreter written in Go.  

Initially started by following [this basic tutorial](https://blog.init-io.net/post/2018/07-01-go-unix-shell/) as a way to learn Go.  
So far it simply executes Linux commands using Go's builtin mechanism for doing that:
```go
	// prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...) // variadic argument expansion

	// set appropriate outputs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// run the command, returning its results and exit status
	return cmd.Run()
```
There's support for changing directory using the standard Bash `cd` command, implemented using Go's `os.Chdir()`.   

I've added customisable prompt support to replace the default prompt string with an arbitrary string (type `setPrompt` followed by your new prompt message *without* double quotes)

Builds can be done with Go's own build system, however there is support for [Task](https://taskfile.dev/), a simpler alternative to Make.  

`task build` will create an executable called `go-shell` in `bin/`.  
