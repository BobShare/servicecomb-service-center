// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package counter

import (
	"context"
	"github.com/apache/servicecomb-service-center/datasource"
	"github.com/apache/servicecomb-service-center/datasource/etcd/kv"
	"github.com/apache/servicecomb-service-center/datasource/etcd/sd"
	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/pkg/registry"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/astaxie/beego"
)

var (
	SharedServiceIds util.ConcurrentMap
)

// ServiceIndexEventHandler counting the number of services
// Deprecated: Use metrics instead.
type ServiceIndexEventHandler struct {
}

func (h *ServiceIndexEventHandler) Type() sd.Type {
	return kv.ServiceIndex
}

func (h *ServiceIndexEventHandler) OnEvent(evt sd.KvEvent) {
	key := core.GetInfoFromSvcIndexKV(evt.KV.Key)
	if core.IsShared(key) {
		SharedServiceIds.Put(key.Tenant+core.SPLIT+evt.KV.Value.(string), struct{}{})
		return
	}

	switch evt.Type {
	case registry.EVT_INIT, registry.EVT_CREATE:
		GetCounters().OnCreate(h.Type(), key.Tenant)
	case registry.EVT_DELETE:
		GetCounters().OnDelete(h.Type(), key.Tenant)
	default:
	}
}

func NewServiceIndexEventHandler() *ServiceIndexEventHandler {
	return &ServiceIndexEventHandler{}
}

// InstanceEventHandler counting the number of instances
// Deprecated: Use metrics instead.
type InstanceEventHandler struct {
	SharedServiceIds map[string]struct{}
}

func (h *InstanceEventHandler) Type() sd.Type {
	return kv.INSTANCE
}

func (h *InstanceEventHandler) OnEvent(evt sd.KvEvent) {
	serviceID, _, domainProject := core.GetInfoFromInstKV(evt.KV.Key)
	key := domainProject + core.SPLIT + serviceID
	if _, ok := SharedServiceIds.Get(key); ok {
		return
	}

	switch evt.Type {
	case registry.EVT_INIT, registry.EVT_CREATE:
		if domainProject == core.RegistryDomainProject {
			resp, err := datasource.Instance().GetService(context.Background(), &registry.GetServiceRequest{ServiceId: serviceID})
			if err != nil || resp.Service == nil {
				log.Errorf(err, "GetService[%s] failed", key)
				return
			}
			if core.IsShared(proto.MicroServiceToKey(domainProject, resp.Service)) {
				SharedServiceIds.Put(key, struct{}{})
				return
			}
		}
		GetCounters().OnCreate(h.Type(), domainProject)
	case registry.EVT_DELETE:
		GetCounters().OnDelete(h.Type(), domainProject)
	}
}

func NewInstanceEventHandler() *InstanceEventHandler {
	return &InstanceEventHandler{SharedServiceIds: make(map[string]struct{})}
}

func RegisterCounterListener(pluginName string) {
	if pluginName != beego.AppConfig.DefaultString("quota_plugin", "buildin") {
		return
	}
	sd.AddEventHandler(NewServiceIndexEventHandler())
	sd.AddEventHandler(NewInstanceEventHandler())
}
