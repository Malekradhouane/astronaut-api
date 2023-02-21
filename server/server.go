package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ory/viper"
	"github/malekradhouane/test-cdi/route"
)

//Server is the http layer of the app
type Server struct {
	Cfg         *Config
	AstronautActions *route.AstronautActions
}

//Config server config
type Config struct {
	Port string
}

//NewConfig constructs a new config
func NewConfig() *Config {
	viper.BindEnv("PORT", "PORT")
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		Port: port,
	}
}

func NewServer(cfg *Config,
	aa *route.AstronautActions,
) *Server {
	return &Server{Cfg: cfg,
		AstronautActions: aa,
	}
}

//Run setup the app with all dependencies
func (s *Server) Run() error {
	r := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AddAllowHeaders("Authorization")
	corsCfg.AllowAllOrigins = true
	r.Use(cors.New(corsCfg))
	r.GET("/status", route.GetStatus)

	users := r.Group("/astronauts")
	{
		users.POST("", s.AstronautActions.Create)
		users.GET("", s.AstronautActions.List)
		users.GET("/:id", s.AstronautActions.Get)
		users.PATCH("/:id", s.AstronautActions.Update)
		users.DELETE("/:id", s.AstronautActions.Delete)
	}

	return r.Run(":" + s.Cfg.Port)
}

//Shutdown shutdowns the server
func (s *Server) Shutdown(ctx context.Context) {
	//Clean up here
}
