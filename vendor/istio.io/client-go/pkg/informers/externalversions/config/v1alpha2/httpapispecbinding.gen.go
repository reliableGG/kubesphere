// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	configv1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	internalinterfaces "istio.io/client-go/pkg/informers/externalversions/internalinterfaces"
	v1alpha2 "istio.io/client-go/pkg/listers/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HTTPAPISpecBindingInformer provides access to a shared informer and lister for
// HTTPAPISpecBindings.
type HTTPAPISpecBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.HTTPAPISpecBindingLister
}

type hTTPAPISpecBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHTTPAPISpecBindingInformer constructs a new informer for HTTPAPISpecBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHTTPAPISpecBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHTTPAPISpecBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHTTPAPISpecBindingInformer constructs a new informer for HTTPAPISpecBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHTTPAPISpecBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha2().HTTPAPISpecBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha2().HTTPAPISpecBindings(namespace).Watch(options)
			},
		},
		&configv1alpha2.HTTPAPISpecBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *hTTPAPISpecBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHTTPAPISpecBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hTTPAPISpecBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha2.HTTPAPISpecBinding{}, f.defaultInformer)
}

func (f *hTTPAPISpecBindingInformer) Lister() v1alpha2.HTTPAPISpecBindingLister {
	return v1alpha2.NewHTTPAPISpecBindingLister(f.Informer().GetIndexer())
}