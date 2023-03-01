// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"mime"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

func decodeLoginPostResponse(resp *http.Response) (res LoginPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response LoginPostOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &LoginPostBadRequest{}, nil
	case 404:
		// Code 404.
		return &LoginPostNotFound{}, nil
	case 500:
		// Code 500.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper LoginPostApplicationJSONInternalServerError
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotRetryAfterVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							wrapperDotRetryAfterVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
						return nil
					}); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 503:
		// Code 503.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper LoginPostApplicationJSONServiceUnavailable
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotRetryAfterVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							wrapperDotRetryAfterVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
						return nil
					}); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeUserGetIDGetResponse(resp *http.Response) (res UserGetIDGetRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response User
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &UserGetIDGetBadRequest{}, nil
	case 404:
		// Code 404.
		return &UserGetIDGetNotFound{}, nil
	case 500:
		// Code 500.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper UserGetIDGetApplicationJSONInternalServerError
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotRetryAfterVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							wrapperDotRetryAfterVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
						return nil
					}); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 503:
		// Code 503.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper UserGetIDGetApplicationJSONServiceUnavailable
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotRetryAfterVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							wrapperDotRetryAfterVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
						return nil
					}); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeUserRegisterPostResponse(resp *http.Response) (res UserRegisterPostRes, err error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response UserRegisterPostOK
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			return &response, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 400:
		// Code 400.
		return &UserRegisterPostBadRequest{}, nil
	case 500:
		// Code 500.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper UserRegisterPostApplicationJSONInternalServerError
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotRetryAfterVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							wrapperDotRetryAfterVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
						return nil
					}); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	case 503:
		// Code 503.
		ct, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return res, errors.Wrap(err, "parse media type")
		}
		switch {
		case ct == "application/json":
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				return res, err
			}
			d := jx.DecodeBytes(buf)

			var response R5xx
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				if err := d.Skip(); err != io.EOF {
					return errors.New("unexpected trailing data")
				}
				return nil
			}(); err != nil {
				err = &ogenerrors.DecodeBodyError{
					ContentType: ct,
					Body:        buf,
					Err:         err,
				}
				return res, err
			}
			var wrapper UserRegisterPostApplicationJSONServiceUnavailable
			wrapper.Response = response
			h := uri.NewHeaderDecoder(resp.Header)
			// Parse "Retry-After" header.
			{
				cfg := uri.HeaderParameterDecodingConfig{
					Name:    "Retry-After",
					Explode: false,
				}
				if err := func() error {
					if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
						var wrapperDotRetryAfterVal int
						if err := func() error {
							val, err := d.DecodeValue()
							if err != nil {
								return err
							}

							c, err := conv.ToInt(val)
							if err != nil {
								return err
							}

							wrapperDotRetryAfterVal = c
							return nil
						}(); err != nil {
							return err
						}
						wrapper.RetryAfter.SetTo(wrapperDotRetryAfterVal)
						return nil
					}); err != nil {
						return err
					}
					return nil
				}(); err != nil {
					return res, errors.Wrap(err, "parse Retry-After header")
				}
			}
			return &wrapper, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
