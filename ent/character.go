// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samuelagm/moviex/ent/character"
)

// Character is the model entity for the Character schema.
type Character struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Height holds the value of the "height" field.
	Height string `json:"height,omitempty"`
	// Mass holds the value of the "mass" field.
	Mass string `json:"mass,omitempty"`
	// HairColor holds the value of the "hair_color" field.
	HairColor string `json:"hair_color,omitempty"`
	// SkinColor holds the value of the "skin_color" field.
	SkinColor string `json:"skin_color,omitempty"`
	// EyeColor holds the value of the "eye_color" field.
	EyeColor string `json:"eye_color,omitempty"`
	// BirthYear holds the value of the "birth_year" field.
	BirthYear string `json:"birth_year,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Films holds the value of the "films" field.
	Films []string `json:"films,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Edited holds the value of the "edited" field.
	Edited time.Time `json:"edited,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CharacterQuery when eager-loading is set.
	Edges CharacterEdges `json:"edges"`
}

// CharacterEdges holds the relations/edges for other nodes in the graph.
type CharacterEdges struct {
	// Film holds the value of the film edge.
	Film []*Movie `json:"film,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FilmOrErr returns the Film value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) FilmOrErr() ([]*Movie, error) {
	if e.loadedTypes[0] {
		return e.Film, nil
	}
	return nil, &NotLoadedError{edge: "film"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Character) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case character.FieldFilms:
			values[i] = new([]byte)
		case character.FieldID:
			values[i] = new(sql.NullInt64)
		case character.FieldName, character.FieldHeight, character.FieldMass, character.FieldHairColor, character.FieldSkinColor, character.FieldEyeColor, character.FieldBirthYear, character.FieldGender, character.FieldURL:
			values[i] = new(sql.NullString)
		case character.FieldCreated, character.FieldEdited:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Character", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Character fields.
func (c *Character) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case character.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case character.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case character.FieldHeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				c.Height = value.String
			}
		case character.FieldMass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mass", values[i])
			} else if value.Valid {
				c.Mass = value.String
			}
		case character.FieldHairColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hair_color", values[i])
			} else if value.Valid {
				c.HairColor = value.String
			}
		case character.FieldSkinColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field skin_color", values[i])
			} else if value.Valid {
				c.SkinColor = value.String
			}
		case character.FieldEyeColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field eye_color", values[i])
			} else if value.Valid {
				c.EyeColor = value.String
			}
		case character.FieldBirthYear:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birth_year", values[i])
			} else if value.Valid {
				c.BirthYear = value.String
			}
		case character.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				c.Gender = value.String
			}
		case character.FieldFilms:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field films", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Films); err != nil {
					return fmt.Errorf("unmarshal field films: %w", err)
				}
			}
		case character.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				c.Created = value.Time
			}
		case character.FieldEdited:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field edited", values[i])
			} else if value.Valid {
				c.Edited = value.Time
			}
		case character.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				c.URL = value.String
			}
		}
	}
	return nil
}

// QueryFilm queries the "film" edge of the Character entity.
func (c *Character) QueryFilm() *MovieQuery {
	return (&CharacterClient{config: c.config}).QueryFilm(c)
}

// Update returns a builder for updating this Character.
// Note that you need to call Character.Unwrap() before calling this method if this Character
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Character) Update() *CharacterUpdateOne {
	return (&CharacterClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Character entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Character) Unwrap() *Character {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Character is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Character) String() string {
	var builder strings.Builder
	builder.WriteString("Character(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(c.Height)
	builder.WriteString(", ")
	builder.WriteString("mass=")
	builder.WriteString(c.Mass)
	builder.WriteString(", ")
	builder.WriteString("hair_color=")
	builder.WriteString(c.HairColor)
	builder.WriteString(", ")
	builder.WriteString("skin_color=")
	builder.WriteString(c.SkinColor)
	builder.WriteString(", ")
	builder.WriteString("eye_color=")
	builder.WriteString(c.EyeColor)
	builder.WriteString(", ")
	builder.WriteString("birth_year=")
	builder.WriteString(c.BirthYear)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(c.Gender)
	builder.WriteString(", ")
	builder.WriteString("films=")
	builder.WriteString(fmt.Sprintf("%v", c.Films))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(c.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("edited=")
	builder.WriteString(c.Edited.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(c.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Characters is a parsable slice of Character.
type Characters []*Character

func (c Characters) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
