package BysideApi

import (
	"Byside/service/controller/noteCtrl"
	noteDaoModel "Byside/service/dao/daoModels/note"
	boNote "Byside/service/internal/model/bo/note"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
)

func NewNote(pack noteApiPack) {
	c := &noteApi{pack: pack}
	group := pack.Root.Group("note")
	{
		group.POST("updateNote", c.updateNote)
	}

}

type noteApiPack struct {
	dig.In
	NoteCtrl noteCtrl.NoteCtrl
	Root     *gin.RouterGroup
}

type noteApi struct {
	pack noteApiPack
}

func (api *noteApi) updateNote(ctx *gin.Context) {

	form := struct {
		IsUpsert bool `valid:"-" form:"is_upsert"`
	}{}

	if err := ctx.BindQuery(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}

	if _, err := govalidator.ValidateStruct(form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var models []*noteDaoModel.PriceRecord

	if err := ctx.BindJSON(&models); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	query := &boNote.UpdateArgs{
		BulkPriceRecordArgs: models,
		IsUpsert:            form.IsUpsert,
	}

	err := api.pack.NoteCtrl.Update(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	return
}
