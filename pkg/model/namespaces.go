package model

import (
	"time"

	"github.com/sirupsen/logrus"

	"git.containerum.net/ch/kube-api/pkg/kubernetes"
	kube_types "github.com/containerum/kube-client/pkg/model"
	api_core "k8s.io/api/core/v1"
)

const defaultNSCPU = 100
const defaultNSRAM = 100

func AddNamespacesWithoutRQ(kube *kubernetes.Kube, nsList *kube_types.NamespacesList, allns interface{}) *kube_types.NamespacesList {
	objects := allns.(*api_core.NamespaceList)
	for _, ns := range objects.Items {
		count := 1
		exists := false
		for _, oldns := range nsList.Namespaces {
			if oldns.ID == ns.Name {
				exists = true
				break
			}
		}
		if !exists {
			podsList, err := kube.GetPodList(ns.Name, "")
			if err == nil {
				podsListT := podsList.(*api_core.PodList)
				podCount := len(podsListT.Items)
				if podCount > 1 {
					count = podCount
				}
			} else {
				logrus.Warn(err)
			}

			owner := ns.GetObjectMeta().GetLabels()["owner"]
			createdAt := ns.ObjectMeta.CreationTimestamp.UTC().Format(time.RFC3339)
			newNs := kube_types.Namespace{
				Owner:     owner,
				ID:        ns.Name,
				CreatedAt: &createdAt,
				Resources: kube_types.Resources{
					Hard: kube_types.Resource{
						CPU:    uint(defaultNSCPU * count),
						Memory: uint(defaultNSRAM * count),
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

func GetNamespacesWithoutQuota(kube *kubernetes.Kube, rqList *kube_types.NamespacesList, allns interface{}) *kube_types.NamespacesList {
	objects := allns.(*api_core.NamespaceList)
	var nsworq kube_types.NamespacesList
	for _, ns := range objects.Items {
		count := 1
		exists := false
		for _, oldns := range rqList.Namespaces {
			if oldns.ID == ns.Name {
				exists = true
				break
			}
		}
		if !exists {
			podsList, err := kube.GetPodList(ns.Name, "")
			if err == nil {
				podsListT := podsList.(*api_core.PodList)
				podCount := len(podsListT.Items)
				if podCount > 1 {
					count = podCount
				}
			} else {
				logrus.Warn(err)
			}

			newNs := kube_types.Namespace{
				ID: ns.Name,
				Resources: kube_types.Resources{
					Hard: kube_types.Resource{
						CPU:    uint(defaultNSCPU * count),
						Memory: uint(defaultNSRAM * count),
					},
				},
			}
			nsworq.Namespaces = append(nsworq.Namespaces, newNs)
		}
	}
	return &nsworq
}
