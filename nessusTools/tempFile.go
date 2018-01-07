package nessusTools

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
)

//TempFileName generates a random file name
func TempFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	ex, _ := os.Executable()
	return filepath.Join(filepath.Dir(ex), hex.EncodeToString(randBytes)+".nessus")
}
