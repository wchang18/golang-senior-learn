// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"user-admin/internal/model"
)

type (
	IStore interface {
		GetUser(userName string) (user model.User, err error)
		SetUser(user model.User) (err error)
	}
)

var (
	localStore IStore
)

func Store() IStore {
	if localStore == nil {
		panic("implement not found for interface IStore, forgot register?")
	}
	return localStore
}

func RegisterStore(i IStore) {
	localStore = i
}
