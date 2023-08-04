// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/apache/dubbo-admin/pkg/rule/apis/dubbo.apache.org/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TagRouteLister helps list TagRoutes.
// All objects returned here must be treated as read-only.
type TagRouteLister interface {
	// List lists all TagRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TagRoute, err error)
	// TagRoutes returns an object that can list and get TagRoutes.
	TagRoutes(namespace string) TagRouteNamespaceLister
	TagRouteListerExpansion
}

// tagRouteLister implements the TagRouteLister interface.
type tagRouteLister struct {
	indexer cache.Indexer
}

// NewTagRouteLister returns a new TagRouteLister.
func NewTagRouteLister(indexer cache.Indexer) TagRouteLister {
	return &tagRouteLister{indexer: indexer}
}

// List lists all TagRoutes in the indexer.
func (s *tagRouteLister) List(selector labels.Selector) (ret []*v1beta1.TagRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TagRoute))
	})
	return ret, err
}

// TagRoutes returns an object that can list and get TagRoutes.
func (s *tagRouteLister) TagRoutes(namespace string) TagRouteNamespaceLister {
	return tagRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TagRouteNamespaceLister helps list and get TagRoutes.
// All objects returned here must be treated as read-only.
type TagRouteNamespaceLister interface {
	// List lists all TagRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TagRoute, err error)
	// Get retrieves the TagRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.TagRoute, error)
	TagRouteNamespaceListerExpansion
}

// tagRouteNamespaceLister implements the TagRouteNamespaceLister
// interface.
type tagRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TagRoutes in the indexer for a given namespace.
func (s tagRouteNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.TagRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TagRoute))
	})
	return ret, err
}

// Get retrieves the TagRoute from the indexer for a given namespace and name.
func (s tagRouteNamespaceLister) Get(name string) (*v1beta1.TagRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("tagroute"), name)
	}
	return obj.(*v1beta1.TagRoute), nil
}