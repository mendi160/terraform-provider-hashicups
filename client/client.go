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
	GetGroup(domain string, groupName string) (*zms.Group, error)
	DeleteGroup(domain string, groupName string, auditRef string) error
	PutGroup(domain string, groupName string, auditRef string, group *zms.Group) error
	PutGroupMembership(domain string, groupName string, memberName zms.GroupMemberName, auditRef string, membership *zms.GroupMembership) error
	DeleteGroupMembership(domain string, groupName string, member zms.GroupMemberName, auditRef string) error
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

func (c Client) PutGroupMembership(domain string, groupName string, memberName zms.GroupMemberName, auditRef string, membership *zms.GroupMembership) error {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.PutGroupMembership(zms.DomainName(domain), zms.EntityName(groupName), memberName, auditRef, membership)
}

func (c Client) DeleteGroupMembership(domain string, groupName string, member zms.GroupMemberName, auditRef string) error {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.DeleteGroupMembership(zms.DomainName(domain), zms.EntityName(groupName), member, auditRef)
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

func (c Client) PutGroup(domain string, groupName string, auditRef string, group *zms.Group) error {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.PutGroup(zms.DomainName(domain), zms.EntityName(groupName), auditRef, group)
}

func (c Client) DeleteGroup(domain string, groupName string, auditRef string) error {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.DeleteGroup(zms.DomainName(domain), zms.EntityName(groupName), auditRef)
}

func (c Client) GetGroup(domain string, groupName string) (*zms.Group, error) {
	zmsClient := zms.NewClient(c.Url, c.Transport)
	return zmsClient.GetGroup(zms.DomainName(domain), zms.EntityName(groupName), nil, nil)
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
