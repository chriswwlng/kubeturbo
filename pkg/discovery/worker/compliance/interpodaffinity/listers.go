package interpodaffinity

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// NodeInfoLister interface represents anything that can list/get NodeInfo objects from node name.
type NodeInfoLister interface {
	// List returns the list of NodeInfos.
	List() ([]*NodeInfo, error)
	// HavePodsWithAffinityList returns the list of NodeInfos of nodes with pods with affinity terms.
	HavePodsWithAffinityList() ([]*NodeInfo, error)
	// HavePodsWithRequiredAntiAffinityList returns the list of NodeInfos of nodes with pods with required anti-affinity terms.
	HavePodsWithRequiredAntiAffinityList() ([]*NodeInfo, error)
	// Get returns the NodeInfo of the given node name.
	Get(nodeName string) (*NodeInfo, error)
}

// NamespaceLister helps list Namespaces.
// All objects returned here must be treated as read-only.
type NamespaceLister interface {
	// List lists all Namespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Namespace, err error)
	// Get retrieves the Namespace from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Namespace, error)
}
