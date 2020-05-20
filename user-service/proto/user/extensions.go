package laracom_user_service

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (m *User) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.NewV4()
	return scope.SetColumn("Id", uid.String())
}
