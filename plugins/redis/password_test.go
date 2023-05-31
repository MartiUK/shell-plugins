package redis

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestPasswordProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, Password().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Password: "pjtxpc2gaddifapjvalggspojexample",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"REDISCLI_AUTH": "pjtxpc2gaddifapjvalggspojexample",
				},
			},
		},
	})
}

func TestPasswordImporter(t *testing.T) {
	plugintest.TestImporter(t, Password().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"REDISCLI_AUTH": "pjtxpc2gaddifapjvalggspojexample",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Password: "pjtxpc2gaddifapjvalggspojexample",
					},
				},
			},
		},
	})
}
