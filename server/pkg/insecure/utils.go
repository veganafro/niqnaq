package insecure

import (
	"log"
	"crypto/tls"
	"crypto/x509"
)

// GetCertPool - Read and return the insecure SSL certificate
func GetCertPool() (*x509.CertPool) {
	certPool := x509.NewCertPool()
	status := certPool.AppendCertsFromPEM([]byte(Cert))
	if !status {
		log.Fatal("Failed to append certificates from PEM | ", status)
	}
	return certPool
}

// GetKeyPair - Read and return the insecure SSL certificate and key
func GetKeyPair() (*tls.Certificate) {
	pair, error := tls.X509KeyPair([]byte(Cert), []byte(Key))
	if error != nil {
		log.Fatal("Failed to create key pair | ", error)
	}
	return &pair
}

