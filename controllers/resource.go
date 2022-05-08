// controllers/resource.go

package controllers

import (
	"github.com/FLM210/demo-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	Gohtpserver         = "dlbppx.online/Gohttpserver"
	GohtpserverLabelkey = "app"
)

func MutateDeployment(Gohttpserver *v1alpha1.Gohtpserver, dep *appsv1.Deployment) {
	dep.Labels = map[string]string{
		GohtpserverLabelkey: "gohtpserver",
	}
	dep.Spec = appsv1.DeploymentSpec{
		Replicas: Gohttpserver.Spec.Replicas,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					GohtpserverLabelkey: Gohttpserver.Spec.Name,
				},
				Name: Gohttpserver.Spec.Name,
			},
			Spec: corev1.PodSpec{
				Containers: newContainers(Gohttpserver),
			},
		},
	}
}

func newContainers(Gohttpserver *v1alpha1.Gohtpserver) []corev1.Container {
	return []corev1.Container{
		corev1.Container{
			Name:  Gohttpserver.Spec.Name,
			Image: Gohttpserver.Spec.Image,
			Ports: []corev1.ContainerPort{
				corev1.ContainerPort{
					Name:          "listenport",
					ContainerPort: 80,
				},
			},
		},
	}
}

func MutateSvc(Gohttpserver *v1alpha1.Gohtpserver, svc *corev1.Service) {
	svc.Labels = map[string]string{
		GohtpserverLabelkey: Gohttpserver.Spec.Name,
	}
	svc.Spec = corev1.ServiceSpec{
		Selector: map[string]string{
			GohtpserverLabelkey: Gohttpserver.Spec.Name,
		},
		Ports: []corev1.ServicePort{
			corev1.ServicePort{
				Name: "listenport",
				Port: 80,
			},
		},
	}
}
