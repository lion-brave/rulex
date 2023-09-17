package sqlitedao

import (
	"runtime"

	"github.com/hootrhino/rulex/core"

	"github.com/hootrhino/rulex/glogger"
	dao "github.com/hootrhino/rulex/plugin/http_server/dao"
	"github.com/hootrhino/rulex/plugin/http_server/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Sqlite dao.DAO

/*
*
* Sqlite 数据持久层
*
 */
type SqliteDAO struct {
	name string   // 框架可以根据名称来选择不同的数据库驱动,为以后扩展准备
	db   *gorm.DB // Sqlite 驱动
}

/*
*
* 新建一个SqliteDAO
*
 */
func Load(dbPath string) {
	Sqlite = &SqliteDAO{name: "SqliteDAO"}
	Sqlite.Init(dbPath)
}

/*
*
* 初始化DAO
*
 */
func (s *SqliteDAO) Init(dbPath string) error {
	var err error
	if core.GlobalConfig.AppDebugMode {
		s.db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: false,
		})
	} else {
		s.db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Error),
			SkipDefaultTransaction: false,
		})
	}
	if err != nil {
		glogger.GLogger.Fatal(err)
	}
	s.RegisterModel()
	return err
}

/*
*
* 停止
*
 */
func (s *SqliteDAO) Stop() {
	s.db = nil
	runtime.GC()
}

/*
*
* 返回数据库查询句柄
*
 */
func (s *SqliteDAO) DB() *gorm.DB {
	return s.db
}

/*
*
* 返回名称
*
 */
func (s *SqliteDAO) Name() string {
	return s.name
}

/*
*
* 注册数据模型
*
 */
func (s *SqliteDAO) RegisterModel() {
	s.DB().AutoMigrate(
		&model.MInEnd{},
		&model.MOutEnd{},
		&model.MRule{},
		&model.MUser{},
		&model.MDevice{},
		&model.MGoods{},
		&model.MApp{},
		&model.MAiBase{},
		&model.MModbusPointPosition{},
		&model.MVisual{},
		&model.MGenericGroup{},
		&model.MGenericGroupRelation{},
		&model.MProtocolApp{},
		&model.MNetworkConfig{},
		&model.MWifiConfig{},
		&model.MDataSchema{},
		&model.MSiteConfig{},
		&model.MIpRoute{},
		&model.MCronTask{},
		&model.MCronResult{},
	)
}
