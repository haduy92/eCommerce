package service

import (
	"fmt"

	"eCommerce/domain/dto"
	"eCommerce/domain/entity"
	"eCommerce/errs"
	"eCommerce/infrastructure/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PersonService interface {
	Get(*string) (*dto.PersonGetDto, error)
	GetAll() ([]*dto.PersonGetDto, error)
	Create(*dto.PersonCreateDto) (*uuid.UUID, error)
	Update(*string, *dto.PersonUpdateDto) error
	Delete(*string) error
}

type personService struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) PersonService {
	return &personService{
		repo: repo,
	}
}

func (service *personService) Get(id *string) (*dto.PersonGetDto, error) {
	guid, err := uuid.Parse(*id)
	if err != nil {
		return nil, errs.E(errs.Validation, errs.Parameter("id"), "incorrect format")
	}

	exists, err := service.repo.Exists(&guid)
	if err != nil {
		return nil, errs.E(errs.Internal, err)
	}
	if !*exists {
		return nil, errs.E(errs.NotExist, errs.Parameter("id"), fmt.Sprintf("\"%s\" is not found", *id))
	}

	person, err := service.repo.Get(&guid)
	if err != nil {
		return nil, errs.E(errs.Internal, err)
	} else {
		return &dto.PersonGetDto{
			ID:        person.ID,
			Name:      person.Name,
			Email:     person.Email,
			Role:      string(person.Role),
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
			DeletedAt: person.DeletedAt.Time,
		}, nil
	}
}

func (service *personService) GetAll() ([]*dto.PersonGetDto, error) {
	persons, err := service.repo.GetAll()
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.PersonGetDto, len(persons))
	for i := range persons {
		dto := &dto.PersonGetDto{}
		dto.ID = persons[i].ID
		dto.Name = persons[i].Name
		dto.Email = persons[i].Email
		dto.Role = persons[i].Role
		dto.CreatedAt = persons[i].CreatedAt
		dto.UpdatedAt = persons[i].UpdatedAt
		dto.DeletedAt = persons[i].DeletedAt.Time
		dtos[i] = dto
	}
	return dtos, nil
}

func (service *personService) Create(dto *dto.PersonCreateDto) (*uuid.UUID, error) {
	exists, err := service.repo.ExistsByEmail(&dto.Email)
	if err != nil {
		return nil, errs.E(errs.Internal, err)
	}
	if *exists {
		return nil, errs.E(errs.Exist, errs.Parameter("email"), fmt.Sprintf("\"%s\" is already exists", dto.Email))
	}

	person := entity.Person{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: HashPassword([]byte(dto.Password)),
		Role:     dto.Role,
	}
	id, err := service.repo.Create(&person)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (service *personService) Update(id *string, dto *dto.PersonUpdateDto) error {
	guid, err := uuid.Parse(*id)
	if err != nil {
		return errs.E(errs.Validation, errs.Parameter("id"), "incorrect format")
	}

	exists, err := service.repo.Exists(&guid)
	if err != nil {
		return errs.E(errs.Internal, err)
	}
	if !*exists {
		return errs.E(errs.NotExist, errs.Parameter("id"), fmt.Sprintf("\"%s\" is not found", *id))
	}

	existsByName, err := service.repo.ExistsByEmailExcludeID(&guid, &dto.Email)
	if err != nil {
		return errs.E(errs.Internal, err)
	}
	if *existsByName {
		return errs.E(errs.Exist, errs.Parameter("email"), fmt.Sprintf("\"%s\" is already exists", dto.Email))
	}

	person, err := service.repo.Get(&guid)
	if err != nil {
		return err
	}
	person.Name = dto.Name
	person.Email = dto.Email

	if err := service.repo.Update(person); err != nil {
		return err
	}
	return nil
}

func (service *personService) Delete(id *string) error {
	guid, err := uuid.Parse(*id)
	if err != nil {
		return errs.E(errs.Validation, errs.Parameter("id"), "incorrect format")
	}

	exists, err := service.repo.Exists(&guid)
	if err != nil {
		return errs.E(errs.Internal, err)
	}
	if !*exists {
		return errs.E(errs.NotExist, errs.Parameter("id"), fmt.Sprintf("\"%s\" is not found", *id))
	}

	if err := service.repo.Delete(&guid); err != nil {
		return err
	}
	return nil
}

func HashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		errs.E(errs.Internal, err)
	}
	return string(hash)
}
