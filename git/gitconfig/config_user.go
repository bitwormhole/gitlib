package gitconfig

// 定义仓库配置名称 (user.*)
const (
	AuthorName        KeyTemplate = "author.name"
	AuthorEmail       KeyTemplate = "author.email"
	CommitterName     KeyTemplate = "committer.name"
	CommitterEmail    KeyTemplate = "committer.email"
	UserName          KeyTemplate = "user.name"
	UserEmail         KeyTemplate = "user.email"
	UserUseConfigOnly KeyTemplate = "user.useConfigOnly"
	UserSigningKey    KeyTemplate = "user.signingKey"
)

// UserProperties ...
type UserProperties struct {
	access PropertyAccess
}

// AuthorEmail ...
func (inst *UserProperties) AuthorEmail() StringProperty {
	return inst.access.ForString(AuthorEmail, "", "")
}

// AuthorName ...
func (inst *UserProperties) AuthorName() StringProperty {
	return inst.access.ForString(AuthorName, "", "")
}

// CommitterEmail ...
func (inst *UserProperties) CommitterEmail() StringProperty {
	return inst.access.ForString(CommitterEmail, "", "")
}

// CommitterName ...
func (inst *UserProperties) CommitterName() StringProperty {
	return inst.access.ForString(CommitterName, "", "")
}

// UserEmail ...
func (inst *UserProperties) UserEmail() StringProperty {
	return inst.access.ForString(UserEmail, "", "")
}

// UserName ...
func (inst *UserProperties) UserName() StringProperty {
	return inst.access.ForString(UserName, "", "")
}

// UserUseConfigOnly ...
func (inst *UserProperties) UserUseConfigOnly() BooleanProperty {
	return inst.access.ForBool(UserUseConfigOnly, "", false)
}

// UserSigningKey ...
func (inst *UserProperties) UserSigningKey() StringProperty {
	return inst.access.ForString(UserSigningKey, "", "")
}
