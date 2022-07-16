package service

import (
	"dousheng/global"
	"dousheng/models"

	"github.com/gin-gonic/gin"
)

func GetVedioList(c *gin.Context) {
	FeedReq := models.FeedReq{}
	c.ShouldBind(&FeedReq) //接收请求数据
	var vedio []models.Vedio
	var vedioResp []models.VedioResp
	var user models.User
	db := global.DB
	max := 30

	db.Find(&vedio) //从数据中查找vedio，构造vedioResp
	if len(vedio) > max {
		vedio = vedio[:max]
	}
	for _, v := range vedio {
		db.Where("user_id = ?", v.UID).Find(&user)
		userResp := models.GetUserResp(user, false)
		media := models.GetVedioResp(v, userResp, false)
		vedioResp = append(vedioResp, media)
	}

	c.JSON(200, gin.H{ //返回vedioResp
		"status_code": 0,
		"status_msg":  "success",
		"vedio_list":  vedioResp,
	})
}
