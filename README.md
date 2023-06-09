# go-cardinality
go-cardinality is a Go library that calculates the cardinality and distinct count of values in a dataset, providing efficient and accurate estimations.

# Usage
- Retrieve all unique values of a specific field in a dataset, which is useful for creating enums or generating dimension tables.
- Analyze the distribution of values within a particular field in a dataset to gain insights into the most frequently occurring values.

```go
fields := naive.DistinctCount(Movie{}, movies, "Year", "Genres")
genres, err := fields.GetField("Genres")
genres.PrettyPrint()
```
```text
Comedy = 350
Drama = 338
Thriller = 194
Horror = 162
Action = 162
Romance = 117
...
```

Check examples here : [Examples](example/README.md)

# Progress

## Features

- [x] Naive approach using a map data structure
- [ ] Naive approach with concurrent processing for improved performance
- [ ] Naive approach with RxGo
- [ ] API for calculating cardinality in a list of objects
- [ ] Use HyperLogLog algorithm for accurate estimation of distinct values

## Types implemented
- [x] Type `int`
- [x] Type `string`
- [x] Type `[]string`
- [x] Type `[]int`

## Tests
- Run unit tests
```shell
make test
```
- Run benchmark tests
```shell
make bench
```
- Run coverage
```shell
make coverage
```

# Considerations
- Usage of reflect library : even if greatly discouraged, we needed to use it (moderately) to build generic methods based on struct fields and types (schema).

# Glossary
- **Cardinality** is a mathematical term. It translates into the number of elements in a set. In databases, cardinality refers to the relationships between the data in two database tables. Cardinality defines how many instances of one entity are related to instances of another entity.
- **Distribution** refers to the way data values are spread or organized within a dataset. It describes the frequency or occurrence of different values or groups of values in a specific attribute or field.

# Inspiration

- HyperLogLog algorithm: is a probabilistic data structure that can be used to approximate the number of distinct elements in a data set.
- Cardinality aggregation: https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-cardinality-aggregation.html

# Resources
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations-metrics-cardinality-aggregation.html
- https://towardsdatascience.com/count-distinct-metrics-at-scale-95a394c03f1
- https://pkg.go.dev/github.com/datadog/hyperloglog
- https://vertabelo.com/blog/cardinality-in-data-modeling/
- https://go.dev/blog/laws-of-reflection
