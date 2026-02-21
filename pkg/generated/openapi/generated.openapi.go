// Code generated manually. DO NOT EDIT via generator.

package openapi

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		// BlanketOps types
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Environment":       schemaEnvironment(ref),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentList":   schemaEnvironmentList(ref),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentSpec":   schemaEnvironmentSpec(ref),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentStatus": schemaEnvironmentStatus(ref),
		// apimachinery meta v1
		"io.k8s.apimachinery.pkg.apis.meta.v1.APIGroupList":              schemaSimpleObject("APIGroupList"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.APIGroup":                  schemaSimpleObject("APIGroup"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.APIVersions":               schemaSimpleObject("APIVersions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.APIResourceList":           schemaSimpleObject("APIResourceList"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.APIResource":               schemaSimpleObject("APIResource"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.Status":                    schemaSimpleObject("Status"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails":             schemaSimpleObject("StatusDetails"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.StatusCause":               schemaSimpleObject("StatusCause"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.WatchEvent":                schemaSimpleObject("WatchEvent"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta":                schemaSimpleObject("ObjectMeta"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta":                  schemaSimpleObject("ListMeta"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.Condition":                 schemaSimpleObject("Condition"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions":             schemaSimpleObject("DeleteOptions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.CreateOptions":             schemaSimpleObject("CreateOptions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.UpdateOptions":             schemaSimpleObject("UpdateOptions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.PatchOptions":              schemaSimpleObject("PatchOptions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.GetOptions":                schemaSimpleObject("GetOptions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.ListOptions":               schemaSimpleObject("ListOptions"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference":            schemaSimpleObject("OwnerReference"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector":             schemaSimpleObject("LabelSelector"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement":  schemaSimpleObject("LabelSelectorRequirement"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry":        schemaSimpleObject("ManagedFieldsEntry"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1":                  schemaSimpleObject("FieldsV1"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.Time":                      schemaSimpleObject("Time"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.Duration":                  schemaSimpleObject("Duration"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.Patch":                     schemaSimpleObject("Patch"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.RootPaths":                 schemaSimpleObject("RootPaths"),
		"io.k8s.apimachinery.pkg.apis.meta.v1.ServerAddressByClientCIDR": schemaSimpleObject("ServerAddressByClientCIDR"),
		// apimachinery runtime
		"io.k8s.apimachinery.pkg.runtime.RawExtension": schemaSimpleObject("RawExtension"),
		// apimachinery version
		"io.k8s.apimachinery.pkg.version.Info": schemaVersionInfo(ref),
	}
}
func schemaSimpleObject(description string) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: description,
			},
		},
	}
}
func schemaVersionInfo(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "Info contains versioning information.",
				Properties: map[string]spec.Schema{
					"major":        {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"minor":        {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"gitVersion":   {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"gitCommit":    {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"gitTreeState": {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"buildDate":    {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"goVersion":    {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"compiler":     {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
					"platform":     {SchemaProps: spec.SchemaProps{Type: []string{"string"}}},
				},
			},
		},
	}
}

func schemaEnvironment(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "Environment is a BlanketOps environment resource.",
			},
		},
	}
}

func schemaEnvironmentList(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "EnvironmentList contains a list of Environment resources.",
			},
		},
	}
}

func schemaEnvironmentSpec(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "EnvironmentSpec is the spec for an Environment.",
			},
		},
	}
}

func schemaEnvironmentStatus(_ common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type:        []string{"object"},
				Description: "EnvironmentStatus is the status for an Environment.",
			},
		},
	}
}
