package service

import (
	"eCommerce/model/dto"
	"eCommerce/model/entity"
	"eCommerce/repository/mock"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gotest.tools/assert"
)

type PersonServiceTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	mock       *mock.MockPersonRepository
	service    PersonService
	entity     *entity.Person
	entityList []*entity.Person
	dto        *dto.PersonGetDto
	dtoList    []*dto.PersonGetDto
}

// this function executes before the test suite begins execution
func (s *PersonServiceTestSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.mock = mock.NewMockPersonRepository(s.ctrl)
	s.service = NewPersonService(s.mock)
}

func TestPersonServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PersonServiceTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (s *PersonServiceTestSuite) BeforeTest(_, _ string) {
	s.entity = &entity.Person{
		ID:        uuid.New(),
		Name:      "John",
		Email:     "john@example.com",
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.dto = &dto.PersonGetDto{
		ID:        s.entity.ID,
		Name:      s.entity.Name,
		Email:     s.entity.Email,
		Role:      s.entity.Role,
		CreatedAt: s.entity.CreatedAt,
		UpdatedAt: s.entity.UpdatedAt,
	}

	for i := 0; i < 5; i++ {
		s.entityList = append(s.entityList, &entity.Person{
			ID:        uuid.New(),
			Name:      fmt.Sprintf("John_%d", i),
			Email:     fmt.Sprintf("john_%d@example.com", i),
			Role:      fmt.Sprintf("admin_%d", i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
	for i := 0; i < 5; i++ {
		s.dtoList = append(s.dtoList, &dto.PersonGetDto{
			ID:        s.entityList[i].ID,
			Name:      s.entityList[i].Name,
			Email:     s.entityList[i].Email,
			Role:      s.entityList[i].Role,
			CreatedAt: s.entityList[i].CreatedAt,
			UpdatedAt: s.entityList[i].UpdatedAt,
		})
	}
}

func (s *PersonServiceTestSuite) TestGet_IdExisted_ShouldReturnData() {
	// Arrange
	id := uuid.New()
	idStr := id.String()
	exists := true
	s.mock.EXPECT().Exists(&id).Return(&exists, nil)
	s.mock.EXPECT().Get(&id).Return(s.entity, nil)

	// Action
	res, err := s.service.Get(&idStr)

	// Assert
	assert.NilError(s.T(), err)
	assert.DeepEqual(s.T(), s.dto, res)
}

func (s *PersonServiceTestSuite) TestGet_IdNotExisted_ShouldReturnData() {
	// Arrange
	id := uuid.New()
	strID := id.String()
	s.mock.EXPECT().Exists(&id).Return(FALSE(), nil)

	// Action
	res, err := s.service.Get(&strID)

	// Assert
	assert.Error(s.T(), err, fmt.Sprintf("\"%s\" is not found", strID))
	assert.Assert(s.T(), res == nil)
}

func TRUE() *bool {
	b := true
	return &b
}

func FALSE() *bool {
	b := false
	return &b
}
