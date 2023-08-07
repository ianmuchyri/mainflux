// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package sdk_test

import (
	"fmt"
	"testing"

	"github.com/mainflux/mainflux"
	"github.com/mainflux/mainflux/pkg/errors"
	sdk "github.com/mainflux/mainflux/pkg/sdk/go"
	thingsclients "github.com/mainflux/mainflux/things/clients"
	"github.com/mainflux/mainflux/things/clients/mocks"
	gmocks "github.com/mainflux/mainflux/things/groups/mocks"
	"github.com/mainflux/mainflux/things/policies"
	thingspmocks "github.com/mainflux/mainflux/things/policies/mocks"
	usersclients "github.com/mainflux/mainflux/users/clients"
	cmocks "github.com/mainflux/mainflux/users/clients/mocks"
	"github.com/mainflux/mainflux/users/jwt"
	userspmocks "github.com/mainflux/mainflux/users/policies/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	cRepo := new(mocks.Repository)
	gRepo := new(gmocks.Repository)
	uauth := cmocks.NewAuthService(users, map[string][]cmocks.SubjectSet{adminID: {uadminPolicy}})
	thingCache := mocks.NewCache()
	policiesCache := thingspmocks.NewCache()
	tokenizer := jwt.NewRepository([]byte(secret), accessDuration, refreshDuration)

	thingsPRepo := new(thingspmocks.Repository)
	psvc := policies.NewService(uauth, thingsPRepo, policiesCache, idProvider)

	thSvc := thingsclients.NewService(uauth, psvc, cRepo, gRepo, thingCache, idProvider)
	ths := newThingsServer(thSvc, psvc)
	defer ths.Close()

	usersPRepo := new(userspmocks.Repository)
	usSvc := usersclients.NewService(cRepo, usersPRepo, tokenizer, emailer, phasher, idProvider, passRegex)
	usClSv := newClientServer(usSvc)
	defer usClSv.Close()

	sdkConf := sdk.Config{
		ThingsURL:       ths.URL,
		UsersURL:        usClSv.URL,
		MsgContentType:  contentType,
		TLSVerification: false,
	}

	mfsdk := sdk.NewSDK(sdkConf)
	cases := map[string]struct {
		service     string
		empty       bool
		description string
		status      string
		err         errors.SDKError
	}{
		"get things service health check": {
			service:     "things",
			empty:       false,
			err:         nil,
			description: "things service",
			status:      "pass",
		},
		"get users service health check": {
			service:     "users",
			empty:       false,
			err:         nil,
			description: "users service",
			status:      "pass",
		},
	}
	for desc, tc := range cases {
		h, err := mfsdk.Health(tc.service)
		assert.Equal(t, tc.err, err, fmt.Sprintf("%s: expected error %s, got %s", desc, tc.err, err))
		assert.Equal(t, tc.status, h.Status, fmt.Sprintf("%s: expected %s status, got %s", desc, tc.status, h.Status))
		assert.Equal(t, tc.empty, h.Version == "", fmt.Sprintf("%s: expected non-empty version", desc))
		assert.Equal(t, mainflux.Commit, h.Commit, fmt.Sprintf("%s: expected non-empty commit", desc))
		assert.Equal(t, tc.description, h.Description, fmt.Sprintf("%s: expected proper description, got %s", desc, h.Description))
		assert.Equal(t, mainflux.BuildTime, h.BuildTime, fmt.Sprintf("%s: expected default epoch date, got %s", desc, h.BuildTime))
	}
}
