package main

import (
	"fmt"
	"github.com/twingdev/go-aphf/signatures"
)

func main() {
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
	if err != nil {
		fmt.Println("Error converting signature to string:", err)
		return
	}
	fmt.Println("Signature as string:", signatureStr)

	// Convert string back to signature
	recoveredSignature, err := signatures.StringToSignature(signatureStr)
	if err != nil {
		fmt.Println("Error converting string to signature:", err)
		return
	}

	// Verify the signature
	if signatures.VerifySignature(data, recoveredSignature, publicKey, trapdoor) {
		fmt.Println("Signature verified successfully")
	} else {
		fmt.Println("Failed to verify signature")
	}
}
