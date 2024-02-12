# Jargon

### Embeddings
**Problem**: We have arbitrary-dimension data, such as descriptions for clubs, or searches for 
events. Given a piece of this arbitrary-dimension data (search, club desc.) we want to find 
other arbitrary-dimension data that is similar to it; think 2 club descriptions where both clubs 
are acapella groups, 2 search queries that are both effectively looking for professional 
fraternities, etc. **Solution**: Transform the arbitrary-dimension data to fixed-dimension data, 
say, a vector of floating-point numbers that is *n*-elements large. Make the transformation in 
such a way that similar arbitrary-dimension pieces of data will also have similar 
fixed-dimension data, i.e vectors that are close together (think Euclidean distance). **How do 
we do this transformation**: Train a machine learning model on large amounts of text, and then 
use the model to make vectors. **So what's an embedding?** Formally, when we 
refer to the embedding of a particular object, we refer to the vector created by feeding that 
object through the machine-learning model.

This is arguably the most complex/unintuitive part of understanding search, so here are some extra 
resources:
- [What are embeddings?](https://www.cloudflare.com/learning/ai/what-are-embeddings/) 
- [fastai book - Chapters 10 and 12 are both about natural language processing](https://github.com/fastai/fastbook)
- [Vector Embeddings for Developers: The Basics](https://www.pinecone.io/learn/vector-embeddings-for-developers/)

### OpenAI API
**Problem:**: We need a machine learning model to create the embeddings. **Solution:** Use 
OpenAI's api to create the embeddings for us; we send text over a REST api and we get a back a 
vector that represents that text's embedding.

### PineconeDB
**Problem**: We've created a bunch of embeddings for our club descriptions (or event 
descriptions, etc.), we now need a place to store them and a way to search through them (with an 
embedding for a search query) **Solution**: PineconeDB is a vector database that allows us to 
upload our embeddings and then query them by giving a vector to find similar ones to.

# How to create searchable objects for fun and fame and profit

```golang
package search

// in backend/search/searchable.go
type Searchable interface {
	SearchId() string
	Namespace() string
	EmbeddingString() string
}

// in backend/search/pinecone.go
type PineconeClientInterface interface {
	Upsert(item Searchable) *errors.Error
	Delete(item Searchable) *errors.Error
	Search(item Searchable, topK int) ([]string, *errors.Error)
}
```

1. Implement the `Searchable` interface on whatever model you want to make searchable. 
   `Searchable` requires 3 methods:
    - `SearchId()`: This should return a unique id that can be used to store a model entry's 
      embedding (if you want to store it at all) in PineconeDB. In practice, this should be the 
      entry's UUID.
    - `Namespace()`: Namespaces are to PineconeDB what tables are to PostgreSQL. Searching in 
      one namespace will only retrieve vectors in that namespace. In practice, this should be 
      unique to the model type (i.e `Club`, `Event`, etc.)
    - `EmbeddingString()`: This should return the string you want to feed into the OpenAI API 
      and create an embedding for. In practice, create a string with the fields you think will 
      affect the embedding all appended together, and/or try repeating a field multiple times in 
      the string to see if that gives a better search experience.
2. Use a `PineconeClientInterface` and call `Upsert` with your searchable object to send it to the 
   database, and `Delete` with your searchable object to delete it from the database. Upserts 
   should be done on creation and updating of a model entry,  and deletes should be done on 
   deleting of a model entry. In practice, a `PineconeClientInterface` should be passed in to 
   the various services in `backend/server.go`, similar to how `*gorm.DB` and `*validator.
   Validator` instances are passed in.

# How to search for fun and fame and profit

TODO: (probably create a searchable object that just uses namespace and embeddingstring, pass to 
pineconeclient search)