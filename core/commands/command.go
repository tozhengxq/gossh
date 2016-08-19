package commands

import (
	sshd "gossh/core/sshd"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Conn struct {
	UseKey  bool
	User    string
	Host    string
	Port    int
	Connset string
}

func (cmd *Conn) Runcmd(c string) {

	//connset := cmd.SetCon()
	//获取ssh session
	if cmd.UseKey {
		session, err := sshd.ConnectWithKey(cmd.User, cmd.Host, cmd.Port, cmd.Connset)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Run(c)
	} else {
		session, err := sshd.ConnectWithPd(cmd.User, cmd.Connset, cmd.Host, cmd.Port)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Run(c)
	}

}

func (cmd *Conn) RunTerminal(c string) {
	//connset := cmd.SetCon()

	if cmd.UseKey {
		session, err := sshd.ConnectWithKey(cmd.User, cmd.Host, cmd.Port, cmd.Connset)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()

		fd := int(os.Stdin.Fd())
		oldState, err := terminal.MakeRaw(fd)
		if err != nil {
			panic(err)
		}
		defer terminal.Restore(fd, oldState)

		// excute command
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Stdin = os.Stdin

		termWidth, termHeight, err := terminal.GetSize(fd)
		if err != nil {
			panic(err)
		}

		// Set up terminal modes
		modes := ssh.TerminalModes{
			ssh.ECHO:          1,     // enable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}

		// Request pseudo terminal
		if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
			log.Fatal(err)
		}

		session.Run(c)
	} else {
		session, err := sshd.ConnectWithPd(cmd.User, cmd.Connset, cmd.Host, cmd.Port)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()

		fd := int(os.Stdin.Fd())
		oldState, err := terminal.MakeRaw(fd)
		if err != nil {
			panic(err)
		}
		defer terminal.Restore(fd, oldState)

		// excute command
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Stdin = os.Stdin

		termWidth, termHeight, err := terminal.GetSize(fd)
		if err != nil {
			panic(err)
		}

		// Set up terminal modes
		modes := ssh.TerminalModes{
			ssh.ECHO:          1,     // enable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}

		// Request pseudo terminal
		if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
			log.Fatal(err)
		}

		session.Run(c)
	}
}
