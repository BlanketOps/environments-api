// Code generated manually. DO NOT EDIT via generator.
package openapi

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		// ── environments.blanketops.dev ──────────────────────────────
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Environment":        schemaSimpleObject("Environment"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentList":    schemaSimpleObject("EnvironmentList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentSpec":    schemaSimpleObject("EnvironmentSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.EnvironmentStatus":  schemaSimpleObject("EnvironmentStatus"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Build":              schemaSimpleObject("Build"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildList":          schemaSimpleObject("BuildList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildSpec":          schemaSimpleObject("BuildSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildStatus":        schemaSimpleObject("BuildStatus"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildTrigger":       schemaSimpleObject("BuildTrigger"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildTriggerList":   schemaSimpleObject("BuildTriggerList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildTriggerSpec":   schemaSimpleObject("BuildTriggerSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.BuildTriggerStatus": schemaSimpleObject("BuildTriggerStatus"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Deployment":         schemaSimpleObject("Deployment"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.DeploymentList":     schemaSimpleObject("DeploymentList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.DeploymentSpec":     schemaSimpleObject("DeploymentSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.DeploymentStatus":   schemaSimpleObject("DeploymentStatus"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Package":            schemaSimpleObject("Package"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.PackageList":        schemaSimpleObject("PackageList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.PackageSpec":        schemaSimpleObject("PackageSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.PackageStatus":      schemaSimpleObject("PackageStatus"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.Route":              schemaSimpleObject("Route"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.RouteList":          schemaSimpleObject("RouteList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.RouteSpec":          schemaSimpleObject("RouteSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.RouteStatus":        schemaSimpleObject("RouteStatus"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.ServiceUnit":        schemaSimpleObject("ServiceUnit"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.ServiceUnitList":    schemaSimpleObject("ServiceUnitList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.ServiceUnitSpec":    schemaSimpleObject("ServiceUnitSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/environments/v1alpha1.ServiceUnitStatus":  schemaSimpleObject("ServiceUnitStatus"),
		// ── events.blanketops.dev ────────────────────────────────────
		"github.com/ntlaletsi70/blanketops-environments-api/api/events/v1alpha1.GitHubEvent":       schemaSimpleObject("GitHubEvent"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/events/v1alpha1.GitHubEventList":   schemaSimpleObject("GitHubEventList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/events/v1alpha1.GitHubEventSpec":   schemaSimpleObject("GitHubEventSpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/events/v1alpha1.GitHubEventStatus": schemaSimpleObject("GitHubEventStatus"),
		// ── sources.blanketops.dev ───────────────────────────────────
		"github.com/ntlaletsi70/blanketops-environments-api/api/sources/v1alpha1.GitRepository":       schemaSimpleObject("GitRepository"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/sources/v1alpha1.GitRepositoryList":   schemaSimpleObject("GitRepositoryList"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/sources/v1alpha1.GitRepositorySpec":   schemaSimpleObject("GitRepositorySpec"),
		"github.com/ntlaletsi70/blanketops-environments-api/api/sources/v1alpha1.GitRepositoryStatus": schemaSimpleObject("GitRepositoryStatus"),
		// ── apimachinery ─────────────────────────────────────────────
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
		"io.k8s.apimachinery.pkg.runtime.RawExtension":                   schemaSimpleObject("RawExtension"),
		"io.k8s.apimachinery.pkg.version.Info":                           schemaVersionInfo(ref),
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
