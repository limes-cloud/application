package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/biz"
)

func IsInit(ctx kratosx.Context) bool {
	db := ctx.DB().Migrator()
	return db.HasTable(biz.Channel{}) &&
		db.HasTable(biz.App{}) &&
		db.HasTable(biz.ExtraField{}) &&
		db.HasTable(biz.User{}) &&
		db.HasTable(biz.UserApp{}) &&
		db.HasTable(biz.UserExtra{})
}

func Init(ctx kratosx.Context, _ *config.Config) {
	db := ctx.DB()
	_ = db.Set("gorm:table_options", "COMMENT='协议信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Agreement{})
	_ = db.Set("gorm:table_options", "COMMENT='场景信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Scene{})
	_ = db.Set("gorm:table_options", "COMMENT='场景协议信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.AgreementScene{})
	_ = db.Set("gorm:table_options", "COMMENT='授权信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Auth{})
	_ = db.Set("gorm:table_options", "COMMENT='授权渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Channel{})
	_ = db.Set("gorm:table_options", "COMMENT='应用信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.App{})
	_ = db.Set("gorm:table_options", "COMMENT='应用接口' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.AppInterface{})
	_ = db.Set("gorm:table_options", "COMMENT='应用渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.AppChannel{})
	_ = db.Set("gorm:table_options", "COMMENT='字段信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.ExtraField{})
	_ = db.Set("gorm:table_options", "COMMENT='用户信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.User{})
	_ = db.Set("gorm:table_options", "COMMENT='用户渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.UserChannel{})
	_ = db.Set("gorm:table_options", "COMMENT='用户应用' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.UserApp{})
	_ = db.Set("gorm:table_options", "COMMENT='用户扩展信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.UserExtra{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(db)
}
