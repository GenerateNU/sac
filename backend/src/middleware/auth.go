package middleware

import (
	"backend/src/models"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type JWTTemplate struct {
	Role string `json:"role"`
	Id string `json:"id"`
	AZP string `json:"azp"`
	EXP int `json:"exp"`
	IAT int `json:"iat"`
	ISS string `json:"iss"`
	JTI	string `json:"jti"`
	NBF int `json:"nbf"`
	SUB string `json:"sub"`
}

// A middleware function that checks if the user is authenticated
func AuthenticationMiddleware(c *fiber.Ctx) error {
	
	sessionToken := c.Get("Authorization")[7:]
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

	token, err := parseToken(sessionToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Error parsing token")
	}

	if token.Valid {
		return c.Next()
	}

	return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
}

func authorize(requiredRole models.UserRole, requiredPermissions []string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Extract user information from the request
        user := getUserFromToken(c.Get("Authorization"))


        // Continue to the next middleware or route handler
        return c.Next()
    }
}


func getUserFromToken(tokenString string) *JWT {
	token, err := parseToken(tokenString)
	if err != nil {
		fmt.Println("error parsing token")
	}


	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["role"])
		return &JWT{
			Role: claims["role"].(string),
			Id: claims["id"].(string),
			AZP: claims["azp"].(string),
			EXP: claims["exp"].(int),
			IAT: claims["iat"].(int),
			ISS: claims["iss"].(string),
			JTI: claims["jti"].(string),
			NBF: claims["nbf"].(int),
			SUB: claims["sub"].(string),
		}
		
	} else {
		fmt.Println("error getting claims")
	}

	return nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("sk_test_YOzhvkCIDAK2IcUi5K24naDQh3RoHlTW9xGXzaThNm"), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}