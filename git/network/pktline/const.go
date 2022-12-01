package pktline

// 定义各种 git-pktline 常量
// 参考 https://gitee.com/bitwormhole/git/blob/master/Documentation/gitprotocol-v2.txt

// 定义 git 服务名称
const (
	ServiceGitReceivePack = "git-receive-pack"
	ServiceGitUploadPack  = "git-upload-pack"
)

// 定义 git content-types
const (
	TypeGitUploadPackAdvertisement = "application/x-git-upload-pack-advertisement"
	TypeGitUploadPackRequest       = "application/x-git-upload-pack-request"
	TypeGitUploadPackResult        = "application/x-git-upload-pack-result"
	TypeGitReceivePackRequest      = "application/x-git-receive-pack-request"
	TypeGitReceivePackResult       = "application/x-git-receive-pack-result"
)
