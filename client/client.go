package client

import (
	"github.com/AthenZ/athenz/libs/go/tls/config"
	"net/http"

	"github.com/AthenZ/athenz/clients/go/zms"
)

type ZmsClient interface {
	GetRole(domain string, roleName string) (*zms.Role, error)
	GetDomain(domainName string) (*zms.Domain, error)
	PostSubDomain(parentDomain string, auditRef string, detail *zms.SubDomain) (*zms.Domain, error)
	DeleteSubDomain(parentDomain string, subDomainName string, auditRef string) error
}

type Client struct {
	Url       string
	Transport *http.Transport
}

type ZmsConfig struct {
	Url  string
	Cert string
	Key  string
}

func (c Client) DeleteSubDomain(parentDomain string, subDomainName string, auditRef string) error {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.DeleteSubDomain(zms.DomainName(parentDomain), zms.SimpleName(subDomainName), auditRef)
}
func (c Client) PostSubDomain(parentDomain string, auditRef string, detail *zms.SubDomain) (*zms.Domain, error) {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.PostSubDomain(zms.DomainName(parentDomain), auditRef, detail)
}

func (c Client) GetDomain(domainName string) (*zms.Domain, error) {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.GetDomain(zms.DomainName(domainName))
}

func (c Client) GetRole(domain string, roleName string) (*zms.Role, error) {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.GetRole(zms.DomainName(domain), zms.EntityName(roleName), nil, nil, nil)
}

func NewClient(url string, certFile string, keyFile string) (*Client, error) {
	tlsConfig, err := config.GetTLSConfigFromFiles(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	transport := http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &Client{
		Url:       url,
		Transport: &transport,
	}
	return client, err
}
