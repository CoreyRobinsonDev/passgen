# passgen
Simple password generator CLI. Passwords generated are 16 characters long and contain a random distribution of lowercase letters, uppercase letters, digits, and symbols. 
<br>
[Usage](#Usage) <span>&nbsp;•&nbsp;</span> [Install](#Install)

## Usage
```bash
passgen | wl-copy
```
## Install
Download pre-built binary for your system here [Releases](https://github.com/CoreyRobinsonDev/passgen/releases).

### Compiling from Source
- Clone this repository
```bash
git clone https://github.com/CoreyRobinsonDev/passgen.git
```
- Create **passgen** binary
```bash
cd passgen
go build
```
- Move binary to <code>/usr/local/bin</code> to call it from anywhere in the terminal
```bash
sudo mv ./passgen /usr/local/bin
```
- Confirm that the program was built successfully
```bash
passgen
```

## License
[The Unlicense](./LICENSE)
