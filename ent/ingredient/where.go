// Code generated by ent, DO NOT EDIT.

package ingredient

import (
	"entgo.io/ent/dialect/sql"
	"github.com/syuparn/fridgesim/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Kind applies equality check predicate on the "kind" field. It's identical to KindEQ.
func Kind(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKind), v))
	})
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...string) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKind), v...))
	})
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...string) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKind), v...))
	})
}

// KindGT applies the GT predicate on the "kind" field.
func KindGT(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKind), v))
	})
}

// KindGTE applies the GTE predicate on the "kind" field.
func KindGTE(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKind), v))
	})
}

// KindLT applies the LT predicate on the "kind" field.
func KindLT(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKind), v))
	})
}

// KindLTE applies the LTE predicate on the "kind" field.
func KindLTE(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKind), v))
	})
}

// KindContains applies the Contains predicate on the "kind" field.
func KindContains(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKind), v))
	})
}

// KindHasPrefix applies the HasPrefix predicate on the "kind" field.
func KindHasPrefix(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKind), v))
	})
}

// KindHasSuffix applies the HasSuffix predicate on the "kind" field.
func KindHasSuffix(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKind), v))
	})
}

// KindEqualFold applies the EqualFold predicate on the "kind" field.
func KindEqualFold(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKind), v))
	})
}

// KindContainsFold applies the ContainsFold predicate on the "kind" field.
func KindContainsFold(v string) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKind), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Ingredient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Ingredient) predicate.Ingredient {
	return predicate.Ingredient(func(s *sql.Selector) {
		p(s.Not())
	})
}
