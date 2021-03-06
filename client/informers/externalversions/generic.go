/*
Copyright AppsCode Inc. and Contributors

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

package externalversions

import (
	"fmt"

	v1alpha1 "kubeform.dev/provider-newrelic-api/apis/alert/v1alpha1"
	apiaccessv1alpha1 "kubeform.dev/provider-newrelic-api/apis/apiaccess/v1alpha1"
	applicationv1alpha1 "kubeform.dev/provider-newrelic-api/apis/application/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-newrelic-api/apis/dashboard/v1alpha1"
	entityv1alpha1 "kubeform.dev/provider-newrelic-api/apis/entity/v1alpha1"
	eventsv1alpha1 "kubeform.dev/provider-newrelic-api/apis/events/v1alpha1"
	infrav1alpha1 "kubeform.dev/provider-newrelic-api/apis/infra/v1alpha1"
	insightsv1alpha1 "kubeform.dev/provider-newrelic-api/apis/insights/v1alpha1"
	nrqlv1alpha1 "kubeform.dev/provider-newrelic-api/apis/nrql/v1alpha1"
	onev1alpha1 "kubeform.dev/provider-newrelic-api/apis/one/v1alpha1"
	pluginsv1alpha1 "kubeform.dev/provider-newrelic-api/apis/plugins/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-newrelic-api/apis/service/v1alpha1"
	syntheticsv1alpha1 "kubeform.dev/provider-newrelic-api/apis/synthetics/v1alpha1"
	workloadv1alpha1 "kubeform.dev/provider-newrelic-api/apis/workload/v1alpha1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=alert.newrelic.kubeform.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("channels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().Channels().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("conditions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().Conditions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("mutingrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().MutingRules().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("policies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().Policies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("policychannels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Alert().V1alpha1().PolicyChannels().Informer()}, nil

		// Group=apiaccess.newrelic.kubeform.com, Version=v1alpha1
	case apiaccessv1alpha1.SchemeGroupVersion.WithResource("keys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiaccess().V1alpha1().Keys().Informer()}, nil

		// Group=application.newrelic.kubeform.com, Version=v1alpha1
	case applicationv1alpha1.SchemeGroupVersion.WithResource("settingses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Application().V1alpha1().Settingses().Informer()}, nil

		// Group=dashboard.newrelic.kubeform.com, Version=v1alpha1
	case dashboardv1alpha1.SchemeGroupVersion.WithResource("dashboards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dashboard().V1alpha1().Dashboards().Informer()}, nil

		// Group=entity.newrelic.kubeform.com, Version=v1alpha1
	case entityv1alpha1.SchemeGroupVersion.WithResource("tagses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Entity().V1alpha1().Tagses().Informer()}, nil

		// Group=events.newrelic.kubeform.com, Version=v1alpha1
	case eventsv1alpha1.SchemeGroupVersion.WithResource("tometricsrules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Events().V1alpha1().ToMetricsRules().Informer()}, nil

		// Group=infra.newrelic.kubeform.com, Version=v1alpha1
	case infrav1alpha1.SchemeGroupVersion.WithResource("alertconditions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Infra().V1alpha1().AlertConditions().Informer()}, nil

		// Group=insights.newrelic.kubeform.com, Version=v1alpha1
	case insightsv1alpha1.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Insights().V1alpha1().Events().Informer()}, nil

		// Group=nrql.newrelic.kubeform.com, Version=v1alpha1
	case nrqlv1alpha1.SchemeGroupVersion.WithResource("alertconditions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nrql().V1alpha1().AlertConditions().Informer()}, nil
	case nrqlv1alpha1.SchemeGroupVersion.WithResource("droprules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Nrql().V1alpha1().DropRules().Informer()}, nil

		// Group=one.newrelic.kubeform.com, Version=v1alpha1
	case onev1alpha1.SchemeGroupVersion.WithResource("dashboards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.One().V1alpha1().Dashboards().Informer()}, nil
	case onev1alpha1.SchemeGroupVersion.WithResource("dashboardraws"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.One().V1alpha1().DashboardRaws().Informer()}, nil

		// Group=plugins.newrelic.kubeform.com, Version=v1alpha1
	case pluginsv1alpha1.SchemeGroupVersion.WithResource("alertconditions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Plugins().V1alpha1().AlertConditions().Informer()}, nil

		// Group=service.newrelic.kubeform.com, Version=v1alpha1
	case servicev1alpha1.SchemeGroupVersion.WithResource("levels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Service().V1alpha1().Levels().Informer()}, nil

		// Group=synthetics.newrelic.kubeform.com, Version=v1alpha1
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("alertconditions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().AlertConditions().Informer()}, nil
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("monitors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().Monitors().Informer()}, nil
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("monitorscripts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().MonitorScripts().Informer()}, nil
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("multilocationalertconditions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().MultilocationAlertConditions().Informer()}, nil
	case syntheticsv1alpha1.SchemeGroupVersion.WithResource("securecredentials"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Synthetics().V1alpha1().SecureCredentials().Informer()}, nil

		// Group=workload.newrelic.kubeform.com, Version=v1alpha1
	case workloadv1alpha1.SchemeGroupVersion.WithResource("workloads"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Workload().V1alpha1().Workloads().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
