# go-sfse
SCP File Send Easy at Go.

And this is created to learn Go.
Sorry I am golang noob and code is low quality.

**Work in progress**

## What is this project?

When execute file is executed, Send file at SCP and Execute remote command at SSH.<br>
Those setting is write to config.toml.


## Motivation

This project motivation is want to solve two operation at very easy.

1. Send file at SCP
2. Execute remote command at SSH in  




## Binary Build

If you want to make a binary of macOS or Linux or Windows.

```bash
# build script download
wget https://gist.githubusercontent.com/shinshin86/3962ff8de51465320cd1aca5f8c05671/raw/69e1c3f4ec9260afe5f4d57754e64c96eeb651e8/go_build01.sh

# build (After change to proper permission)
./go_build01.sh sfse.go

```



Binary file export to bin

```bash
go-sfse - bin
        ├── darwin386
        │   └── sfse
        ├── darwin64
        │   └── sfse
        ├── linux386
        │   └── sfse
        ├── linux64
        │   └── sfse
        ├── windows386
        │   └── sfse.exe
        └── windows64
            └── sfse.exe

```




## Develop

TODO : 
About procedure for conducting development.
(Work in progress.)



```bash
# clone
git clone https://github.com/shinshin86/go-sfse.git
cd go-sfse

# config
cp -p config.toml_example config.toml
vim config.toml

# go run
go run sfse.go
```

