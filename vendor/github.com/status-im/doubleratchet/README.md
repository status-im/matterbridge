# doubleratchet

[![Go Report Card](https://goreportcard.com/badge/github.com/status-im/doubleratchet)](https://goreportcard.com/report/github.com/status-im/doubleratchet)
[![Build Status](https://travis-ci.org/status-im/doubleratchet.svg?branch=master)](https://travis-ci.org/status-im/doubleratchet)
[![Coverage Status](https://coveralls.io/repos/github/status-im/doubleratchet/badge.svg?branch=master)](https://coveralls.io/github/status-im/doubleratchet?branch=master)
[![GoDoc](https://godoc.org/github.com/status-im/doubleratchet?status.svg)](https://godoc.org/github.com/status-im/doubleratchet)

[The Double Ratchet Algorithm](https://whispersystems.org/docs/specifications/doubleratchet) is used
by two parties to exchange encrypted messages based on a shared secret key. Typically the parties
will use some key agreement protocol (such as X3DH) to agree on the shared secret key.
Following this, the parties will use the Double Ratchet to send and receive encrypted messages.

The parties derive new keys for every Double Ratchet message so that earlier keys cannot be calculated
from later ones. The parties also send Diffie-Hellman public values attached to their messages.
The results of Diffie-Hellman calculations are mixed into the derived keys so that later keys cannot
be calculated from earlier ones. These properties gives some protection to earlier or later encrypted 
messages in case of a compromise of a party's keys.

## Project status

The library is in beta version and ready for integration into production projects with care.
Let me know if you face any problems or have any questions or suggestions.

## Implementation notes

### The Double Ratchet logic

1. No more than 1000 messages can be skipped in a single chain.
1. Skipped messages from a single ratchet step are deleted after 100 ratchet steps.
1. Both parties' sending and receiving chains are initialized with the shared key so that both
of them could message each other from the very beginning.

### Cryptographic primitives 

1. **GENERATE_DH():** Curve25519
1. **KDF_RK(rk, dh_out):** HKDF with SHA-256
1. **KDF_CK(ck):** HMAC with SHA-256 and constant inputs
1. **ENCRYPT(mk, pt, associated_data):** AES-256-CTR with HMAC-SHA-256 and IV derived alongside an encryption key

## Installation

    go get github.com/status-im/doubleratchet

then `cd` into the project directory and install dependencies:

    glide up
    
If `glide` is not installed, [install it](https://github.com/Masterminds/glide).

## Usage

### Basic usage example

```go
package main

import (
	"fmt"
	"log"

	"github.com/status-im/doubleratchet"
)

func main() {
	// The shared key both parties have already agreed upon before the communication.
	sk := [32]byte{
		0xeb, 0x8, 0x10, 0x7c, 0x33, 0x54, 0x0, 0x20,
		0xe9, 0x4f, 0x6c, 0x84, 0xe4, 0x39, 0x50, 0x5a,
		0x2f, 0x60, 0xbe, 0x81, 0xa, 0x78, 0x8b, 0xeb,
		0x1e, 0x2c, 0x9, 0x8d, 0x4b, 0x4d, 0xc1, 0x40,
	}

	// Diffie-Hellman key pair generated by one of the parties during key exchange or
	// by any other means. The public key MUST be sent to another party for initialization
	// before the communication begins.
	keyPair, err := doubleratchet.DefaultCrypto{}.GenerateDH()
	if err != nil {
		log.Fatal(err)
	}

	// Bob MUST be created with the shared secret and a DH key pair.
	bob, err := doubleratchet.New([]byte("bob-session-id"), sk, keyPair, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Alice MUST be created with the shared secret and Bob's public key.
	alice, err := doubleratchet.NewWithRemoteKey([]byte("alice-session-id"), sk, keyPair.PublicKey(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Alice can now encrypt messages under the Double Ratchet session.
	m, err := alice.RatchetEncrypt([]byte("Hi Bob!"), nil)

	if err != nil {
		log.Fatal(err)
	}

	// Which Bob can decrypt.
	plaintext, err := bob.RatchetDecrypt(m, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(plaintext))
}
```

### Options

Additional options can be passed to constructors to customize the algorithm behavior:

```go
doubleratchet.New(
    sk, keyPair,
    
    // Your own cryptography supplement implementing doubleratchet.Crypto.
    WithCrypto(c),
    
    // Custom storage for skipped keys implementing doubleratchet.KeysStorage.
    WithKeysStorage(ks),
    
    // The maximum number of skipped keys. Error will be raised in an attempt to store more keys
    // in a single chain while decrypting.
    WithMaxSkip(1200),
    
    // The number of Diffie-Hellman ratchet steps skipped keys will be stored.
    WithMaxKeep(90),
)
```

## License

MIT