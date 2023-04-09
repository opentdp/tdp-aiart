package passport

import (
	"tdp-aiart/module/model/artimg"
	"tdp-aiart/module/model/user"
)

// 统计信息

func Summary(userId uint) map[string]any {

	artimgCount, _ := artimg.Count(&artimg.FetchAllParam{UserId: userId})
	userInfo, _ := user.Fetch(&user.FetchParam{Id: userId})

	return map[string]any{
		"Artimg":      artimgCount,
		"QuotaArtimg": userInfo.QuotaArtimg,
	}

}
