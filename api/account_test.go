package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/imran4u/simple-bank/db/mock"
	db "github.com/imran4u/simple-bank/db/sqlc"
	"github.com/imran4u/simple-bank/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

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

}

func randomAccount() db.Account {
	return db.Account{
		ID:        util.RandomInt(1, 1000),
		Owner:     util.RandomName(),
		Currency:  util.RandomCurrency(),
		Balance:   util.RandomAmount(),
		CreatedAt: time.Now(),
	}
}
