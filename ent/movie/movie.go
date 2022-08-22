// Code generated by ent, DO NOT EDIT.

package movie

const (
	// Label holds the string label denoting the movie type in the database.
	Label = "movie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldEpisodeID holds the string denoting the episode_id field in the database.
	FieldEpisodeID = "episode_id"
	// FieldOpeningCrawl holds the string denoting the opening_crawl field in the database.
	FieldOpeningCrawl = "opening_crawl"
	// FieldDirector holds the string denoting the director field in the database.
	FieldDirector = "director"
	// FieldProducer holds the string denoting the producer field in the database.
	FieldProducer = "producer"
	// FieldReleaseDate holds the string denoting the release_date field in the database.
	FieldReleaseDate = "release_date"
	// FieldCharacters holds the string denoting the characters field in the database.
	FieldCharacters = "characters"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldEdited holds the string denoting the edited field in the database.
	FieldEdited = "edited"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgePeople holds the string denoting the people edge name in mutations.
	EdgePeople = "people"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// Table holds the table name of the movie in the database.
	Table = "movies"
	// PeopleTable is the table that holds the people relation/edge. The primary key declared below.
	PeopleTable = "movie_people"
	// PeopleInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	PeopleInverseTable = "characters"
	// CommentsTable is the table that holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "movie_comments"
)

// Columns holds all SQL columns for movie fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldEpisodeID,
	FieldOpeningCrawl,
	FieldDirector,
	FieldProducer,
	FieldReleaseDate,
	FieldCharacters,
	FieldCreated,
	FieldEdited,
	FieldURL,
}

var (
	// PeoplePrimaryKey and PeopleColumn2 are the table columns denoting the
	// primary key for the people relation (M2M).
	PeoplePrimaryKey = []string{"movie_id", "character_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// EpisodeIDValidator is a validator for the "episode_id" field. It is called by the builders before save.
	EpisodeIDValidator func(int) error
	// DefaultOpeningCrawl holds the default value on creation for the "opening_crawl" field.
	DefaultOpeningCrawl string
)
