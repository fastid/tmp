package services

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCrypt(t *testing.T) {
	cfg, _ := config.New("../../configs/fastid.yml")

	// Logger
	log := logger.New(cfg)

	crypt := NewCryptService(cfg, log)

	cipher, err := crypt.GenerateCipher()
	require.NoError(t, err)
	require.NotEmpty(t, cipher)
}

func TestCryptEncryptDecryptB64(t *testing.T) {
	cfg, _ := config.New("../../configs/fastid.yml")

	// Logger
	log := logger.New(cfg)

	crypt := NewCryptService(cfg, log)
	secretKey, err := crypt.GenerateCipher()
	require.NoError(t, err)
	crypt.SecretKey(secretKey)

	encrypted, err := crypt.EncryptBase64("Hello")
	require.NoError(t, err)
	require.NotEmpty(t, encrypted)

	result, err := crypt.DecryptBase64(encrypted)
	require.NoError(t, err)
	require.Equal(t, result, "Hello")
}

func TestCryptEncryptDecrypt(t *testing.T) {
	cfg, _ := config.New("../../configs/fastid.yml")

	// Logger
	log := logger.New(cfg)

	crypt := NewCryptService(cfg, log)
	secretKey, err := crypt.GenerateCipher()
	require.NoError(t, err)
	crypt.SecretKey(secretKey)

	encrypted, err := crypt.Encrypt("Hello")
	require.NoError(t, err)
	require.NotEmpty(t, encrypted)

	result, err := crypt.Decrypt(encrypted)
	require.NoError(t, err)
	require.Equal(t, result, "Hello")
}
