package controller

import (
	"github.com/HAL-RO-Developer/caseTeamB/controller/response"
	"github.com/HAL-RO-Developer/caseTeamB/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamB/service"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
	"fmt"
)

var Button = buttonimpl{}

type buttonimpl struct {
}

// ボタン押下回数変更
func (b *buttonimpl) DeviceIncrement(c *gin.Context) {
	var msg string
	req, ok := validation.ButtonCheck(c)
	if !ok {
		return
	}

	bocco, find := service.GetDeviceInfoFromDeviceId(req.DeviceId)
	if find {
		data, find := service.GetGoalFromDeviceId(req.DeviceId)
		if !find {
			response.Json(gin.H{"angle": "目標が見つかりません。"}, c)
			return
		}

		success := service.ApprovalGoal(data[0].GoalId, 1)
		if !success {
			response.BadRequest(gin.H{"error": "押下回数を追加できませんでした"}, c)
			return
		}
		// 押下回数追加後データ取得
		data, _ = service.GetGoalFromDeviceId(req.DeviceId)
		// サーボモーターの移動角度計算
		progress := (float64(data[0].Run) / float64(data[0].Criteria)) * 180
		if progress > 180 {
			progress = 180
		}
		message, find := service.GetMessageFromGoal(data[0].GoalId)
		if !find {
			deviceInfo,_ := service.GetDeviceInfoFromDeviceId(req.DeviceId)
			childInfo,_ := service.GetOneChildInfo(deviceInfo[0].Name, deviceInfo[0].ChildId)
			goalInfo, _ := service.GetGoalFromDeviceId(req.DeviceId)
			defMsg,_ := service.GetDefaultMessage(childInfo[0].BirthDay, goalInfo[0].Run, goalInfo[0].Criteria,goalInfo[0].Deadline)
			
			switch defMsg.MsgCondition {
			case 2:
			case 3:
				msg = fmt.Sprintf(defMsg.Message ,childInfo[0].NickName)
			case 4:
				msg = fmt.Sprintf(defMsg.Message ,childInfo[0].NickName, goalInfo[0].Criteria - goalInfo[0].Run)
			case 5:
				msg = fmt.Sprintf(defMsg.Message ,childInfo[0].NickName, goalInfo[0].Run)
			default:
				msg = defMsg.Message
			}
			talkBocco(msg, deviceInfo[0].Name)
			response.Json(gin.H{"angle": int(progress)}, c)
			return
		}
		for i := 0; i < len(message); i++ {
			// 目標の実行回数がメッセージの発信条件を満たした時
			if data[0].Run == message[i].MessageCall {
				talkBocco(message[i].Message, bocco[0].Name)
				break
			}
		}

		response.Json(gin.H{"angle": int(progress)}, c)
		return
	}
	response.Json(gin.H{"angle": "デバイスIDが見つかりません。"}, c)
}

func talkBocco(message string, name string) {
	boccoInfo, find := service.ExisByBoccoAPI(name)
	if !find {
		return
	}
	boccoToken, _ := service.GetBoccoToken(boccoInfo[0].Email, service.APIKEY, boccoInfo[0].Pass)
	roomId, _ := service.GetRoomId(boccoToken)
	uuid := uuid.Must(uuid.NewV4()).String()
	service.SendMessage(uuid, roomId, boccoToken, message)
}
