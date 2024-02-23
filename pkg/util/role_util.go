package util

var ROLE_MAP map[string]string = map[string]string{"Admin": "ADMIN", "User": "USER"}

func ToGraphqlRole(role string) string {
	return ROLE_MAP[role]
}
