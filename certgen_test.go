// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPublicKey(t *testing.T) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ed25519 key: %v", err)
	}

	res := publicKey(privKey)
	edRes, ok := res.(ed25519.PublicKey)
	if !ok || !edRes.Equal(pubKey) {
		t.Errorf("publicKey(ed25519) did not return expected public key")
	}

	ecdsaPriv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ecdsa key: %v", err)
	}
	ecdsaRes := publicKey(ecdsaPriv)
	ecdsaPub, ok := ecdsaRes.(*ecdsa.PublicKey)
	if !ok || !ecdsaPub.Equal(&ecdsaPriv.PublicKey) {
		t.Errorf("publicKey(ecdsa) did not return expected public key")
	}

	if publicKey("invalid") != nil {
		t.Errorf("publicKey(invalid) expected nil, got non-nil")
	}
}

func TestCertificateCreation(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "certgen_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Certgen Test"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}

	pkey := publicKey(priv)
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, pkey, priv)
	if err != nil {
		t.Fatalf("Failed to create certificate: %v", err)
	}

	certFile := filepath.Join(tempDir, "public.crt")
	certOut, err := os.Create(certFile)
	if err != nil {
		t.Fatalf("Failed to create cert file: %v", err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		t.Fatalf("Failed to encode certificate: %v", err)
	}
	_ = certOut.Close()

	certBytes, err := os.ReadFile(certFile)
	if err != nil {
		t.Fatalf("Failed to read created certificate: %v", err)
	}
	block, _ := pem.Decode(certBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		t.Fatalf("Invalid certificate block")
	}

	parsed, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}

	if len(parsed.DNSNames) == 0 || parsed.DNSNames[0] != "localhost" {
		t.Errorf("Expected DNSNames to contain 'localhost', got %v", parsed.DNSNames)
	}
}
