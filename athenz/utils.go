package athenz

import (
	"github.com/AthenZ/athenz/clients/go/zms"
	"strings"
)

func shortName(domainName string, en string, separator string) string {
	shortName := en
	if strings.HasPrefix(shortName, domainName+separator) {
		shortName = shortName[len(domainName)+len(separator):]
	}
	return shortName
}

func splitSubDomainId(subDomainId string) (string, string) {
	return splitId(subDomainId, SUB_DOMAIN_SEPARATOR)
}

func splitId(id, separator string) (string, string) {
	indexOfPrefixEnd := strings.LastIndex(id, separator)
	prefix := id[:indexOfPrefixEnd]
	shortName := id[indexOfPrefixEnd+1:]
	return prefix, shortName
}

func flattenRoleMembers(list []*zms.RoleMember) []interface{} {
	roleMembers := make([]interface{}, 0, len(list))
	for _, m := range list {
		roleMembers = append(roleMembers, string(m.MemberName))
	}
	return roleMembers
}
