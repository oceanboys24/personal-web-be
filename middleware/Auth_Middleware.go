package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)



func CheckToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error" : "Token Not Found",
			})
	}

	// Split Token
	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Invalid Authorization header Format",
		})
	}
	tokenStr := splitToken[1]
	jwtSecret := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(jwtSecret) , nil
	})
	if err != nil || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Expired Token",
		})	
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		ctx.Locals("user_id", claims["sub"])
		ctx.Locals("email", claims["email"]) 
	}

	return ctx.Next()
}