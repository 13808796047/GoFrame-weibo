package upload

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

//var allowed_ext = []string{"png", "jpg", "gif", "jpeg"}
var allowed_ext = garray.NewStrArrayFrom(g.SliceStr{"png", "jpg", "gif", "jpeg"})

func Save(file *ghttp.UploadFile, folder string, file_prefix string) (string, error) {
	folder_name := "uploads/images/" + folder + "/" + gtime.Date()
	upload_path := g.Cfg().GetString("server.ServerRoot") + "/" + folder_name
	extension := gstr.SubStrRune(file.Filename, gstr.PosRRune(file.Filename, ".")+1, gstr.LenRune(file.Filename)-1)
	file.Filename = file_prefix + "_" + gsha1.Encrypt(file_prefix) + "." + extension
	if !allowed_ext.Contains(extension) {
		err := gerror.New("上传文件类型不合法")
		return "", err
	}
	filename, err := file.Save(upload_path, true)
	url := "/" + folder_name + "/" + filename

	if err != nil {
		return "", err
	}
	return url, nil
}
