package service

import (
	"github.com/HAL-RO-Developer/caseTeamB/model"
	"math/rand"
	"time"
)

type Condition struct {
	Birthday time.Time
	Run      int
	Criteria int
	Limit    *time.Time
}

const RUNGOAl = 5 // 目標達成数の条件
const REST = 0.7

// デフォルトメッセージの取得
func GetDefaultMessage(birthday time.Time, run int, criteria int, limit *time.Time) (model.DefaultMessage, bool) {
	var messages []model.DefaultMessage
	var find bool
	var condition Condition
	var useCondition []int

	condition.Birthday = birthday
	condition.Run = run
	condition.Criteria = criteria
	condition.Limit = limit

	// 満たしている条件を取得
	useCondition = getCondition(condition)
	if len(useCondition) != 0 {
		messages, find = getMessage(useCondition[getRandomNo(len(useCondition))])
		if !find {
			return model.DefaultMessage{}, false
		}
		return messages[len(messages)], true
	}
	return model.DefaultMessage{}, false
}

func getMessage(condition int) ([]model.DefaultMessage, bool) {
	var message []model.DefaultMessage
	db.Where("msg_condition = ?", condition).Find(&message)
	return message, len(message) != 0
}

func getRandomNo(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

func getCondition(condition Condition) []int {
	var useCondition []int

	// 目標達成時
	if condition.Run == condition.Criteria {
		if condition.Limit != nil {
			// 達成期限
			if condition.Limit.Before(time.Now()) {
				useCondition = append(useCondition, 1) // 期限切れ
			} else {
				useCondition = append(useCondition, 2) // 期限内
			}
		}
	} else {
		per := float64(condition.Run) / float64(condition.Criteria)
		if per > REST && per < 1.0 {
			useCondition = append(useCondition, 3) // 目標の進捗が7割以上
			useCondition = append(useCondition, 4) // 目標を進捗が7割以上(残り回数を喋る)
		}
		if condition.Run > RUNGOAl {
			useCondition = append(useCondition, 5) // 目標を5回以上実行
		}
		useCondition = append(useCondition, 6) // 条件なし
	}
	return useCondition
}
