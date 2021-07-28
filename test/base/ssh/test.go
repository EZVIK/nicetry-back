package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strconv"
)

type SSHClient struct {
	Client        *ssh.Client
	IPAddress     string
	Username      string
	Password      string
	PrivateSSHKey string
	Port          int

	Session *ssh.Session
}

func (s *SSHClient) Connect() error {

	addr := s.IPAddress + ":" + strconv.Itoa(s.Port)

	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User:            s.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(s.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
		return err
	}

	s.Client = client
	return nil
}

func (s *SSHClient) Exec(cmd string) (res string, err error) {

	if s.Session == nil {
		if err := s.NewSession(); err != nil {
			return "", err
		}
	}

	resByte, err := s.Session.Output(cmd)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to run command, Err:%s", err.Error())
		os.Exit(0)
	}
	output := string(resByte)

	fmt.Println(output)

	return output, nil
}

func (s *SSHClient) NewSession() error {

	session, err := s.Client.NewSession()

	if err != nil {
		return err
	}

	s.Session = session

	return nil
}

func main() {

	s := SSHClient{
		Username:  "root",
		Password:  "elish828MKB",
		IPAddress: "159.75.82.148",
		Port:      22,
	}

	if err := s.Connect(); err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	output, err := s.Exec("ls -a")

	fmt.Println(output, err)

}
