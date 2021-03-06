package grafana

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	grafanav1alpha1 "go.f110.dev/mono/go/pkg/api/grafana/v1alpha1"
	"go.f110.dev/mono/go/pkg/grafana"
	"go.f110.dev/mono/go/pkg/k8s/controllers/controllertest"
)

func TestUserController_ObjectToKeys(t *testing.T) {
	runner, controller := newController(t)

	keys := controller.ObjectToKeys(&grafanav1alpha1.Grafana{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: metav1.NamespaceDefault,
			Name:      "test1",
		},
	})
	require.Len(t, keys, 1)
	assert.Equal(t, "default/test1", keys[0])

	runner.RegisterFixture(&grafanav1alpha1.Grafana{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: metav1.NamespaceDefault,
			Name:      "test1",
		},
		Spec: grafanav1alpha1.GrafanaSpec{
			UserSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "test",
				},
			},
		},
	})
	keys = controller.ObjectToKeys(&grafanav1alpha1.GrafanaUser{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: metav1.NamespaceDefault,
			Name:      "user1",
			Labels: map[string]string{
				"app": "test",
			},
		},
	})
	require.Len(t, keys, 1)
	assert.Equal(t, "default/test1", keys[0])

	keys = controller.ObjectToKeys(&corev1.Service{})
	require.Len(t, keys, 0)
}

func TestUserController_GetObject(t *testing.T) {
	runner, controller := newController(t)

	_, err := controller.GetObject("")
	require.Error(t, err)

	runner.RegisterFixture(&grafanav1alpha1.Grafana{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test1",
			Namespace: metav1.NamespaceDefault,
		},
	})
	obj, err := controller.GetObject("default/test1")
	require.NoError(t, err)
	app, ok := obj.(*grafanav1alpha1.Grafana)
	require.True(t, ok)
	assert.Equal(t, "test1", app.Name)
}

func TestUserController_UpdateObject(t *testing.T) {
	runner, controller := newController(t)

	_, err := controller.UpdateObject(context.Background(), &corev1.Service{})
	require.Error(t, err)

	target := &grafanav1alpha1.Grafana{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test1",
			Namespace: metav1.NamespaceDefault,
		},
		Status: grafanav1alpha1.GrafanaStatus{
			ObservedGeneration: 2,
		},
	}
	runner.RegisterFixture(target)

	app, err := controller.UpdateObject(context.Background(), target)
	require.NoError(t, err)
	runner.AssertAction(t, controllertest.Action{
		Verb:   controllertest.ActionUpdate,
		Object: target,
	})
	assert.Equal(t, target, app)
}

func TestUserController(t *testing.T) {
	runner, controller := newController(t)
	target, fixtures := grafanaFixture()

	mockTransport := httpmock.NewMockTransport()
	controller.transport = mockTransport
	mockTransport.RegisterResponder(
		http.MethodGet,
		"http://grafana.default.svc:3000/api/users",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, []grafana.User{
			{
				Id:    1,
				Email: "user2@example.com",
			},
		}),
	)
	mockTransport.RegisterResponder(
		http.MethodPost,
		"http://grafana.default.svc:3000/api/admin/users",
		httpmock.NewStringResponder(http.StatusOK, ""),
	)
	mockTransport.RegisterResponder(
		http.MethodDelete,
		"http://grafana.default.svc:3000/api/admin/users/1",
		httpmock.NewStringResponder(http.StatusOK, ""),
	)

	user := &grafanav1alpha1.GrafanaUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "user1",
			Namespace: metav1.NamespaceDefault,
			Labels:    target.Spec.UserSelector.MatchLabels,
		},
		Spec: grafanav1alpha1.GrafanaUserSpec{
			Email: "user1@example.com",
		},
	}
	runner.RegisterFixture(user)
	runner.RegisterFixture(fixtures...)

	err := runner.Reconcile(controller, target)
	require.NoError(t, err)

	expect := target.DeepCopy()
	expect.Status.ObservedGeneration = 2
	runner.AssertAction(
		t, controllertest.Action{
			Verb:        controllertest.ActionUpdate,
			Subresource: "status",
			Object:      expect,
		},
	)
	runner.AssertNoUnexpectedAction(t)
}

func newController(t *testing.T) (*controllertest.TestRunner, *UserController) {
	runner := controllertest.NewTestRunner()
	controller, err := NewUserController(
		runner.CoreSharedInformerFactory,
		runner.SharedInformerFactory,
		runner.CoreClient,
		runner.Client,
	)
	require.NoError(t, err)

	return runner, controller
}

func grafanaFixture() (*grafanav1alpha1.Grafana, []runtime.Object) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "admin",
			Namespace: metav1.NamespaceDefault,
		},
		Data: map[string][]byte{
			"password": []byte("foobar"),
		},
	}
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "grafana",
			Namespace: metav1.NamespaceDefault,
		},
	}

	target := &grafanav1alpha1.Grafana{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test",
			Namespace:  metav1.NamespaceDefault,
			Generation: 2,
		},
		Spec: grafanav1alpha1.GrafanaSpec{
			UserSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "test",
				},
			},
			Service: &corev1.LocalObjectReference{
				Name: service.Name,
			},
			AdminPasswordSecret: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: secret.Name,
				},
				Key: "password",
			},
		},
	}
	return target, []runtime.Object{secret, service}
}
