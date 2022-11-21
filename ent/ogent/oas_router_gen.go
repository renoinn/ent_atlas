// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"strings"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/users"
			if l := len("/users"); len(elem) >= l && elem[0:l] == "/users" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch r.Method {
				case "GET":
					s.handleListUsersRequest([0]string{}, w, r)
				case "POST":
					s.handleCreateUsersRequest([0]string{}, w, r)
				default:
					s.notAllowed(w, r, "GET,POST")
				}

				return
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "DELETE":
						s.handleDeleteUsersRequest([1]string{
							args[0],
						}, w, r)
					case "GET":
						s.handleReadUsersRequest([1]string{
							args[0],
						}, w, r)
					case "PATCH":
						s.handleUpdateUsersRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "DELETE,GET,PATCH")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [1]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/users"
			if l := len("/users"); len(elem) >= l && elem[0:l] == "/users" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				switch method {
				case "GET":
					r.name = "ListUsers"
					r.operationID = "listUsers"
					r.args = args
					r.count = 0
					return r, true
				case "POST":
					r.name = "CreateUsers"
					r.operationID = "createUsers"
					r.args = args
					r.count = 0
					return r, true
				default:
					return
				}
			}
			switch elem[0] {
			case '/': // Prefix: "/"
				if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "DELETE":
						// Leaf: DeleteUsers
						r.name = "DeleteUsers"
						r.operationID = "deleteUsers"
						r.args = args
						r.count = 1
						return r, true
					case "GET":
						// Leaf: ReadUsers
						r.name = "ReadUsers"
						r.operationID = "readUsers"
						r.args = args
						r.count = 1
						return r, true
					case "PATCH":
						// Leaf: UpdateUsers
						r.name = "UpdateUsers"
						r.operationID = "updateUsers"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}
