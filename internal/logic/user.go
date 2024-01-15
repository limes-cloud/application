package logic

import (
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/user-center/api/v1"
	"github.com/limes-cloud/user-center/config"
	"github.com/limes-cloud/user-center/internal/model"
	"github.com/limes-cloud/user-center/pkg/util"
)

type User struct {
	conf *config.Config
}

func NewLogic(conf *config.Config) *User {
	return &User{
		conf: conf,
	}
}

// Get 获取用户信息
func (u *User) Get(ctx kratosx.Context, in *v1.GetUserRequest) (*v1.User, error) {
	var err error
	var user model.User

	switch in.Condition.(type) {
	case *v1.GetUserRequest_Id:
		cond := in.Condition.(*v1.GetUserRequest_Id)
		err = user.FindByID(ctx, cond.Id)
	case *v1.GetUserRequest_Email:
		cond := in.Condition.(*v1.GetUserRequest_Email)
		err = user.FindByEmail(ctx, cond.Email)
	case *v1.GetUserRequest_Phone:
		cond := in.Condition.(*v1.GetUserRequest_Phone)
		err = user.FindByPhone(ctx, cond.Phone)
	case *v1.GetUserRequest_Username:
		cond := in.Condition.(*v1.GetUserRequest_Username)
		err = user.FindByUsername(ctx, cond.Username)
	}

	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.User{}
	if err := util.Transform(user, &reply); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Page 分页获取用户
func (u *User) Page(ctx kratosx.Context, in *v1.PageUserRequest) (*v1.PageUserReply, error) {
	user := model.User{}

	list, total, err := user.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.App != nil {
				db = db.InnerJoins("UserApps", ctx.DB().Where("UserApps.app=?", in.App))
			}
			if in.Username != nil {
				db = db.Where("username=?", *in.Username)
			}
			if in.Email != nil {
				db = db.Where("email=?", *in.Email)
			}
			if in.Phone != nil {
				db = db.Where("phone=?", *in.Phone)
			}
			return db
		},
	})
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	reply := v1.PageUserReply{Total: total}
	// 进行数据转换
	if err = util.Transform(list, &reply.List); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	return &reply, nil
}

// Add 添加用户信息
func (u *User) Add(ctx kratosx.Context, in *v1.AddUserRequest) (*v1.AddUserReply, error) {
	user := model.User{}
	// 进行数据转换
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 设置默认值
	user.Status = proto.Bool(true)
	if user.NickName == "" {
		user.NickName = user.RealName
	}

	// 创建用户
	if err := user.Create(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return &v1.AddUserReply{Id: user.ID}, nil
}

// Import 导入用户信息
func (u *User) Import(ctx kratosx.Context, in *v1.ImportUserRequest) (*empty.Empty, error) {
	var list []*model.User
	// 进行数据转换
	if err := util.Transform(in.List, &list); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 设置默认值
	user := model.User{}

	// 创建用户
	if err := user.Import(ctx, list); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Update 添加用户信息
func (u *User) Update(ctx kratosx.Context, in *v1.UpdateUserRequest) (*empty.Empty, error) {
	user := model.User{}
	// 进行数据转换
	if err := util.Transform(in, &user); err != nil {
		return nil, v1.TransformErrorFormat(err.Error())
	}

	// 更新用户
	if err := user.Update(ctx); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Delete 删除用户信息
func (u *User) Delete(ctx kratosx.Context, in *v1.DeleteUserRequest) (*empty.Empty, error) {
	user := model.User{}

	// 删除用户
	if err := user.DeleteById(ctx, in.Id); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	return nil, nil
}

// Enable 启用用户
func (r *User) Enable(ctx kratosx.Context, in *v1.EnableUserRequest) (*empty.Empty, error) {
	nu := model.User{
		BaseModel:   types.BaseModel{ID: in.Id},
		Status:      proto.Bool(true),
		DisableDesc: proto.String(""),
	}

	if err := nu.Update(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Disable 禁用用户
func (r *User) Disable(ctx kratosx.Context, in *v1.DisableUserRequest) (*empty.Empty, error) {
	nu := model.User{
		BaseModel:   types.BaseModel{ID: in.Id},
		Status:      proto.Bool(false),
		DisableDesc: proto.String(in.Desc),
	}

	if err := nu.Update(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Offline 对当前登陆用户进行下线
func (r *User) Offline(ctx kratosx.Context, in *v1.OfflineUserRequest) (*empty.Empty, error) {
	// 查询用户
	user := model.UserApp{}

	list, err := user.FindByUserId(ctx, in.Id)
	if err != nil {
		return nil, v1.NotFoundError()
	}

	// 将用户下线
	for _, app := range list {
		if app.ExpireAt < time.Now().Unix() {
			ctx.JWT().AddBlacklist(app.Token)
		}
	}
	return nil, nil
}
