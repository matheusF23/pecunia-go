package database

import (
	"database/sql"
	"testing"

	"github.com/matheusF23/pecunia-go/domain/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id varchar(255) NOT NULL, name varchar(255) NOT NULL, email varchar(255) NOT NULL, password varchar(255) NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestGivenAnUser_WhenCreate_ThenShouldCreateUser() {
	user, err := entity.NewUser("123", "Aristeu", "aristeu@email.com", "secretPass")
	suite.NoError(err)
	repo := NewUserRepository(suite.Db)
	err = repo.CreateUser(user)
	suite.NoError(err)
	var userCreated entity.User
	err = suite.Db.QueryRow("SELECT id, name, email, password FROM users where id = ?", user.ID).
		Scan(&userCreated.ID, &userCreated.Name, &userCreated.Email, &userCreated.Password)
	suite.NoError(err)
	suite.Equal(user.ID, userCreated.ID)
	suite.Equal(user.Name, userCreated.Name)
	suite.Equal(user.Email, userCreated.Email)
	suite.Equal(user.Password, userCreated.Password)
}
