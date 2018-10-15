/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package oauth

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName       = "oauth.openshift.io"
	LegacyGroupName = ""
)

// SchemeGroupVersion is group version used to register these objects
var (
	SchemeGroupVersion       = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	LegacySchemeGroupVersion = schema.GroupVersion{Group: LegacyGroupName, Version: runtime.APIVersionInternal}

	LegacySchemeBuilder    = runtime.NewSchemeBuilder(addLegacyKnownTypes)
	AddToSchemeInCoreGroup = LegacySchemeBuilder.AddToScheme

	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// LegacyKind takes an unqualified kind and returns back a Group qualified GroupKind
func LegacyKind(kind string) schema.GroupKind {
	return LegacySchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func LegacyResource(resource string) schema.GroupResource {
	return LegacySchemeGroupVersion.WithResource(resource).GroupResource()
}

// IsKindOrLegacy checks if the provided GroupKind matches with the given kind by looking
// up the API group and also the legacy API.
func IsKindOrLegacy(kind string, gk schema.GroupKind) bool {
	return gk == Kind(kind) || gk == LegacyKind(kind)
}

// IsResourceOrLegacy checks if the provided GroupResources matches with the given
// resource by looking up the API group and also the legacy API.
func IsResourceOrLegacy(resource string, gr schema.GroupResource) bool {
	return gr == Resource(resource) || gr == LegacyResource(resource)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&OAuthAccessToken{},
		&OAuthAccessTokenList{},
		&OAuthAuthorizeToken{},
		&OAuthAuthorizeTokenList{},
		&OAuthClient{},
		&OAuthClientList{},
		&OAuthClientAuthorization{},
		&OAuthClientAuthorizationList{},
		&OAuthRedirectReference{},
	)
	return nil
}

func addLegacyKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&OAuthAccessToken{},
		&OAuthAccessTokenList{},
		&OAuthAuthorizeToken{},
		&OAuthAuthorizeTokenList{},
		&OAuthClient{},
		&OAuthClientList{},
		&OAuthClientAuthorization{},
		&OAuthClientAuthorizationList{},
		&OAuthRedirectReference{},
	}
	scheme.AddKnownTypes(LegacySchemeGroupVersion, types...)
	return nil
}