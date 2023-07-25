package auth

import (
	"testing"
)

func TestLoginWithAuth(t *testing.T) {
	gotAuth, err := LoginWithAuth("username", "password") // !! DO NOT PUSH YOUR PASSWORD
	if err != nil {
		// if failed here 'invalid character '<' looking for beginning of value' check username and password
		t.Errorf("LoginWithAuth() error = %v", err)
		return
	} else {
		t.Logf("LoginWithAuth() SUCCESS.")
	}

	err = gotAuth.Login()
	if err != nil {
		t.Errorf("Login() error = %v", err)
		return
	} else {
		t.Logf("Login() SUCCESS.")
	}

	err = gotAuth.Refresh()
	if err != nil {
		t.Errorf("Refresh() error = %v", err)
		return
	} else {
		t.Logf("Refresh() SUCCESS.")
	}

	exp := gotAuth.ExpiredAt()
	if exp < 80000 {
		t.Errorf("ExpiredAt() error = exp < 80000")
		return
	} else {
		t.Logf("ExpiredAt() SUCCESS.")
	}
}
