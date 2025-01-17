/*
Copyright 2020 The cert-manager Authors.

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

package v2

import (
	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/jetstack/cert-manager/internal/api/validation"
)

func ValidateTestType(_ *admissionv1.AdmissionRequest, obj runtime.Object) (field.ErrorList, validation.WarningList) {
	el := field.ErrorList{}
	tt := obj.(*TestType)
	if tt.TestField == DisallowedTestFieldValue {
		el = append(el, field.Invalid(field.NewPath("testField"), tt.TestField, "value not allowed"))
	}
	return el, nil
}
