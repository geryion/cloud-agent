package handler

import (
	"cloud-agent/common/log"
	"fmt"
)

func (cc *CloudHandler) GetMainPage(user string) (string, error) {
	prefix := fmt.Sprintf("CLoud->GetMainPage: %v.", user)
	log.CouldLog.Info(prefix)
	mp, err := Get
}