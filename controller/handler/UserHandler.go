package handler

import (
	"cloud-agent/common/log"
	"fmt"
)

func (nm *CloudHandler) Login(user string, passwd string) (string, error) {
	prefix := fmt.Sprintf("CLoud->Login: %v.", user)
	log.CouldLog.Info(prefix)


	return "", nil
}
