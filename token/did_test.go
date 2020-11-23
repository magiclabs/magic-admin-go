package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDIDToken = "WyIweGFhNTBiZTcwNzI5Y2E3MDViYTdjOGQwMDE4NWM2ZjJkYTQ3OWQwZm" +
	"NkZTUzMTFjYTRjZTViMWJhNzE1YzhhNzIxYzVmMTk0ODQzNGY5NmZmNTc3ZDdiMmI2YWQ4MmQ" +
	"zZGQ1YTI0NTdmZTY5OThiMTM3ZWQ5YmMwOGQzNmU1NDljMWIiLCJ7XCJpYXRcIjoxNTg2NzY0" +
	"MjcwLFwiZXh0XCI6MTExNzM1Mjg1MDAsXCJpc3NcIjpcImRpZDpldGhyOjB4NEI3M0M1ODM3M" +
	"EFFZmNFZjg2QTYwMjFhZkNEZTU2NzM1MTEzNzZCMlwiLFwic3ViXCI6XCJOanJBNTNTY1E4SV" +
	"Y4ME5Kbng0dDNTaGk5LWtGZkY1cWF2RDJWcjBkMWRjPVwiLFwiYXVkXCI6XCJkaWQ6bWFnaWM" +
	"6NzMxODQ4Y2MtMDg0ZS00MWZmLWJiZGYtN2YxMDM4MTdlYTZiXCIsXCJuYmZcIjoxNTg2NzY0" +
	"MjcwLFwidGlkXCI6XCJlYmNjODgwYS1mZmM5LTQzNzUtODRhZS0xNTRjY2Q1Yzc0NmRcIixcI" +
	"mFkZFwiOlwiMHg4NGQ2ODM5MjY4YTFhZjkxMTFmZGVjY2QzOTZmMzAzODA1ZGNhMmJjMDM0NT" +
	"BiN2ViMTE2ZTJmNWZjOGM1YTcyMmQxZmI5YWYyMzNhYTczYzVjMTcwODM5Y2U1YWQ4MTQxYjl" +
	"iNDY0MzM4MDk4MmRhNGJmYmIwYjExMjg0OTg4ZjFiXCJ9Il0="

func TestDIDTokenDecode(t *testing.T) {
	issuer := "did:ethr:0x4B73C58370AEfcEf86A6021afCDe5673511376B2"
	proof := "0xaa50be70729ca705ba7c8d00185c6f2da479d0fcde5311ca4ce5b1ba715c8a721c5" +
		"f1948434f96ff577d7b2b6ad82d3dd5a2457fe6998b137ed9bc08d36e549c1b"

	claim := Claim {
		Iat: 1586764270,
		Ext: 11173528500,
		Nbf: 1586764270,
		Iss: "did:ethr:0x4B73C58370AEfcEf86A6021afCDe5673511376B2",
		Sub: "NjrA53ScQ8IV80NJnx4t3Shi9-kFfF5qavD2Vr0d1dc=",
		Aud: "did:magic:731848cc-084e-41ff-bbdf-7f103817ea6b",
		Tid: "ebcc880a-ffc9-4375-84ae-154ccd5c746d",
		Add: "0x84d6839268a1af9111fdeccd396f303805dca2bc03450b7eb116e2f5fc8c5a722" +
			"d1fb9af233aa73c5c170839ce5ad8141b9b4643380982da4bfbb0b11284988f1b",
	}

	token, err := NewToken(testDIDToken)
	require.NoError(t, err, "can't create new token")

	assert.Equal(t, claim, token.GetClaim())
	assert.Equal(t, proof, token.GetProof())
	assert.Equal(t, issuer, token.GetIssuer())

	assert.NoError(t, token.Validate(), "token is not valid")
}
