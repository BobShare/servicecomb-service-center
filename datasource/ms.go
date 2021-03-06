/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package datasource

import (
	"context"
	pb "github.com/apache/servicecomb-service-center/pkg/registry"
)

// Attention: request validation must be finished before the following interface being invoked!!!
// MetadataManager contains the CRUD of cache metadata
type MetadataManager interface {
	// Microservice management
	RegisterService(ctx context.Context, request *pb.CreateServiceRequest) (*pb.CreateServiceResponse, error)
	GetServices(ctx context.Context, request *pb.GetServicesRequest) (*pb.GetServicesResponse, error)
	GetService(ctx context.Context, request *pb.GetServiceRequest) (*pb.GetServiceResponse, error)

	GetServiceDetail(ctx context.Context, request *pb.GetServiceRequest) (*pb.GetServiceDetailResponse, error)
	GetServicesInfo(ctx context.Context, request *pb.GetServicesInfoRequest) (*pb.GetServicesInfoResponse, error)
	GetApplications(ctx context.Context, request *pb.GetAppsRequest) (*pb.GetAppsResponse, error)

	ExistServiceByID(ctx context.Context, request *pb.GetExistenceByIDRequest) (*pb.GetExistenceByIDResponse, error)
	ExistService(ctx context.Context, request *pb.GetExistenceRequest) (*pb.GetExistenceResponse, error)
	UpdateService(ctx context.Context, request *pb.UpdateServicePropsRequest) (*pb.UpdateServicePropsResponse, error)
	UnregisterService(ctx context.Context, request *pb.DeleteServiceRequest) (*pb.DeleteServiceResponse, error)
	GetDeleteServiceFunc(ctx context.Context, serviceID string, force bool,
		serviceRespChan chan<- *pb.DelServicesRspInfo) func(context.Context)

	// Instance management
	RegisterInstance(ctx context.Context, request *pb.RegisterInstanceRequest) (*pb.RegisterInstanceResponse, error)
	// GetInstances returns instances under the current domain
	GetInstance(ctx context.Context, request *pb.GetOneInstanceRequest) (*pb.GetOneInstanceResponse, error)
	GetInstances(ctx context.Context, request *pb.GetInstancesRequest) (*pb.GetInstancesResponse, error)
	// GetProviderInstances returns instances under the specified domain
	GetProviderInstances(ctx context.Context, request *pb.HeartbeatSetElement) (instances []*pb.MicroServiceInstance, rev string, err error)
	BatchGetProviderInstances(ctx context.Context, request *pb.BatchGetInstancesRequest) (instances []*pb.MicroServiceInstance, rev string, err error)
	// FindInstances returns instances under the specified domain
	FindInstances(ctx context.Context, request *pb.FindInstancesRequest) (*pb.FindInstancesResponse, error)
	UpdateInstanceStatus(ctx context.Context, request *pb.UpdateInstanceStatusRequest) (
		*pb.UpdateInstanceStatusResponse, error)
	UpdateInstanceProperties(ctx context.Context, request *pb.UpdateInstancePropsRequest) (
		*pb.UpdateInstancePropsResponse, error)
	UnregisterInstance(ctx context.Context, request *pb.UnregisterInstanceRequest) (*pb.UnregisterInstanceResponse,
		error)
	Heartbeat(ctx context.Context, request *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error)
	HeartbeatSet(ctx context.Context, request *pb.HeartbeatSetRequest) (*pb.HeartbeatSetResponse, error)
	BatchFind(ctx context.Context, request *pb.BatchFindInstancesRequest) (*pb.BatchFindInstancesResponse, error)

	// Schema management
	ModifySchemas(ctx context.Context, request *pb.ModifySchemasRequest) (*pb.ModifySchemasResponse, error)
	ModifySchema(ctx context.Context, request *pb.ModifySchemaRequest) (*pb.ModifySchemaResponse, error)
	ExistSchema(ctx context.Context, request *pb.GetExistenceRequest) (*pb.GetExistenceResponse, error)
	GetSchema(ctx context.Context, request *pb.GetSchemaRequest) (*pb.GetSchemaResponse, error)
	GetAllSchemas(ctx context.Context, request *pb.GetAllSchemaRequest) (*pb.GetAllSchemaResponse, error)
	DeleteSchema(ctx context.Context, request *pb.DeleteSchemaRequest) (*pb.DeleteSchemaResponse, error)

	// Tag management
	AddTags(ctx context.Context, request *pb.AddServiceTagsRequest) (*pb.AddServiceTagsResponse, error)
	GetTags(ctx context.Context, request *pb.GetServiceTagsRequest) (*pb.GetServiceTagsResponse, error)
	UpdateTag(ctx context.Context, request *pb.UpdateServiceTagRequest) (*pb.UpdateServiceTagResponse, error)
	DeleteTags(ctx context.Context, request *pb.DeleteServiceTagsRequest) (*pb.DeleteServiceTagsResponse, error)

	// White/black list management
	AddRule(ctx context.Context, request *pb.AddServiceRulesRequest) (*pb.AddServiceRulesResponse, error)
	GetRule(ctx context.Context, request *pb.GetServiceRulesRequest) (*pb.GetServiceRulesResponse, error)
	UpdateRule(ctx context.Context, request *pb.UpdateServiceRuleRequest) (*pb.UpdateServiceRuleResponse, error)
	DeleteRule(ctx context.Context, request *pb.DeleteServiceRulesRequest) (*pb.DeleteServiceRulesResponse, error)
}
