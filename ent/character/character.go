// Code generated by ent, DO NOT EDIT.

package character

const (
	// Label holds the string label denoting the character type in the database.
	Label = "character"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldMass holds the string denoting the mass field in the database.
	FieldMass = "mass"
	// FieldHairColor holds the string denoting the hair_color field in the database.
	FieldHairColor = "hair_color"
	// FieldSkinColor holds the string denoting the skin_color field in the database.
	FieldSkinColor = "skin_color"
	// FieldEyeColor holds the string denoting the eye_color field in the database.
	FieldEyeColor = "eye_color"
	// FieldBirthYear holds the string denoting the birth_year field in the database.
	FieldBirthYear = "birth_year"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldFilms holds the string denoting the films field in the database.
	FieldFilms = "films"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldEdited holds the string denoting the edited field in the database.
	FieldEdited = "edited"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeFilm holds the string denoting the film edge name in mutations.
	EdgeFilm = "film"
	// Table holds the table name of the character in the database.
	Table = "characters"
	// FilmTable is the table that holds the film relation/edge. The primary key declared below.
	FilmTable = "movie_people"
	// FilmInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	FilmInverseTable = "movies"
)

// Columns holds all SQL columns for character fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldHeight,
	FieldMass,
	FieldHairColor,
	FieldSkinColor,
	FieldEyeColor,
	FieldBirthYear,
	FieldGender,
	FieldFilms,
	FieldCreated,
	FieldEdited,
	FieldURL,
}

var (
	// FilmPrimaryKey and FilmColumn2 are the table columns denoting the
	// primary key for the film relation (M2M).
	FilmPrimaryKey = []string{"movie_id", "character_id"}
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
