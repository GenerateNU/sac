package helpers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/goccy/go-json"

	"github.com/huandu/go-assert"
)

type TestRequest struct {
	Method  string
	Path    string
	Body    *map[string]interface{}
	Headers *map[string]string
	Role    *models.UserRole
}

func (app TestApp) Send(request TestRequest) (*http.Response, error) {
	address := fmt.Sprintf("%s%s", app.Address, request.Path)

	var req *http.Request

	if request.Body == nil {
		req = httptest.NewRequest(request.Method, address, nil)
	} else {
		bodyBytes, err := json.Marshal(request.Body)
		if err != nil {
			return nil, err
		}

		req = httptest.NewRequest(request.Method, address, bytes.NewBuffer(bodyBytes))

		if request.Headers == nil {
			request.Headers = &map[string]string{}
		}

		if _, ok := (*request.Headers)["Content-Type"]; !ok {
			(*request.Headers)["Content-Type"] = "application/json"
		}
	}

	if request.Headers != nil {
		for key, value := range *request.Headers {
			req.Header.Add(key, value)
		}
	}

	if app.TestUser != nil {
		req.AddCookie(&http.Cookie{
			Name:  "access_token",
			Value: app.TestUser.AccessToken,
		})
	}

	resp, err := app.App.Test(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (request TestRequest) Test(t *testing.T, existingAppAssert *ExistingAppAssert) (ExistingAppAssert, *http.Response) {
	if existingAppAssert == nil {
		app, assert := InitTest(t)

		if request.Role != nil {
			app.Auth(*request.Role)
		}
		existingAppAssert = &ExistingAppAssert{
			App:    app,
			Assert: assert,
		}
	}

	resp, err := existingAppAssert.App.Send(request)

	existingAppAssert.Assert.NilError(err)

	return *existingAppAssert, resp
}

func (request TestRequest) TestOnStatus(t *testing.T, existingAppAssert *ExistingAppAssert, status int) ExistingAppAssert {
	appAssert, resp := request.Test(t, existingAppAssert)

	_, assert := appAssert.App, appAssert.Assert

	assert.Equal(status, resp.StatusCode)

	return appAssert
}

func (request *TestRequest) testOn(t *testing.T, existingAppAssert *ExistingAppAssert, status int, key string, value string) (ExistingAppAssert, *http.Response) {
	appAssert, resp := request.Test(t, existingAppAssert)
	assert := appAssert.Assert

	var respBody map[string]interface{}

	err := json.NewDecoder(resp.Body).Decode(&respBody)

	assert.NilError(err)
	assert.Equal(value, respBody[key].(string))

	assert.Equal(status, resp.StatusCode)
	return appAssert, resp
}

func (request TestRequest) TestOnError(t *testing.T, existingAppAssert *ExistingAppAssert, expectedError errors.Error) ExistingAppAssert {
	appAssert, _ := request.testOn(t, existingAppAssert, expectedError.StatusCode, "error", expectedError.Message)
	return appAssert
}

type ErrorWithTester struct {
	Error  errors.Error
	Tester Tester
}

func (request TestRequest) TestOnErrorAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, errorWithDBTester ErrorWithTester) ExistingAppAssert {
	appAssert, resp := request.testOn(t, existingAppAssert, errorWithDBTester.Error.StatusCode, "error", errorWithDBTester.Error.Message)
	errorWithDBTester.Tester(appAssert.App, appAssert.Assert, resp)
	return appAssert
}

func (request TestRequest) TestOnMessage(t *testing.T, existingAppAssert *ExistingAppAssert, status int, message string) ExistingAppAssert {
	request.testOn(t, existingAppAssert, status, "message", message)
	return *existingAppAssert
}

func (request TestRequest) TestOnMessageAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, status int, message string, dbTester Tester) ExistingAppAssert {
	appAssert, resp := request.testOn(t, existingAppAssert, status, "message", message)
	dbTester(appAssert.App, appAssert.Assert, resp)
	return appAssert
}

type Tester func(app TestApp, assert *assert.A, resp *http.Response)

type TesterWithStatus struct {
	Status int
	Tester
}

func (request TestRequest) TestOnStatusAndDB(t *testing.T, existingAppAssert *ExistingAppAssert, dbTesterStatus TesterWithStatus) ExistingAppAssert {
	appAssert, resp := request.Test(t, existingAppAssert)
	app, assert := appAssert.App, appAssert.Assert

	assert.Equal(dbTesterStatus.Status, resp.StatusCode)

	dbTesterStatus.Tester(app, assert, resp)

	return appAssert
}
