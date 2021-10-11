package main

import (
	"framedemo/framework"
)

func resistRouter(core *framework.Core) {

	core.Get("/admin/login", Login)
	groupUser := core.Group("/user")
	{
		groupUser.Post("/*id", Create)
		groupUser.Put("/*id", Update)
		groupUser.Delete("/*id", Delete)
	}

}
