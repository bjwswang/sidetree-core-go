/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteResponse(t *testing.T) {
	rw := httptest.NewRecorder()
	WriteResponse(rw, http.StatusOK, "content")
	require.Equal(t, http.StatusOK, rw.Code)
	require.Equal(t, "\"content\"\n", rw.Body.String())
	require.Equal(t, "application/did+ld+json", rw.Header().Get("content-type"))
}

func TestWriteError(t *testing.T) {
	rw := httptest.NewRecorder()
	errExpected := errors.New("some error")
	WriteError(rw, http.StatusBadRequest, errExpected)
	require.Equal(t, http.StatusBadRequest, rw.Code)
	require.Equal(t, errExpected.Error(), rw.Body.String())
}
