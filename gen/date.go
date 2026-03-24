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

type DualDate struct {
	Year  int16 `json:"year"`
	Month int8  `json:"month"`
	Day   int8  `json:"day"`
	Dual  bool  `json:"dual"` // Whether Date is a dual date, e.g. 1733/34
}

type DateRange struct {
	Start DualDate `json:"start"`
	End   DualDate `json:"end"`
}

type Date struct {
	// Format ???
	Calendar CalendarCode `json:"calendar"`
	Span     DateRange    `json:"dateval"`
	Modifier ModifierCode `json:"modifier"`
	NewYear  NewYearCode  `json:"newyear"`
	Quality  QualityFlags `json:"quality"`
	SortVal  int          `json:"sortval"`
	Text     string       `json:"text"`
}
