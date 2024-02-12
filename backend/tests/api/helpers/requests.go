package helpers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/GenerateNU/sac/backend/src/errors"
	"github.com/GenerateNU/sac/backend/src/models"

	"github.com/goccy/go-json"
)

type TestRequest struct {
	Method             string
	Path               string
	Body               *map[string]interface{}
	Headers            *map[string]string
	Role               *models.UserRole
	TestUserIDReplaces *string
}

func (app TestApp) Send(request TestRequest) (*http.Response, error) {
	address := fmt.Sprintf("%s%s", app.Address, request.Path)

	var req *http.Request

	if request.TestUserIDReplaces != nil {
		if strings.Contains(request.Path, *request.TestUserIDReplaces) {
			request.Path = strings.Replace(request.Path, *request.TestUserIDReplaces, app.TestUser.UUID.String(), 1)
			address = fmt.Sprintf("%s%s", app.Address, request.Path)
		}
		if request.Body != nil {
			if _, ok := (*request.Body)[*request.TestUserIDReplaces]; ok {
				(*request.Body)[*request.TestUserIDReplaces] = app.TestUser.UUID.String()
			}
		}
	}

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

func (request TestRequest) test(existingAppAssert ExistingAppAssert) (ExistingAppAssert, *http.Response) {
	if existingAppAssert.App.TestUser == nil && request.Role != nil {
		existingAppAssert.App.Auth(*request.Role)
	}

	resp, err := existingAppAssert.App.Send(request)

	existingAppAssert.Assert.NilError(err)

	return existingAppAssert, resp
}

func (existingAppAssert ExistingAppAssert) TestOnStatus(request TestRequest, status int) ExistingAppAssert {
	appAssert, resp := request.test(existingAppAssert)

	appAssert.Assert.Equal(status, resp.StatusCode)

	return appAssert
}

func (request *TestRequest) testOn(existingAppAssert ExistingAppAssert, status int, key string, value string) (ExistingAppAssert, *http.Response) {
	appAssert, resp := request.test(existingAppAssert)

	var respBody map[string]interface{}

	err := json.NewDecoder(resp.Body).Decode(&respBody)

	appAssert.Assert.NilError(err)

	appAssert.Assert.Equal(value, respBody[key].(string))

	appAssert.Assert.Equal(status, resp.StatusCode)

	return appAssert, resp
}

func (existingAppAssert ExistingAppAssert) TestOnError(request TestRequest, expectedError errors.Error) ExistingAppAssert {
	appAssert, _ := request.testOn(existingAppAssert, expectedError.StatusCode, "error", expectedError.Message)
	return appAssert
}

type ErrorWithTester struct {
	Error  errors.Error
	Tester Tester
}

func (existingAppAssert ExistingAppAssert) TestOnErrorAndTester(request TestRequest, errorWithTester ErrorWithTester) ExistingAppAssert {
	appAssert, resp := request.testOn(existingAppAssert, errorWithTester.Error.StatusCode, "error", errorWithTester.Error.Message)
	errorWithTester.Tester(appAssert, resp)
	return appAssert
}

func (existingAppAssert ExistingAppAssert) TestOnMessage(request TestRequest, status int, message string) ExistingAppAssert {
	request.testOn(existingAppAssert, status, "message", message)
	return existingAppAssert
}

func (existingAppAssert ExistingAppAssert) TestOnMessageAndTester(request TestRequest, status int, message string, tester Tester) ExistingAppAssert {
	appAssert, resp := request.testOn(existingAppAssert, status, "message", message)
	tester(appAssert, resp)
	return appAssert
}

type Tester func(eaa ExistingAppAssert, resp *http.Response)

type TesterWithStatus struct {
	Status int
	Tester
}

func (existingAppAssert ExistingAppAssert) TestOnStatusAndTester(request TestRequest, testerStatus TesterWithStatus) ExistingAppAssert {
	appAssert, resp := request.test(existingAppAssert)

	appAssert.Assert.Equal(testerStatus.Status, resp.StatusCode)

	testerStatus.Tester(appAssert, resp)

	return appAssert
}
