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

package http_handlers

import (
	"net/http"

	"github.com/apache/dubbo-kubernetes/pkg/bufman/controllers"
	registryv1alpha1 "github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/go/registry/v1alpha1"
	"github.com/gin-gonic/gin"
)

type authnGroup struct {
	authnController *controllers.AuthnController
}

var AuthnGroup = &authnGroup{
	authnController: controllers.NewAuthnController(),
}

func (group *authnGroup) GetCurrentUser(c *gin.Context) {
	req := &registryv1alpha1.GetCurrentUserRequest{}
	resp, err := group.authnController.GetCurrentUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewHTTPResponse(err))
		return
	}

	// 正常返回
	c.JSON(http.StatusOK, NewHTTPResponse(resp))
}
