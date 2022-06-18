package handler

import (
	"net/http"

	"cloud_disk/core/internal/logic"
	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareFileBasicDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareFileBasicDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShareFileBasicDetailLogic(r.Context(), svcCtx)
		resp, err := l.ShareFileBasicDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
