package permission

import (
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/auth"
	resourcev1 "github.com/limes-cloud/manager/api/manager/resource/v1"

	"github.com/limes-cloud/usercenter/internal/pkg/service"
)

const (
	App = "uc_app"
)

func GetPermission(ctx kratosx.Context, keyword string) (bool, []uint32, error) {
	request, is := http.RequestFromServerContext(ctx)
	if !is {
		return true, nil, nil
	}
	if strings.Contains(request.RequestURI, "/usercenter/client/") {
		return true, nil, nil
	}

	info, err := auth.GetAuthInfo(ctx)
	if err != nil {
		return false, nil, err
	}
	if info.UserId == 1 || info.RoleId == 1 {
		return true, nil, nil
	}

	client, err := service.NewManagerResource(ctx)
	if err != nil {
		return false, nil, err
	}
	reply, err := client.GetResourceScopes(ctx, &resourcev1.GetResourceScopesRequest{
		Keyword: keyword,
		UserId:  info.UserId,
	})

	if err != nil {
		return false, nil, err
	}
	return reply.All, reply.Scopes, nil
}

func GetApp(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := GetPermission(ctx, App)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

func HasApp(ctx kratosx.Context, id uint32) bool {
	all, ids, err := GetPermission(ctx, App)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}
