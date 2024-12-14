package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/imran4u/simple-bank/db/mock"
	db "github.com/imran4u/simple-bank/db/sqlc"
	"github.com/imran4u/simple-bank/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomName(),
		Currency: util.RandomCurrency(),
		Balance:  util.RandomAmount(),
		// CreatedAt: time.Now(),   //ignore time case other it will create problem in equal.
	}
}

func requiredBodyMachAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, gotAccount, account)

}

func TestGetAccount(t *testing.T) {
	accout := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // this is important

	store := mockdb.NewMockStore(ctrl)
	//create stub
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(accout.ID)).
		Times(1).
		Return(accout, nil)

	//start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/account/%d", accout.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)
	require.Equal(t, http.StatusOK, recorder.Code)
	requiredBodyMachAccount(t, recorder.Body, accout)

}

// Table drivern test set to cover all possible senario

// func TestGetAccountAppPossiblies(t *testing.T){
// 	account := randomAccount()

// 	//Ananomous struct
// 	testCase := []struct{
// 		name string,
// 		accountId int64,
// 		buildStub func(store *mockdb.MockStore),
// 		// checkResponse func(t *testing.T, )

// 	}{

// 	}

// }
