package datacatalog

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"

	componentApi "github.com/opendatahub-io/opendatahub-operator/v2/api/components/v1alpha1"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/actions/deploy"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/actions/gc"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/actions/status/releases"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/reconciler"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/metadata/labels"
)

func (s *componentHandler) NewComponentReconciler(ctx context.Context, mgr ctrl.Manager) error {
	_, err := reconciler.ReconcilerFor(mgr, &componentApi.DataCatalog{}).
		WithAction(initialize).
		WithAction(createFeatureStoreCR).
		WithAction(releases.NewAction()).
		WithAction(deploy.NewAction(
			deploy.WithCache(),
			deploy.WithLabel(labels.ODH.Component(ComponentName), labels.True),
		)).
		WithAction(gc.NewAction()).
		Build(ctx)

	if err != nil {
		return err
	}

	return nil
}
