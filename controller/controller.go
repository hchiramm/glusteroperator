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

package controller

import (
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	tprv1 "github.com/hchiramm/glusteroperator/apis/tpr/v1"
	"github.com/hchiramm/glusteroperator/nodeagent"
)

type GlusterController struct {
	GlusterClient *rest.RESTClient
	GlusterScheme *runtime.Scheme
}

// Run starts a gluster resource controller
func (c *GlusterController) Run(ctx <-chan struct{}) error {
	glog.Infof("Watch gluster objects\n")

	// Watch gluster objects
	source := cache.NewListWatchFromClient(
		c.GlusterClient,
		tprv1.GlusterClusterResourcePlural,
		apiv1.NamespaceAll,
		fields.Everything())

	_, controller := cache.NewInformer(
		source,

		// The object type.
		&tprv1.GlusterCluster{},

		// resyncPeriod
		// Every resyncPeriod, all resources in the cache will retrigger events.
		// Set to 0 to disable the resync.
		0,

		// Your custom resource event handlers.
		cache.ResourceEventHandlerFuncs{
			AddFunc:    c.onAdd,
			UpdateFunc: c.onUpdate,
			DeleteFunc: c.onDelete,
		})

	go controller.Run(ctx)
	return nil
}

func (c *GlusterController) onAdd(obj interface{}) {
	gcluster := obj.(*tprv1.GlusterCluster)
	glog.Infof("[CONTROLLER] OnAdd %s", gcluster.ObjectMeta.SelfLink)

	if gcluster.Spec.Node != "" {
		err := nodeagent.FetchNodeName()
		if err != nil {
			glog.Warningf("failed to get node %s, err: %v", gcluster.Spec.Node, err)
		} else {
			glog.Infof("Node %s belong to gluster cluster", gcluster.Spec.Node)
		}
	}

}

func (c *GlusterController) onUpdate(oldObj, newObj interface{}) {
	oldGlusterNode := oldObj.(*tprv1.GlusterCluster)
	newGlusterNode := newObj.(*tprv1.GlusterCluster)
	glog.Infof("[CONTROLLER] OnUpdate oldObj: %s\n", oldGlusterNode.ObjectMeta.SelfLink)
	glog.Infof("[CONTROLLER] OnUpdate newObj: %s\n", newGlusterNode.ObjectMeta.SelfLink)
}

func (c *GlusterController) onDelete(obj interface{}) {
	gcluster := obj.(*tprv1.GlusterCluster)
	glog.Infof("[CONTROLLER] OnDelete %s\n", gcluster.ObjectMeta.SelfLink)
}
