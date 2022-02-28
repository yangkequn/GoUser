package logic

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/jpeg"
	"net/http"

	"user/internal/svc"

	"github.com/nfnt/resize"
	"github.com/yangkequn/GoTools"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) PutUserAvatarLogic {
	return PutUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func CompressImageResource(data []byte) ([]byte, error) {
	var MaxHeight uint = 128

	if len(data) < 5*1024 {
		return data, nil
	}

	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// 修改图片的大小
	m := resize.Resize(0, MaxHeight, img, resize.Lanczos3)
	buf := bytes.Buffer{}

	// 修改图片的质量
	err = jpeg.Encode(&buf, m, &jpeg.Options{Quality: 60})
	if err != nil {
		return data, err
	}

	if buf.Len() > len(data) {
		return data, nil
	}
	return buf.Bytes(), nil
}
func (l *PutUserAvatarLogic) PutUserAvatar(r *http.Request) error {
	Uid := l.ctx.Value("uid").(string)
	uid := GoTools.StringToInt64(Uid)
	if uid == 0 {
		return nil
	}

	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		return err
	}
	if fileHeader == nil || fileHeader.Size > 5*1024*1024 {
		return errors.New("bad file size")
	}
	defer file.Close()
	avatarBytes := make([]byte, fileHeader.Size)
	n, err := file.Read(avatarBytes)
	if int64(n) != fileHeader.Size || err != nil {
		return errors.New("bad file upload")
	}

	u, errUser := l.svcCtx.UsersModel.FindOne(uid)
	if errUser != nil {
		return errUser
	}
	if u.Id > 0 {
		compressed, errCompress := CompressImageResource(avatarBytes)
		//_,err:=base64.StdEncoding.Decode(dst,req.Avatar)
		if errCompress != nil {
			return errCompress
		}
		u.Avatar = string(compressed)
		l.svcCtx.UsersModel.Update(u)
		return nil
	}
	return nil
}
