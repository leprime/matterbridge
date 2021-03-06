// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ACLType undocumented
type ACLType string

const (
	// ACLTypeVUser undocumented
	ACLTypeVUser ACLType = "user"
	// ACLTypeVGroup undocumented
	ACLTypeVGroup ACLType = "group"
	// ACLTypeVEveryone undocumented
	ACLTypeVEveryone ACLType = "everyone"
	// ACLTypeVEveryoneExceptGuests undocumented
	ACLTypeVEveryoneExceptGuests ACLType = "everyoneExceptGuests"
)

var (
	// ACLTypePUser is a pointer to ACLTypeVUser
	ACLTypePUser = &_ACLTypePUser
	// ACLTypePGroup is a pointer to ACLTypeVGroup
	ACLTypePGroup = &_ACLTypePGroup
	// ACLTypePEveryone is a pointer to ACLTypeVEveryone
	ACLTypePEveryone = &_ACLTypePEveryone
	// ACLTypePEveryoneExceptGuests is a pointer to ACLTypeVEveryoneExceptGuests
	ACLTypePEveryoneExceptGuests = &_ACLTypePEveryoneExceptGuests
)

var (
	_ACLTypePUser                 = ACLTypeVUser
	_ACLTypePGroup                = ACLTypeVGroup
	_ACLTypePEveryone             = ACLTypeVEveryone
	_ACLTypePEveryoneExceptGuests = ACLTypeVEveryoneExceptGuests
)
