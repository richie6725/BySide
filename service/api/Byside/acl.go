package BysideApi

import (
	"Byside/service/controller/aclCtrl"
	aclDaoModel "Byside/service/dao/daoModels/acl"
	boAcl "Byside/service/internal/model/bo/acl"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"time"
)

func NewAcl(pack aclApiPack) {
	c := &aclApi{pack: pack}
	group := pack.Root.Group("acl")
	{
		group.GET("status", c.getAcl)
		group.POST("status", c.updateAcl)
	}

}

type aclApiPack struct {
	dig.In
	AclCtrl aclCtrl.AclCtrl
	Root    *gin.RouterGroup
}

type aclApi struct {
	pack aclApiPack
}

func (api *aclApi) getAcl(ctx *gin.Context) {
	//form := struct {
	//	Username string `form:"username" valid:"required"`
	//	Password string `form:"password" valid:"required"`
	//}{}
	//if err := ctx.BindQuery(&form); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
	//	return
	//}
	//if _, err := govalidator.ValidateStruct(form); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
	//	return
	//}
	form := struct {
		Username string `json:"username" valid:"required"`
		Password string `json:"password" valid:"required"`
	}{}

	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	user := &boAcl.GetArgs{
		User: aclDaoModel.User{
			Username: form.Username,
			Password: form.Password,
		}}

	result, err := api.pack.AclCtrl.Get(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (api *aclApi) updateAcl(ctx *gin.Context) {
	//form := struct {
	//	Username string   `form:"username" valid:"required"`
	//	Password string   `form:"password" valid:"required"`
	//	Roles    []string `form:"roles" valid:"required"`
	//	CreateAt int64    `form:"createAt" valid:"required"`
	//	UpdateAt int64    `form:"updateAt" valid:"required"`
	//}{}
	//if err := ctx.BindQuery(&form); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
	//	return
	//}
	//if _, err := govalidator.ValidateStruct(form); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
	//	return
	//}
	form := struct {
		Username string   `json:"username" valid:"required"`
		Password string   `json:"password" valid:"required"`
		Roles    []string `json:"roles" valid:"required"`
		CreateAt int64    `json:"createAt" valid:"required"`
		UpdateAt int64    `json:"updateAt" valid:"required"`
	}{}

	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	user := &boAcl.UpdateArgs{
		User: aclDaoModel.User{
			Username:  form.Username,
			Password:  form.Password,
			Roles:     form.Roles,
			CreatedAt: time.Unix(form.CreateAt, 0),
			UpdatedAt: time.Unix(form.UpdateAt, 0),
		},
	}

	result, err := api.pack.AclCtrl.Update(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4. 成功回應
	ctx.JSON(http.StatusOK, result)
}
