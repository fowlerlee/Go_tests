package proj3

import (
	
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"example/test/chatroom"
	"fmt"
	"net/http"

	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func MakePeerSwarm() {
	node, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	var chatty chatroom.ChatServer
	http.ListenAndServe(":8080", chatty)

	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	m := "the hash string for bytes"
	hash := sha256.Sum256([]byte(m))
	hashSign, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", hashSign)

	verified := ecdsa.VerifyASN1(&privKey.PublicKey, hash[:], hashSign)
	fmt.Println("signature verified: ", verified)

	// start a libp2p node with default settings
	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/2000"))
	if err != nil {
		panic(err)
	}

	// print the node's listening addresses
	fmt.Println("Listen addresses:", node.Addrs())

	pubs, err := pubsub.NewGossipSub(ctx, host)
	if err != nil {
		panic(err)
	}
	fmt.Println("the pubsub is: ‰v", pubs)

	// shut the node down
	if err := node.Close(); err != nil {
		panic(err)
	}
}

