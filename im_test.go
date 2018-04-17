package taobao

import (
	"fmt"
	"testing"
)

func TestTaoBao_OpenIMGetUsers(t *testing.T) {
	fmt.Println("===== OpenIMGetUsers =====")

	var p = OpenIMGetUsersParam{}
	p.UserIds = []string{"56e279cfc77b930ae86b52e7", "55ff574b4ea0461440000003"}

	var rsp, err = client.OpenIMGetUsers(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}

func TestTaoBao_OpenIMAddUsers(t *testing.T) {
	fmt.Println("===== OpenIMAddUsers =====")

	var p = OpenIMAddUsersParam{}

	var u1 = OpenIMUserInfo{}
	u1.UserId = "56e279cfc77b930ae86b52e7"
	u1.Password = "a6facfa821ba92c3c17f4c3fae5442c2"
	u1.Nick = "我是管理员"
	u1.IconURL = "http://7xjcby.com2.z0.glb.qiniucdn.com/files%2Favatar%2F1%2F1_20150108073305311.jpg"

	var u2 = OpenIMUserInfo{}
	u2.UserId = "admin6"
	u2.Password = "123456"
	u2.Nick = "我应该换个名字啦"
	u2.IconURL = "http://7xjcby.com2.z0.glb.qiniucdn.com/files%2Favatar%2F1%2F1_20150123012518056.jpg"

	p.AddOpenIMUser(&u1)
	p.AddOpenIMUser(&u2)

	var rsp, err = client.OpenIMAddUsers(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}

func TestTaoBao_OpenIMUpdateUsers(t *testing.T) {
	fmt.Println("===== OpenIMUpdateUsers =====")

	var p = OpenIMUpdateUsersParam{}

	var u1 = OpenIMUserInfo{}
	u1.UserId = "admin6"
	u1.Password = "123456"

	var u2 = OpenIMUserInfo{}
	u2.UserId = "admin511"
	u2.Password = "123456"

	p.AddOpenIMUser(&u1)
	p.AddOpenIMUser(&u2)

	var rsp, err = client.OpenIMUpdateUsers(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}

func TestTaoBao_OpenIMDeleteUsers(t *testing.T) {
	fmt.Println("===== OpenIMDeleteUsers =====")
	var p = OpenIMDeleteUsersParam{}
	p.UserIds = []string{"56e279cfc77b930ae86b52e7xv", "admin6x"}

	var rsp, err = client.OpenIMDeleteUsers(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
