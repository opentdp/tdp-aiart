package passport

import (
	"tdp-aiart/module/model/artwork"
	"tdp-aiart/module/model/user"
)

// 统计信息

func Summary(userId uint) map[string]any {

	artworkCount, _ := artwork.Count(&artwork.FetchAllParam{UserId: userId})
	userInfo, _ := user.Fetch(&user.FetchParam{Id: userId})

	return map[string]any{
		"Artwork":      artworkCount,
		"ArtworkQuota": userInfo.ArtworkQuota,
	}

}
