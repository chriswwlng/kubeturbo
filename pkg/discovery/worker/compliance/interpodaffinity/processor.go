/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package interpodaffinity

import (
	"github.com/turbonomic/kubeturbo/pkg/discovery/repository"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/klog/v2"
)

// InterPodAffinityProcessor processes inter pod affinities
type InterPodAffinityProcessor struct {
	nodeInfoLister NodeInfoLister
	nsLister       NamespaceLister
}

// New initializes a new plugin and returns it.
func New(nodeInfoLister NodeInfoLister, nsLister NamespaceLister,
	clusterSummary repository.ClusterSummary) (*InterPodAffinityProcessor, error) {
	pl := &InterPodAffinityProcessor{
		nodeInfoLister: nodeInfoLister,
		nsLister:       nsLister,
	}

	return pl, nil
}

// Updates Namespaces with the set of namespaces identified by NamespaceSelector.
// If successful, NamespaceSelector is set to nil.
// The assumption is that the term is for an incoming pod, in which case
// namespaceSelector is either unrolled into Namespaces (and so the selector
// is set to Nothing()) or is Empty(), which means match everything. Therefore,
// there when matching against this term, there is no need to lookup the existing
// pod's namespace labels to match them against term's namespaceSelector explicitly.
func (pr *InterPodAffinityProcessor) mergeAffinityTermNamespacesIfNotEmpty(at *AffinityTerm) error {
	if at.NamespaceSelector.Empty() {
		return nil
	}
	ns, err := pr.nsLister.List(at.NamespaceSelector)
	if err != nil {
		return err
	}
	for _, n := range ns {
		at.Namespaces.Insert(n.Name)
	}
	at.NamespaceSelector = labels.Nothing()
	return nil
}

// GetNamespaceLabelsSnapshot returns a snapshot of the labels associated with
// the namespace.
func GetNamespaceLabelsSnapshot(ns string, nsLister NamespaceLister) (nsLabels labels.Set) {
	podNS, err := nsLister.Get(ns)
	if err == nil {
		// Create and return snapshot of the labels.
		return labels.Merge(podNS.Labels, nil)
	}
	klog.V(3).InfoS("getting namespace, assuming empty set of namespace labels", "namespace", ns, "err", err)
	return
}

func (pr *InterPodAffinityProcessor) ProcessAffinities() map[string][]string {
	return nil
}
