package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	common "github.com/hootrhino/rulex/plugin/http_server/common"
	"github.com/hootrhino/rulex/plugin/http_server/model"
	"github.com/hootrhino/rulex/plugin/http_server/service"
	"github.com/hootrhino/rulex/utils"
)

type MGenericGroupVo struct {
	UUID   string `json:"uuid" validate:"required"` // 名称
	Name   string `json:"name" validate:"required"` // 名称
	Type   string `json:"type" validate:"required"` // 组的类型, DEVICE: 设备分组, VISUAL: 大屏分组
	Parent string `json:"parent"`                   // 上级, 如果是0表示根节点
}
type MGenericGroupRelationVo struct {
	Gid string `json:"gid" validate:"required"` // 分组ID
	Rid string `json:"rid" validate:"required"` // 被绑定方
}

/*
*
* 新建大屏
*
 */
func CreateGroup(c *gin.Context, hh *HttpApiServer) {
	vvo := MGenericGroupVo{}
	if err := c.ShouldBindJSON(&vvo); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	if !utils.SContains([]string{"VISUAL", "DEVICE"}, vvo.Type) {
		c.JSON(common.HTTP_OK, common.Error400(fmt.Errorf("invalid type [%s]", vvo.Type)))
	}
	Model := model.MGenericGroup{
		UUID:   utils.GroupUuid(),
		Name:   vvo.Name,
		Type:   vvo.Type,
		Parent: "0",
	}
	if err := service.InsertGenericGroup(&Model); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	c.JSON(common.HTTP_OK, common.Ok())

}

/*
*
* 更新大屏
*
 */
func UpdateGroup(c *gin.Context, hh *HttpApiServer) {
	vvo := MGenericGroupVo{}
	if err := c.ShouldBindJSON(&vvo); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	Model := model.MGenericGroup{
		UUID:   vvo.UUID,
		Name:   vvo.Name,
		Type:   vvo.Type,
		Parent: "0",
	}
	if err := service.UpdateGenericGroup(&Model); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	c.JSON(common.HTTP_OK, common.Ok())
}

/*
*
* 删除大屏
*
 */
func DeleteGroup(c *gin.Context, hh *HttpApiServer) {
	uuid, _ := c.GetQuery("uuid")
	count, err := service.CheckBindResource(uuid)
	if err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	// 删除之前需要判断一下是不是有子元
	if count > 0 {
		msg := fmt.Errorf("group:%s have binding to other resources", uuid)
		c.JSON(common.HTTP_OK, common.Error400(msg))
		return
	}
	if err := service.DeleteGenericGroup(uuid); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
	}
	c.JSON(common.HTTP_OK, common.Ok())

}

/*
*
* 大屏列表
*
 */
func ListGroup(c *gin.Context, hh *HttpApiServer) {
	visuals := []MGenericGroupVo{}
	for _, vv := range service.AllGenericGroup() {
		visuals = append(visuals, MGenericGroupVo{
			UUID:   vv.UUID,
			Name:   vv.Name,
			Type:   vv.Type,
			Parent: vv.Parent,
		})
	}
	c.JSON(common.HTTP_OK, common.OkWithData(visuals))

}

/*
*
* 大屏详情
*
 */
func GroupDetail(c *gin.Context, hh *HttpApiServer) {
	uuid, _ := c.GetQuery("uuid")
	mG, err := service.GetGenericGroupWithUUID(uuid)
	if err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
	}
	vo := MGenericGroupVo{
		UUID:   mG.UUID,
		Name:   mG.Name,
		Type:   mG.Type,
		Parent: mG.Parent,
	}
	c.JSON(common.HTTP_OK, common.OkWithData(vo))
}

/*
*
* 绑定资源
*
 */
func BindResource(c *gin.Context, hh *HttpApiServer) {
	type FormDto struct {
		Gid string `json:"gid"`
		Rid string `json:"rid"`
	}
	form := FormDto{}
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	if count, err := service.CheckAlreadyBinding(form.Gid, form.Rid); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	} else {
		if count > 0 {
			msg := fmt.Errorf("[%s] already binding to group [%s]", form.Rid, form.Gid)
			c.JSON(common.HTTP_OK, common.Error400(msg))
			return
		}
	}
	if err := service.BindResource(form.Gid, form.Rid); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	c.JSON(common.HTTP_OK, common.Ok())

}

/*
*
* 取消绑定
*
 */
func UnBindResource(c *gin.Context, hh *HttpApiServer) {
	gid, _ := c.GetQuery("gid")
	rid, _ := c.GetQuery("rid")
	if err := service.UnBindResource(gid, rid); err != nil {
		c.JSON(common.HTTP_OK, common.Error400(err))
		return
	}
	c.JSON(common.HTTP_OK, common.Ok())

}

/*
*
* 设备
*
 */
func FindDeviceByGroup(c *gin.Context, hh *HttpApiServer) {
	Type, _ := c.GetQuery("type")
	Gid, _ := c.GetQuery("gid")
	vv2 := []model.MDevice{}

	if Type == "DEVICE" {
		_, MDevices := service.FindByType(Gid, Type)
		for _, mG := range MDevices {
			vv2 = append(vv2, model.MDevice{
				UUID: mG.UUID,
				Name: mG.Name,
				Type: mG.Type,
			})
		}
		c.JSON(common.HTTP_OK, common.OkWithData(vv2))
		return
	}
	c.JSON(common.HTTP_OK, common.Error400(fmt.Errorf("unsupported group type:%s", Type)))
}

/*
*
* 大屏
*
 */
func FindVisualByGroup(c *gin.Context, hh *HttpApiServer) {
	Type, _ := c.GetQuery("type")
	Gid, _ := c.GetQuery("gid")
	vv1 := []model.MVisual{}
	if Type == "VISUAL" {
		MVisuals, _ := service.FindByType(Gid, Type)
		for _, mG := range MVisuals {
			vv1 = append(vv1, model.MVisual{
				UUID: mG.UUID,
				Name: mG.Name,
				Type: mG.Type,
			})
		}
		c.JSON(common.HTTP_OK, common.OkWithData(vv1))
		return
	}
	c.JSON(common.HTTP_OK, common.Error400(fmt.Errorf("unsupported group type:%s", Type)))
}