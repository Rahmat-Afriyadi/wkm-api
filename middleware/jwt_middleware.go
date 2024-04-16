package middleware

import (
	"database/sql"
	"strings"
	"wkm/config"
	"wkm/repository"
	"wkm/utils"

	"github.com/gofiber/fiber/v2"
)

func DeserializeUser(c *fiber.Ctx) error {
	var conn *sql.DB = config.GetConnection()
	defer conn.Close()
	userRepository := repository.NewUserRepository(conn)
	var access_token string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		access_token = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		access_token = c.Cookies("access_token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenClaims, err := utils.ValidateToken(access_token, "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDTkJEMGg0T0RlU3dxNFpwSDJqQVVXeURVdQpXTWxvVmVPdFJCbzI5b1NxUUlrNUpKUlo0ekpGR2dpNGtXeGpiWXpmNmRDMkViSHRSK202MjN0bFdOYkt1SXJJClhCdk1OdW1LREx4b2o3dU5VRVZyUVVFWVZhOVNuTlMxbDNVOEtRUW0xWGlhTGplSGFrN0M4WGYvcFBBTWpYc3AKd1A0TWU2S2lUTVFxTkRITjNRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==")
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	user := userRepository.FindById(tokenClaims.UserID)

	c.Locals("user", user)

	return c.Next()
}
