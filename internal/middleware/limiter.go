package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/ratelimit"
	"time"
)

var prevRequestTime = time.Now()

type Limiter struct {
	RL ratelimit.Limiter
}

// 限流
func (l *Limiter) Take() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ips := ctx.IPs()
		var ip string
		if len(ips) > 0 && ips[0] != "" {
			ip = ips[len(ips)-1]
		} else {
			ip = ctx.IP()
		}
		//global.Logger.Info("Request IP:", ip)
		_ = len(ip)
		l.SetPrevTime(l.RL.Take())
		ctx.Next()
		return nil
	}
}

func (l *Limiter) GetPrevTime() time.Time {
	return prevRequestTime
}

func (l *Limiter) SetPrevTime(time2 time.Time) {
	prevRequestTime = time2
}
