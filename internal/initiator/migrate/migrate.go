package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"github.com/limes-cloud/user-center/internal/biz/agreement"
	"github.com/limes-cloud/user-center/internal/biz/app"
	"github.com/limes-cloud/user-center/internal/biz/channel"
	"github.com/limes-cloud/user-center/internal/biz/user"
	"github.com/limes-cloud/user-center/internal/config"
)

func IsInit(ctx kratosx.Context) bool {
	// db := ctx.DB().Migrator()
	// return db.HasTable(other.Channel{}) &&
	//	db.HasTable(other.App{}) &&
	//	db.HasTable(other.ExtraField{}) &&
	//	db.HasTable(other.User{}) &&
	//	db.HasTable(other.UserApp{}) &&
	//	db.HasTable(other.UserExtra{})
	return false
}

func Init(ctx kratosx.Context, _ *config.Config) {
	db := ctx.DB()

	_ = db.Set("gorm:table_options", "COMMENT='协议内容' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(agreement.Content{})
	_ = db.Set("gorm:table_options", "COMMENT='协议场景' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(agreement.Scene{})
	_ = db.Set("gorm:table_options", "COMMENT='协议场景-内容' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(agreement.SceneContent{})
	_ = db.Set("gorm:table_options", "COMMENT='授权渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(channel.Channel{})
	_ = db.Set("gorm:table_options", "COMMENT='应用信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(app.App{})
	_ = db.Set("gorm:table_options", "COMMENT='用户信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(user.User{})
	_ = db.Set("gorm:table_options", "COMMENT='用户应用' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(user.UserApp{})
	_ = db.Set("gorm:table_options", "COMMENT='用户扩展信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(user.UserExtra{})
	_ = db.Set("gorm:table_options", "COMMENT='用户授权' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(user.Auth{})

	// _ = db.Set("gorm:table_options", "COMMENT='协议信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.Agreement{})
	// _ = db.Set("gorm:table_options", "COMMENT='场景信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.Scene{})
	// _ = db.Set("gorm:table_options", "COMMENT='场景协议信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.AgreementScene{})
	// _ = db.Set("gorm:table_options", "COMMENT='授权信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.Auth{})
	// _ = db.Set("gorm:table_options", "COMMENT='授权渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.Channel{})
	// _ = db.Set("gorm:table_options", "COMMENT='应用信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.App{})
	// _ = db.Set("gorm:table_options", "COMMENT='应用接口' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.AppInterface{})
	// _ = db.Set("gorm:table_options", "COMMENT='应用渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.AppChannel{})
	// _ = db.Set("gorm:table_options", "COMMENT='字段信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.ExtraField{})
	// _ = db.Set("gorm:table_options", "COMMENT='用户信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.User{})
	// _ = db.Set("gorm:table_options", "COMMENT='用户渠道' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.UserChannel{})
	// _ = db.Set("gorm:table_options", "COMMENT='用户应用' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.UserApp{})
	// _ = db.Set("gorm:table_options", "COMMENT='用户扩展信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(other.UserExtra{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(db)
}
