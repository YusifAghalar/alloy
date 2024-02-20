package controller

import (
	"sync"

	"github.com/grafana/river/ast"
)

// CustomComponentRegistry holds custom component definitions that are available in the context.
type CustomComponentRegistry struct {
	parent *CustomComponentRegistry // nil if root config

	mut      sync.RWMutex
	declares map[string]ast.Body // customComponentName: template
}

// NewCustomComponentRegistry creates a new CustomComponentRegistry with a parent.
// parent can be nil.
func NewCustomComponentRegistry(parent *CustomComponentRegistry) *CustomComponentRegistry {
	return &CustomComponentRegistry{
		parent:   parent,
		declares: make(map[string]ast.Body),
	}
}

// registerDeclare stores a local declare block.
func (s *CustomComponentRegistry) registerDeclare(declare *ast.BlockStmt) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.declares[declare.Label] = declare.Body
}