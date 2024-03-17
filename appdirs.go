package appdirs

type AppConf struct {
	name    string
	author  string
	version string
	roaming bool
}

func (conf AppConf) UserDataDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) SiteDataDir() (string, error) {
	return conf.siteDataDir()
}

func (conf AppConf) GlobalDataDir() (string, error) {
	return conf.globalDataDir()
}

func (conf AppConf) UserConfigDir() (string, error) {
	return conf.userConfigDir()
}

func (conf AppConf) SiteConfigDir() (string, error) {
	return conf.siteConfigDir()
}

func (conf AppConf) GlobalConfigDir() (string, error) {
	return conf.globalConfigDir()
}

func (conf AppConf) UserCacheDir() (string, error) {
	return conf.userCacheDir()
}

func (conf AppConf) UserStateDir() (string, error) {
	return conf.userStateDir()
}

func (conf AppConf) UserLogDir() (string, error) {
	return conf.userLogDir()
}
