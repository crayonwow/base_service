package user

import (
	"testing"
	"time"
)

func TestName123(t *testing.T) {
	//r := response{
	//	Success: false,
	//	Errors: []errResponse{
	//		{
	//			Code:    "0",
	//			Context: "unexpected error",
	//		},
	//	},
	//}
	r := createUserRequest{
		Name:        "na",
		DateOfBirth: time.Now(),
	}
	b, _ := r.MarshalJSON()
	t.Log(string(b))
}
