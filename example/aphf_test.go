package main

import (
	"fmt"
	"github.com/twingdev/go-aphf/signatures"
	"github.com/zeebo/assert"
	"testing"
)

func TestSignature(t *testing.T) {
	// Example usage within your application
	trapdoor, _, publicKey, g1, err := signatures.GenerateKeys()
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	data := []byte("Example data for signing")
	signature, _, _ := signatures.NewSignature(data, g1, trapdoor, publicKey)

	// Convert signature to string
	signatureStr, err := signatures.SignatureToString(signature)
	assert.NoError(t, err)
	fmt.Println("Signature as string:", signatureStr)

	// Convert string back to signature
	recoveredSignature, err := signatures.StringToSignature(signatureStr)
	assert.NoError(t, err)

	// Verify the signature
	if signatures.VerifySignature(data, recoveredSignature, publicKey, trapdoor) {
		fmt.Println("Signature verified successfully")
	} else {
		fmt.Println("Failed to verify signature")
	}
}

func TestSerializer(t *testing.T) {
	trapdoor, _, publicKey, g1, err := signatures.GenerateKeys()
	if err != nil {
		// Properly handle the error, depending on your application's needs
		panic("failed to generate keys: " + err.Error())
	}

	data := []byte("Example data for signing")
	signature, _, _ := signatures.NewSignature(data, g1, trapdoor, publicKey)

	signatureStr, err := signatures.SerializeSignature(signature)
	if err != nil {
		panic("failed to serialize signature: " + err.Error())
	}

	// Output the string to console or log, depending on your needs
	fmt.Println("Serialized Signature:", signatureStr)

	recoveredSignature, err := signatures.DeserializeSignature(signatureStr)
	if err != nil {
		panic("failed to deserialize signature: " + err.Error())
	}

	// Verify the signature to ensure round-trip integrity
	if valid := signatures.VerifySignature(data, recoveredSignature, publicKey, trapdoor); !valid {
		panic("failed to verify signature")
	}

	fmt.Println("Signature verified successfully")
}
