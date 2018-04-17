package taobao

//import (
//	"fmt"
//	"testing"
//)
//
//func TestOpenIMAddUsersParam(t *testing.T) {
//	var p = OpenIMUpdateUsersParam{}
//
//	var u1 = OpenIMUserInfo{}
//	u1.UserId = "56e279cfc77b930ae86b52e7"
//	u1.Password = "a6facfa821ba92c3c17f4c3fae5442c2"
//	u1.Nick = "我是管理员"
//	u1.IconURL = "http://7xjcby.com2.z0.glb.qiniucdn.com/files%2Favatar%2F1%2F1_20150108073305311.jpg"
//
//	var u2 = OpenIMUserInfo{}
//	u2.UserId = "admin5"
//	u2.Password = "123456"
//	u2.Nick = "我应该换个名字啦"
//	u2.IconURL = "http://7xjcby.com2.z0.glb.qiniucdn.com/files%2Favatar%2F1%2F1_20150123012518056.jpg"
//
//	p.AddOpenIMUser(&u1)
//	p.AddOpenIMUser(&u2)
//
//	fmt.Println("===== OpenIMAddUserParam =====")
//	fmt.Println(RequestWithKey("23201003", "1e2dfd16981d75142597fd10131b17b5", p))
//}
//
//func TestOpenIMUpdateUsersParam(t *testing.T) {
//	var p = params.OpenIMUpdateUsersParam{}
//
//	var u1 = params.OpenIMUser{}
//	u1.UserId = "admin4"
//	u1.Password = "123456"
//
//	var u2 = params.OpenIMUser{}
//	u2.UserId = "admin5"
//	u2.Password = "123456"
//
//	p.AddOpenIMUser(&u1)
//	p.AddOpenIMUser(&u2)
//
//	fmt.Println("===== OpenIMUpdateUsersParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}

//func TestOpenIMGetUsersParam(t *testing.T) {
//	var p = OpenIMGetUsersParam{}
//	p.UserIds = []string{"56e279cfc77b930ae86b52e7", "55ff574b4ea0461440000003"}
//	fmt.Println("===== OpenIMDeleteUserParam =====")
//	fmt.Println(RequestWithKey("23201003", "1e2dfd16981d75142597fd10131b17b5", p))
//}

//func TestOpenIMDeleteUsersParam(t *testing.T) {
//	var p = OpenIMDeleteUsersParam{}
//	p.UserIds = []string{"56e279cfc77b930ae86b52e7", "55ff574b4ea0461440000003"}
//	fmt.Println("===== OpenIMDeleteUserParam =====")
//	fmt.Println(RequestWithKey("23201003", "1e2dfd16981d75142597fd10131b17b5", p))
//}

//func TestOpenIMPushCustomMsgParam(t *testing.T) {
//	var p = params.OpenIMPushCustomMsgParam{}
//	p.FromUser = "admin2"
//	p.ToUsers = []string{"admin"}
//	p.Summary = "推送内容"
//	p.Data = "消息内容"
//	p.SetApsAlert("aa")
//	p.SetApsBadge(3)
//
//	fmt.Println("===== OpenIMPushCustomMsgParam =====")
//	fmt.Println(taobao.RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}

//func TestOpenIMPushMsgParam(t *testing.T) {
//	var p = OpenIMPushMsgParam{}
//	p.FromUser = "55ff574b4ea0461440000003"
//	p.ToUsers = []string{"56e279cfc77b930ae86b52e7"}
//	p.Context = "我是不是应该说点什么啊,继续吹.....劳而无功sdsdfsf"
//
//	fmt.Println("===== TestOpenIMPushMsgParam =====")
//	fmt.Println(RequestWithKey("23201003", "1e2dfd16981d75142597fd10131b17b5", p))
//}

//func TestOpenIMCreateTribeParam(t *testing.T) {
//	var p = OpenIMCreateTribeParam{}
//
//	var u = &OpenIMUser{}
//	u.AppKey = "23274732"
//	u.TaoBaoAccount = false
//	u.UserId = "55ff574b4ea0461440000003"
//	p.User = u
//
//	p.TribeName = "test trib"
//	p.Notice = "open la"
//	p.TribeType = "0"
//
//	fmt.Println("===== TestOpenIMCreateTribeParam =====")
//	fmt.Println(RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}

//func TestOpenIMGetTribeInfoParam(t *testing.T) {
//	var p = OpenIMGetTribeInfoParam{}
//	var u = &OpenIMUser{}
//	u.AppKey = "23274732"
//	u.TaoBaoAccount = false
//	u.UserId = "55ff574b4ea0461440000003"
//	p.User = u
//	p.TribeId = "1675657139"
//
//	fmt.Println("===== TestOpenIMGetTribeInfoParam =====")
//	fmt.Println(RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}