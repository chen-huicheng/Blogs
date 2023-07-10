package main

import (
	"context"
	"fmt"
)

// TimerTask 触发事件
type TimerTask struct {
	Task  string
	Param map[string]string
}

// 事件类型
const (
	Remind         = "remind"
	AppointCreated = "appoint_created"
)

/*
 * ISender 接口
 */

type ISender interface {
	GetUsers(ctx context.Context) ([]string, error)
	SendMsg(ctx context.Context, openID string) error
	Name() string
}

// Handler 创建发送消息的具体类
func Handler(ctx context.Context, e TimerTask) error {
	sender, _ := CreateSender(e.Task)
	return Service(ctx, sender)
}

// Service 使用 ISender 发送消息
func Service(ctx context.Context, sender ISender) error {
	openIDs, _ := sender.GetUsers(ctx)
	name := sender.Name()
	var failIDs []string
	var successIDs []string
	var noNeedSend []string
	for _, openID := range openIDs {
		if hasSend(name, openID) {
			noNeedSend = append(noNeedSend, openID)
			continue
		}
		err := sender.SendMsg(ctx, openID)
		if err != nil {
			failIDs = append(failIDs, openID)
			continue
		}
		successIDs = append(successIDs, openID)
		markSend(name, openID)
	}
	// 记录发送结果
	fmt.Printf("发送失败的 OpenID:%+v", failIDs)
	fmt.Printf("发送成功的 OpenID:%+v", successIDs)
	fmt.Printf("已发送无需发送的 OpenID:%+v", noNeedSend)
	if len(failIDs) != 0 {
		return fmt.Errorf("%s消息发送失败", name)
	}
	return nil
}

// CreateSender 根据 event.Type 创建对应的具体类
func CreateSender(t string) (ISender, error) {
	var sender ISender
	switch t {
	case Remind:
		sender = &SenderRemind{}
	case AppointCreated:
		sender = &SenderRemind{}
	default:
		return nil, fmt.Errorf("没有类型[%s]接口", t)
	}
	return sender, nil
}

/*
 * SenderRemind 实现 ISender 接口
 */
type User struct {
	OpenID string
	// ...
}

// RemindMsg 包含提醒消息所需字段
type RemindMsg struct {
	OpenID string
	// ...
}
type SenderRemind struct {
	UserMap map[string]RemindMsg
}

func (s *SenderRemind) GetUsers(ctx context.Context) (map[string]msg, error) {
	// RPC get users who need to send messages
	users := []User{}
	openIDs := []string{}
	// RPC get RemindMsg by user
	s.UserMap = make(map[string]RemindMsg)
	for _, u := range users {
		s.UserMap[u.OpenID] = RemindMsg{}
		openIDs = append(openIDs, u.OpenID)
	}
	if (){
		msg:=s.shortMsg()

	}
	
	return openIDs, nil
}
func (s *SenderRemind) SendMsg(ctx context.Context, openID string) error {
	// SendShortMsg

	// SendFeishuMsg

	// SendEmailMsg
	return nil
}

func (s *SenderRemind) Name() string {
	return Remind
}

func hasSend(t string, openID string) bool {
	return true
}
func markSend(t string, openID string) {
	return
}

func main() {

}

/* 定时消息类型
预约前一日提醒  问卷测评二次提醒消息  线上咨询预约二次提醒  线下检查预约二次提醒 报告解读预约二次提醒 日程提前十五分钟提醒

抽象：给【特定人】发送【特定类型】消息。 消息包含【短信】、【飞书】、【消息队列】。

接口：type ISender interface {
		GetUsers(ctx context.Context) ([]string, error) 
		handleMsg
		handleFeishu

		sendMsg()
	 }

	根据特定类型提醒实现 ISender 接口。

	var as ASender


	openIDs:=a.GetUsers()
	for openID in openIDs{
		a.SendMsg(openID)
	}
*/
