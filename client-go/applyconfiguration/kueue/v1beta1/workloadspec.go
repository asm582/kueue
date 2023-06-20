/*
Copyright 2022 The Kubernetes Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// WorkloadSpecApplyConfiguration represents an declarative configuration of the WorkloadSpec type for use
// with apply.
type WorkloadSpecApplyConfiguration struct {
	PodSets           []PodSetApplyConfiguration `json:"podSets,omitempty"`
	QueueName         *string                    `json:"queueName,omitempty"`
	PriorityClassName *string                    `json:"priorityClassName,omitempty"`
	Priority          *int32                     `json:"priority,omitempty"`
}

// WorkloadSpecApplyConfiguration constructs an declarative configuration of the WorkloadSpec type for use with
// apply.
func WorkloadSpec() *WorkloadSpecApplyConfiguration {
	return &WorkloadSpecApplyConfiguration{}
}

// WithPodSets adds the given value to the PodSets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PodSets field.
func (b *WorkloadSpecApplyConfiguration) WithPodSets(values ...*PodSetApplyConfiguration) *WorkloadSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPodSets")
		}
		b.PodSets = append(b.PodSets, *values[i])
	}
	return b
}

// WithQueueName sets the QueueName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QueueName field is set to the value of the last call.
func (b *WorkloadSpecApplyConfiguration) WithQueueName(value string) *WorkloadSpecApplyConfiguration {
	b.QueueName = &value
	return b
}

// WithPriorityClassName sets the PriorityClassName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PriorityClassName field is set to the value of the last call.
func (b *WorkloadSpecApplyConfiguration) WithPriorityClassName(value string) *WorkloadSpecApplyConfiguration {
	b.PriorityClassName = &value
	return b
}

// WithPriority sets the Priority field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Priority field is set to the value of the last call.
func (b *WorkloadSpecApplyConfiguration) WithPriority(value int32) *WorkloadSpecApplyConfiguration {
	b.Priority = &value
	return b
}
