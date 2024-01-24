# Using the Integration Testing Helpers

The integration testing helpers are a set of functions that reduce the boilerplate code required to write integration tests. They are located in the `backend/tests/helpers.go`.

## Modeling a Request with `TestRequest`

You can model a request with the `TestRequest` struct:

```go
type TestRequest struct {
 Method  string
 Path    string
 Body    *map[string]interface{}
 Headers *map[string]string
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
   "category_id": 1,
  },
}
```

This handles a lot of the logic for you, for example, if the body is not nil, it will be marshalled into JSON and the `Content-Type` header will be set to `application/json`.

## Testing that a Request Returns a XXX Status Code

Say you want to test hitting the `[APP_ADDRESS]/health` endpoint with a GET request returns a `200` status code.

```go
TestRequest{
  Method: fiber.MethodGet,
  Path:   "/health",
 }.TestOnStatus(t, nil, fiber.StatusOK).Close()
```

## Testing that a Request Returns a XXX Status Code and Assert Something About the Database

Say you want to test that a creating a catgory with POST `[APP_ADDRESS]/api/v1/categories/` returns a `201`

```go
TestRequest{
  Method: fiber.MethodPost,
  Path:   "/api/v1/categories/",
  Body: SampleCategoryFactory(),
 }.TestOnStatusAndDB(t, nil,
  DBTesterWithStatus{
   Status:   fiber.StatusCreated,
   DBTester: AssertSampleCategoryBodyRespDB,
  },
 ).Close()
```

### DBTesters

Often times there are common assertions you want to make about the database, for example, if the object in the response is the same as the object in the database. We can create a lambda function that takes in the `TestApp`, `*assert.A`, and `*http.Response` and makes the assertions we want. We can then pass this function to the `DBTesterWithStatus` struct.

```go
func AssertSampleCategoryBodyRespDB(app TestApp, assert *assert.A, resp *http.Response) {
 AssertCategoryWithIDBodyRespDB(app, assert, resp, 1, SampleCategoryFactory())
}

func AssertCategoryWithIDBodyRespDB(app TestApp, assert *assert.A, resp *http.Response, id uint, body *map[string]interface{}) {
 var respCategory models.Category

 err := json.NewDecoder(resp.Body).Decode(&respCategory)

 assert.NilError(err)

 var dbCategory models.Category

 err = app.Conn.First(&dbCategory, id).Error

 assert.NilError(err)

 assert.Equal(dbCategory.ID, respCategory.ID)
 assert.Equal(dbCategory.Name, respCategory.Name)

 assert.Equal((*body)["name"].(string), dbCategory.Name)
}
```

### Existing App Asserts

Since the test suite creates a new database for each test, we can have a deterministic database state for each test. However, what if we have a multi step test that depends on the previous steps database state? That is where `ExistingAppAssert` comes in! This will allow us to keep using the database from a previous step in the test.

Consider this example, to create a tag, we need to create a category first. This is a multi step test, so we need to use `ExistingAppAssert` to keep the database state from the previous step.

```go
TestRequest{
  Method: fiber.MethodPost,
  Path:   "/api/v1/categories/",
  Body: SampleCategoryFactory(),
 }.TestOnStatusAndDB(t, nil,
  DBTesterWithStatus{
   Status:   fiber.StatusCreated,
   DBTester: AssertSampleCategoryBodyRespDB,
  },
 )

TestRequest{
  Method: fiber.MethodPost,
  Path:   "/api/v1/tags/",
  Body: SampleTagFactory(),
 }.TestOnStatusAndDB(t, &appAssert,
  DBTesterWithStatus{
   Status:   fiber.StatusCreated,
   DBTester: AssertSampleTagBodyRespDB,
  },
 ).Close()
```

### Why Close?

This closes the connection to the database. This is important because if you don't close the connection, we will run out of available connections and the tests will fail. **Call this on the last test request of a test**

## Testing that a Request Returns the Correct Error (Status Code and Message), and Assert Something About the Database

Say you want to test a bad request to POST `[APP_ADDRESS]/api/v1/categories/` endpoint returns a `400` status code, the message is `failed to process the request`, and that a category was not created. We can leverage our errors defined in the error package to do this!

```go
 TestRequest{
  Method: fiber.MethodPost,
  Path:   "/api/v1/categories/",
  Body: &map[string]interface{}{
   "name": 1231,
  },
 }.TestOnStatusMessageAndDB(t, nil,
  ErrorWithDBTester{
   Error:    errors.FailedToParseRequestBody,
   DBTester: AssertNoCategories,
  },
 ).Close()
```
