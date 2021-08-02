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

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-newrelic-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Channels returns a ChannelInformer.
	Channels() ChannelInformer
	// Conditions returns a ConditionInformer.
	Conditions() ConditionInformer
	// MutingRules returns a MutingRuleInformer.
	MutingRules() MutingRuleInformer
	// Policies returns a PolicyInformer.
	Policies() PolicyInformer
	// PolicyChannels returns a PolicyChannelInformer.
	PolicyChannels() PolicyChannelInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Channels returns a ChannelInformer.
func (v *version) Channels() ChannelInformer {
	return &channelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Conditions returns a ConditionInformer.
func (v *version) Conditions() ConditionInformer {
	return &conditionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MutingRules returns a MutingRuleInformer.
func (v *version) MutingRules() MutingRuleInformer {
	return &mutingRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Policies returns a PolicyInformer.
func (v *version) Policies() PolicyInformer {
	return &policyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PolicyChannels returns a PolicyChannelInformer.
func (v *version) PolicyChannels() PolicyChannelInformer {
	return &policyChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
