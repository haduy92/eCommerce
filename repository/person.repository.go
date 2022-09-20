package repository

import (
	"time"

	"eCommerce/model/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonRepository interface {
	Get(*uuid.UUID) (*entity.Person, error)
	Search(*string) ([]*entity.Person, error)
	Create(*entity.Person) (*uuid.UUID, error)
	Update(*entity.Person) error
	Delete(*uuid.UUID) error
	Exists(*uuid.UUID) (*bool, error)
	ExistsByEmail(*string) (*bool, error)
	ExistsByEmailExcludeID(*uuid.UUID, *string) (*bool, error)
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	return &personRepository{
		db: db,
	}
}

func (repo *personRepository) Get(id *uuid.UUID) (*entity.Person, error) {
	var person *entity.Person
	result := repo.db.Where("id = ?", id).First(&person)
	return person, result.Error
}

func (repo *personRepository) Search(q *string) ([]*entity.Person, error) {
	var persons []*entity.Person
	var query = repo.db
	if q != nil {
		query = query.Where("name like ? or email like ?", "%"+*q+"%", "%"+*q+"%")
	}
	result := query.Find(&persons)
	return persons, result.Error
}

func (repo *personRepository) Create(person *entity.Person) (*uuid.UUID, error) {
	person.ID = uuid.New()
	person.CreatedAt = time.Now().UTC()
	if err := repo.db.Save(&person).Error; err != nil {
		return nil, err
	}
	return &person.ID, nil
}

func (repo *personRepository) Update(person *entity.Person) error {
	var oldPerson *entity.Person
	if err := repo.db.Find(&oldPerson, person.ID).Error; err != nil {
		return err
	} else if err = repo.db.Model(&oldPerson).Updates(map[string]interface{}{
		"Name":      person.Name,
		"Email":     person.Email,
		"Password":  person.Password,
		"Role":      person.Role,
		"UpdatedAt": time.Now().UTC(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *personRepository) Delete(id *uuid.UUID) error {
	var person *entity.Person
	repo.db.Delete(&entity.Person{}, 10)
	if err := repo.db.Where("ID = ?", id).Delete(&person).Error; err != nil {
		return err
	}
	return nil
}

func (repo *personRepository) Exists(id *uuid.UUID) (*bool, error) {
	var found bool
	res := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM person WHERE id = ? AND deleted_at IS NULL) AS found", id).Scan(&found)
	if res.Error != nil {
		return nil, res.Error
	}
	return &found, nil
}

func (repo *personRepository) ExistsByEmail(email *string) (*bool, error) {
	var found bool
	res := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM person WHERE email = ? AND deleted_at IS NULL) AS found", email).Scan(&found)
	if res.Error != nil {
		return nil, res.Error
	}
	return &found, nil
}

func (repo *personRepository) ExistsByEmailExcludeID(id *uuid.UUID, email *string) (*bool, error) {
	var found bool
	res := repo.db.Raw("SELECT EXISTS(SELECT 1 FROM person WHERE email = ? AND id <> ? AND deleted_at IS NULL) AS found", email, id).Scan(&found)
	if res.Error != nil {
		return nil, res.Error
	}
	return &found, nil
}
