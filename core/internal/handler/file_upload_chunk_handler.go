package handler

import (
	"cloud_disk/core/helper"
	"errors"
	"net/http"

	"cloud_disk/core/internal/logic"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		//判断关键字是否存在
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
			return
		}
		if r.PostForm.Get("upload_id") == "" {
			httpx.Error(w, errors.New("uploadId is empty"))
			return
		}
		if r.PostForm.Get("part_number") == "" {
			httpx.Error(w, errors.New("partNumber is empty"))
			return
		}
		etag, err := helper.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp = new(types.FileUploadChunkReply)
		resp.Etag = etag
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
