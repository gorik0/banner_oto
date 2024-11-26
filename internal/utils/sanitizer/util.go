package sanitizer

import (
	"banners_oto/internal/entity"

	"github.com/microcosm-cc/bluemonday"
)

// Позже, если будем делать ACL, скорей всего, добавим сюда и другие функции, но пока что только о пользователе меняются данные

func User(u *entity.User) *entity.User {
	sanitizer := bluemonday.UGCPolicy()
	u.Name = sanitizer.Sanitize(u.Name)
	u.Phone = sanitizer.Sanitize(u.Phone)
	u.Email = sanitizer.Sanitize(u.Email)
	u.Address = sanitizer.Sanitize(u.Address)
	u.ImgUrl = sanitizer.Sanitize(u.ImgUrl)
	return u
}

func Address(address string) string {
	sanitizer := bluemonday.UGCPolicy()
	return sanitizer.Sanitize(address)
}
