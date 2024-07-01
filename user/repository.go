package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// Digunakan untuk inisialisasi instance baru dari struct repository dengan parameter db
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Kita membuat function yang bernama Save untuk tipe repository dan memiliki balikan User dan error
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// Note :
// Repository, R besar bersifat public, r kecil bersifat private
