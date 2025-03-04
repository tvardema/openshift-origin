// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	operatorv1 "github.com/openshift/api/operator/v1"
	applyconfigurationsoperatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// MachineConfigurationsGetter has a method to return a MachineConfigurationInterface.
// A group's client should implement this interface.
type MachineConfigurationsGetter interface {
	MachineConfigurations() MachineConfigurationInterface
}

// MachineConfigurationInterface has methods to work with MachineConfiguration resources.
type MachineConfigurationInterface interface {
	Create(ctx context.Context, machineConfiguration *operatorv1.MachineConfiguration, opts metav1.CreateOptions) (*operatorv1.MachineConfiguration, error)
	Update(ctx context.Context, machineConfiguration *operatorv1.MachineConfiguration, opts metav1.UpdateOptions) (*operatorv1.MachineConfiguration, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, machineConfiguration *operatorv1.MachineConfiguration, opts metav1.UpdateOptions) (*operatorv1.MachineConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*operatorv1.MachineConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*operatorv1.MachineConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *operatorv1.MachineConfiguration, err error)
	Apply(ctx context.Context, machineConfiguration *applyconfigurationsoperatorv1.MachineConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *operatorv1.MachineConfiguration, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, machineConfiguration *applyconfigurationsoperatorv1.MachineConfigurationApplyConfiguration, opts metav1.ApplyOptions) (result *operatorv1.MachineConfiguration, err error)
	MachineConfigurationExpansion
}

// machineConfigurations implements MachineConfigurationInterface
type machineConfigurations struct {
	*gentype.ClientWithListAndApply[*operatorv1.MachineConfiguration, *operatorv1.MachineConfigurationList, *applyconfigurationsoperatorv1.MachineConfigurationApplyConfiguration]
}

// newMachineConfigurations returns a MachineConfigurations
func newMachineConfigurations(c *OperatorV1Client) *machineConfigurations {
	return &machineConfigurations{
		gentype.NewClientWithListAndApply[*operatorv1.MachineConfiguration, *operatorv1.MachineConfigurationList, *applyconfigurationsoperatorv1.MachineConfigurationApplyConfiguration](
			"machineconfigurations",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *operatorv1.MachineConfiguration { return &operatorv1.MachineConfiguration{} },
			func() *operatorv1.MachineConfigurationList { return &operatorv1.MachineConfigurationList{} },
		),
	}
}
