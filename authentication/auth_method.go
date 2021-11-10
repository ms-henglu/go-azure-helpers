package authentication

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
)

type authMethod interface {
	build(b Builder) (authMethod, error)
	getADALToken(ctx context.Context, sender autorest.Sender, oauthConfig *OAuthConfig, endpoint string) (autorest.Authorizer, error)
	getMSALToken(ctx context.Context, sender autorest.Sender, oauthConfig *OAuthConfig, endpoint string) (autorest.Authorizer, error)
	isApplicable(b Builder) bool
	name() string
	populateConfig(c *Config) error
	validate() error
}
