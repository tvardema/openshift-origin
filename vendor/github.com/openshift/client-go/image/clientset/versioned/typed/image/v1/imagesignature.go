// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	imagev1 "github.com/openshift/api/image/v1"
	scheme "github.com/openshift/client-go/image/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gentype "k8s.io/client-go/gentype"
)

// ImageSignaturesGetter has a method to return a ImageSignatureInterface.
// A group's client should implement this interface.
type ImageSignaturesGetter interface {
	ImageSignatures() ImageSignatureInterface
}

// ImageSignatureInterface has methods to work with ImageSignature resources.
type ImageSignatureInterface interface {
	Create(ctx context.Context, imageSignature *imagev1.ImageSignature, opts metav1.CreateOptions) (*imagev1.ImageSignature, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	ImageSignatureExpansion
}

// imageSignatures implements ImageSignatureInterface
type imageSignatures struct {
	*gentype.Client[*imagev1.ImageSignature]
}

// newImageSignatures returns a ImageSignatures
func newImageSignatures(c *ImageV1Client) *imageSignatures {
	return &imageSignatures{
		gentype.NewClient[*imagev1.ImageSignature](
			"imagesignatures",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *imagev1.ImageSignature { return &imagev1.ImageSignature{} },
		),
	}
}
