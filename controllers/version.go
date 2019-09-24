package controllers

import (
	"github.com/oldthreefeng/devops-api/common"
)

// Get 获取程序版本号
func (v *VersionController) Get() {
	v.JsonOK("Get App Version", common.GetVersion(), true)
}
