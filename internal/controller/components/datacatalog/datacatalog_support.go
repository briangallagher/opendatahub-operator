package datacatalog

import (
	componentApi "github.com/opendatahub-io/opendatahub-operator/v2/api/components/v1alpha1"
	"github.com/opendatahub-io/opendatahub-operator/v2/internal/controller/status"
)

const (
	ComponentName = componentApi.DataCatalogComponentName

	ReadyConditionType = componentApi.DataCatalogKind + status.ReadySuffix

	DefaultCatalogImage    = "quay.io/rhoai/feast-server-catalog:latest"
	DefaultRegistrySecret  = "feast-catalog-db"
	FeatureStoreCRName     = "data-catalog"
)
