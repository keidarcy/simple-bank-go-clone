package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/keidarcy/simple-bank/db/mock"
	db "github.com/keidarcy/simple-bank/db/sqlc"
	"github.com/keidarcy/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	store := mockdb.NewMockStore(ctrl)

	account1 := randomAccount()
	account2 := randomAccount()

	amount := util.RandomInt(10, 1000)

	account1.Currency = "USD"
	account2.Currency = "USD"

	store.
		EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
		Times(1).
		Return(account1, nil)

	store.
		EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account2.ID)).
		Times(1).
		Return(account2, nil)

	arg := db.TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	}
	store.
		EXPECT().
		TransferTx(gomock.Any(), gomock.Eq(arg)).
		Times(1).
		Return(db.TransferTxResult{
			FromAccount: account1,
			ToAccount:   account2,
		}, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := "/transfer"
	body := transferRequest{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
		Currency:      "USD",
	}
	requestBody, err := json.Marshal(body)
	require.NoError(t, err)

	request, err := http.NewRequest("POST", url, bytes.NewReader(requestBody))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	data, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	var gotTransfer db.TransferTxResult

	err = json.Unmarshal(data, &gotTransfer)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, recorder.Code)
	require.Equal(t, gotTransfer.FromAccount.ID, account1.ID)
	require.Equal(t, gotTransfer.ToAccount.ID, account2.ID)

}
