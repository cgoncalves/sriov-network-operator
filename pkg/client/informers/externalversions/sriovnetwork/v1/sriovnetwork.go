/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	sriovnetworkv1 "github.com/openshift/sriov-network-operator/api/v1"
	versioned "github.com/openshift/sriov-network-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift/sriov-network-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/sriov-network-operator/pkg/client/listers/sriovnetwork/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SriovNetworkInformer provides access to a shared informer and lister for
// SriovNetworks.
type SriovNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SriovNetworkLister
}

type sriovNetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSriovNetworkInformer constructs a new informer for SriovNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSriovNetworkInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSriovNetworkInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSriovNetworkInformer constructs a new informer for SriovNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSriovNetworkInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SriovnetworkV1().SriovNetworks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SriovnetworkV1().SriovNetworks(namespace).Watch(context.TODO(), options)
			},
		},
		&sriovnetworkv1.SriovNetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *sriovNetworkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSriovNetworkInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sriovNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sriovnetworkv1.SriovNetwork{}, f.defaultInformer)
}

func (f *sriovNetworkInformer) Lister() v1.SriovNetworkLister {
	return v1.NewSriovNetworkLister(f.Informer().GetIndexer())
}
