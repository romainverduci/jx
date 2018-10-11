// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkinsiov1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx/pkg/client/listers/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ComplianceCheckInformer provides access to a shared informer and lister for
// ComplianceChecks.
type ComplianceCheckInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ComplianceCheckLister
}

type complianceCheckInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewComplianceCheckInformer constructs a new informer for ComplianceCheck type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewComplianceCheckInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredComplianceCheckInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredComplianceCheckInformer constructs a new informer for ComplianceCheck type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredComplianceCheckInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().ComplianceChecks(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().ComplianceChecks(namespace).Watch(options)
			},
		},
		&jenkinsiov1.ComplianceCheck{},
		resyncPeriod,
		indexers,
	)
}

func (f *complianceCheckInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredComplianceCheckInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *complianceCheckInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.ComplianceCheck{}, f.defaultInformer)
}

func (f *complianceCheckInformer) Lister() v1.ComplianceCheckLister {
	return v1.NewComplianceCheckLister(f.Informer().GetIndexer())
}
