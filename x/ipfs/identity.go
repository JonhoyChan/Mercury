package ipfs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/ipfs/go-ipfs-config"
)

func init() {
	peer.AdvancedEnableInlining = false
}

// PubKey wraps IPFS's underlying PubKey dependency
type PubKey interface {
	crypto.PubKey
}

func IdentityFromKey(privkey []byte) (config.Identity, error) {
	ident := config.Identity{}
	sk, err := crypto.UnmarshalPrivateKey(privkey)
	if err != nil {
		return ident, err
	}
	skbytes, err := sk.Bytes()
	if err != nil {
		return ident, err
	}
	ident.PrivKey = base64.StdEncoding.EncodeToString(skbytes)

	id, err := peer.IDFromPublicKey(sk.GetPublic())
	if err != nil {
		return ident, err
	}
	ident.PeerID = id.Pretty()
	return ident, nil
}

func IdentityKeyFromSeed(seed []byte, bits int) ([]byte, error) {
	hm := hmac.New(sha256.New, []byte("Mercury seed"))
	_, err := hm.Write(seed)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(hm.Sum(nil))
	sk, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, bits, reader)
	if err != nil {
		return nil, err
	}
	encodedKey, err := sk.Bytes()
	if err != nil {
		return nil, err
	}
	return encodedKey, nil
}
