package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (s Service) connectSftp() (*ssh.Client, *sftp.Client) {
	// SSHクライアントの作成
	sshClient := s.createSshClient()

}

func (s Service) createSshClient() *ssh.Client {
	config, addr := s.createSshClientConfig()
}

func (s Service) createSshClientConfig() (ssh.ClientConfig, string) {
	signer := s.getSshKey()

	// 環境変数のセット
	var (
		userId = os.Getenv("UserId")
		ip     = os.Getenv("Ip")
		port   = os.Getenv("Port")
	)

	// 認証情報などを設定
	config := ssh.ClientConfig{
		// SSH ユーザ名
		User: userId,
		// 認証方式
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout: 5 * time.Second,
	}

	// アドレスの設定
	addr := fmt.Sprintf("%s:%s", ip, port)

	return config, addr
}

// 認証に必要な鍵情報の取得
func (s Service) getSshKey() ssh.Signer {
	buf, err := ioutil.ReadFile(os.Getenv("SshKeyPath"))
	if err != nil {
		s.logger.Fatalf("failed to read SshKeyPath: %v", err)
	}

	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		s.logger.Fatalf("failed to parse SshKey: %v", err)
	}

	return key
}
