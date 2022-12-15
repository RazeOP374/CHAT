package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUser map[int]*UserProcess
}

// 初始化
func init() {
	userMgr = &UserMgr{
		onlineUser: make(map[int]*UserProcess, 1024),
	}
}

// online添加
func (this *UserMgr) AddonlineUser(up *UserProcess) {
	this.onlineUser[up.UserId] = up
}

// delete online
func (this *UserMgr) DeleteUser(userId int) {
	delete(this.onlineUser, userId)
}

func (this *UserMgr) GetonlineUser() map[int]*UserProcess {
	return this.onlineUser
}
func (this *UserMgr) GetonlineById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUser[userId]
	if !ok {
		err = fmt.Errorf("ID不存在", userId)
		return
	}
	return
}
