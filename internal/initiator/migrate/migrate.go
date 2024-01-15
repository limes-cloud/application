package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/model"
)

func IsInit(ctx kratosx.Context) bool {
	db := ctx.DB().Migrator()
	return db.HasTable(model.User{}) &&
		db.HasTable(model.UserApp{}) &&
		db.HasTable(model.UserExtra{})
}

func Init(ctx kratosx.Context, _ *config.Config) {
	db := ctx.DB()
	_ = db.Set("gorm:table_options", "COMMENT='用户信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.User{})
	_ = db.Set("gorm:table_options", "COMMENT='用户应用信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.UserApp{})
	_ = db.Set("gorm:table_options", "COMMENT='用户扩展信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.UserExtra{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(ctx.DB())
}
