// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package tokeninvalidator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTokenInvalidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResourceManager Controller TokenInvalidator Suite")
}
