// Package runtime contains BloxOne DDI API helper
// runtime abstractions for internal client use
package runtime

import (
	"path"
	"strings"
)

// TrimIDPrefix removes app ID and resource type prefixes from the ID property.
// If no prefix is found, TrimIDPrefix will return the original id value.
// More about BloxOne DDI resource IDs:
// https://github.com/infobloxopen/atlas-app-toolkit/tree/v1.1.2/rpc/resource#resource
func TrimIDPrefix(pathPattern string, id string) string {
	if !path.IsAbs(id) {
		id = "/" + id
	}
	prefix := pathPattern
	for !strings.HasPrefix(id, prefix) {
		prefix = prefix[0 : len(prefix)-1]
		if len(prefix) == 0 {
			return id
		}
	}
	if strings.HasPrefix(pathPattern, prefix+"{id}") {
		return strings.TrimPrefix(id, prefix)
	} else {
		return strings.TrimPrefix(id, "/")
	}

}
