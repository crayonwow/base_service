//go:generate easyjson
//go:generate mockgen -source=models.go -destination=models_mocks.go -package=user

package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"

	cachePkg "github.com/crayonwow/base_service/pkg/cache"
)

// errors
var (
	ErrUserNotFound = errors.New("user not found")

	errInvalidID   = errors.New("invalid id")
	errInvalidName = errors.New("invalid name")
	errInvalidDOB  = errors.New("invalid date of birth")

	errFailedToWriteResponse = errors.New("failed to write response")
	errEmptyUserID           = errors.New("empty user id")
)

const (
	hardFailError = `{"success":false,"errors":[{"code":"0","context":"unexpected error"}]}`
	dobFormat     = "2006-01-02"
)

type (
	userService interface {
		Create(ctx context.Context, user User) (int64, error)
		Get(ctx context.Context, userID int64) (User, error)
		Update(ctx context.Context, user User) error
		Delete(ctx context.Context, userID int64) error
	}
)

//easyjson:json
type (
	createUserRequest struct {
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
	}

	createUserResponse struct {
		ID int64 `json:"id"`
	}

	getUserResponse struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
	}

	updateUserRequest struct {
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
	}

	User struct {
		ID          int64
		Name        string
		DateOfBirth time.Time
	}

	userCacheKey struct {
		str string
		num int64
	}

	response struct {
		Success bool            `json:"success"`
		Body    json.RawMessage `json:"body,omitempty"`
		Errors  []errResponse   `json:"errors,omitempty"`
	}

	errResponse struct {
		Code    string `json:"code"`
		Context string `json:"context"`
	}
)

func (u User) Key() cachePkg.Key {
	return newUserCacheKey(u.ID)
}

func (u User) Validate() error {
	if l := utf8.RuneCountInString(u.Name); l > 40 || l < 4 {
		return validationError{errInvalidName}
	}

	if u.DateOfBirth.After(time.Now().AddDate(-14, 0, 0)) {
		return validationError{errInvalidDOB}
	}

	return nil
}

func newUserCacheKey(userID int64) *userCacheKey {
	return &userCacheKey{
		str: strconv.FormatInt(userID, 10),
		num: userID,
	}
}

func (u userCacheKey) String() string {
	return u.str
}

func (u userCacheKey) Int64() int64 {
	return u.num
}

func (c createUserRequest) toUser() (User, error) {
	t, err := time.Parse(dobFormat, c.DateOfBirth)
	if err != nil {
		return User{}, fmt.Errorf("time parse: %w", err)
	}

	return User{
		Name:        c.Name,
		DateOfBirth: t,
	}, nil
}

func (c updateUserRequest) toUser() (User, error) {
	t, err := time.Parse(dobFormat, c.DateOfBirth)
	if err != nil {
		return User{}, fmt.Errorf("time parse: %w", err)
	}

	return User{
		Name:        c.Name,
		DateOfBirth: t,
	}, nil
}
