package noteCtrl

import (
	noteMongoDao "Byside/service/dao/mongoDao/note"
	boNote "Byside/service/internal/model/bo/note"
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

type noteCtrl struct {
	pack noteCtrlPack
}

type noteCtrlPack struct {
	dig.In
	MongoByside *mongo.Database `name:"mongo_byside"`
	RedisByside *redis.Client   `name:"redis_byside"`
}

type NoteCtrl interface {
	Update(ctx context.Context, args *boNote.UpdateArgs) error
}

func NewNote(pack noteCtrlPack) NoteCtrl {
	return &noteCtrl{
		pack: pack,
	}
}

func (ctrl *noteCtrl) Update(ctx context.Context, args *boNote.UpdateArgs) error {
	noteDao := noteMongoDao.New(ctrl.pack.MongoByside)

	err := noteDao.Update(ctx, args.BulkPriceRecordArgs, args.IsUpsert)
	if err != nil {
		return err
	}

	return nil
}
