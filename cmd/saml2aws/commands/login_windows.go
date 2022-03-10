package commands

import (
	"github.com/gngj/saml2aws/v2/helper/credentials"
	"github.com/gngj/saml2aws/v2/helper/wincred"
)

func init() {
	credentials.CurrentHelper = &wincred.Wincred{}
}
