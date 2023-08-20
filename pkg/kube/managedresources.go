/*
Copyright 2018 FairwindsOps Inc.

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

package kube

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	automationv1alpha1 "github.com/kubensync/operator/api/v1alpha1"
)

func GetManagedResources(ctx context.Context) (automationv1alpha1.ManagedResourceList, error) {
	list := automationv1alpha1.ManagedResourceList{}

	client, err := getMRDefClient()
	if err != nil {
		return list, err
	}

	err = client.Get().Resource("managedresources").Do(ctx).Into(&list)

	return list, err
}

func getMRDefClient() (*rest.RESTClient, error) {
	_ = automationv1alpha1.AddToScheme(scheme.Scheme)
	clientConfig := config.GetConfigOrDie()
	clientConfig.ContentConfig.GroupVersion = &automationv1alpha1.GroupVersion
	clientConfig.APIPath = "/apis"
	clientConfig.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	clientConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	return rest.UnversionedRESTClientFor(clientConfig)
}
