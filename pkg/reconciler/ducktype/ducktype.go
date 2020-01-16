/*
Copyright 2020 The Knative Authors

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

package ducktype

import (
	"context"
	"github.com/n3wscott/discovery/pkg/apis/discovery/v1alpha1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/klog"
	"knative.dev/pkg/reconciler"
	"sort"
)

// ReconcileKind implements Interface
func (r *Reconciler) ReconcileKind(ctx context.Context, dt *v1alpha1.DuckType) reconciler.Event {
	if dt.GetDeletionTimestamp() != nil {
		// Check for a DeletionTimestamp.  If present, elide the normal reconcile logic.
		// When a controller needs finalizer handling, it would go here.
		return nil
	}
	dt.Status.InitializeConditions()

	/// By query

	kinds := make(map[string]*v1beta1.CustomResourceDefinition, 0)

	for _, st := range dt.Spec.SelectorType {
		crds, err := r.getCRDsWith(st.Selector)
		if err != nil {
			return err
		}
		for _, crd := range crds {
			key := crd.Name
			if _, found := kinds[key]; !found {
				kinds[key] = crd
			}
		}
	}

	gvrks := make([]v1alpha1.GroupVersionResourceKind, 0)
	for _, crd := range kinds {
		gvrks = append(gvrks, CRDToGVRK(crd))
	}

	/// By ref

	for _, gvrk := range dt.Spec.RefsList {
		// TODO we should query and test that the Ref is installed and works on this cluster.
		gvrks = append(gvrks, gvrk)
	}

	// Sort and store.

	sort.Sort(ByGR(gvrks))
	dt.Status.DuckList = gvrks
	dt.Status.DuckCount = len(gvrks)

	dt.Status.MarkDucksAvailable()
	return nil
}

// ByGR implements sort.Interface for []v1alpha1.GroupVersionResourceKind based on
// the group and resource fields.
type ByGR []v1alpha1.GroupVersionResourceKind

func (a ByGR) Len() int      { return len(a) }
func (a ByGR) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByGR) Less(i, j int) bool {
	keyI := a[i].Group + a[i].Resource
	keyJ := a[j].Group + a[j].Resource
	return keyI < keyJ
}

func CRDToGVRK(crd *v1beta1.CustomResourceDefinition) v1alpha1.GroupVersionResourceKind {
	for _, v := range crd.Spec.Versions {
		if !v.Served {
			continue
		}

		return v1alpha1.GroupVersionResourceKind{
			Group:    crd.Spec.Group,
			Version:  v.Name,
			Resource: crd.Spec.Names.Plural,
			Kind:     crd.Spec.Names.Kind,
		}
	}
	return v1alpha1.GroupVersionResourceKind{}
}

// Ducks returns CRDs labeled as given.
// labelSelector should be in the form "duck.knative.dev/source=true"
func (r *Reconciler) getCRDsWith(labelSelector string) ([]*v1beta1.CustomResourceDefinition, error) {
	ls, err := labels.Parse(labelSelector)
	if err != nil {
		return nil, err
	}

	list, err := r.CRDLister.List(ls)
	if err != nil {
		klog.Errorf("failed to list customresourcedefinitions, %v", err)
		return nil, err
	}

	return list, nil
}
