package appdirs

import "path/filepath"

func (conf AppConf) userDataDir() (string, error) {
	base, err := getHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "Library", "Application Support", conf.Name, conf.Version)
}

func (conf AppConf) siteDataDir() (string, error) {
	return filepath.Join("/Library", "Application Support", conf.Name, conf.Version), nil
}

func (conf AppConf) globalDataDir() (string, error) {
	return conf.siteDataDir()
}

func (conf AppConf) userConfigDir() (string, error) {
	base, err := getHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "Library", "Preferences", conf.Name, conf.Version)
}

func (conf AppConf) siteConfigDir() (string, error) {
	return filepath.Join("/Library", "Preferences", conf.Name, conf.Version)
}

func (conf AppConf) globalConfigDir() (string, error) {
	return conf.siteConfigDir()
}

func (conf AppConf) globalCacheDir() (string, error) {
	return filepath.Join("/Library", "Caches", conf.Name, conf.Version)
}

func (conf AppConf) userStateDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) userLogDir() (string, error) {
	base, err := getHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "Library", "Logs", conf.Name, conf.Version)
}
