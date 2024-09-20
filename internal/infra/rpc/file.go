package rpc

import (
	"sync"

	"github.com/limes-cloud/kratosx"
	file "github.com/limes-cloud/resource/api/resource/file/v1"

	"github.com/limes-cloud/application/api/application/errors"
)

const (
	Resource = "Resource"
)

var (
	fileIns  *File
	fileOnce sync.Once
)

type File struct {
}

func NewFile() *File {
	fileOnce.Do(func() {
		fileIns = &File{}
	})
	return fileIns
}

func (i File) client(ctx kratosx.Context) (file.FileClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Resource)
	if err != nil {
		return nil, errors.ResourceServerError()
	}
	return file.NewFileClient(conn), nil
}

func (i File) GetFileURL(ctx kratosx.Context, sha string) string {
	client, err := i.client(ctx)
	if err != nil {
		ctx.Logger().Warnw("msg", "connect resource server error", "err", err.Error())
		return ""
	}
	reply, err := client.GetFile(ctx, &file.GetFileRequest{Sha: &sha})
	if err != nil {
		ctx.Logger().Warnw("msg", "get resource file error", "err", err.Error())
		return ""
	}
	return reply.Url
}
