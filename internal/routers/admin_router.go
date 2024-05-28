package routers

import (
	"github.com/PLDao/gin-frame/internal/controller/admin_v1"
	"github.com/PLDao/gin-frame/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetAdminApiRoute(e *gin.Engine) {
	// version 1
	v1 := e.Group("api/v1")
	{
		// 无需校验权限
		// 预先生成管理员然后通过这个接口拿到access_token
		loginC := admin_v1.NewLoginController()
		v1.POST("admin/login", loginC.Login)

		// social interface
		social := v1.Group("social")
		{
			socialUser := admin_v1.NewUserSocialController()
			social.GET("list", socialUser.GetAllUserSocial)
			social.POST("add", socialUser.AddUserSocial)
			social.PUT("update", socialUser.UpdateSocial)
		}

		// 需要校验权限
		reqAuth := v1.Group("", middleware.AdminAuthHandler())
		{
			// 管理员用户
			adminUser := reqAuth.Group("admin-user")
			{
				r := admin_v1.NewAdminUserController()
				// 获取用户信息
				adminUser.GET("get", r.GetUserInfo)
				adminUser.POST("add", r.AddUser)
			}

			// api权限管理
			permissions := reqAuth.Group("permission")
			{
				r := admin_v1.NewPermissionController()
				permissions.POST("edit", r.Edit)
				permissions.GET("list", r.List)
			}

			// 菜单管理
			menu := reqAuth.Group("menu")
			{
				r := admin_v1.NewAdminUserController()
				menu.GET("get", r.GetUserInfo)
			}

			// 角色管理
			role := reqAuth.Group("role")
			{
				r := admin_v1.NewAdminUserController()
				role.GET("get", r.GetUserInfo)
			}

			// 用户组管理
			group := reqAuth.Group("group")
			{
				r := admin_v1.NewAdminUserController()
				group.GET("get", r.GetUserInfo)
			}
		}
	}
}
