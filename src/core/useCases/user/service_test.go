package user

import (
	"errors"
	entity "go-api/src/core/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func newFixtureUser() *entity.User {
// 	return &entity.User{
// 		ID:        entity.NewID(),
// 		Name:      "oloco",
// 		Email:     "oloco@oloco.com",
// 		Password:  "213216574894",
// 		CreatedAt: time.Now(),
// 	}
// }

func newMockUser() *entity.User {
	user, _ := entity.NewUser("test", "test", "test")
	return user
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetById(id entity.ID) (*entity.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*entity.User), arg.Error(1)
}

func Test_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUser()
		mockRepo := new(MockRepository)

		mockRepo.On("GetById").Return(userMock, nil)

		testService := NewService(mockRepo)

		result, err := testService.GetUser(userMock.ID)

		assert.Nil(t, err)
		assert.Equal(t, userMock.ID, result.ID)
		assert.Equal(t, userMock.Name, result.Name)
		assert.Equal(t, userMock.Email, result.Email)
		assert.Equal(t, userMock.Password, result.Password)
		assert.Equal(t, userMock.CreatedAt, result.CreatedAt)
	})

	t.Run("error internal", func(t *testing.T) {
		userMock := newMockUser()

		mockRepo := new(MockRepository)

		errMock := errors.New("err")

		mockRepo.On("GetById").Return(userMock, errMock)

		testService := NewService(mockRepo)

		_, err := testService.GetUser(userMock.ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, errMock)
	})

	t.Run("error user not found", func(t *testing.T) {
		userMock := newMockUser()
		userMockResult := &entity.User{ID: uuid.Nil}

		mockRepo := new(MockRepository)

		mockRepo.On("GetById").Return(userMockResult, nil)

		testService := NewService(mockRepo)

		_, err := testService.GetUser(userMock.ID)

		assert.NotNil(t, err)
		assert.Equal(t, err, entity.ErrUserNotFound)
	})
}
