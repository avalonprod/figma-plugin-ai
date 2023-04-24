package app

import (
	"context"
	"fmt"
	"gasstrem-api-ai/internal/adapters/openai"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func CorsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	openai := openai.NewOpenaiService(os.Getenv("OPENAI_API_KEY"))

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		CorsMiddleware,
	)

	router.POST("/img-generation/message", func(c *gin.Context) {
		var message MessageResponse

		if err := c.BindJSON(&message); err != nil {
			fmt.Println(err)
		}

		res, err := openai.ImgGenerator.NewImg(message.Message)
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, res)

	})
	router.POST("/text-generation/message", func(c *gin.Context) {
		var message MessageResponse

		if err := c.BindJSON(&message); err != nil {
			fmt.Println(err)
		}

		res, err := openai.TextGenerator.NewMessages(message.Message)
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, MessageResponse{
			Message: res,
		})

	})
	cfg := Config{
		HTTP: HTTPConfig{
			Port:           "3030",
			ReadTimeout:    40 * time.Second,
			WriteTimeout:   40 * time.Second,
			MaxHeaderBytes: 10,
		},
	}
	srv := NewServer(cfg, router)
	err = srv.Run()

	if err != nil {
		fmt.Println(err)
	}
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderBytes << 25,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
