package service

import (
	"math/rand"
	"time"

	"github.com/HAL-RO-Developer/caseTeamB/model"
)

var RegistInfo []deviceInfo

type deviceInfo struct {
	Name    string
	ChildId int
	GoalId  string
	Pin     string
}

// デバイス情報新規登録
func CreateDevice(name string, goalId string) (string, bool) {
	var info deviceInfo
	goalData, find := GetOneGoal(goalId)
	if !find {
		return "目標が見つかりません。", false
	}

	info.Name = name
	info.ChildId = goalData.ChildId
	info.GoalId = goalId
	info.Pin = createPin()
	RegistInfo = append(RegistInfo, info)

	return info.Pin, true
}

// macAddr&デバイスID登録
func RegistrationDevice(pin string) (string, bool) {
	buf, find := GetPin(pin)
	if !find {
		return "pinが見つかりませんでした。", false
	}
	RegistInfo = PinRemove(pin)

	deviceId := CreateDeviceId()

	device := model.Device{
		Name:     buf.Name,
		ChildId:  buf.ChildId,
		DeviceId: deviceId,
	}

	err := db.Create(&device).Error
	if err != nil {
		return "デバイスIDが登録できませんでした。", false
	}

	err = UpdateGoal(buf.GoalId, deviceId)
	if err != nil {
		return "デバイスIDが登録に失敗しました。", false
	}
	return deviceId, true
}

// デバイスID作成
func CreateDeviceId() string {
	var deviceId string
	for {
		deviceId = createUuid(12, []rune("ABCDEFGHRJKLNMOPQRSTUPWXYZabcdefghijklmnopqrstuvwxyz0123456789"))
		_, find := GetDeviceInfoFromDeviceId(deviceId)
		if !find {
			break
		}
	}
	return deviceId
}

// ランダム文字列作成
func createUuid(length int, letters []rune) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ピン作成
func createPin() string {
	var pin string
	for {
		pin = createUuid(4, []rune("0123456789"))
		_, find := GetPin(pin)
		if !find {
			break
		}
	}
	return pin
}

// データベースからデバイス情報取得
func GetDeviceInfoFromDeviceId(deviceId string) ([]model.Device, bool) {
	var devices []model.Device
	db.Where("device_id = ?", deviceId).Find(&devices)
	return devices, len(devices) != 0
}

// 配列からPin検索
func GetPin(pin string) (deviceInfo, bool) {
	for _, v := range RegistInfo {
		if v.Pin == pin {
			return v, true
		}
	}
	return deviceInfo{}, false
}

func ExisByMac(mac string) bool {
	var devices []model.Device
	db.Where("mac = ?", mac).Find(&devices)
	return len(devices) != 0
}

// データベースからデバイステーブルの情報取得(ユーザー名から)
func GetDeviceId(name string) ([]model.Device, bool) {
	var devices []model.Device
	db.Where("name = ?", name).Find(&devices)
	return devices, len(devices) != 0
}

// デバイス情報取得(ユーザー名と子どもIDから)
func GetDeviceIdFromChildId(name string, childId int) ([]model.Device, bool) {
	var devices []model.Device
	db.Where("name = ? and child_id = ?", name, childId).Find(&devices)
	return devices, len(devices) != 0
}

// 指定されたデバイスIDの削除
func DeleteDeviceId(name string, buttonId string) bool {
	var devices model.Device
	db.Where("name = ? and device_id = ?", name, buttonId).First(&devices)
	if devices.DeviceId == "" {
		return false
	}

	db.Delete(devices)
	return true
}

// 子どものデバイスIDの削除
func DeleteDeviceIdFromChild(name string, childId int) bool {
	var devices model.Device
	db.Where("name = ? and child_id = ?", name, childId).First(&devices)
	if devices.DeviceId == "" {
		return false
	}

	db.Delete(devices)
	return true
}

// 1ユーザーの最初のボタンIDの削除
func DeleteButtonFirst(name string) bool {
	var devices model.Device
	err := db.Where("name = ?", name).First(&devices).Error
	if err != nil {
		return false
	}
	db.Delete(devices)
	return true
}

// スライスの中身削除
func PinRemove(pin string) []deviceInfo {
	result := []deviceInfo{}
	for _, v := range RegistInfo {
		if v.Pin != pin {
			result = append(result, v)
		}
	}
	return result
}
