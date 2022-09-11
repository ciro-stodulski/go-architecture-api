package notificaitonpb

import (
	"context"
	"go-clean-api/cmd/interface/grpc/notification/pb"
	"go-clean-api/cmd/main/container"
	"go-clean-api/cmd/shared/mocks"

	listuserusecasemock "go-clean-api/cmd/shared/mocks/core/use-cases/list-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServiceGrpc_FindUser_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUseCase := new(listuserusecasemock.MockUseCase)
		userMock := mocks.NewMockUser()

		mockUseCase.On("ListUsers").Return(userMock, nil)
		//

		// test func
		testPb := New(&container.Container{
			ListUsersUseCase: mockUseCase,
		})

		pb := &pb.ResquestNotification{List: &pb.List{Name: "", Describe: ""}}
		ctx := context.Background()

		result, err := testPb.Verify(ctx, pb)
		//

		// asserts
		assert.Nil(t, err)
		mockUseCase.AssertCalled(t, "ListUsers")
		assert.NotNil(t, result)
		//
	})
}
