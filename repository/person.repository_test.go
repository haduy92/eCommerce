package repository

import (
	"database/sql"
	"eCommerce/model/entity"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

type PersonRepositoryTestSuite struct {
	suite.Suite
	db      *gorm.DB
	mock    sqlmock.Sqlmock
	repo    PersonRepository
	person  *entity.Person
	persons []*entity.Person
	columns []string
}

// this function executes before the test suite begins execution
func (s *PersonRepositoryTestSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	dialector := postgres.New(postgres.Config{
		DriverName:           "postgres",
		DSN:                  "sqlmock_db_0",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)

	s.repo = NewPersonRepository(s.db)
	s.columns = []string{"id", "name", "email", "password", "role", "created_at", "updated_at", "deleted_at"}
}

// This will run right before the test starts
// and receives the suite and test names as input
func (s *PersonRepositoryTestSuite) BeforeTest(_, _ string) {
	s.person = &entity.Person{
		ID:        uuid.New(),
		Name:      "John",
		Email:     "john@example.com",
		Role:      "admin",
		Password:  "password",
		CreatedAt: time.Now(),
	}

	for i := 1; i < 2; i++ {
		s.persons = append(s.persons, &entity.Person{
			ID:        uuid.New(),
			Name:      fmt.Sprintf("John_%d", i),
			Email:     fmt.Sprintf("john_%d@example.com", i),
			Role:      fmt.Sprintf("admin_%d", i),
			Password:  fmt.Sprintf("password_%d", i),
			CreatedAt: time.Now(),
		})
	}
}

// This will run after test finishes
// and receives the suite and test names as input
func (s *PersonRepositoryTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestPersonRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PersonRepositoryTestSuite))
}

func (s *PersonRepositoryTestSuite) TestGet_IdExisted_ShouldReturnData() {
	// Arrange
	rows := sqlmock.NewRows(s.columns).
		AddRow(s.person.ID, s.person.Name, s.person.Email, s.person.Password, s.person.Role, s.person.CreatedAt, s.person.UpdatedAt, s.person.DeletedAt)
	sql := regexp.QuoteMeta(`SELECT * FROM "person" WHERE id = $1 AND "person"."deleted_at" IS NULL ORDER BY "person"."id" LIMIT 1`)
	s.mock.ExpectQuery(sql).WithArgs(s.person.ID).WillReturnRows(rows)

	// Action
	res, err := s.repo.Get(&s.person.ID)

	// Assert
	assert.NilError(s.T(), err)
	assert.DeepEqual(s.T(), s.person, res)
}

func (s *PersonRepositoryTestSuite) TestGet_IdNotExisted_ShouldReturnError() {
	// Arrange
	id := uuid.New()
	rows := sqlmock.NewRows(s.columns)
	sql := regexp.QuoteMeta(`SELECT * FROM "person" WHERE id = $1 AND "person"."deleted_at" IS NULL ORDER BY "person"."id" LIMIT 1`)
	s.mock.ExpectQuery(sql).WithArgs(id).WillReturnRows(rows)

	// Action
	res, err := s.repo.Get(&id)

	// Assert
	assert.Error(s.T(), err, "record not found")
	assert.DeepEqual(s.T(), &entity.Person{}, res)
}

func (s *PersonRepositoryTestSuite) TestGetAll_HasRows_ShouldReturnData() {
	rows := sqlmock.NewRows(s.columns)
	for _, row := range s.persons {
		rows.AddRow(row.ID, row.Name, row.Email, row.Password, row.Role, row.CreatedAt, row.UpdatedAt, row.DeletedAt)
	}

	sql := regexp.QuoteMeta(`SELECT * FROM "person" WHERE "person"."deleted_at" IS NULL`)

	s.mock.ExpectQuery(sql).WillReturnRows(rows)

	res, err := s.repo.GetAll()
	assert.NilError(s.T(), err)
	assert.DeepEqual(s.T(), s.persons, res)
}

// Other tests for PersonRepository
// ...
