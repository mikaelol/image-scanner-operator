// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// PolicyReportApplyConfiguration represents an declarative configuration of the PolicyReport type for use
// with apply.
type PolicyReportApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Source                           *string                                      `json:"source,omitempty"`
	Scope                            *corev1.ObjectReference                      `json:"scope,omitempty"`
	ScopeSelector                    *v1.LabelSelectorApplyConfiguration          `json:"scopeSelector,omitempty"`
	Configuration                    *PolicyReportConfigurationApplyConfiguration `json:"configuration,omitempty"`
	Summary                          *PolicyReportSummaryApplyConfiguration       `json:"summary,omitempty"`
	Results                          []PolicyReportResultApplyConfiguration       `json:"results,omitempty"`
}

// PolicyReport constructs an declarative configuration of the PolicyReport type for use with
// apply.
func PolicyReport(name, namespace string) *PolicyReportApplyConfiguration {
	b := &PolicyReportApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("PolicyReport")
	b.WithAPIVersion("reports.x-k8s.io/v1beta2")
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithKind(value string) *PolicyReportApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithAPIVersion(value string) *PolicyReportApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithName(value string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithGenerateName(value string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithNamespace(value string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithUID(value types.UID) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithResourceVersion(value string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithGeneration(value int64) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithCreationTimestamp(value metav1.Time) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *PolicyReportApplyConfiguration) WithLabels(entries map[string]string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *PolicyReportApplyConfiguration) WithAnnotations(entries map[string]string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *PolicyReportApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *PolicyReportApplyConfiguration) WithFinalizers(values ...string) *PolicyReportApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *PolicyReportApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithSource(value string) *PolicyReportApplyConfiguration {
	b.Source = &value
	return b
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithScope(value corev1.ObjectReference) *PolicyReportApplyConfiguration {
	b.Scope = &value
	return b
}

// WithScopeSelector sets the ScopeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScopeSelector field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithScopeSelector(value *v1.LabelSelectorApplyConfiguration) *PolicyReportApplyConfiguration {
	b.ScopeSelector = value
	return b
}

// WithConfiguration sets the Configuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Configuration field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithConfiguration(value *PolicyReportConfigurationApplyConfiguration) *PolicyReportApplyConfiguration {
	b.Configuration = value
	return b
}

// WithSummary sets the Summary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Summary field is set to the value of the last call.
func (b *PolicyReportApplyConfiguration) WithSummary(value *PolicyReportSummaryApplyConfiguration) *PolicyReportApplyConfiguration {
	b.Summary = value
	return b
}

// WithResults adds the given value to the Results field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Results field.
func (b *PolicyReportApplyConfiguration) WithResults(values ...*PolicyReportResultApplyConfiguration) *PolicyReportApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResults")
		}
		b.Results = append(b.Results, *values[i])
	}
	return b
}