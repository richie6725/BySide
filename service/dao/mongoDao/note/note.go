package note

import (
	noteDaoModel "Byside/service/dao/daoModels/note"
	"Byside/service/dao/mongoDao"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const noteCollection = "note"

type NoteDao interface {
	Update(ctx context.Context, models []*noteDaoModel.PriceRecord, isUpsert bool) error
}

type noteDao struct {
	collection *mongo.Collection
}

func New(db *mongo.Database) NoteDao {
	dao := &noteDao{
		collection: db.Collection(noteCollection),
	}
	return dao
}

func (dao *noteDao) Update(ctx context.Context, models []*noteDaoModel.PriceRecord, isUpsert bool) error {
	writes := make([]mongo.WriteModel, len(models))
	for i := range writes {
		filter := mongoDao.NewMatchBuilder().
			AddEqual(noteDaoModel.FieldMarketName, models[i].Market.Name).
			AddEqual(noteDaoModel.FieldProductID, models[i].Product.ProductID).Generate()
		if len(filter) == 0 {
			return fmt.Errorf("invalid filter: market=%s productID=%s", models[i].Market.Name, models[i].Product.ProductID)
		}
		doc := bson.M{"$set": models[i]}
		update := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(doc).SetUpsert(isUpsert)
		writes[i] = update
	}
	_, err := dao.collection.BulkWrite(ctx, writes, options.BulkWrite().SetOrdered(false))
	if err != nil {
		return err
	}

	return nil
}
