package clients

// Action 表示动作
type Action string

// 定义各种动作
const (
	ActionFetch Action = "fetch"
	ActionPush  Action = "push"
)

func (a Action) String() string {
	return string(a)
}

////////////////////////////////////////////////////////////////////////////////

// CommonActions ...
type CommonActions struct {
}

// LoadRemoteRefs ...
func (inst *CommonActions) LoadRemoteRefs(c *Context) *RemoteRefsLoader {
	return &RemoteRefsLoader{context: c}
}

// LoadLocalConfig ...
func (inst *CommonActions) LoadLocalConfig(c *Context) *LocalConfigLoader {
	return &LocalConfigLoader{context: c}
}

// LoadRemoteConfig ...
func (inst *CommonActions) LoadRemoteConfig(c *Context) *RemoteConfigLoader {
	return &RemoteConfigLoader{context: c}
}

// LoadAdvertisement ...
func (inst *CommonActions) LoadAdvertisement(c *Context) *AdvertisementLoader {
	return &AdvertisementLoader{context: c}
}
