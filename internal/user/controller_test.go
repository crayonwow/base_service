package user

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

type testCase struct {
}

func TestController(t *testing.T) {
	testCases := []struct {
		name             string
		serviceMock      func(*testing.T) MockuserService
		handler          func(*controller) http.HandlerFunc
		reqBody          string
		code             int
		expectedResponse json.Marshaler
		actualResponse   json.Unmarshaler
		comparer         func(t *testing.T, expected, actual interface{})
	}{
		{
			name: "success",
			serviceMock: func(t *testing.T) MockuserService {
				m := NewMockuserService(gomock.NewController(t))
				m.EXPECT().Create(gomock.Any(), User{
					Name:        "test_name",
					DateOfBirth: time.Time{},
				}).Return(1, nil)
			},
		},
	}
}
