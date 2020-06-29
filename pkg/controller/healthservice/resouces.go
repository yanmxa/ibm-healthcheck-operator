//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package healthservice

import (
	"strconv"

	"github.com/IBM/ibm-healthcheck-operator/pkg/apis/operator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
)

var defaultResources = &HealthServiceResources{
	requestsCpu:    resource.NewMilliQuantity(50, resource.DecimalSI),      // 50m,
	requestsMemory: resource.NewQuantity(64*1024*1024, resource.BinarySI),  // 64Mi
	limitsCpu:      resource.NewMilliQuantity(500, resource.DecimalSI),     // 500m
	limitsMemory:   resource.NewQuantity(512*1024*1024, resource.BinarySI), // 512Mi
}

type HealthServiceResources struct {
	requestsCpu    *resource.Quantity
	requestsMemory *resource.Quantity
	limitsCpu      *resource.Quantity
	limitsMemory   *resource.Quantity
}

func (r *ReconcileHealthService) desiredResources(res *v1alpha1.Resources) (*HealthServiceResources, error) {
	var (
		requestCpu    *resource.Quantity
		requestMemory *resource.Quantity
		limitsCpu     *resource.Quantity
		limitsMemory  *resource.Quantity
	)
	if reqCpu, err := strconv.ParseInt(res.Requests.Cpu, 10, 64); err != nil {
		return defaultResources, err
	} else {
		requestCpu = resource.NewMilliQuantity(reqCpu, resource.DecimalSI)
	}

	if reqMemory, err := strconv.ParseInt(res.Requests.Memory, 10, 64); err != nil {
		return defaultResources, err
	} else {
		requestMemory = resource.NewQuantity(reqMemory*1024*1024, resource.BinarySI)
	}

	if limCpu, err := strconv.ParseInt(res.Limits.Cpu, 10, 64); err != nil {
		return defaultResources, err
	} else {
		limitsCpu = resource.NewMilliQuantity(limCpu, resource.DecimalSI)
	}

	if limMemory, err := strconv.ParseInt(res.Limits.Memory, 10, 64); err != nil {
		return defaultResources, err
	} else {
		limitsMemory = resource.NewQuantity(limMemory*1024*1024, resource.BinarySI)
	}

	return &HealthServiceResources{
		requestsCpu:    requestCpu,
		requestsMemory: requestMemory,
		limitsCpu:      limitsCpu,
		limitsMemory:   limitsMemory,
	}, nil
}
