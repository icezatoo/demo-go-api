package repository

import (
	"errors"
	"reflect"
	"testing"

	dto "github.com/icezatoo/demo-go-api/pkg/dto/user"
	mocks "github.com/icezatoo/demo-go-api/pkg/mocks/repository"
	model "github.com/icezatoo/demo-go-api/pkg/model"
)

func TestRepository_CreateUser(t *testing.T) {

	type args struct {
		request *dto.CreateUserRequest
	}

	payload := &dto.CreateUserRequest{
		Email:    "test@gmail.com",
		FullName: "test1",
		LastName: "test2",
		Enabled:  true,
		Password: "123456",
	}

	mockReturnUser := &model.User{
		ID:       "0000-0000-0000-000",
		Email:    "test@gmail.com",
		FullName: "test1",
		LastName: "test2",
		Enabled:  true,
		Password: "123456",
	}

	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "success_create_user",
			args: args{
				request: payload,
			},
			want:    mockReturnUser,
			wantErr: false,
		},
		{
			name: "failed_create_user",
			args: args{
				request: payload,
			},
			want:    &model.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.UserRepository{}

			if !tt.wantErr {
				repo.On("CreateUser", tt.args.request).Return(tt.want, nil)
			} else {
				repo.On("CreateUser", tt.args.request).Return(tt.want, errors.New("Failed to create user"))
			}

			got, err := repo.CreateUser(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateUser() error = %v, wantErr %v\n", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreateUser() = %v, want %v\n", got, tt.want)
			}
		})
	}
}
