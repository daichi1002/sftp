package service

import (
	"os"
	"path/filepath"
	"sftp/constant"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func (s Service) getSalesData(sshClient *ssh.Client, sftpClient *sftp.Client) {
	// 取得対象となるファイルのパスを設定
	serverFilePath := filepath.Join(constant.SERVER_DIR, constant.FILE_NAME)
	localFilePath := filepath.Join(constant.TMP_DIR, constant.FILE_NAME)

	serverFile, err := sftpClient.Open(serverFilePath)

	if err != nil {
		s.logger.Fatalf("Failed to get file : %v", err)
	}

	localFile, err := os.Create(localFilePath) // ローカル
	if err != nil {
		s.logger.Fatalf("Failed to create file : %v", err)
	}

	_, err = serverFile.WriteTo(localFile)

	if err != nil {
		s.logger.Fatalf("Failed to write file : %v", err)
	}

	defer localFile.Close()
	defer serverFile.Close()
}
