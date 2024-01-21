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
  Method: "POST",
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
  Method: "GET",
  Path:   "/health",
 }.TestOnStatus(t, nil, 200)
```

## Testing that a Request Returns a XXX Status Code and Assert Something About the Database

Say you want to test that a creating a catgory with POST `[APP_ADDRESS]/api/v1/categories/` returns a `201`

```go
TestRequest{
  Method: "POST",
  Path:   "/api/v1/categories/",
  Body: &map[string]interface{}{
   "category_name": categoryName,
  },
 }.TestOnStatusAndDB(t, nil,
  DBTesterWithStatus{
   Status:   201,
   DBTester: AssertRespCategorySameAsDBCategory,
  },
 )
```

### DBTesters

Often times there are common assertions you want to make about the database, for example, if the object in the response is the same as the object in the database. We can create a lambda function that takes in the `TestApp`, `*assert.A`, and `*http.Response` and makes the assertions we want. We can then pass this function to the `DBTesterWithStatus` struct.

```go
var AssertRespCategorySameAsDBCategory = func(app TestApp, assert *assert.A, resp *http.Response) {
 var respCategory models.Category

 err := json.NewDecoder(resp.Body).Decode(&respCategory)

 assert.NilError(err)

 dbCategory, err := transactions.GetCategory(app.Conn, respCategory.ID)

 assert.NilError(err)

 assert.Equal(dbCategory, &respCategory)
}
```

### Existing App Asserts

Since the test suite creates a new database for each test, we can have a deterministic database state for each test. However, what if we have a multi step test that depends on the previous steps database state? That is where `ExistingAppAssert` comes in! This will allow us to keep using the database from a previous step in the test.

Consider this example, to create a tag, we need to create a category first. This is a multi step test, so we need to use `ExistingAppAssert` to keep the database state from the previous step.

```go
appAssert := TestRequest{
  Method: "POST",
  Path:   "/api/v1/categories/",
  Body: &map[string]interface{}{
   "category_name": categoryName,
  },
 }.TestOnStatusAndDB(t, nil,
  DBTesterWithStatus{
   Status:   201,
   DBTester: AssertRespCategorySameAsDBCategory,
  },
 )

TestRequest{
  Method: "POST",
  Path:   "/api/v1/tags/",
  Body: &map[string]interface{}{
   "name":        tagName,
   "category_id": 1,
  },
 }.TestOnStatusAndDB(t, &appAssert,
  DBTesterWithStatus{
   Status:   201,
   DBTester: AssertRespTagSameAsDBTag,
  },
 )
```

## Testing that a Request Returns a XXX Status Code, Assert Something About the Message, and Assert Something About the Database

Say you want to test a bad request to POST `[APP_ADDRESS]/api/v1/categories/` endpoint returns a `400` status code, the message is `failed to process the request`, and that a category was not created.

```go
TestRequest{
  Method: "POST",
  Path:   "/api/v1/categories/",
  Body: &map[string]interface{}{
   "category_name": 1231,
  },
 }.TestOnStatusMessageAndDB(t, nil,
  StatusMessageDBTester{
   MessageWithStatus: MessageWithStatus{
    Status:  400,
    Message: "failed to process the request",
   },
   DBTester: AssertNoCategories,
  },
 )
```
