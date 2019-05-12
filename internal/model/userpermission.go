package model

import (
	"sort"
	"strings"
)

type UserPermission struct {
	parent      *UserPermission
	permissions []string
}

func (p *UserPermission) AddPermission(permission string) {
	for _, p := range p.permissions {
		if p == permission {
			return
		}
	}
	p.permissions = append(p.permissions, permission)
	if p.parent != nil {
		p.parent.AddPermission(permission)
	}
}

func (p *UserPermission) AddPermissions(permissions ...string) {
	for _, permission := range permissions {
		p.AddPermission(permission)
	}
}

func (p *UserPermission) RemovePermission(permission string) {
	for i, perm := range p.permissions {
		if perm == permission {
			p.permissions = append(p.permissions[:i], p.permissions[i+1:]...)
			break
		}
	}
	if p.parent != nil {
		p.parent.RemovePermission(permission)
	}
}

func (p *UserPermission) GetPermissions() []string {
	if !sort.IsSorted(sort.StringSlice(p.permissions)) {
		sort.Slice(p.permissions, func(i, j int) bool {
			return strings.Compare(p.permissions[i], p.permissions[j]) < 0
		})

	}
	return p.permissions
}

func (p *UserPermission) SetParent(parent *UserPermission) {
	p.parent = parent
	p.parent.AddPermissions(p.permissions...)
}
