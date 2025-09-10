package acl

import "Byside/service/dao/daoModels/acl"

type GetArgs struct {
	User aclDaoModel.User
}

type GetReply struct {
	User aclDaoModel.User
}
type UpdateArgs struct {
	Query aclDaoModel.Query
}

type UpdateReply struct{}
