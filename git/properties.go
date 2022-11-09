package git

// Properties 表示属性文档格式，例如 .git/config 文件
type Properties struct {
	table map[string]string
}

func (inst *Properties) getTab() map[string]string {
	t := inst.table
	if t == nil {
		t = make(map[string]string)
		inst.table = t
	}
	return t
}

// Get ...
func (inst *Properties) Get(name string) string {
	t := inst.getTab()
	return t[name]
}

// Set ...
func (inst *Properties) Set(name, value string) {
	t := inst.getTab()
	t[name] = value
}

// Clear ...
func (inst *Properties) Clear() {
	inst.table = nil // make(map[string]string)
}

// Trim ...
func (inst *Properties) Trim() {
	const empty = ""
	src := inst.table
	dst := make(map[string]string)
	for k, v := range src {
		if k == empty || v == empty {
			continue
		}
		dst[k] = v
	}
	inst.table = dst
}

// Export ...
func (inst *Properties) Export(dst map[string]string) map[string]string {
	const empty = ""
	if dst == nil {
		dst = make(map[string]string)
	}
	src := inst.table
	for k, v := range src {
		if k == empty || v == empty {
			continue
		}
		dst[k] = v
	}
	return dst
}

// Import ...
func (inst *Properties) Import(src map[string]string) {
	const empty = ""
	dst := inst.table
	if dst == nil {
		dst = make(map[string]string)
	}
	for k, v := range src {
		if k == empty || v == empty {
			continue
		}
		dst[k] = v
	}
	inst.table = dst
}
