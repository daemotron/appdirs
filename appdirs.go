// Package appdirs helps to determine appropriate platform-specific system directories.
// This includes directories where to place cache files, user data, configuration, etc.
//
// Currently, this package supports the following operating systems:
//
//   - Linux
//   - Windows
//
// macOS is currently not supported (but planned for the future). While most unix systems should work with the present
// implementation, OS specific flavours (such as FreeBSD's /usr/local/etc directory) have not been implemented (yet).
//
// This package is inspired by ActiveState's now deprecated appdirs Python package.
package appdirs

// An AppConf instance represents an application.
// The fields in the struct can be filled in before calling
// any of its application-specific directory methods.
//
//   - Name indicates the application name and is mandatory
//   - Author indicates the application author or publisher and is optional
//   - Version is a version string and is optional
//   - Roaming is a flag indicating whether to use the roaming or local variant (Windows only)
//   - Multi is a flag indicating whether to return only one, or all possible paths separated by the OS specific separator
type AppConf struct {
	Name    string
	Author  string
	Version string
	Roaming bool
	Multi   bool
}

// UserDataDir returns the full path to the user-specific directory for this application
//
// Typical user data directories are:
//
//	macOS:                  ~/Library/Application Support/<AppName>
//	Unix:                   ~/.local/share/<AppName>    # or in $XDG_DATA_HOME, if defined
//	Windows (not roaming):  C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>
//	Windows (roaming):      C:\Users\<username>\AppData\Roaming\<AppAuthor>\<AppName>
//
// For Unix, we follow the XDG spec and support $XDG_DATA_HOME.
// That means, by default "~/.local/share/<AppName>".
func (conf AppConf) UserDataDir() (string, error) {
	return conf.userDataDir()
}

// SiteDataDir returns the full path to the user-shared data dir for this application.
//
// Typical site data directories are:
//
//	macOS:      /Library/Application Support/<AppName>
//	Unix:       /usr/local/share/<AppName> or /usr/share/<AppName>
//	Windows:    C:\ProgramData\<AppAuthor>\<AppName>   # Hidden, but writeable on Win 7.
//
// For Unix, this is using the $XDG_DATA_DIRS[0] default.
func (conf AppConf) SiteDataDir() (string, error) {
	return conf.siteDataDir()
}

// GlobalDataDir returns the full path to the global data dir for this application.
//
// Typical site data directories are:
//
//	macOS:      /Library/Application Support/<AppName>
//	Unix:       /var/lib/<AppName>
//	Windows:    C:\ProgramData\<AppAuthor>\<AppName>   # Hidden, but writeable on Win 7.
//
// This method is mostly geared towards Unix; on Windows it is identical to [SiteDataDir].
func (conf AppConf) GlobalDataDir() (string, error) {
	return conf.globalDataDir()
}

// UserConfigDir returns the full path to the user-specific config dir for this application.
//
// Typical user config directories are:
//
//	macOS:      ~/Library/Preferences/<AppName>
//	Unix:       ~/.config/<AppName>     # or in $XDG_CONFIG_HOME, if defined
//	Windows:    same as user_data_dir
//
// For Unix, we follow the XDG spec and support $XDG_CONFIG_HOME.
// That means, by default "~/.config/<AppName>".
func (conf AppConf) UserConfigDir() (string, error) {
	return conf.userConfigDir()
}

// SiteConfigDir returns the full path to the user-shared config dir for this application.
//
// Typical site config directories are:
//
//	macOS:      same as [SiteDataDir]
//	Unix:       /etc/xdg/<AppName> or $XDG_CONFIG_DIRS[i]/<AppName> for each value in $XDG_CONFIG_DIRS
//	Windows:    same as [SiteDataDir]
//
// For Unix, this is using the $XDG_CONFIG_DIRS[0] default, if conf.Multi = False
func (conf AppConf) SiteConfigDir() (string, error) {
	return conf.siteConfigDir()
}

// GlobalConfigDir returns the full path to the global config dir for this application.
//
// Typical global config directories are:
//
//	macOS:      same as [SiteConfigDir]
//	Unix:       /etc/<AppName>
//	Windows:    same as [SiteConfigDir]
//
// This method is mostly geared towards Unix; on Windows it is identical to [SiteDataDir].
func (conf AppConf) GlobalConfigDir() (string, error) {
	return conf.globalConfigDir()
}

// UserCacheDir returns the full path to the user-specific cache dir for this application.
//
// Typical user cache directories are:
//
//	macOS:      ~/Library/Caches/<AppName>
//	Unix:       ~/.cache/<AppName> (XDG default)
//	Windows:    C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Cache
//
// For Unix, we follow the XDG spec and support $XDG_CACHE_HOME.
// That means, by default "~/.cache/<AppName>"
func (conf AppConf) UserCacheDir() (string, error) {
	return conf.userCacheDir()
}

// GlobalCacheDir returns the full path to the global cache dir for this application.
//
// Typical global cache directories are:
//
//	macOS:      to be determined
//	Unix:       /var/cache/<AppName>
//	Windows:    same as [UserCacheDir]
//
// This method is mostly geared towards Unix; on Windows it is identical to [SiteDataDir].
func (conf AppConf) GlobalCacheDir() (string, error) {
	return conf.globalCacheDir()
}

// UserStateDir returns the full path to the user-specific state dir for this application.
//
// Typical user state directories are:
//
//	macOS:     same as [UserDataDir]
//	Unix:      ~/.local/state/<AppName>   # or in $XDG_STATE_HOME, if defined
//	Windows:   same as [UserDataDir]
//
// For Unix, we follow this Debian proposal <https://wiki.debian.org/XDGBaseDirectorySpecification#state>
// to extend the XDG spec and support $XDG_STATE_HOME.
// That means, by default "~/.local/state/<AppName>".
func (conf AppConf) UserStateDir() (string, error) {
	return conf.userStateDir()
}

// UserLogDir returns the full path to the user-specific log dir for this application.
//
// Typical user log directories are:
//
//	macOS:      ~/Library/Logs/<AppName>
//	Unix:       ~/.cache/<AppName>/log  # or under $XDG_CACHE_HOME if defined
//	Windows:    C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Logs
func (conf AppConf) UserLogDir() (string, error) {
	return conf.userLogDir()
}
