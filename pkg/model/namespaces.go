package model

import (
	"time"

	kube_types "github.com/containerum/kube-client/pkg/model"
	api_core "k8s.io/api/core/v1"
)

const defaultNSCPU = 100
const defaultNSRAM = 100

func AddNamespacesWithoutRQ(nsList *kube_types.NamespacesList, allns interface{}) *kube_types.NamespacesList {
	objects := allns.(*api_core.NamespaceList)
	for _, ns := range objects.Items {
		exists := false
		for _, oldns := range nsList.Namespaces {
			if oldns.ID == ns.Name {
				exists = true
				break
			}
		}
		if !exists {
			owner := ns.GetObjectMeta().GetLabels()["owner"]
			createdAt := ns.ObjectMeta.CreationTimestamp.UTC().Format(time.RFC3339)
			newNs := kube_types.Namespace{
				Owner:     owner,
				ID:        ns.Name,
				CreatedAt: &createdAt,
				Resources: kube_types.Resources{
					Hard: kube_types.Resource{
						CPU:    defaultNSCPU,
						Memory: defaultNSRAM,
					},
				},
			}
			if owner == "" {
				newNs.Owner = "00000000-0000-0000-0000-000000000000"
			}
			nsList.Namespaces = append(nsList.Namespaces, newNs)
		}
	}
	return nsList
}
