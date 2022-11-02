package api

import (
	"testing"

	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/JILSE7/simplebank/utils"
)

func TestGetAccountAPI(t *testing.T) {
	/*
		 	account := randomAccount()

			ctrl := gomock.NewController(t)

			// store := mockdb.NewMockStore(ctrl)
			defer ctrl.Finish()

			// build stubs
			store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(account, nil)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%d", account.ID)

			request, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			// check response
			require.Equal(t, http.StatusOK, recorder.Code)
	*/
}

func randomAccount() db.Account {
	return db.Account{
		ID:       utils.RandomInt(1, 14),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

}
