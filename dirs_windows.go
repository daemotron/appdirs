package appdirs

import (
	"os"
	"path/filepath"
)

func (conf AppConf) userDataDir() (string, error) {
	var err error
	var base string
	if conf.Roaming {
		base, err = os.UserConfigDir()
	} else {
		base, err = os.UserCacheDir()
	}
	if err != nil {
		return "", err
	}
	return filepath.Join(base, conf.Author, conf.Name), err
}

func (conf AppConf) siteDataDir() (string, error) {
	base := os.Getenv("ALLUSERSPROFILE")
	if base == "" {
		return "", ErrAllUsersProfileNotDefined
	}
	return filepath.Join(os.Getenv("ALLUSERSPROFILE"), conf.Author, conf.Name), nil
}

func (conf AppConf) globalDataDir() (string, error) {
	return conf.siteDataDir()
}

func (conf AppConf) userConfigDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) siteConfigDir() (string, error) {
	base, err := conf.siteDataDir()
	if err != nil {
		return "", err
	}
	if conf.Version != "" {
		return filepath.Join(base, conf.Version), nil
	}
	return base, nil
}

func (conf AppConf) globalConfigDir() (string, error) {
	return conf.siteConfigDir()
}

func (conf AppConf) userCacheDir() (string, error) {
	var base string
	var err error
	base, err = os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, conf.Author, conf.Name, "Cache"), err
}

func (conf AppConf) userStateDir() (string, error) {
	return conf.userDataDir()
}

func (conf AppConf) userLogDir() (string, error) {
	base, err := conf.userDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "Logs"), nil
}
