package card

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Service(t *testing.T) {
	tests := []struct {
		name           string
		repositoryFunc func(mock *MockRepository)
		validaterFunc  func(mock *MockValidater)
		wantErr        bool
	}{
		{
			name: "ok",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name:           "validation error",
			repositoryFunc: func(m *MockRepository) {},
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			wantErr: true,
		},
		{
			name: "create card error",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := NewMockRepository(ctrl)
			validater := NewMockValidater(ctrl)

			tc.repositoryFunc(repo)
			tc.validaterFunc(validater)

			s := NewService(repo, validater)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			form := Form{
				Word:          "reject",
				Transcription: "|rɪˈdʒekt|",
				Translation:   "отклонять",
				UserID:        1,
			}

			_, err := s.Create(ctx, &form)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}

func Test_Find_Service(t *testing.T) {
	tests := []struct {
		name           string
		repositoryFunc func(m *MockRepository)
		wantErr        bool
	}{
		{
			name: "ok",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&Card{}, nil)
			},
		},
		{
			name: "internal error",
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&Card{}, errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			repo := NewMockRepository(ctrl)
			tc.repositoryFunc(repo)

			s := NewService(repo, nil)

			_, err := s.Find(ctx, 1)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}

func Test_Update_Service(t *testing.T) {
	tests := []struct {
		name           string
		validaterFunc  func(mock *MockValidater)
		repositoryFunc func(mock *MockRepository)
		wantErr        bool
	}{
		{
			name: "ok",
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&Card{}, nil)
				m.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "validation error",
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			repositoryFunc: func(m *MockRepository) {},
			wantErr:        true,
		},
		{
			name: "find card error",
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&Card{}, errors.New("mock error"))
			},
			wantErr: true,
		},
		{
			name: "update card error",
			validaterFunc: func(m *MockValidater) {
				m.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			},
			repositoryFunc: func(m *MockRepository) {
				m.EXPECT().Find(gomock.Any(), gomock.Any()).Return(&Card{}, nil)
				m.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("mock error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := NewMockRepository(ctrl)
			validater := NewMockValidater(ctrl)

			tc.repositoryFunc(repo)
			tc.validaterFunc(validater)

			s := NewService(repo, validater)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			form := Form{
				Word:          "reject",
				Transcription: "|rɪˈdʒekt|",
				Translation:   "отклонять",
				UserID:        1,
			}

			_, err := s.Update(ctx, 1, &form)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Nil(t, err)
		})
	}
}
