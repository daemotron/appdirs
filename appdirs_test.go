package appdirs_test

import (
	"testing"

	"github.com/daemotron/appdirs"
)

func setup() appdirs.AppConf {
	return appdirs.AppConf{Name: "apptest", Author: "daemotron", Version: "1.0", Multi: true}
}

func TestUserDataDir(t *testing.T) {
	app := setup()
	res, err := app.UserDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`User Data Dir: %v`, res)
}

func TestSiteDataDir(t *testing.T) {
	app := setup()
	res, err := app.SiteDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`Site Data Dir: %v`, res)
}

func TestGlobalDataDir(t *testing.T) {
	app := setup()
	res, err := app.GlobalDataDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`Global Data Dir: %v`, res)
}

func TestUserConfigDir(t *testing.T) {
	app := setup()
	res, err := app.UserConfigDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`User Config Dir: %v`, res)
}

func TestSiteConfigDir(t *testing.T) {
	app := setup()
	res, err := app.SiteConfigDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`Site Config Dir: %v`, res)
}

func TestGlobalConfigDir(t *testing.T) {
	app := setup()
	res, err := app.GlobalConfigDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`Global Config Dir: %v`, res)
}

func TestUserStateDir(t *testing.T) {
	app := setup()
	res, err := app.UserStateDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`User State Dir: %v`, res)
}

func TestUserCacheDir(t *testing.T) {
	app := setup()
	res, err := app.UserCacheDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`User Cache Dir: %v`, res)
}

func TestGlobalCacheDir(t *testing.T) {
	app := setup()
	res, err := app.GlobalCacheDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`Global Cache Dir: %v`, res)
}

func TestUserLogDir(t *testing.T) {
	app := setup()
	res, err := app.UserLogDir()
	if err != nil {
		t.Fatalf(`Unexpected error: %v`, err)
	}
	t.Logf(`User Log Dir: %v`, res)
}
