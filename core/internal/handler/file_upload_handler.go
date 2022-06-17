package handler

import (
	"cloud_disk/core/helper"
	"cloud_disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud_disk/core/internal/logic"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		bytes := make([]byte, fileHeader.Size)
		_, err = file.Read(bytes)
		if err != nil {
			return
		}
		//生成文件hash值
		hash := fmt.Sprintf("%x",md5.Sum(bytes))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if has {
			httpx.OkJson(w,&types.FileUploadReply{
				Identity: rp.Identity,
				Name: rp.Name,
				Ext: rp.Ext,
			})
			return
		}
		//如果文件不存在，则上传文件到cos
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}
		//往logic层传递文件信息 request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
