# Legion
A fast, easy to use, command line image to ascii converter.

![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/hrszpuk/legion/master)
![GitHub issues](https://img.shields.io/github/issues/hrszpuk/legion)

## Installation
The installation of Legion is broken down into three distinct parts: downloading project files, building project files, and installing the executable.
Please ensure you have a Go (Golang) compiler, and git installed on your computer.

### Downloading project files
Project files can be downloaded onto your computer using `git`. Please use the commands below to download the source code.
```
git clone https://github.com/hrszpuk/legion.git
```
After running this command you should have a folder called "legion". Ensure you open that folder using `cd legion` or opening the directory in a GUI file explorer.

### Building project files
Next we need to build the project files into an executable file that can be ran on your computer. We do this using the Go compiler.
```
go build -o legion main.go
```
After this an executable file called "legion" should be created. Ensure the executable has permissions to run by using the following command:
```
sudo chmod +x legion
```

### Installing the executable
Now we have the executable you should be able to run legion using `./legion` while in the same directory as the executable file. In order to have this file accessible to your across the system you must either add the folder contraining the legion executable to your PATH variable or move legion into `/usr/bin` with your other system executables.

## Usage
After installation using the `legion` executable is as easy as using any other command line utility. <br>
You can use the following to convert image files (jpg, png, gif) to ascii:
```
legion path/to/file
```
In addition, legion provides two flags `-info` and `-v` that can be used to get additional file information, and additional processing information. These flags must be placed before the `path/to/file` as shown below.
```
legion -info path/to/file
```
```
legion -v path/to/file
```

## Contribution
This is one of the first project on this account and I appricate any feedback you may have about the project itself. if you have any issues or cool ideas please create a GitHub issue and I will try to resolve any issues swiftly. There will be issue templates for problems and feature suggestions. 


![Twitter Follow](https://img.shields.io/twitter/follow/hrszpuk?style=social)
![GitHub followers](https://img.shields.io/github/followers/hrszpuk?style=social)
