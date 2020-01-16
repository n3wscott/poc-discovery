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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/tracker"

	client "github.com/n3wscott/discovery/pkg/client/injection/client"
	informer "github.com/n3wscott/discovery/pkg/client/injection/informers/discovery/v1alpha1/ducktype"
	"knative.dev/pkg/client/injection/apiextensions/informers/apiextensions/v1beta1/customresourcedefinition"
	"knative.dev/pkg/injection/clients/dynamicclient"
)

const (
	controllerAgentName = "ducktype-controller"
)

// NewController returns a new DuckType controller.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	dtInformer := informer.Get(ctx)

	crdInformer := customresourcedefinition.Get(ctx)

	c := &Reconciler{
		Client:        client.Get(ctx),
		DynamicClient: dynamicclient.Get(ctx),
		Lister:        dtInformer.Lister(),
		CRDLister:     crdInformer.Lister(),
		Recorder: record.NewBroadcaster().NewRecorder(
			scheme.Scheme, corev1.EventSource{Component: controllerAgentName}),
	}
	impl := controller.NewImpl(c, logger, "DuckTypes")

	logger.Info("Setting up event handlers")

	dtInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	// Watch custom resource definitions.
	grDt := func(obj interface{}) {
		impl.GlobalResync(dtInformer.Informer())
	}
	crdInformer.Informer().AddEventHandler(controller.HandleAll(grDt))

	c.Tracker = tracker.New(impl.EnqueueKey, controller.GetTrackerLease(ctx))
	return impl
}
