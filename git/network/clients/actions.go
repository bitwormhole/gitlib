package clients

// Actions ...
type Actions struct {
}

// LoadRemoteRefs ...
func (inst *Actions) LoadRemoteRefs(c *Context) *RemoteRefsLoader {
	return &RemoteRefsLoader{context: c}
}

// LoadRemoteConfig ...
func (inst *Actions) LoadRemoteConfig(c *Context) *RemoteConfigLoader {
	return &RemoteConfigLoader{context: c}
}

// LoadAdvertisement ...
func (inst *Actions) LoadAdvertisement(c *Context) *AdvertisementLoader {
	return &AdvertisementLoader{context: c}
}
