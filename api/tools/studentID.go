package tools

import (
	"go-admin/common/global"
	"sync"
)

var studentIdLock sync.Mutex

/**
获取最新的学号
*/
func GetNewStudentId() int {
	studentIdLock.Lock()

	id := global.StudentId + 1

	global.StudentId = id

	studentIdLock.Unlock()
	return id
}
