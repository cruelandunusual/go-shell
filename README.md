# go-shell - a minimal shell written in Go
![Linux](https://img.shields.io/badge/-Linux-grey?logo=linux)
![macOS](https://img.shields.io/badge/-macOS-black?logo=apple)
![wsl](https://img.shields.io/badge/-wsl-red)
<br>
<img style="vertical-align: middle; height: 60px; width: 60px;" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" />
<!--
<div align="left">
<img style="vertical-align: middle; height: 20px; width: 59px;" src="https://img.shields.io/badge/-Linux-grey?logo=linux" />
<img style="vertical-align: middle; height: 20px; width: 59px;" src="https://img.shields.io/badge/-macOS-black?logo=apple" />
<img style="vertical-align: middle; height: 20px; width: 59px;" src="https://img.shields.io/badge/-Windows-red" /><br><br>
<img style="vertical-align: middle; height: 40px; width: 40px;" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" />
-->
## Description

A simple shell interpreter written in Go.  

This project initially started by following [this basic tutorial](https://blog.init-io.net/post/2018/07-01-go-unix-shell/) as a way to learn Go.  
So far it mainly executes Linux commands using Go's builtin `Run()` mechanism from the `os/exec` package:
```go
	// prepare the command to execute
	cmd := exec.Command(args[0], args[1:]...) // variadic argument expansion

	// set appropriate outputs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// run the command, returning its results and exit status
	return cmd.Run()
```
I've enhanced the mechanism for changing directory with the standard Bash `cd` command, so that it now supports analogous commands from other shells, e.g. `Set-Location`, `chdir`, among others.  

I've added the hostname to the prompt, as well as providing a shell builtin to customise the rest of the prompt string with an arbitrary string;  
enter `setPrompt` followed by your new prompt message *without* double quotes, e.g.:  
```
setPrompt my custom prompt message
```
> [!TIP]  
> Builds can be done with Go's own build system, however there is support for [Task](https://taskfile.dev/), a simpler alternative to Make, which allows for more customisable options than Go's build system provides.  
> 
> Follow the [instructions to install Task](https://taskfile.dev/installation/);  
> then executing `task build` will create an executable called `go-shell` in `bin/`.  
