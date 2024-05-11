# Using the Integration Testing Helpers

The integration testing helpers are a set of functions that reduce the boilerplate code required to write integration tests. They are located in the `backend/tests/api/helpers.go`.

## Modeling a Request with `TestRequest`

You can model a request with the `TestRequest` struct:

```go
type TestRequest struct {
 Method             string
 Path               string
 Body               *map[string]interface{}
 Headers            *map[string]string
 Role               *models.UserRole
 TestUserIDReplaces *string
}
```

Since `Body` and `Headers` are pointers, if they don't set them when creating a `TestRequest`, they will be `nil`.

Here is an example of creating a `TestRequest`, notice how instead of saying `Headers: nil`, we can simply omit the `Headers` field.

```go
TestRequest{
  Method: fiber.MethodPost,
  Path:   "/api/v1/tags/",
  Body: &map[string]interface{}{
   "name":        tagName,
   "category_id": uuid.New(),
  },
}
```

This handles a lot of the logic for you, for example, if the body is not nil, it will be marshalled into JSON and the `Content-Type` header will be set to `application/json`.

Role represents the desired role for the request to be sent with, see `models.UserRole` for the available roles. This allows us to test our permission based authentication. Since this is a pointer, if it is nil, the request will be sent without a role.

TestUserIDReplaces allows us to use the userID of the created user in our tests. For example, say the path is `/api/v1/users/:userID/`, we can use `h.StringToPointer(":userID")` to replace the path param with the dynamically generated userID. The `:` prefix indicates to attempt to replace in the path. If there isn't a match in the path, nothing will happen. Else, without a `:` prefix, it will attempt to replace the corresponding JSON key in the path. For example, if the userID is passed in the body, use `h.StringToPointer("user_id")` and it will attempt the value of the key `user_id` in the body. Again, if there isn't a match, nothing will happen.

## Initializing a Test

Simply call `h.InitTest(t)` to initialize a test. This will create a new database for the test and return an `ExistingAppAssert` struct that can be reused throughout the sequence of requests. The `h.InitTest(...)` function also accepts optional `TestAppConfigurator`s that allow you to configure the test app before the test starts. You can define additional configurators in the `backend/tests/api/helpers/test.go` file.

## Existing App Asserts

Since the test suite creates a new database for each test, we can have a deterministic database state for each test. However, what if we have a multi step test that depends on the previous steps database state? That is where `ExistingAppAssert` comes in! This will allow us to keep using the database from a previous step in the test. They also allow us to configure & use (or not use) a user across a sequence of requests. Additionally, the signature of tests allow for a chaining of tests, so you can run multiple tests in a sequence.

```go
appAssert.TestOnStatus(
  h.TestRequest{
   Method:             fiber.MethodPost,
   Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
   Role:               &models.Super,
   TestUserIDReplaces: h.StringToPointer(":userID"),
  },
  http.StatusCreated,
 ).TestOnStatusAndTester(
  h.TestRequest{
   Method:             fiber.MethodDelete,
   Path:               fmt.Sprintf("/api/v1/users/:userID/follower/%s", clubUUID),
   Role:               &models.Super,
   TestUserIDReplaces: h.StringToPointer(":userID"),
  },
  h.TesterWithStatus{
   Status: http.StatusNoContent,
   Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
    var user models.User

    err := eaa.App.Conn.Where("id = ?", eaa.App.TestUser.UUID).Preload("Follower").First(&user)

    eaa.Assert.NilError(err)

    eaa.Assert.Equal(1, len(user.Follower))

    var club models.Club

    err = eaa.App.Conn.Where("id = ?", clubUUID).Preload("Follower").First(&club)

    eaa.Assert.NilError(err)

    eaa.Assert.Equal(0, len(club.Follower))
   },
  },
 ).Close()
  ```

## Testers

Often times there are common assertions you want to make about the database, for example, if the object in the response is the same as the object in the database. We can create a lambda function that takes in the `TestApp`, `*assert.A`, and `*http.Response` and makes the assertions we want. We can then pass this function to the `TesterWithStatus` struct.

```go
func AssertSampleCategoryBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response) uuid.UUID {
 return AssertCategoryBodyRespDB(eaa, resp, SampleCategoryFactory())
}

func AssertCategoryBodyRespDB(eaa h.ExistingAppAssert, resp *http.Response, body *map[string]interface{}) uuid.UUID {
 var respCategory models.Category

 err := json.NewDecoder(resp.Body).Decode(&respCategory)

 eaa.Assert.NilError(err)

 var dbCategories []models.Category

 err = eaa.App.Conn.Find(&dbCategories).Error

 eaa.Assert.NilError(err)

 eaa.Assert.Equal(1, len(dbCategories))

 dbCategory := dbCategories[0]

 eaa.Assert.Equal(dbCategory.ID, respCategory.ID)
 eaa.Assert.Equal(dbCategory.Name, respCategory.Name)

 eaa.Assert.Equal((*body)["name"].(string), dbCategory.Name)

 return dbCategory.ID
}
```

## Why Close?

This closes the connection to the database. This is important because if you don't close the connection, we will run out of available connections and the tests will fail. **Always call `.close()` on the last test request of a test.** If you want to test a sequence of requests, call `.close()` on the last request of the sequence.

## Testing that a Request Returns a XXX Status Code

Say you want to test hitting the `[APP_ADDRESS]/health` endpoint with a GET request returns a `200` status code.

```go
 h.InitTest(t).TestOnStatus(
  h.TestRequest{
   Method: fiber.MethodGet,
   Path:   "/health",
  },
  http.StatusOK,
 ).Close()
```

## Testing that a Request Returns a XXX Status Code and Assert Something About the Database

Say you want to test that a creating a catgory with POST `[APP_ADDRESS]/api/v1/categories/` returns a `201`

```go
existingAppAssert.TestOnStatusAndTester(
  h.TestRequest{
   Method: fiber.MethodPost,
   Path:   "/api/v1/categories/",
   Body:   SampleCategoryFactory(),
   Role:   &models.Super,
  },
  h.TesterWithStatus{
   Status: http.StatusCreated,
   Tester: func(eaa h.ExistingAppAssert, resp *http.Response) {
    sampleCategoryUUID = AssertSampleCategoryBodyRespDB(eaa, resp)
   },
  },
 )
```

## Testing that a Request Returns the Correct Error (Status Code and Message), and Assert Something About the Database

Say you want to test a bad request to POST `[APP_ADDRESS]/api/v1/categories/` endpoint returns a `400` status code, the message is `failed to process the request`, and that a category was not created. We can leverage our errors defined in the error package to do this!

```go
 h.InitTest(t).TestOnStatusAndTester(
  h.TestRequest{
   Method: fiber.MethodPost,
   Path:   "/api/v1/categories/",
   Body: &map[string]interface{}{
    "name": 1231,
   },
   Role: &models.Super,
  },
  h.ErrorWithTester{
   Error:  errors.FailedToParseRequestBody,
   Tester: AssertNoCategories,
  },
 ).Close()
```
