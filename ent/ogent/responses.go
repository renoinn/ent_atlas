// Code generated by ent, DO NOT EDIT.

package ogent

import "github.com/renoinn/ent_atlas/ent"

func NewUsersCreate(e *ent.Users) *UsersCreate {
	if e == nil {
		return nil
	}
	var ret UsersCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Email = e.Email
	ret.Introduction = e.Introduction
	return &ret
}

func NewUsersCreates(es []*ent.Users) []UsersCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersCreate, len(es))
	for i, e := range es {
		r[i] = NewUsersCreate(e).Elem()
	}
	return r
}

func (u *UsersCreate) Elem() UsersCreate {
	if u == nil {
		return UsersCreate{}
	}
	return *u
}

func NewUsersList(e *ent.Users) *UsersList {
	if e == nil {
		return nil
	}
	var ret UsersList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Email = e.Email
	ret.Introduction = e.Introduction
	return &ret
}

func NewUsersLists(es []*ent.Users) []UsersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersList, len(es))
	for i, e := range es {
		r[i] = NewUsersList(e).Elem()
	}
	return r
}

func (u *UsersList) Elem() UsersList {
	if u == nil {
		return UsersList{}
	}
	return *u
}

func NewUsersRead(e *ent.Users) *UsersRead {
	if e == nil {
		return nil
	}
	var ret UsersRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Email = e.Email
	ret.Introduction = e.Introduction
	return &ret
}

func NewUsersReads(es []*ent.Users) []UsersRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersRead, len(es))
	for i, e := range es {
		r[i] = NewUsersRead(e).Elem()
	}
	return r
}

func (u *UsersRead) Elem() UsersRead {
	if u == nil {
		return UsersRead{}
	}
	return *u
}

func NewUsersUpdate(e *ent.Users) *UsersUpdate {
	if e == nil {
		return nil
	}
	var ret UsersUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Email = e.Email
	ret.Introduction = e.Introduction
	return &ret
}

func NewUsersUpdates(es []*ent.Users) []UsersUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UsersUpdate, len(es))
	for i, e := range es {
		r[i] = NewUsersUpdate(e).Elem()
	}
	return r
}

func (u *UsersUpdate) Elem() UsersUpdate {
	if u == nil {
		return UsersUpdate{}
	}
	return *u
}