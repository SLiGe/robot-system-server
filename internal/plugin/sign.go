package plugin

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/service"
)

type SignInPlugin struct {
	signInService service.SignInService
}

func NewSignInPlugin(signInService service.SignInService) SignInPlugin {
	return SignInPlugin{signInService: signInService}
}

func (s *SignInPlugin) Code() string {
	return "SignIn"
}

func (s *SignInPlugin) Do() {
	zero.OnFullMatch("签到").Handle(func(ctx *zero.Ctx) {
		s.signInService.AddSignInPoints(v1.AddSignPointsRequest{Points: 30, QQ: "357078415"})
		ctx.Send("签到成功")
	})
}
