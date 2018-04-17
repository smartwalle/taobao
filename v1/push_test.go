package taobao

//import (
//	"fmt"
//	"testing"
//)
//
//func TestCloudPushNoticeToIOSParam(t *testing.T) {
//	var p = CloudPushNoticeToIOSParam{}
//	p.Summary = "golang"
//	p.Target = "all"
//	p.Environment = "DEV"
//	p.TargetValue = "e"
//	p.AddParam("badge", 3)
//	p.AddParam("sound", "v")
//	p.AddParam("aaa", "eee")
//
//	fmt.Println("===== TestCloudPushNoticeToIOSParam =====")
//	fmt.Println(RequestWithKey("23201003", "1e2dfd16981d75142597fd10131b17b5", p))
//}
//
//func TestCloudPushNoticeToAndroidParam(t *testing.T) {
//	var p = CloudPushNoticeToAndroidParam{}
//	p.Title = "ttt"
//	p.Summary = "aaaa"
//	p.Target = "all"
//	p.TargetValue = "e"
//
//	fmt.Println("===== TestCloudPushNoticeToAndroidParam =====")
//	fmt.Println(RequestWithKey("23201003", "1e2dfd16981d75142597fd10131b17b5", p))
//}

//func TestCloudPushMessageToIOSParam(t *testing.T) {
//	var p = CloudPushMessageToIOSParam{}
//	p.Body = "中文"
//	p.Target = "all"
//	p.TargetValue = "e"
//
//	fmt.Println("===== TestCloudPushMessageToIOSParam =====")
//	fmt.Println(RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}

//func TestCloudPushParam(t *testing.T) {
//	var p = CloudPushParam{}
//	p.Title = "标题"
//	p.Body = "hiaa"
//	p.Summery = "summery"
//	p.DeviceType = 0
//	p.Target = "all"
//	p.TargetValue = "e"
//	p.IOSBadge = "1"
//	p.Type = 1
//	p.Remind = true
//	p.StoreOffline = false
//	p.Timeout = 3
//	p.AddiOSExtParam("k", "v")
//	fmt.Println("===== TestCloudPushParam =====")
//	fmt.Println(RequestWithKey("23274732", "8eed038bab11210aa8099d8e450cd631", p))
//}