// Code generated by goctl. DO NOT EDIT.
package types

type ShareFileBasicDetailRequest struct {
	Identity string `json:"identity"` //用户资源表identity
}

type ShareFileBasicDetailReplay struct {
	RepositoryIdentity string `json:"repository_identity"` //资源表identity
	Name               string `json:"name"`                //资源名称
	Size               int64  `json:"size"`                //资源大小
	Ext                string `json:"ext"`                 //资源扩展名
	Path               string `json:"path"`                //资源路径
}

type UserShareCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"` //文件仓库标识
	ExpiredTime            int    `json:"expired_time"`             //过期时间
}

type UserShareCreateReply struct {
	Identity string `json:"identity"` //文件分享标识
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`        //文件标识
	ParentIdentity string `json:"parent_identity"` //父文件夹身份标识
}

type UserFileMoveReply struct {
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteReply struct {
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parentId"` //文件夹层级
	Name     string `json:"name"`     //文件夹名字
}

type UserFolderCreateReply struct {
	Identity string `json:"identity"` //文件夹标识
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateReply struct {
}

type UserFileListRequest struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFileListReply struct {
	List  []*UserFile `json:"list,optional"`
	Count int64       `json:"count,optional"`
}

type UserFile struct {
	Id                 int64  `json:"id,optional"`
	Identity           string `json:"identity,optional"`
	RepositoryIdentity string `json:"repository_identity,optional"`
	Name               string `json:"name,optional"`
	Exc                string `json:"exc,optional"`
	Size               int64  `json:"size,optional"`
	Path               string `json:"path,optional"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveReply struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterReplay struct {
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReplay struct {
	Token string `json:"token"`
}

type DetailRequest struct {
	Identity string `json:"identity"`
}

type DetailReplay struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendReplay struct {
}
