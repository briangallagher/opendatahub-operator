package datacatalog

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	componentApi "github.com/opendatahub-io/opendatahub-operator/v2/api/components/v1alpha1"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/cluster"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/controller/types"
)

func initialize(_ context.Context, rr *types.ReconciliationRequest) error {
	rr.Manifests = nil
	return nil
}

func createFeatureStoreCR(ctx context.Context, rr *types.ReconciliationRequest) error {
	dc, ok := rr.Instance.(*componentApi.DataCatalog)
	if !ok {
		return fmt.Errorf("instance is not a DataCatalog")
	}

	appNamespace, err := cluster.ApplicationNamespace(ctx, rr.Client)
	if err != nil {
		return fmt.Errorf("failed to determine application namespace: %w", err)
	}

	catalogImage := dc.Spec.CatalogImage
	if catalogImage == "" {
		catalogImage = DefaultCatalogImage
	}

	registrySecret := dc.Spec.RegistrySecretName
	if registrySecret == "" {
		registrySecret = DefaultRegistrySecret
	}

	fs := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "feast.dev/v1",
			"kind":       "FeatureStore",
			"metadata": map[string]interface{}{
				"name":      FeatureStoreCRName,
				"namespace": appNamespace,
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by": "opendatahub-operator",
					"app.kubernetes.io/part-of":    "data-catalog",
				},
			},
			"spec": map[string]interface{}{
				"feastProject": "",
				"catalog": map[string]interface{}{
					"enabled": true,
					"ssar": map[string]interface{}{
						"apiGroup":  "datacatalog.opendatahub.io",
						"resources": []interface{}{"namespaces", "tables", "volumes"},
					},
				},
				"services": map[string]interface{}{
					"registry": map[string]interface{}{
						"local": map[string]interface{}{
							"server": map[string]interface{}{
								"image":   catalogImage,
								"restAPI": true,
							},
							"persistence": map[string]interface{}{
								"store": map[string]interface{}{
									"type": "sql",
									"secretRef": map[string]interface{}{
										"name": registrySecret,
									},
									"secretKeyName": "sql",
								},
							},
						},
					},
				},
				"authz": map[string]interface{}{
					"kubernetes": map[string]interface{}{},
				},
			},
		},
	}

	fs.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "feast.dev",
		Version: "v1",
		Kind:    "FeatureStore",
	})

	return rr.AddResources(fs)
}
