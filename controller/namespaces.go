// Copyright 2016 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Handlers for namespace updates.  Keeps an index of namespace
// annotations

package main

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/controller"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/wait"
	"k8s.io/kubernetes/pkg/watch"
)

func initNamespaceInformer() {
	namespaceInformer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options api.ListOptions) (runtime.Object, error) {
				return kubeClient.Core().Namespaces().List(options)
			},
			WatchFunc: func(options api.ListOptions) (watch.Interface, error) {
				return kubeClient.Core().Namespaces().Watch(options)
			},
		},
		&api.Namespace{},
		controller.NoResyncPeriodFunc(),
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    namespaceChanged,
		UpdateFunc: namespaceUpdated,
		DeleteFunc: namespaceChanged,
	})

	go namespaceInformer.GetController().Run(wait.NeverStop)
	go namespaceInformer.Run(wait.NeverStop)
}

func namespaceUpdated(_ interface{}, obj interface{}) {
	namespaceChanged(obj)
}

func namespaceChanged(obj interface{}) {
	indexMutex.Lock()
	defer indexMutex.Unlock()

	ns := obj.(*api.Namespace)

	pods := podInformer.GetStore().List()
	for _, podobj := range pods {
		pod := podobj.(*api.Pod)

		if ns.Name == pod.ObjectMeta.Namespace {
			podChangedLocked(pod)
		}
	}
}
