package usage

import "os"

func KeyRPCFromEnv() string {
	key := os.Getenv("TEST_RPC_MINECRAFT_KEY")
	if key == "" {
		panic("key is empty")
	}
	return key
}
