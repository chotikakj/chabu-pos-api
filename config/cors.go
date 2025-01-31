package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var CorsConfigDefault = cors.Config{
	Next: nil,
	// AllowOrigins: "*",
	AllowOrigins:     "http://localhost:5173, http://127.0.0.1:5173, http://127.0.0.1:4173, http://127.0.0.1, http://192.168.1.35:5173, https://iamkkltto.onrender.com, http://localhost:3000, http://127.0.0.1:3000, https://zlfn8k1s-5173.asse.devtunnels.ms, https://chabu-pos.onrender.com",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           30000,
}
