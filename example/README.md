# Example

## Usage

```go
fields, errs := naive.DistinctCount(Movie{}, movies, "Year", "Genres")
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
Science Fiction = 83
Animated = 79
Biography = 59
Fantasy = 55
Supernatural = 54
Crime = 51
Musical = 45
Adventure = 43
Superhero = 39
Documentary = 29
Sports = 29
Mystery = 27
Historical = 26
War = 25
Teen = 20
Erotic = 17
Slasher = 17
Family = 14
Western = 13
Spy = 12
Noir = 10
Independent = 7
Live Action = 7
Satire = 7
Short = 6
Martial Arts = 5
Political = 4
Legal = 3
Performance = 3
Disaster = 3
Dance = 3
Found Footage = 3
```