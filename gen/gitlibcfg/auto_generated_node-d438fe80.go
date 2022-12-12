// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gitlibcfg

import (
	clients0x465781 "github.com/bitwormhole/gitlib/git/network/clients"
	http0xe484f4 "github.com/bitwormhole/gitlib/git/network/http"
	servers0xb5845d "github.com/bitwormhole/gitlib/git/network/servers"
	http4git0x537f63 "github.com/bitwormhole/gitlib/git/support/net/http4git"
	services0x887aef "github.com/bitwormhole/gitlib/git/support/services"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComMainClientImpl struct {
	instance *clients0x465781.MainClientImpl
	 markup0x23084a.Component `id:"git-main-client"`
	ClientRegistryList []clients0x465781.ClientRegistry `inject:".git-client-registry"`
}


type pComGitClient struct {
	instance *http0xe484f4.GitClient
	 markup0x23084a.Component `class:"git-client-registry"`
}


type pComMainServerImpl struct {
	instance *servers0xb5845d.MainServerImpl
	 markup0x23084a.Component `id:"git-main-server"`
	ServerRegistryList []servers0xb5845d.ServerRegistry `inject:".git-server-registry"`
}


type pComHTTPGitConnector struct {
	instance *http4git0x537f63.HTTPGitConnector
	 markup0x23084a.Component `class:"pktline-connector-registry"`
}


type pComGitFetchService struct {
	instance *services0x887aef.GitFetchService
	 markup0x23084a.Component `class:"git-instruction-registry"`
	MainClient clients0x465781.MainClient `inject:"#git-main-client"`
}


type pComGitPushService struct {
	instance *services0x887aef.GitPushService
	 markup0x23084a.Component `class:"git-instruction-registry"`
	MainClient clients0x465781.MainClient `inject:"#git-main-client"`
}

