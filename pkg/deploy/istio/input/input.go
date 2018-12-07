package input

import (
	"github.com/rancher/rio/types/apis/networking.istio.io/v1alpha3"
	projectv1 "github.com/rancher/rio/types/apis/project.rio.cattle.io/v1"
	"k8s.io/api/core/v1"
)

type IstioDeployment struct {
	LBNamespace     *v1.Namespace
	VirtualServices []*v1alpha3.VirtualService
	PublicDomains   []*projectv1.PublicDomain
	Secret          *v1.Secret
}
