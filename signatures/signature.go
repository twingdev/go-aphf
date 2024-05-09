package signatures

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/cloudflare/bn256"
	"github.com/zeebo/blake3"
	"math/big"
)

func GenerateKeys() (*big.Int, *big.Int, *bn256.G2, *bn256.G1, error) {
	privateKey, g1, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	publicKey := new(bn256.G2).ScalarMult(new(bn256.G2).ScalarBaseMult(big.NewInt(1)), privateKey)

	trapdoor, _ := rand.Int(rand.Reader, bn256.Order)

	return trapdoor, privateKey, publicKey, g1, nil

}

func sign(data []byte, privateKey *bn256.G1, trapdoor *big.Int) *bn256.G1 {
	hasher := blake3.New()
	hasher.Write(data) // Base data hashing

	// Influence the hash with the trapdoor. Example: combine data hash with trapdoor hash
	trapdoorData := trapdoor.Bytes()
	hasher.Write(trapdoorData) // Use trapdoor to influence the hash

	dataHash := new(big.Int).SetBytes(hasher.Sum(nil))          // Convert hash to big.Int
	signature := new(bn256.G1).ScalarMult(privateKey, dataHash) // Sign using the modified hash

	return signature
}

func verify(data []byte, signature *bn256.G1, publicKey *bn256.G2, trapdoor *big.Int) bool {
	hasher := blake3.New()
	hasher.Write(data)

	// Optionally use the trapdoor if provided for verification
	if trapdoor != nil {
		trapdoorData := trapdoor.Bytes()
		hasher.Write(trapdoorData)
	}

	dataHash := new(big.Int).SetBytes(hasher.Sum(nil))

	hashPoint := new(bn256.G1).ScalarBaseMult(dataHash)
	pairSig := bn256.Pair(signature, new(bn256.G2).ScalarBaseMult(big.NewInt(1)))
	pairData := bn256.Pair(hashPoint, publicKey)

	return pairSig.String() == pairData.String()
}

func NewSignature(data []byte, privateKey *bn256.G1, trapdoor *big.Int, publicKey *bn256.G2) (*bn256.G1, *bn256.G2, *big.Int) {

	signature := sign(data, privateKey, trapdoor)
	fmt.Println("Signature:", signature)
	return signature, publicKey, trapdoor
}

func VerifySignature(data []byte, signature *bn256.G1, publicKey *bn256.G2, trapdoor *big.Int) bool {
	valid := verify(data, signature, publicKey, trapdoor)
	fmt.Println("Signature valid:", valid)
	return valid
}

func SignatureToString(signature *bn256.G1) (string, error) {
	// Marshal the signature into a byte slice
	data := signature.Marshal()

	// Encode the byte slice to a Base64 string
	encodedString := base64.StdEncoding.EncodeToString(data)
	return encodedString, nil
}

func StringToSignature(encodedString string) (*bn256.G1, error) {
	// Decode the Base64 string to a byte slice
	data, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return nil, err
	}

	// Unmarshal the byte slice into a G1 point
	signature := new(bn256.G1)
	if _, err := signature.Unmarshal(data); err != nil {
		return nil, err
	}

	return signature, nil
}
