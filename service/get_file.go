package service

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (s Service) getFile(sshCliend *ssh.Client, sftpClient *sftp.Client) (file string) {
	return ""
}