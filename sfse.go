package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

type Config struct {
	Server   ServerConfig
	Command  Command
	Sendfile Sendfile
	Destpath Destpath
}

type ServerConfig struct {
	Host string
	Port string
	User string
	Key  string
}

type Command struct {
	Hello string
}

type Sendfile struct {
	Zip string
}

type Destpath struct {
	Test string
}

func LoadConfig() *Config {
	// read config
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	// server
	fmt.Println("--------------> Connection Server")
	fmt.Printf("Host is :%s\n", config.Server.Host)
	fmt.Printf("Post is :%s\n", config.Server.Port)
	fmt.Printf("User is :%s\n", config.Server.User)
	fmt.Printf("Key  is :%s\n", config.Server.Key)

	// command
	fmt.Println("--------------> Exdcution Commmand")
	fmt.Printf("Command is :%s\n", config.Command.Hello)

	// send file
	fmt.Println("--------------> Send file")
	fmt.Printf("Sendfile is :%s\n", config.Sendfile.Zip)

	// dest file
	fmt.Println("--------------> Dest path")
	fmt.Printf("Dest Path is :%s\n", config.Destpath.Test)

	return &config
}

func SetSSHConfig(config *Config) *ssh.ClientConfig {
	buf, err := ioutil.ReadFile(config.Server.Key)
	if err != nil {
		panic(err)
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		panic(err)
	}

	sshconf := &ssh.ClientConfig{
		User: config.Server.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return sshconf
}

func SCPRun(config *Config, sshconf *ssh.ClientConfig) int {
	fmt.Println("----------> File send : START")

	var connstr = config.Server.Host + ":" + config.Server.Port
	conn, err := ssh.Dial("tcp", connstr, sshconf)

	fmt.Println("-------->Access to remote server")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect ssh: %v", err)
		return 1
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open new session: %v", err)
		return 1
	}
	defer session.Close()

	filepath := config.Sendfile.Zip
	destpath := config.Destpath.Test

	cperr := scp.CopyPath(filepath, destpath, session)
	fmt.Println(cperr)

	fmt.Println("----------> File send : END")

	return 0
}

func CMDRun(config *Config, sshconf *ssh.ClientConfig) int {
	fmt.Println("----------> Execute command at connect ssh server : START")

	var connstr = config.Server.Host + ":" + config.Server.Port
	conn, err := ssh.Dial("tcp", connstr, sshconf)

	fmt.Println("-------->Access to remote server")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect ssh: %v", err)
		return 1
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open new session: %v", err)
		return 1
	}
	defer session.Close()

	// execute command
	cmderr := session.Run(config.Command.Hello)
	if cmderr != nil {
		fmt.Fprintf(os.Stderr, "cannot execute command: %v", err)
		return 1
	}

	fmt.Println("----------> Execute command at connect ssh server : END")
	return 0
}

func main() {
	config := LoadConfig()
	sshconf := SetSSHConfig(config)
	fmt.Println(SCPRun(config, sshconf))

	fmt.Println("----------> Successful!!")
	fmt.Println("----------> Please Enter key...")
	bufio.NewScanner(os.Stdin).Scan()

	os.Exit(CMDRun(config, sshconf))
}
