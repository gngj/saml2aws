package commands

import (
	"github.com/gngj/saml2aws/v2/helper/credentials"
	"github.com/gngj/saml2aws/v2/helper/linuxkeyring"
)

func init() {
	if keyringHelper, err := linuxkeyring.NewKeyringHelper(); err == nil {
		credentials.CurrentHelper = keyringHelper
	}
}
