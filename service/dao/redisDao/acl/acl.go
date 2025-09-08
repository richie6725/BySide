package acl

import (
	"Byside/service/dao/daoModels/acl"
	"Byside/service/internal/tools/compress"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type AclRedisDao interface {
	Get(ctx context.Context, username string) (*aclDaoModel.User, error)
	Set(ctx context.Context, acl aclDaoModel.User, expiration time.Duration) error
}

func New(r *redis.Client) AclRedisDao {
	dao := &aclRedisDao{
		prefixKey: "byside:acl:",
		client:    r,
	}
	return dao
}

type aclRedisDao struct {
	prefixKey string
	client    *redis.Client
}

func (dao *aclRedisDao) buildAclKey(username string) string {
	return fmt.Sprintf("%s:%s", dao.prefixKey, username)
}

func (dao *aclRedisDao) Get(ctx context.Context, username string) (*aclDaoModel.User, error) {
	key := dao.buildAclKey(username)
	a, err := dao.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, nil
	}

	acl, err := dao.unCompressAcl(a)
	if err != nil {
		return nil, err
	}

	return acl, nil
}

func (dao *aclRedisDao) Set(ctx context.Context, acl aclDaoModel.User, expiration time.Duration) error {

	key := dao.buildAclKey(acl.Username)

	compressed, err := dao.compressAcl(&acl)
	if err != nil {
		return err
	}

	return dao.client.Set(ctx, key, compressed, expiration).Err()

}

func (dao *aclRedisDao) compressAcl(acl *aclDaoModel.User) ([]byte, error) {
	value, err := json.Marshal(acl)
	if err != nil {
		return nil, err
	}
	compressed, _ := compress.CompressBytes(value)
	return compressed, nil
}

func (dao *aclRedisDao) unCompressAcl(data []byte) (*aclDaoModel.User, error) {
	uncompressed, _ := compress.UncompressBytes(data)
	var result aclDaoModel.User
	if err := json.Unmarshal(uncompressed, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
