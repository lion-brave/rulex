package httpserver

import (
	"context"
	"net/http"
	"rulex/statistics"
	"rulex/x"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ngaut/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/square/go-jose.v2/json"
	"gorm.io/gorm"
)

const API_ROOT string = "/api/v1/"
const DASHBOARD_ROOT string = "/dashboard/v1/"

type HttpApiServer struct {
	Port       int
	Root       string
	sqliteDb   *gorm.DB
	ginEngine  *gin.Engine
	ruleEngine *x.RuleEngine
}

func NewHttpApiServer(port int, root string, e *x.RuleEngine) *HttpApiServer {
	return &HttpApiServer{Port: port, Root: root, ruleEngine: e}
}
func (hh *HttpApiServer) Load() *x.XPluginEnv {
	return x.NewXPluginEnv()
}

//
func (hh *HttpApiServer) Init(env *x.XPluginEnv) error {
	gin.SetMode(gin.ReleaseMode)
	hh.ginEngine = gin.New()
	hh.ginEngine.Use(Authorize())
	hh.InitDb()
	hh.ginEngine.LoadHTMLGlob(hh.Root)
	ctx := context.Background()
	go func(ctx context.Context, port int) {
		hh.ginEngine.Run(":" + strconv.Itoa(port))
	}(ctx, hh.Port)
	return nil
}
func (hh *HttpApiServer) Install(env *x.XPluginEnv) (*x.XPluginMetaInfo, error) {
	return &x.XPluginMetaInfo{
		Name:     "HttpApiServer",
		Version:  "0.0.1",
		Homepage: "www.ezlinker.cn",
		HelpLink: "www.ezlinker.cn",
		Author:   "wwhai",
		Email:    "cnwwhai@gmail.com",
		License:  "MIT",
	}, nil
}

//
//
func (hh *HttpApiServer) Start(env *x.XPluginEnv) error {
	hh.ginEngine.GET(DASHBOARD_ROOT, func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{})
	})
	hh.ginEngine.GET(API_ROOT+"plugins", func(c *gin.Context) {
		cros(c)
		c.PureJSON(http.StatusOK, gin.H{
			"plugins": hh.ruleEngine.GetPlugins(),
		})
	})
	hh.ginEngine.GET(API_ROOT+"system", func(c *gin.Context) {
		cros(c)
		//
		percent, _ := cpu.Percent(time.Second, false)
		memInfo, _ := mem.VirtualMemory()
		parts, _ := disk.Partitions(true)
		diskInfo, _ := disk.Usage(parts[0].Mountpoint)
		c.JSON(http.StatusOK, gin.H{
			"diskInfo":   diskInfo.UsedPercent,
			"memInfo":    memInfo.UsedPercent,
			"cpuPercent": percent[0],
			"os":         runtime.GOOS,
			"arch":       runtime.GOARCH,
			"cpus":       runtime.GOMAXPROCS(0)})
	})
	//
	hh.ginEngine.GET(API_ROOT+"inends", func(c *gin.Context) {
		cros(c)
		c.JSON(http.StatusOK, gin.H{"inends": hh.ruleEngine.AllInEnd()})
	})
	//
	hh.ginEngine.GET(API_ROOT+"outends", func(c *gin.Context) {
		cros(c)
		c.JSON(http.StatusOK, gin.H{"outends": hh.ruleEngine.AllOutEnd()})
	})
	//
	hh.ginEngine.GET(API_ROOT+"rules", func(c *gin.Context) {
		cros(c)
		c.JSON(http.StatusOK, gin.H{"rules": hh.ruleEngine.AllRule()})
	})
	//
	hh.ginEngine.GET(API_ROOT+"statistics", func(c *gin.Context) {
		cros(c)
		c.JSON(http.StatusOK, gin.H{"statistics": statistics.AllStatistics()})
	})
	//
	// Create InEnd
	//
	hh.ginEngine.POST(API_ROOT+"inends", func(c *gin.Context) {
		cros(c)
		type Form struct {
			Type        string                 `json:"type" binding:"required"`
			Name        string                 `json:"name" binding:"required"`
			Description string                 `json:"description"`
			Config      map[string]interface{} `json:"config" binding:"required"`
		}
		form := Form{}
		err0 := c.ShouldBindJSON(&form)
		if err0 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err0.Error()})
		} else {
			configJson, err1 := json.Marshal(form.Config)
			if err1 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err1.Error()})
			} else {
				uuid := x.MakeUUID("INEND")
				hh.InsertMInEnd(&MInEnd{
					UUID:        uuid,
					Type:        form.Type,
					Name:        form.Name,
					Description: form.Description,
					Config:      string(configJson),
				})
				err := hh.LoadNewestInEnd(uuid)
				if err != nil {

					c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				} else {
					c.JSON(http.StatusOK, gin.H{"msg": "create success"})
				}
			}
		}
	})
	//
	// Create OutEnd
	//
	hh.ginEngine.POST(API_ROOT+"outends", func(c *gin.Context) {
		cros(c)
		type Form struct {
			Type        string                 `json:"type" binding:"required"`
			Name        string                 `json:"name" binding:"required"`
			Description string                 `json:"description"`
			Config      map[string]interface{} `json:"config" binding:"required"`
		}
		form := Form{}
		err0 := c.ShouldBindJSON(&form)
		if err0 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err0.Error()})
		} else {
			configJson, err1 := json.Marshal(form.Config)
			if err1 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"msg": err1.Error()})
			} else {
				// TODO : Load newest OutEnd
				uuid := x.MakeUUID("OUTEND")
				hh.InsertMOutEnd(&MOutEnd{
					UUID:        uuid,
					Type:        form.Type,
					Name:        form.Name,
					Description: form.Description,
					Config:      string(configJson),
				})
				err := hh.LoadNewestOutEnd(uuid)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
				} else {
					c.JSON(http.StatusOK, gin.H{"msg": "create success"})
				}
			}
		}
	})
	// Create rule
	hh.ginEngine.POST(API_ROOT+"rules", func(c *gin.Context) {
		cros(c)
		type Form struct {
			From        string `json:"from" binding:"required"`
			Name        string `json:"name" binding:"required"`
			Description string `json:"description"`
			Actions     string `json:"actions"`
			Success     string `json:"success"`
			Failed      string `json:"failed"`
		}
		form := Form{}
		err0 := c.ShouldBindJSON(&form)
		if err0 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err0.Error()})
		} else {
			rule := x.NewRule(nil,
				form.Name,
				form.Description,
				nil,
				form.Success,
				form.Actions,
				form.Failed)
			if len(strings.Split(form.From, ",")) > 0 {
				for _, id := range strings.Split(form.From, ",") {
					// must be: 111,222,333... style
					if id != "" {
						if hh.ruleEngine.GetInEnd(id) == nil {
							c.JSON(http.StatusBadRequest, gin.H{"msg": "inend not exists:" + id})
							return
						}
					} else {
						c.JSON(http.StatusOK, gin.H{"msg": "invalid 'from' string format:" + form.From})
						return
					}
				}
				if err1 := x.VerifyCallback(rule); err1 != nil {
					c.JSON(http.StatusBadRequest, gin.H{"msg": err1.Error()})
				} else {
					mRule := &MRule{
						Name:        form.Name,
						Description: form.Description,
						From:        form.From,
						Success:     form.Success,
						Failed:      form.Failed,
						Actions:     form.Actions,
					}
					hh.InsertMRule(mRule)
					rule := x.NewRule(hh.ruleEngine,
						mRule.Name,
						mRule.Description,
						strings.Split(mRule.From, ","),
						mRule.Success,
						mRule.Actions,
						mRule.Failed)
					if err := hh.ruleEngine.LoadRule(rule); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
					} else {
						c.JSON(http.StatusOK, gin.H{"msg": "create success"})
					}
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"msg": "from can't empty"})
			}
		}
	})

	//
	hh.ginEngine.DELETE(API_ROOT+"rules", func(c *gin.Context) {
		cros(c)
		ruleId, exists := c.GetQuery("id")
		if exists {
			hh.ruleEngine.RemoveRule(ruleId)
			c.JSON(http.StatusOK, gin.H{})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "rule not exists"})
		}
	})
	//
	log.Info("Http web dashboard started on:http://127.0.0.1:2580" + DASHBOARD_ROOT)
	return nil
}

func (hh *HttpApiServer) Uninstall(env *x.XPluginEnv) error {
	return nil
}
func (hh *HttpApiServer) Clean() {
}
