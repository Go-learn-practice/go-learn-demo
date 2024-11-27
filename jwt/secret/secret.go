package secret

const KEYPATH = "keys"

type OutSecret struct {
	Secret         string
	PublicKeyFile  string
	PrivateKeyFile string
}

type Secret interface {
	Generate() (*OutSecret, error)
}
