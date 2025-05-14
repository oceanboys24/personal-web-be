package visitor

import (
	"context"
	"fmt"
	"log"
	"personal-web-be/config"
	"time"

	"github.com/gofiber/fiber/v2"
)


var c = context.Background()


func CountVisitorHandler(ctx *fiber.Ctx) error  {
	ip := ctx.IP()
	var key = "visitor:" + ip
	count , err := config.RedisClient.Incr(c, key).Result()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if count == 1 {
		err = config.RedisClient.Expire(c, key, 24*time.Hour).Err()
		if err != nil {
			log.Println("Fail set EXPIRE for visitor_count:", err)
		}
	}
	return ctx.SendString(fmt.Sprintf("you're visitor %d today", count))
}

func GetVisitorHandler(ctx *fiber.Ctx) error  {
	ip := ctx.IP() 
	key := "visitor:" + ip 

	count, err := config.RedisClient.Get(c, key).Int()
	if err != nil {
		if err.Error() == "redis: nil" {
			count = 0
		} else {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return ctx.Status(200).JSON(fiber.Map{
		"ip":    ip,
		"count": count,
	})
}