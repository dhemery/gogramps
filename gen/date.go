package gen

// Modifier Enum
// MOD_NONE = 0  # CODE
// MOD_BEFORE = 1
// MOD_AFTER = 2
// MOD_ABOUT = 3
// MOD_RANGE = 4
// MOD_SPAN = 5
// MOD_TEXTONLY = 6
// MOD_FROM = 7
// MOD_TO = 8

// Quality bit flags
// QUAL_NONE = 0  # BITWISE
// QUAL_ESTIMATED = 1
// QUAL_CALCULATED = 2
// # QUAL_INTERPRETED = 4 unused in source!!

// Calendar enum
// CAL_GREGORIAN = 0  # CODE
// CAL_JULIAN = 1
// CAL_HEBREW = 2
// CAL_FRENCH = 3
// CAL_PERSIAN = 4
// CAL_ISLAMIC = 5
// CAL_SWEDISH = 6
// CALENDARS = range(7)

// NewYear enum
// NEWYEAR_JAN1 = 0  # CODE
// NEWYEAR_MAR1 = 1
// NEWYEAR_MAR25 = 2
// NEWYEAR_SEP1 = 3

type DualDate struct {
	Year  int16 `json:"year"`
	Month int8  `json:"month"`
	Day   int8  `json:"day"`
	Dual  bool  `json:"dual"` // Whether Date is a dual date, e.g. 1733/34
}

type DateSpan struct {
	Start DualDate `json:"start"`
	End   DualDate `json:"end"`
}

type Date struct {
	// Format ???
	Calendar int      `json:"calendar"`
	Span     DateSpan `json:"dateval"`
	Modifier int      `json:"modifier"`
	NewYear  int      `json:"newyear"`
	Quality  int      `json:"quality"`
	SortVal  int      `json:"sortval"`
	Text     string   `json:"text"`
}
