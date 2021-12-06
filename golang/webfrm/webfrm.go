package webfrm

import (
	bloglog "blog/log"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

//ServerInstanceCfg 配置服务中保存的服务实例的信息
type ServerInstanceCfg struct {
	InstanceName string                   `json:"instance_name"`
	Kind         string                   `json:"kind"`
	SDHosts      []string                 `json:"sd_hosts"`
	Services     []string                 `json:"services"`
	TLS          bool                     `json:"tls"` //tls目前先使用false，certFile 和keyfile 传的是路径，需要整理一个方案出来
	HostName     string                   `json:"host_name"`
	Port         int                      `json:"port"`
	IsDebug      bool                     `json:"is_debug"`
	SupportSR    bool                     `json:"support_sr"`   //是否支持服务注册
	SupportCors  bool                     `json:"support_cors"` //是否支持跨域请求
	LocalEnable  bool                     `json:"local_enable"` //本地配置是否直接作为服务配置，true，不会去config服务获取
	LoggerConfig bloglog.FileLoggerConfig `json:"log"`
}

// Webfrm
type Webfrm struct {
	Router *gin.Engine
	Logger bloglog.Logger
	Cfg    *ServerInstanceCfg
	Server *http.Server
}

func NewWebfrm(srvKind string, logger *bloglog.Logger) (*Webfrm, error) {
	serverInsCfg, err := loadServerInstanceConfig(srvKind+"/serverinstancecfg.json", logger)
	if nil != err {
		(*logger).ErrErr("Failed at LoadServerInstanceConfig", err)
		return nil, err
	}
	fileLogger, err := newServerLogger(serverInsCfg)
	if nil != err {
		(*logger).ErrErr("Failed at newServerLogger", err)
		return nil, err
	}
	if serverInsCfg.IsDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	webFrm := &Webfrm{
		Router: gin.Default(),
		Cfg:    serverInsCfg,
		Logger: fileLogger,
	}

	if serverInsCfg.SupportCors {
		webFrm.Router.Use(Cors())
	}
	return webFrm, nil
}

//Cors 路由跨域支持
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT,OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-CAXA-Auth, X-CAXA-Lang")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func (wf *Webfrm) Start(logger *bloglog.Logger) error {
	wf.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", wf.Cfg.Port),
		Handler: wf.Router,
	}

	go func(srv *http.Server, tls bool, logger *bloglog.Logger) {
		if tls {
			path, _ := GetCurrentDirectory()
			path += "/../cfg/pki/"
			if err := srv.ListenAndServeTLS(path+"server.crt", path+"server.key"); nil != err && !errors.Is(err, http.ErrServerClosed) {
				(*logger).ErrErr("Listen and server Tls", err)
				return
			}
		} else {
			if err := srv.ListenAndServe(); nil != err && !errors.Is(err, http.ErrServerClosed) {
				(*logger).ErrErr("Listen and server", err)
				return
			}
		}
	}(wf.Server, wf.Cfg.TLS, logger)

	(*logger).Info("Server is running")
	return nil
}

//WaitForExit 等待服务退出 和svcfrm 类似，调用一个即可
func (wf *Webfrm) WaitForExit(logger *bloglog.Logger) {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	(*logger).Info("Shutting down server...")
}

//Stop 停止Web服务器
func (wf *Webfrm) Stop(ctx context.Context, logger *bloglog.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := wf.Server.Shutdown(ctx); err != nil {
		(*logger).ErrErr("Server forced to shutdown", err)
	}
}

//GetCurrentDirectory 获取服务所在文件夹目录
func GetCurrentDirectory() (string, error) {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1), nil
}

//loadServerInstanceConfig 加载本地的服务实例信息
func loadServerInstanceConfig(cfgPath string, logger *bloglog.Logger) (*ServerInstanceCfg, error) {
	path, _ := GetCurrentDirectory()
	path += "/../cfg/"
	path += cfgPath
	data, err := ioutil.ReadFile(path)
	if err != nil {
		(*logger).ErrErr("ioutil.ReadFile failed >>"+cfgPath, err)
		return nil, err
	}
	var serverCfg ServerInstanceCfg
	err = json.Unmarshal(data, &serverCfg)
	if err != nil {
		(*logger).ErrErr("json Unmarshal failed >>"+string(data), err)
		return nil, err
	}
	return &serverCfg, nil
}

func newServerLogger(cfg *ServerInstanceCfg) (bloglog.Logger, error) {
	var logCfg bloglog.FileLoggerConfig
	logCfg.Console = false
	logCfg.ErrorFileName = cfg.Kind + "_err.log"
	logCfg.InfoFileName = cfg.Kind + "_info.log"
	path, _ := GetCurrentDirectory()
	path += "/../journal"
	logCfg.LogPath = path
	logCfg.MaxAge = cfg.LoggerConfig.MaxAge
	logCfg.MaxBackups = cfg.LoggerConfig.MaxBackups
	logCfg.MaxSize = cfg.LoggerConfig.MaxSize
	logger, err := bloglog.NewFileLogger(&logCfg)
	if err != nil {
		return nil, fmt.Errorf("new file logger : %w", err)
	}
	return logger, nil
}
