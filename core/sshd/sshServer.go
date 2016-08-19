package sshd

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func ConnectWithKey(user, host string, port int, keypath string) (*ssh.Session, error) {

	var session *ssh.Session
	//key, err := ioutil.ReadFile("/Users/zhengxq/.ssh/id_rsa")
	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
	}

	// Connect to the remote server and perform the SSH handshake.
	addr := fmt.Sprintf("%s:%d", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	//defer client.Close()
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil

}

func ConnectWithPd(user, password, host string, port int) (*ssh.Session, error) {
	var (
		authpd       []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	// get auth method
	authpd = []ssh.AuthMethod{
		ssh.Password(password),
	}

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    authpd,
		Timeout: 30 * time.Second,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}
