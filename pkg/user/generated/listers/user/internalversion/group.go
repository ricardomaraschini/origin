// This file was automatically generated by lister-gen

package internalversion

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GroupLister helps list Groups.
type GroupLister interface {
	// List lists all Groups in the indexer.
	List(selector labels.Selector) (ret []*user.Group, err error)
	// Groups returns an object that can list and get Groups.
	Groups(namespace string) GroupNamespaceLister
	GroupListerExpansion
}

// groupLister implements the GroupLister interface.
type groupLister struct {
	indexer cache.Indexer
}

// NewGroupLister returns a new GroupLister.
func NewGroupLister(indexer cache.Indexer) GroupLister {
	return &groupLister{indexer: indexer}
}

// List lists all Groups in the indexer.
func (s *groupLister) List(selector labels.Selector) (ret []*user.Group, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*user.Group))
	})
	return ret, err
}

// Groups returns an object that can list and get Groups.
func (s *groupLister) Groups(namespace string) GroupNamespaceLister {
	return groupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GroupNamespaceLister helps list and get Groups.
type GroupNamespaceLister interface {
	// List lists all Groups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*user.Group, err error)
	// Get retrieves the Group from the indexer for a given namespace and name.
	Get(name string) (*user.Group, error)
	GroupNamespaceListerExpansion
}

// groupNamespaceLister implements the GroupNamespaceLister
// interface.
type groupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Groups in the indexer for a given namespace.
func (s groupNamespaceLister) List(selector labels.Selector) (ret []*user.Group, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*user.Group))
	})
	return ret, err
}

// Get retrieves the Group from the indexer for a given namespace and name.
func (s groupNamespaceLister) Get(name string) (*user.Group, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(user.Resource("group"), name)
	}
	return obj.(*user.Group), nil
}
