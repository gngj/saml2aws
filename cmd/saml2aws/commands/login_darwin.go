// +build darwin,cgo

package commands

import (
	"github.com/gngj/saml2aws/v2/helper/credentials"
	"github.com/gngj/saml2aws/v2/helper/osxkeychain"
)

func init() {
	credentials.CurrentHelper = &osxkeychain.Osxkeychain{}
}
