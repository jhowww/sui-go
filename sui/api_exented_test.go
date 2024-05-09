package sui_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/howjmay/sui-go/models"
	"github.com/howjmay/sui-go/sui"
	"github.com/howjmay/sui-go/sui/conn"
	"github.com/howjmay/sui-go/sui_signer"
	"github.com/howjmay/sui-go/sui_types"
	"github.com/stretchr/testify/require"
)

func TestGetDynamicFieldObject(t *testing.T) {
	t.Skip("FIXME")
	api := sui.NewSuiClient(conn.TestnetEndpointUrl)
	parentObjectID, err := sui_types.SuiAddressFromHex("0x1719957d7a2bf9d72459ff0eab8e600cbb1991ef41ddd5b4a8c531035933d256")
	require.NoError(t, err)
	type args struct {
		ctx            context.Context
		parentObjectID *sui_types.ObjectID
		name           *sui_types.DynamicFieldName
	}
	tests := []struct {
		name    string
		args    args
		want    *models.SuiObjectResponse
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:            context.TODO(),
				parentObjectID: parentObjectID,
				name: &sui_types.DynamicFieldName{
					Type:  "address",
					Value: "0xf9ed7d8de1a6c44d703b64318a1cc687c324fdec35454281035a53ea3ba1a95a",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := api.GetDynamicFieldObject(tt.args.ctx, tt.args.parentObjectID, tt.args.name)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetDynamicFieldObject() error: %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%#v", got)
			},
		)
	}
}

func TestGetDynamicFields(t *testing.T) {
	api := sui.NewSuiClient(conn.DevnetEndpointUrl)
	parentObjectID, err := sui_types.SuiAddressFromHex("0x1719957d7a2bf9d72459ff0eab8e600cbb1991ef41ddd5b4a8c531035933d256")
	require.NoError(t, err)
	limit := uint(5)
	type args struct {
		ctx            context.Context
		parentObjectID *sui_types.ObjectID
		cursor         *sui_types.ObjectID
		limit          *uint
	}
	tests := []struct {
		name    string
		args    args
		want    *models.DynamicFieldPage
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:            context.TODO(),
				parentObjectID: parentObjectID,
				cursor:         nil,
				limit:          &limit,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := api.GetDynamicFields(tt.args.ctx, tt.args.parentObjectID, tt.args.cursor, tt.args.limit)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetDynamicFields() error: %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Log(got)
			},
		)
	}
}

func TestGetOwnedObjects(t *testing.T) {
	api := sui.NewSuiClient(conn.TestnetEndpointUrl)
	obj, err := sui_types.SuiAddressFromHex("0x2")
	require.NoError(t, err)
	query := models.SuiObjectResponseQuery{
		Filter: &models.SuiObjectDataFilter{
			Package: obj,
			// StructType: "0x2::coin::Coin<0x2::sui::SUI>",
		},
		Options: &models.SuiObjectDataOptions{
			ShowType: true,
		},
	}
	limit := uint(1)
	objs, err := api.GetOwnedObjects(context.Background(), sui_signer.TEST_ADDRESS, &query, nil, &limit)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(objs.Data), int(limit))
}

func TestQueryEvents(t *testing.T) {
	api := sui.NewSuiClient(conn.DevnetEndpointUrl)
	limit := uint(10)
	type args struct {
		ctx             context.Context
		query           *models.EventFilter
		cursor          *models.EventId
		limit           *uint
		descendingOrder bool
	}
	tests := []struct {
		name    string
		args    args
		want    *models.EventPage
		wantErr bool
	}{
		{
			name: "test for query events",
			args: args{
				ctx: context.TODO(),
				query: &models.EventFilter{
					Sender: sui_signer.TEST_ADDRESS,
				},
				cursor:          nil,
				limit:           &limit,
				descendingOrder: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := api.QueryEvents(
					tt.args.ctx,
					tt.args.query,
					tt.args.cursor,
					tt.args.limit,
					tt.args.descendingOrder,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("QueryEvents() error: %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Log(got)
			},
		)
	}
}

func TestQueryTransactionBlocks(t *testing.T) {
	api := sui.NewSuiClient(conn.DevnetEndpointUrl)
	limit := uint(10)
	type args struct {
		ctx             context.Context
		query           *models.SuiTransactionBlockResponseQuery
		cursor          *sui_types.TransactionDigest
		limit           *uint
		descendingOrder bool
	}
	tests := []struct {
		name    string
		args    args
		want    *models.TransactionBlocksPage
		wantErr bool
	}{
		{
			name: "test for queryTransactionBlocks",
			args: args{
				ctx: context.TODO(),
				query: &models.SuiTransactionBlockResponseQuery{
					Filter: &models.TransactionFilter{
						FromAddress: sui_signer.TEST_ADDRESS,
					},
					Options: &models.SuiTransactionBlockResponseOptions{
						ShowInput:   true,
						ShowEffects: true,
					},
				},
				cursor:          nil,
				limit:           &limit,
				descendingOrder: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := api.QueryTransactionBlocks(
					tt.args.ctx,
					tt.args.query,
					tt.args.cursor,
					tt.args.limit,
					tt.args.descendingOrder,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("QueryTransactionBlocks() error: %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%#v", got)
			},
		)
	}
}

func TestResolveNameServiceAddress(t *testing.T) {
	api := sui.NewSuiClient(conn.MainnetEndpointUrl)
	addr, err := api.ResolveNameServiceAddress(context.Background(), "2222.sui")
	require.NoError(t, err)
	require.Equal(t, "0x6174c5bd8ab9bf492e159a64e102de66429cfcde4fa883466db7b03af28b3ce9", addr.String())

	_, err = api.ResolveNameServiceAddress(context.Background(), "2222.suijjzzww")
	require.ErrorContains(t, err, "not found")
}

func TestResolveNameServiceNames(t *testing.T) {
	api := sui.NewSuiClient(conn.MainnetEndpointUrl)
	owner := AddressFromStrMust("0x57188743983628b3474648d8aa4a9ee8abebe8f6816243773d7e8ed4fd833a28")
	namePage, err := api.ResolveNameServiceNames(context.Background(), owner, nil, nil)
	require.NoError(t, err)
	require.NotEmpty(t, namePage.Data)
	t.Log(namePage.Data)

	owner = AddressFromStrMust("0x57188743983628b3474648d8aa4a9ee8abebe8f681")
	namePage, err = api.ResolveNameServiceNames(context.Background(), owner, nil, nil)
	require.NoError(t, err)
	require.Empty(t, namePage.Data)
}

func TestSubscribeEvent(t *testing.T) {
	t.Skip("passed at local side, but failed on GitHub")
	api := sui.NewSuiWebsocketClient(conn.MainnetWebsocketEndpointUrl)

	type args struct {
		ctx      context.Context
		filter   *models.EventFilter
		resultCh chan models.SuiEvent
	}
	tests := []struct {
		name    string
		args    args
		want    *models.EventPage
		wantErr bool
	}{
		{
			name: "test for filter events",
			args: args{
				ctx: context.TODO(),
				filter: &models.EventFilter{
					Package: sui_types.MustPackageIDFromHex("0x000000000000000000000000000000000000000000000000000000000000dee9"),
				},
				resultCh: make(chan models.SuiEvent),
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := api.SubscribeEvent(
					tt.args.ctx,
					tt.args.filter,
					tt.args.resultCh,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("SubscribeEvent() error: %v, wantErr %v", err, tt.wantErr)
					return
				}
				cnt := 0
				for results := range tt.args.resultCh {
					fmt.Println("results: ", results)
					// FIXME we need to check finite number request in details
					cnt++
					if cnt > 3 {
						break
					}
				}
			},
		)
	}
}
