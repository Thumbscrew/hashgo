package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
)

// Client can be initialized with an io.Writer to track progress. When it is set,
// an io.MultiWriter will be used to write to it.
type Client struct {
	Progress io.Writer
}

// Hash returns the checksum of the data from the provided io.Reader.
func (c *Client) Hash(h hash.Hash, r io.Reader) (string, error) {
	var writer io.Writer
	if c.Progress != nil {
		writer = io.MultiWriter(h, c.Progress)
	} else {
		writer = h
	}

	if _, err := io.Copy(writer, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Sha1 returns the SHA1 checksum of the data from the provided io.Reader.
func (c *Client) Sha1(r io.Reader) (string, error) {
	h := sha1.New()

	return c.Hash(h, r)
}

// Sha224 returns the SHA224 checksum of the data from the provided io.Reader.
func (c *Client) Sha224(r io.Reader) (string, error) {
	h := sha256.New224()

	return c.Hash(h, r)
}

// Sha256 returns the SHA256 checksum of the data from the provided io.Reader.
func (c *Client) Sha256(r io.Reader) (string, error) {
	h := sha256.New()

	return c.Hash(h, r)
}

// Sha384 returns the SHA384 checksum of the data from the provided io.Reader.
func (c *Client) Sha384(r io.Reader) (string, error) {
	h := sha512.New384()

	return c.Hash(h, r)
}

// Sha512 returns the SHA512 checksum of the data from the provided io.Reader.
func (c *Client) Sha512(r io.Reader) (string, error) {
	h := sha512.New()

	return c.Hash(h, r)
}

// Md5 returns the MD5 checksum of the data from the provided io.Reader.
func (c *Client) Md5(r io.Reader) (string, error) {
	h := md5.New()

	return c.Hash(h, r)
}
