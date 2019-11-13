/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha2

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	v1alpha3 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha3"
)

// Convert_v1alpha3_VPCSpec_To_v1alpha2_VPCSpec is an autogenerated conversion function.
func Convert_v1alpha3_VPCSpec_To_v1alpha2_VPCSpec(in *v1alpha3.VPCSpec, out *VPCSpec, s conversion.Scope) error { // nolint
	if err := autoConvert_v1alpha3_VPCSpec_To_v1alpha2_VPCSpec(in, out, s); err != nil {
		return err
	}
	// Discard Ipv6CidrBlock and EnableIPv6 fields
	return nil
}
