// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/newsletter"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Newsletter is the model entity for the Newsletter schema.
type Newsletter struct {
	config `json:"-"`
	// ID of the ent.
	// The unique identifier of the entity.
	ID uuid.UUID `json:"id,omitempty"`
	// The time when the entity was updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The time when the entity was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Newsletter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case newsletter.FieldUpdatedAt, newsletter.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case newsletter.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Newsletter", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Newsletter fields.
func (n *Newsletter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case newsletter.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				n.ID = *value
			}
		case newsletter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case newsletter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Newsletter.
// Note that you need to call Newsletter.Unwrap() before calling this method if this Newsletter
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Newsletter) Update() *NewsletterUpdateOne {
	return NewNewsletterClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Newsletter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Newsletter) Unwrap() *Newsletter {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Newsletter is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Newsletter) String() string {
	var builder strings.Builder
	builder.WriteString("Newsletter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Newsletters is a parsable slice of Newsletter.
type Newsletters []*Newsletter