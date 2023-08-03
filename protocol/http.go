package protocol

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"restful-api-demo/apps"
	"restful-api-demo/conf"
	"time"
)

type HTTPService struct {
	server *http.Server
	l      logger.Logger
	c      *conf.Config
	r      *gin.Engine
}

func NewHttpService() *HTTPService {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.C().App.HttpAddr(),
		Handler:           r,
	}
	return &HTTPService{
		r:      r,
		server: server,
		l:      zap.L().Named("HTTP Service"),
		c:      conf.C(),
	}
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由, 注册给gin
	apps.InitGin(s.r)

	s.l.Infof("成功加载的gin apps: %s", apps.LoadedGinApps())
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	// 启动 HTTP服务
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed { //如果是 正常监听到关闭信号
			s.l.Info("HTTP Service is stopped")
		} else {
			s.l.Errorf("start service error, %v", err.Error())
			return err
		}
		return nil
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
