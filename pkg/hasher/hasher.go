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

type Client struct {
	Progress io.Writer
}

func (c *Client) hash(h hash.Hash, r io.Reader) (string, error) {
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

func (c *Client) Sha1(r io.Reader) (string, error) {
	h := sha1.New()

	return c.hash(h, r)
}

func (c *Client) Sha224(r io.Reader) (string, error) {
	h := sha256.New224()

	return c.hash(h, r)
}

func (c *Client) Sha256(r io.Reader) (string, error) {
	h := sha256.New()

	return c.hash(h, r)
}

func (c *Client) Sha384(r io.Reader) (string, error) {
	h := sha512.New384()

	return c.hash(h, r)
}

func (c *Client) Sha512(r io.Reader) (string, error) {
	h := sha512.New()

	return c.hash(h, r)
}

func (c *Client) Md5(r io.Reader) (string, error) {
	h := md5.New()

	return c.hash(h, r)
}
