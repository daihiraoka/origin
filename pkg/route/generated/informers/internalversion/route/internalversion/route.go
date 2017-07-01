// This file was automatically generated by informer-gen

package internalversion

import (
	route "github.com/openshift/origin/pkg/route/apis/route"
	internalinterfaces "github.com/openshift/origin/pkg/route/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/route/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/route/generated/listers/route/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// RouteInformer provides access to a shared informer and lister for
// Routes.
type RouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.RouteLister
}

type routeInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newRouteInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Route().Routes(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Route().Routes(v1.NamespaceAll).Watch(options)
			},
		},
		&route.Route{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *routeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&route.Route{}, newRouteInformer)
}

func (f *routeInformer) Lister() internalversion.RouteLister {
	return internalversion.NewRouteLister(f.Informer().GetIndexer())
}