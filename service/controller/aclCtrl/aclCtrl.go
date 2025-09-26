package aclCtrl

import (
	aclDaoModel "Byside/service/dao/daoModels/acl"
	aclMongoDao "Byside/service/dao/mongoDao/acl"
	aclRedisDao "Byside/service/dao/redisDao/acl"
	boAcl "Byside/service/internal/model/bo/acl"
	"Byside/service/internal/utils"
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"time"
)

type aclCtrl struct {
	pack aclCtrlPack
}

type aclCtrlPack struct {
	dig.In
	MongoByside *mongo.Database `name:"mongo_byside"`
	RedisByside *redis.Client   `name:"redis_byside"`
}

type AclCtrl interface {
	Get(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetReply, error)
	GetLogin(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetLoginReply, error)
	Update(ctx context.Context, args *boAcl.UpdateArgs) error
}

func NewAcl(pack aclCtrlPack) AclCtrl {
	return &aclCtrl{
		pack: pack,
	}
}

func (ctrl *aclCtrl) Get(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetReply, error) {
	aclDao := aclMongoDao.New(ctrl.pack.MongoByside)
	reply := &boAcl.GetReply{}

	user, err := aclDao.Get(ctx, args.User)
	if err != nil {
		return nil, err
	}

	if user != nil {
		reply.User = user
		return reply, nil
	}

	return nil, nil
}
func (ctrl *aclCtrl) GetLogin(ctx context.Context, args *boAcl.GetArgs) (*boAcl.GetLoginReply, error) {
	aclDao := aclMongoDao.New(ctrl.pack.MongoByside)
	aclRao := aclRedisDao.New(ctrl.pack.RedisByside)

	session, err := aclRao.Get(ctx, args.User.Username, args.User.Token)
	if err != nil {
		return nil, err
	}
	if session != nil {
		return &boAcl.GetLoginReply{Session: session}, nil
	}

	user, err := aclDao.Get(ctx, args.User)
	if err != nil {
		return nil, err
	}

	if user != nil {
		token := utils.GenerateToken()

		session := &aclDaoModel.UserSession{
			Username: user.Username,
			Token:    token,
		}

		if err := aclRao.Set(ctx, session, time.Minute*30); err != nil {
			return nil, err
		}

		return &boAcl.GetLoginReply{Session: session}, nil
	}

	return nil, nil
}

func (ctrl *aclCtrl) Update(ctx context.Context, args *boAcl.UpdateArgs) error {

	aclDao := aclMongoDao.New(ctrl.pack.MongoByside)

	err := aclDao.Update(ctx, aclDaoModel.Query{args.Query.BulkUserArgs, args.Query.CreatedAt})
	if err != nil {
		return err
	}

	return nil
}
