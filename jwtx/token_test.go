package jwtx

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

type AccessTokenTest struct {
	UserID   string
	UserName string
}

func TestNewJwtToken(t *testing.T) {
	a := NewJwtToken[AccessTokenTest]()
	require.Equal(t, "*jwtx.jwtToken[jwtx.AccessTokenTest]", reflect.TypeOf(a).String())
}

func Test_jwtToken_GenerateToken(t1 *testing.T) {
	a := NewJwtToken[AccessTokenTest](
		WithIssuer("test"),
		WithAudience("test"),
		WithExpired(time.Hour),
		WithSecretKey("test"),
	)
	token, err := a.GenerateToken(AccessTokenTest{UserID: "test", UserName: "test"})
	require.NoError(t1, err)

	fmt.Println(token)
}

func TestName(t *testing.T) {
	p := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiYXVkIjpbInRlc3QiXSwiZXhwIjoxNjUxNjcyODQ3LCJpYXQiOjE2NTE2NjkyNDcsIlBheWxvYWQiOnsiVXNlcklEIjoidGVzdCIsIlVzZXJOYW1lIjoidGVzdCJ9fQ.rZJdL32R2QYjo1PNI2UqJEd_suCboCaC38NQjJ5x0Gs"
	s := base64.URLEncoding.EncodeToString([]byte(p))
	
	fmt.Println(s)
}
