package partyHandler

import "github.com/google/uuid"

var (
	vampParty = party{
		Title:       "80's Vampire Night Club",
		Id:          uuid.NewString(),
		TimeOfDay:   []string{"Evening", "Night"},
		Theme:       "Gothic vampire",
		Decorations: nil,
		Drinks:      nil,
		Food:        nil,
	}
	gnomeParty = party{
		Title:       "Gnomes Frollicing in the Garden",
		Id:          uuid.NewString(),
		TimeOfDay:   []string{"Morning", "Afternoon", "Evening"},
		Theme:       "Fantasy cozy gnomes",
		Decorations: nil,
		Drinks:      nil,
		Food:        nil,
	}
)

func GetParties(req partyRequest) ([]party, error) {
	return []party{gnomeParty, vampParty}, nil
}
