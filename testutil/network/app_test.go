package network

// import (
// 	"context"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/ChihuahuaChain/chihuahua/app"
// 	"github.com/ChihuahuaChain/chihuahua/testutil/signer"
// 	"github.com/stretchr/testify/require"
// 	"github.com/tendermint/spm/cosmoscmd"
// 	"google.golang.org/grpc"
// )

// func TestIntegration(t *testing.T) {
// 	const testAcc = "test-account"

// 	// create a network and a keyring signer
// 	net, kr := New(t, DefaultConfig(), testAcc)

// 	time.Sleep(3 * time.Second)

// 	// check liveness before testing
// 	height, err := net.LatestHeight()
// 	require.NoError(t, err)
// 	require.Greater(t, height, int64(0))

// 	// use the first validator as the node to access the network
// 	val := net.Validators[0]

// 	// create the signer that we will use to create and sign a message
// 	sgnr := signer.NewKeyringSigner(
// 		kr,
// 		cosmoscmd.MakeEncodingConfig(app.ModuleBasics),
// 		testAcc,
// 		val.Ctx.Config.ChainID(),
// 	)

// 	port := strings.Replace(val.RPCAddress, "tcp://0.0.0.0:", "", 1)

// 	conn, err := grpc.Dial("127.0.0.1:"+port, grpc.WithInsecure())
// 	require.NoError(t, err)

// 	err = sgnr.QueryAccountNumber(context.TODO(), conn)
// 	require.NoError(t, err)
// }
