package auth

import (
	"time"

	"GoTenancy/config"
	"GoTenancy/config/auth/themes"
	"GoTenancy/config/bindatafs"
	"github.com/qor/render"

	//"GoTenancy/config/bindatafs"
	"GoTenancy/config/db"
	"GoTenancy/models/users"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
)

var (

	// Auth 初始化用于认证的 Auth
	Auth = themes.New(&auth.Config{
		DB:         db.DB,
		Mailer:     config.Mailer,
		Render:     render.New(&render.Config{AssetFileSystem: bindatafs.AssetFS.NameSpace("admin"), DefaultLayout: "layout"}),
		UserModel:  users.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority 初始化用于认证的 Authority
	Authority = authority.New(&authority.Config{
		Auth: Auth,
	})
)

func init() {

	Auth.RegisterProvider(github.New(&config.Config.Github))
	Auth.RegisterProvider(google.New(&config.Config.Google))
	Auth.RegisterProvider(facebook.New(&config.Config.Facebook))
	Auth.RegisterProvider(twitter.New(&config.Config.Twitter))

	Authority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})
}
