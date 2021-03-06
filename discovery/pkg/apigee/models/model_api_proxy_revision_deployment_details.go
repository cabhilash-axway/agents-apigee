/*
 * Deployments
 *
 * Manage API proxy and shared flow deployments.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// ApiProxyRevisionDeploymentDetails API proxy revision deployment details in an environment.
type ApiProxyRevisionDeploymentDetails struct {
	// API proxy revision deployment details in the environment.
	Environment []ApiProxyRevisionDeploymentDetailsEnvironment `json:"environment,omitempty"`
	// Name of the API proxy.
	Name string `json:"name,omitempty"`
	// Name of the organization.
	Organization string `json:"organization,omitempty"`
}
