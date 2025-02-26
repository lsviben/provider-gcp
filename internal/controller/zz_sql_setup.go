/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/upbound/provider-gcp/internal/controller/sql/database"
	databaseinstance "github.com/upbound/provider-gcp/internal/controller/sql/databaseinstance"
	sourcerepresentationinstance "github.com/upbound/provider-gcp/internal/controller/sql/sourcerepresentationinstance"
	sslcert "github.com/upbound/provider-gcp/internal/controller/sql/sslcert"
	user "github.com/upbound/provider-gcp/internal/controller/sql/user"
)

// Setup_sql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		databaseinstance.Setup,
		sourcerepresentationinstance.Setup,
		sslcert.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
