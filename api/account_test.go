package api

import (
	"bytes"
	"database/sql"
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
	//Check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requiredBodyMachAccount(t, recorder.Body, accout) // response body will install in recorder.Body field.

}

// Table drivern test set to cover all possible senario

func TestGetAccountAppPossiblies(t *testing.T) {
	account := randomAccount()

	//Ananomous struct
	testCase := []struct {
		name          string
		accountId     int64
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "Ok",
			accountId: account.ID,
			buildStub: func(store *mockdb.MockStore) {
				//create stub
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//Check response
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMachAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "Not found",
			accountId: account.ID,
			buildStub: func(store *mockdb.MockStore) {
				//create stub
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//Check response
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "Internal server error",
			accountId: account.ID,
			buildStub: func(store *mockdb.MockStore) {
				//create stub
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				//Check response
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // this is important

			store := mockdb.NewMockStore(ctrl)
			//create stub
			tc.buildStub(store)

			//start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/account/%d", tc.accountId)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			//check Response
			tc.checkResponse(t, recorder)

		})

	}

}
