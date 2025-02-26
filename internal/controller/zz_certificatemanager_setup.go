/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	certificate "github.com/upbound/provider-gcp/internal/controller/certificatemanager/certificate"
	certificatemap "github.com/upbound/provider-gcp/internal/controller/certificatemanager/certificatemap"
	certificatemapentry "github.com/upbound/provider-gcp/internal/controller/certificatemanager/certificatemapentry"
	dnsauthorization "github.com/upbound/provider-gcp/internal/controller/certificatemanager/dnsauthorization"
)

// Setup_certificatemanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_certificatemanager(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		certificate.Setup,
		certificatemap.Setup,
		certificatemapentry.Setup,
		dnsauthorization.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
