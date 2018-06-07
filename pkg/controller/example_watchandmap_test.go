/*
Copyright 2017 The Kubernetes Authors.

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

package controller_test

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/kubernetes-sigs/controller-runtime/pkg/config"
	"github.com/kubernetes-sigs/controller-runtime/pkg/controller"
	"github.com/kubernetes-sigs/controller-runtime/pkg/controller/types"
	"github.com/kubernetes-sigs/controller-runtime/pkg/inject/run"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func ExampleGenericController_WatchTransformationOf() {
	// One time setup for program
	flag.Parse()
	informerFactory := config.GetKubernetesInformersOrDie()
	if err := controller.AddInformerProvider(&corev1.Pod{}, informerFactory.Core().V1().Pods()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}
	if err := controller.AddInformerProvider(&appsv1.ReplicaSet{}, informerFactory.Apps().V1().ReplicaSets()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}

	// Per-controller setup
	c := &controller.GenericController{
		Reconcile: func(key types.ReconcileKey) error {
			fmt.Printf("Reconciling Pod %s\n", key)
			return nil
		},
	}
	err := c.Watch(&appsv1.ReplicaSet{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = c.WatchTransformationOf(&corev1.Pod{},
		func(i interface{}) string {
			p, ok := i.(*corev1.Pod)
			if !ok {
				return ""
			}

			// Find the parent key based on the name
			return p.Namespace + "/" + strings.Split(p.Name, "-")[0]
		},
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	controller.AddController(c)

	// One time for program
	controller.RunInformersAndControllers(run.CreateRunArguments())
}

func ExampleGenericController_WatchTransformationsOf() {
	// One time setup for program
	flag.Parse()
	informerFactory := config.GetKubernetesInformersOrDie()
	if err := controller.AddInformerProvider(&corev1.Pod{}, informerFactory.Core().V1().Pods()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}
	if err := controller.AddInformerProvider(&appsv1.ReplicaSet{}, informerFactory.Apps().V1().ReplicaSets()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}

	// Per-controller setup
	c := &controller.GenericController{
		Reconcile: func(key types.ReconcileKey) error {
			fmt.Printf("Reconciling Pod %s\n", key)
			return nil
		},
	}
	err := c.Watch(&appsv1.ReplicaSet{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = c.WatchTransformationsOf(&corev1.Pod{},
		func(i interface{}) []string {
			p, ok := i.(*corev1.Pod)
			if !ok {
				return []string{}
			}

			// Find multiple parents based off the name
			return []string{
				p.Namespace + "/" + strings.Split(p.Name, "-")[0] + "-parent-1",
				p.Namespace + "/" + strings.Split(p.Name, "-")[0] + "-parent-2",
			}
		},
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	controller.AddController(c)

	// One time for program
	controller.RunInformersAndControllers(run.CreateRunArguments())
}

func ExampleGenericController_WatchTransformationKeyOf() {
	// One time setup for program
	flag.Parse()
	informerFactory := config.GetKubernetesInformersOrDie()
	if err := controller.AddInformerProvider(&corev1.Pod{}, informerFactory.Core().V1().Pods()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}
	if err := controller.AddInformerProvider(&appsv1.ReplicaSet{}, informerFactory.Apps().V1().ReplicaSets()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}

	// Per-controller setup
	c := &controller.GenericController{
		Reconcile: func(key types.ReconcileKey) error {
			fmt.Printf("Reconciling Pod %s\n", key)
			return nil
		},
	}
	err := c.Watch(&appsv1.ReplicaSet{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = c.WatchTransformationKeyOf(&corev1.Pod{},
		func(i interface{}) types.ReconcileKey {
			p, ok := i.(*corev1.Pod)
			if !ok {
				return types.ReconcileKey{}
			}

			// Find multiple parents based off the name
			return types.ReconcileKey{p.Namespace, strings.Split(p.Name, "-")[0]}
		},
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	controller.AddController(c)

	// One time for program
	controller.RunInformersAndControllers(run.CreateRunArguments())
}

func ExampleGenericController_WatchTransformationKeysOf() {
	// One time setup for program
	flag.Parse()
	informerFactory := config.GetKubernetesInformersOrDie()
	if err := controller.AddInformerProvider(&corev1.Pod{}, informerFactory.Core().V1().Pods()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}
	if err := controller.AddInformerProvider(&appsv1.ReplicaSet{}, informerFactory.Apps().V1().ReplicaSets()); err != nil {
		log.Fatalf("Could not set informer %v", err)
	}

	// Per-controller setup
	c := &controller.GenericController{
		Reconcile: func(key types.ReconcileKey) error {
			fmt.Printf("Reconciling Pod %s\n", key)
			return nil
		},
	}
	err := c.Watch(&appsv1.ReplicaSet{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = c.WatchTransformationKeysOf(&corev1.Pod{},
		func(i interface{}) []types.ReconcileKey {
			p, ok := i.(*corev1.Pod)
			if !ok {
				return []types.ReconcileKey{}
			}

			// Find multiple parents based off the name
			return []types.ReconcileKey{
				{p.Namespace, strings.Split(p.Name, "-")[0] + "-parent-1"},
				{p.Namespace, strings.Split(p.Name, "-")[0] + "-parent-2"},
			}
		},
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	controller.AddController(c)

	// One time for program
	controller.RunInformersAndControllers(run.CreateRunArguments())
}
