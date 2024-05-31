package server

import (
	"context"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"
	"robot-system-server/internal/plugin"
	_ "robot-system-server/internal/plugin"
	"robot-system-server/pkg/log"
)

type QQBotServer struct {
	log *log.Logger
}

var op *zero.Config

func NewQQBotServer(logger *log.Logger, signIn plugin.SignInPlugin) *QQBotServer {

	op = &zero.Config{
		NickName:      []string{"bot"},
		CommandPrefix: "/",
		SuperUsers:    []int64{357078415},
		Driver: []zero.Driver{
			// 正向 WS
			driver.NewWebSocketClient("ws://172.16.0.107:3001", ""),
			driver.NewWebSocketClient("ws://172.16.0.107:6701", ""),
		},
	}
	signIn.Do()
	return &QQBotServer{
		log: logger,
	}
}

func (j *QQBotServer) Start(ctx context.Context) error {
	j.log.Info("!!!QQ Bot Start Success!!!")
	zero.Run(op)
	return nil
}
func (j *QQBotServer) Stop(ctx context.Context) error {
	return nil
}
