package gen

type ModifierCode int
type QualityFlags int
type CalendarCode int
type NewYearCode int

const (
	ModNone = ModifierCode(iota)
	ModBefore
	ModAfter
	ModAbout
	ModRange
	ModSpan
	ModTextOnly
	ModFrom
	ModTo
)

const (
	QualNone = QualityFlags((0x1 << iota) >> 1)
	QualEstimated
	QualCalculated
	QualInterpreted
)

const (
	CalGregorian = CalendarCode(iota)
	CalJulian
	CalHebrew
	CalFrench
	CalPersian
	CalIslamic
	CalSwedish
)

const (
	NewYearJan1 = NewYearCode(iota)
	NewYearMar1
	NewYearMar25
	NewYearSep1
)

type Date struct {
	Year  int16 `json:"year,omitzero"`
	Month int8  `json:"month,omitzero"`
	Day   int8  `json:"day,omitzero"`
	Dual  bool  `json:"dual,omitzero"` // Whether Date is a dual date, e.g. 1733/34
}

type DateBounds struct {
	Start Date `json:"start,omitzero"`
	End   Date `json:"end,omitzero"`
}

type DateValue struct {
	// Format ???
	Calendar CalendarCode `json:"calendar,omitzero"`
	Bounds   DateBounds   `json:"dateval,omitzero"`
	Modifier ModifierCode `json:"modifier,omitzero"`
	NewYear  NewYearCode  `json:"newyear,omitzero"`
	Quality  QualityFlags `json:"quality,omitzero"`
	SortVal  int          `json:"sortval,omitzero"`
	Text     string       `json:"text,omitzero"`
}
