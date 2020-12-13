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

	binddnsv1 "github.com/bind-dns/binddns-operator/pkg/apis/binddns/v1"
	versioned "github.com/bind-dns/binddns-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/bind-dns/binddns-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/bind-dns/binddns-operator/pkg/generated/listers/binddns/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DnsRuleInformer provides access to a shared informer and lister for
// DnsRules.
type DnsRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DnsRuleLister
}

type dnsRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDnsRuleInformer constructs a new informer for DnsRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDnsRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDnsRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDnsRuleInformer constructs a new informer for DnsRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDnsRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BinddnsV1().DnsRules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BinddnsV1().DnsRules(namespace).Watch(context.TODO(), options)
			},
		},
		&binddnsv1.DnsRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *dnsRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDnsRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dnsRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&binddnsv1.DnsRule{}, f.defaultInformer)
}

func (f *dnsRuleInformer) Lister() v1.DnsRuleLister {
	return v1.NewDnsRuleLister(f.Informer().GetIndexer())
}
