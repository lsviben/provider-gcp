/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	backupbackupplan "github.com/upbound/provider-gcp/internal/controller/gke/backupbackupplan"
)

// Setup_gke creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_gke(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupbackupplan.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
