/*
 * Copyright (C) 2022 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */

package tasks

import (
	"github.com/onsi/gomega"
	"intel/amber/kbs/v1/repository/mocks"
	"testing"
)

func TestCreateAdminUser(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	var userStore *mocks.MockUserStore = mocks.NewFakeUserStore()
	ac := CreateAdminUser{
		AdminUsername: "testAdmin",
		AdminPassword: "testPassword",
		UserStore:     userStore,
	}

	err := ac.CreateAdminUser()
	g.Expect(err).NotTo(gomega.HaveOccurred())
}

func TestCreateAdminUserWithInvalidCreds(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	var userStore *mocks.MockUserStore = mocks.NewFakeUserStore()
	ac := CreateAdminUser{
		AdminUsername: "",
		AdminPassword: "",
		UserStore:     userStore,
	}

	err := ac.CreateAdminUser()
	g.Expect(err).To(gomega.HaveOccurred())
}