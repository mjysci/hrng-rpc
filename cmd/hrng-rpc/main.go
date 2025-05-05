package main

import (
	"crypto/fips140"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/mjysci/hrng-rpc/internal/config"
	"github.com/mjysci/hrng-rpc/internal/service"
	"github.com/mjysci/hrng-rpc/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	randomGenerator := utils.NewRandomGenerator()

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	randomService := service.NewRandomService(randomGenerator, cfg)
	if err := s.RegisterService(randomService, ""); err != nil {
		log.Fatalf("Error registering service: %v", err)
	}

	r := gin.Default()

	trustedProxies := cfg.Server.TrustedProxies
	if len(trustedProxies) == 0 {
		trustedProxies = []string{"127.0.0.1", "::1"}
	}
	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	r.POST("/json-rpc", func(c *gin.Context) {
		s.ServeHTTP(c.Writer, c.Request)
	})

	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)

	// Check hardware RNG availability on Linux
	if _, err := os.Stat("/sys/class/misc/hw_random/rng_current"); err == nil {
		content, err := os.ReadFile("/sys/class/misc/hw_random/rng_current")
		if err == nil && strings.TrimSpace(string(content)) == "none" {
			log.Println("WARNING: Hardware RNG is not available. Random numbers will not be hardware-generated.")
		}
	}

	if !fips140.Enabled() {
		log.Println("WARNING: FIPS 140-3 mode is NOT enabled")
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
