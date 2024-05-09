package signatures

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/cloudflare/bn256"
)

// SerializeSignature converts a bn256.G1 point to a hex string.
func SerializeSignature(signature *bn256.G1) (string, error) {
	if signature == nil {
		return "", errors.New("signature cannot be nil")
	}

	data := signature.Marshal()
	return hex.EncodeToString(data), nil
}

// DeserializeSignature converts a hex string back into a bn256.G1 point.
func DeserializeSignature(encodedString string) (*bn256.G1, error) {
	data, err := hex.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}

	signature := new(bn256.G1)
	if _, err := signature.Unmarshal(data); err != nil {
		return nil, err
	}

	return signature, nil
}

// ExampleUsage Example usage of serialize and unserialize functionss
func ExampleUsage() {
	trapdoor, _, publicKey, g1, err := GenerateKeys()
	if err != nil {
		// Properly handle the error, depending on your application's needs
		panic("failed to generate keys: " + err.Error())
	}

	data := []byte("Example data for signing")
	signature := sign(data, g1, trapdoor)

	signatureStr, err := SerializeSignature(signature)
	if err != nil {
		panic("failed to serialize signature: " + err.Error())
	}

	// Output the string to console or log, depending on your needs
	fmt.Println("Serialized Signature:", signatureStr)

	recoveredSignature, err := DeserializeSignature(signatureStr)
	if err != nil {
		panic("failed to deserialize signature: " + err.Error())
	}

	// Verify the signature to ensure round-trip integrity
	if valid := verify(data, recoveredSignature, publicKey, trapdoor); !valid {
		panic("failed to verify signature")
	}

	fmt.Println("Signature verified successfully")
}
