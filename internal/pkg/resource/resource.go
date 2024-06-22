package resource

import (
	"github.com/limes-cloud/kratosx"
	file "github.com/limes-cloud/resource/api/resource/file/v1"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/usercenter/internal/pkg/service"
)

func GetURLBySha(ctx kratosx.Context, sha string) string {
	client, err := service.NewResourceFile(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "init resource client error", "err", err.Error())
		return ""
	}
	reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: proto.String(sha)})
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource sha error", "err", err.Error())
		return ""
	}

	return reply.Url
}
