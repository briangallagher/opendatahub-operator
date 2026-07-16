package v1alpha1

import (
	"github.com/opendatahub-io/opendatahub-operator/v2/api/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DataCatalogComponentName = "datacatalog"
	DataCatalogInstanceName  = "default-" + "datacatalog"
	DataCatalogKind          = "DataCatalog"
)

type DataCatalogCommonSpec struct {
	// CatalogImage is the container image for the catalog-enabled Feast server.
	// +optional
	CatalogImage string `json:"catalogImage,omitempty"`
	// RegistrySecretName is the name of the Secret containing the PostgreSQL DSN
	// for the shared catalog registry.
	// +optional
	RegistrySecretName string `json:"registrySecretName,omitempty"`
}

type DataCatalogSpec struct {
	DataCatalogCommonSpec `json:",inline"`
}

type DataCatalogCommonStatus struct {
	common.ComponentReleaseStatus `json:",inline"`
}

type DataCatalogStatus struct {
	common.Status            `json:",inline"`
	DataCatalogCommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'default-datacatalog'",message="DataCatalog name must be default-datacatalog"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`,description="Ready"
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].reason`,description="Reason"

// DataCatalog is the Schema for the DataCatalogs API
type DataCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataCatalogSpec   `json:"spec,omitempty"`
	Status DataCatalogStatus `json:"status,omitempty"`
}

func (c *DataCatalog) GetStatus() *common.Status {
	return &c.Status.Status
}

func (c *DataCatalog) GetConditions() []common.Condition {
	return c.Status.GetConditions()
}

func (c *DataCatalog) SetConditions(conditions []common.Condition) {
	c.Status.SetConditions(conditions)
}

func (c *DataCatalog) GetReleaseStatus() *[]common.ComponentRelease {
	return &c.Status.Releases
}

func (c *DataCatalog) SetReleaseStatus(releases []common.ComponentRelease) {
	c.Status.Releases = releases
}

// +kubebuilder:object:root=true

// DataCatalogList contains a list of DataCatalog
type DataCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataCatalog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataCatalog{}, &DataCatalogList{})
}

type DSCDataCatalog struct {
	common.ManagementSpec    `json:",inline"`
	DataCatalogCommonSpec `json:",inline"`
}

type DSCDataCatalogStatus struct {
	common.ManagementSpec       `json:",inline"`
	*DataCatalogCommonStatus `json:",inline"`
}
