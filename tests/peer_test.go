package peer

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gladiusio/gladius-controld/pkg/blockchain"
	"github.com/gladiusio/gladius-controld/pkg/p2p/message"
	"github.com/gladiusio/gladius-controld/pkg/p2p/peer"
	"github.com/spf13/viper"
)

const (
	numOfPeers = 20
)

func TestMain(m *testing.M) {
	// Setup
	os.Setenv("GLADIUSBASE", ".") // So that we don't get conflicts with our install
	createWallets()               // Create the same wallets for every run
	viper.SetDefault("P2P.VerifyOverride", true)

	// Run the tests
	retCode := m.Run()

	// Teardown
	deleteWallets() // Deletes all wallets in the wallets dir

	// Exit with the test status
	os.Exit(retCode)
}

func createWallets() {
	for i := 0; i < numOfPeers; i++ {
		// Create a new directory for the wallet
		walletDir := fmt.Sprintf("./wallets/g%d", i)
		err := os.MkdirAll(walletDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func deleteWallets() {
	for i := 0; i < numOfPeers; i++ {
		err := os.RemoveAll("./wallets/g" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func buildPeers() []*peer.Peer {
	peers := make([]*peer.Peer, 0)
	for i := 0; i < numOfPeers; i++ {
		walletDir := fmt.Sprintf("./wallets/g%d", i)
		// New keystore pointing to that directory
		ks := keystore.NewKeyStore(
			walletDir,
			keystore.LightScryptN,
			keystore.LightScryptP,
		)

		// Create account manager pointing to that keystore
		ga := blockchain.NewGladiusAccountManagerCustomKeystore(ks)
		ga.CreateAccount("password")
		// Kinda hacky but lets us configure each peer
		viper.SetDefault("P2P.BindAddress", "localhost")
		viper.SetDefault("P2P.BindPort", 7946+i)
		peers = append(peers, peer.New(ga))
	}
	time.Sleep(1 * time.Second) // Sleep to let all the peers start
	return peers
}

func killPeers(peers []*peer.Peer) {
	time.Sleep(1 * time.Second)
	for _, p := range peers {
		p.Stop()
	}
}

func buildNetwork(peers []*peer.Peer, t *testing.T) {
	// Let the first node be a seed node
	for i := 1; i < numOfPeers; i++ {
		err := peers[i].Join([]string{fmt.Sprintf("kcp://localhost:%d", 7946)})
		if err != nil {
			t.Errorf("node %d couldn't join network: error was: %s", i, err.Error())
		}
	}
}

func signTestMessage(peers []*peer.Peer, t *testing.T) {
	// Go through each peer and sign and post an update message
	for i := 0; i < numOfPeers; i++ {
		peers[i].UnlockWallet("password")
		sm, err := peers[i].SignMessage(message.New(
			[]byte(`{"node" : {"ip_address": "localhost"}}`),
		))
		if err != nil {
			t.Errorf("node %d couldn't sign message: error was: %s", i, err.Error())
			t.Fail()
		}

		err = peers[i].UpdateAndPushState(sm)
		if err != nil {
			t.Errorf("node %d couldn't update state: error was: %s", i, err.Error())
			t.Fail()
		}
	}
}

func TestSuccessfulJoin(t *testing.T) {
	peers := buildPeers()
	defer killPeers(peers)

	buildNetwork(peers, t)
}

func TestCorrectNumberOfNodesInState(t *testing.T) {
	peers := buildPeers()
	defer killPeers(peers)

	buildNetwork(peers, t)      // Build the network (let nodes join)
	time.Sleep(1 * time.Second) // Sleep to let the dht do its magic

	signTestMessage(peers, t)          // Sign a message so we have some state
	time.Sleep(100 * time.Millisecond) // Wait for state to update

	for i := 1; i < numOfPeers; i++ {
		numOfNodesInState := len(peers[i].GetState().NodeDataMap)
		if numOfNodesInState != numOfPeers {
			t.Errorf("there were %d nodes in state, expected %d",
				numOfNodesInState,
				numOfPeers,
			)
		}
	}

}

func TestStateEquality(t *testing.T) {
	peers := buildPeers()
	defer killPeers(peers)

	buildNetwork(peers, t)      // Build the network (let nodes join)
	time.Sleep(1 * time.Second) // Sleep to let the dht do its magic

	signTestMessage(peers, t)          // Sign a message so we have some state
	time.Sleep(100 * time.Millisecond) // Wait for state to update

	for i := 0; i < numOfPeers; i++ {
		for i2 := 0; i2 < numOfPeers; i2++ {
			p1 := peers[i].GetState().NodeDataMap
			p2 := peers[i2].GetState().NodeDataMap

			if !reflect.DeepEqual(p1, p2) {
				t.FailNow()
			}
		}

	}
}
