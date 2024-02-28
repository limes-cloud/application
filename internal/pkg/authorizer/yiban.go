package authorizer

import (
	"errors"
	"time"

	"github.com/forgoer/openssl"
	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
)

type yb struct {
}

func init() {
	register("yb", &yb{})
}

type yiBanAccessTokenInfo struct {
	// VisitTime int64 `json:"visit_time"`

	// VisitUser struct {
	//	UserId string `json:"userid"`
	// } `json:"visit_user"`
	VisitOauth struct {
		AccessToken  string `json:"access_token"`
		TokenExpires string `json:"token_expires"`
	} `json:"visit_oauth"`
}

type yiBanAuthInfo struct {
	Status string `json:"status"`
	Info   struct {
		MsgCN        string `json:"msgCN"`
		YbUserid     string `json:"yb_userid"`
		YbUsername   string `json:"yb_username"`
		YbUsernick   string `json:"yb_usernick"`
		YbSex        string `json:"yb_sex"`
		YbBirthday   string `json:"yb_birthday"`
		YbMoney      string `json:"yb_money"`
		YbExp        string `json:"yb_exp"`
		YbUserhead   string `json:"yb_userhead"`
		YbRegtime    string `json:"yb_regtime"`
		YbSchoolid   string `json:"yb_schoolid"`
		YbSchoolname string `json:"yb_schoolname"`
	} `json:"info"`
}

func (y yb) Name() string {
	return "易班"
}

func (y yb) GetAccessToken(ctx kratosx.Context, req GetAccessTokenRequest) (*GetAccessTokenReply, error) {
	src := util.HexToByte(req.Code)
	iv := []byte(req.Ak)
	key := []byte(req.Sk)

	body, _ := openssl.AesCBCDecrypt(src, key, iv, openssl.PKCS7_PADDING)
	if body == nil {
		return nil, errors.New("decrypt error")
	}

	// 解析数据
	res := yiBanAccessTokenInfo{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	return &GetAccessTokenReply{
		Token:  res.VisitOauth.AccessToken,
		Expire: time.Duration(util.ToInt64(res.VisitOauth.TokenExpires)),
	}, nil
}

func (y yb) GetAuthInfo(ctx kratosx.Context, req GetAuthInfoRequest) (*GetAuthInfoReply, error) {
	url := "https://openapi.yiban.cn/user/me?access_token=" + req.Token
	resp, err := ctx.Http().Get(url)
	if err != nil {
		return nil, err
	}

	data := yiBanAuthInfo{}
	if err := resp.Result(&data); err != nil {
		return nil, err
	}

	if data.Status == "error" {
		return nil, errors.New(data.Info.MsgCN)
	}

	return &GetAuthInfoReply{AuthId: data.Info.YbUserid}, nil
}