package main
import (
	"fmt"
	sshClientDomain "../domain/client"
	"bytes"
	// "../contract/request"
	clientDomain "../domain/client"
	executerDomain "../domain/executer"
	"golang.org/x/crypto/ssh"
 )

 type SshClient struct {
}
//  type ISshClient interface {
// 	GetConnection(host string, username string, password string) ssh.Client
//  }


func InitSshClient() sshClientDomain.ISshClient {
	sshClient := &SshClient{}
	return sshClient
 }
 func (sc *SshClient) GetConnection(host string, username string, password string) ssh.Client {
	config := &ssh.ClientConfig{
	   User: username,
	   Auth: []ssh.AuthMethod{
		  ssh.Password(password),
	   },
	   HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
 sshConn, err := ssh.Dial("tcp", host, config)
	if err != nil {
	   fmt.Println(err.Error())
	} 
	return *sshConn
 }

//  type ISshCommandExecuter interface {
// 	RunSshCommand(host string, username string, password string, cmdName string, cmd string) string
//  }

 type SshCommandExecuter struct {
	sshClient         clientDomain.ISshClient
 }
 func InitSshExecuter(sc clientDomain.ISshClient) executerDomain.ISshCommandExecuter {
	executer := &SshCommandExecuter{
	   sshClient:         sc,
	}
	return executer
 }
 func (e *SshCommandExecuter) RunSshCommand(host string, username string, password string, cmd string) string {
	client := e.sshClient.GetConnection(host, username, password)
	sshSession, err := client.NewSession()
	if err != nil {
	   fmt.Println(err.Error())
	}
	defer sshSession.Close()
	defer client.Close()
	var stdoutBuf bytes.Buffer
	
	sshSession.Stdout = &stdoutBuf
	sshSession.Run(cmd)
	
	return stdoutBuf.String()
 }

 func main()  {
	sshClient := InitSshClient()
	sshCommandExecuter :=InitSshExecuter(sshClient)
	cmdResult := sshCommandExecuter.RunSshCommand("192.168.100.9:22", "swamym", "pramati123","virsh  create")
	//-insta -n myRHELVM7  --os-type=Linux --os-variant=rhel6 --ram=2048 --vcpus=2 --disk path=/var/lib/libvirt/images/myRHELVM7.img,bus=virtio,size=10 --cdrom /home/swamym/Downloads/ubuntu-18.04.4-desktop-amd64.iso --network bridge:virbr0")
	fmt.Println(cmdResult)
	
 }