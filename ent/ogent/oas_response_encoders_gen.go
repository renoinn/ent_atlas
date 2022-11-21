// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreateUsersResponse(response CreateUsersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UsersCreate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeDeleteUsersResponse(response DeleteUsersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteUsersNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeListUsersResponse(response ListUsersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ListUsersOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeReadUsersResponse(response ReadUsersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UsersRead:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUpdateUsersResponse(response UpdateUsersRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UsersUpdate:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R400:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R404:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R409:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		span.SetStatus(codes.Error, http.StatusText(409))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *R500:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))
		e := jx.GetEncoder()

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
