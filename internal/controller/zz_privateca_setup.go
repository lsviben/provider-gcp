/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	capool "github.com/upbound/provider-gcp/internal/controller/privateca/capool"
	capooliammember "github.com/upbound/provider-gcp/internal/controller/privateca/capooliammember"
	certificate "github.com/upbound/provider-gcp/internal/controller/privateca/certificate"
	certificateauthority "github.com/upbound/provider-gcp/internal/controller/privateca/certificateauthority"
	certificatetemplate "github.com/upbound/provider-gcp/internal/controller/privateca/certificatetemplate"
	certificatetemplateiammember "github.com/upbound/provider-gcp/internal/controller/privateca/certificatetemplateiammember"
)

// Setup_privateca creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_privateca(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		capool.Setup,
		capooliammember.Setup,
		certificate.Setup,
		certificateauthority.Setup,
		certificatetemplate.Setup,
		certificatetemplateiammember.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
