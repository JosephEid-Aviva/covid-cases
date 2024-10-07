# Aviva Zero User Story - A sorted list of dates by case volume

The API currently only fetches the first page of data from the public API endpoint.

The aims of this user story are to:

- get the first 10 pages of results
- return a list of results sorted by case counts descending e.g.

```sh
  [
    {
    "date": "2020-03-03",
    "count": 1300
    },
    {
    "date": "2020-03-09",
    "count": 1200
    },
    {
    "date": "2020-03-01",
    "count": 1100
    },
    {
    "date": "2020-02-28",
    "count": 879
    },
  ]
```

## What we would like to see as part of your solution

- Make use of Go's concurrency features

  - Can you make use of go routines and channels to optimise the code?

- Testing

  - How would you test this code?
