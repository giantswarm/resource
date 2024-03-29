package appresourcetest

import (
	"time"

	applicationv1alpha1 "github.com/giantswarm/apiextensions-application/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewAppCRFromFilled returns a *applicationv1alpha1.App object with the fully
// populated spec. It also takes modifyFunc which allows to modify the object
// before returning. This is useful for testing. Example usage:
//
//	appCR := NewAppCRFromFilled(func(v *applicationv1alpha1.App) {
//		v.Name = "test-app-2"
//		v.Spec.Catalog = "test-spec-catalog-2"
//	})
func NewAppCRFromFilled(modifyFunc func(v *applicationv1alpha1.App)) *applicationv1alpha1.App {
	appCR := &applicationv1alpha1.App{
		TypeMeta: applicationv1alpha1.NewAppTypeMeta(),
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-app",
			Namespace: "test-namespace",
		},
		Spec: applicationv1alpha1.AppSpec{
			Catalog: "test-spec-catalog",
			Config: applicationv1alpha1.AppSpecConfig{
				ConfigMap: applicationv1alpha1.AppSpecConfigConfigMap{
					Name:      "test-spec-config-configmap-name",
					Namespace: "test-spec-config-configmap-namespace",
				},
				Secret: applicationv1alpha1.AppSpecConfigSecret{
					Name:      "test-spec-config-secret-name",
					Namespace: "test-spec-config-secret-namespace",
				},
			},
			KubeConfig: applicationv1alpha1.AppSpecKubeConfig{
				Secret: applicationv1alpha1.AppSpecKubeConfigSecret{
					Name:      "test-spec-kubeconfig-secret-name",
					Namespace: "test-spec-kubeconfig-secret-namespace",
				},
			},
			Name:      "test-spec-name",
			Namespace: "test-spec-namespace",
			UserConfig: applicationv1alpha1.AppSpecUserConfig{
				ConfigMap: applicationv1alpha1.AppSpecUserConfigConfigMap{
					Name:      "test-spec-userconfig-configmap-name",
					Namespace: "test-spec-userconfig-configmap-namespace",
				},
				Secret: applicationv1alpha1.AppSpecUserConfigSecret{
					Name:      "test-spec-userconfig-secret-name",
					Namespace: "test-spec-userconfig-secret-namespace",
				},
			},
			Version: "test-spec-version",
		},
		Status: applicationv1alpha1.AppStatus{
			AppVersion: "test-status",
			Release: applicationv1alpha1.AppStatusRelease{
				LastDeployed: metav1.Time{
					Time: time.Date(2019, 2, 12, 12, 4, 0, 0, time.UTC),
				},
				Status: "test-status-release-status",
			},
			Version: "test-status-version",
		},
	}

	modifyFunc(appCR)
	return appCR
}
