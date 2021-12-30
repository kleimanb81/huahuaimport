package signer

// func TestQueryAccountNumber(t *testing.T) {
// 	t.Skip("no local connection to app and no funds in wallet")
// 	const accName = "test"
// 	testRing := testutil.GenerateKeyring(t, accName)

// 	encCfg := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

// 	k := NewKeyringSigner(testRing, encCfg, accName)

// 	RPCAddress := "127.0.0.1:9090"

// 	rpcClient, err := grpc.Dial(RPCAddress, grpc.WithInsecure())
// 	require.NoError(t, err)
// 	err = k.QueryAccountNumber(context.TODO(), rpcClient)
// 	require.NoError(t, err)
// }
