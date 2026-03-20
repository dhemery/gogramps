package gramps

// Record is a mix-in for the fields common to all records in a Gramps database.
type Record struct {
	// Identifes the record in the table.
	Handle string `xml:"handle,attr"`
	// The date and time of the latest update,
	// represented as a number of seconds since the Unix epoch.
	Change uint64 `xml:"change,attr"`
}

// PrimaryRecord is a mix-in for the fields common to all Gramps primary record types.
type PrimaryRecord struct {
	Record
	Privacy
	Tags
	// The object's Gramps ID.
	ID string `xml:"id,attr"`
}

type Citation struct {
	PrimaryRecord
	Page       string      `xml:"page"`
	Confidence uint8       `xml:"confidence"`
	Date      []DateVal   `xml:"datestr"`
	Attributes []Attribute `xml:"attribute"`
	Media      []MediaRef  `xml:"objref"`
	Notes      []NoteRef   `xml:"noteref"`
	Sources    []SourceRef `xml:"sourceref"`
}

type Event struct {
	PrimaryRecord
	Type        string        `xml:"type"`
	Date       []DateVal     `xml:"dateval"`
	Place       PlaceRef      `xml:"place"`
	Description string        `xml:"description"`
	Attributes  []Attribute   `xml:"attribute"`
	Citations   []CitationRef `xml:"citationref"`
	Notes       []NoteRef     `xml:"noteref"`
	Media       []MediaRef    `xml:"objref"`
}

type Family struct {
	PrimaryRecord
	Rel        FamilyType    `xml:"rel"`
	Father     PersonRef     `xml:"father"`
	Mother     PersonRef     `xml:"mother"`
	Children   []PersonRef   `xml:"childref"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Events     []EventRef    `xml:"eventref"`
	Media      []MediaRef    `xml:"objref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type Media struct {
	PrimaryRecord
	File       MediaFile     `xml:"file"`
	Date      []DateVal     `xml:"dateval"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type Note struct {
	PrimaryRecord
	Type string `xml:"type,attr"`
	Text string `xml:"text"`
}

type Person struct {
	PrimaryRecord
	Unknown
	Privacy
	Gender        string         `xml:"gender"`
	Names         []PersonName   `xml:"name"`
	Addresses     []Address      `xml:"address"`
	Associations  []PersonRef    `xml:"personref"`
	ChildOf       []FamilyRef    `xml:"childof"`
	ParentIn      []FamilyRef    `xml:"parentin"`
	Attributes    []Attribute    `xml:"attribute"`
	Citations     []CitationRef  `xml:"citationref"`
	Events        []EventRef     `xml:"eventref"`
	LdsOrdinances []LDSOrdinance `xml:"lds_ord"`
	Media         []MediaRef     `xml:"objref"`
	Notes         []NoteRef      `xml:"noteref"`
	URLs          []URL          `xml:"url"`
}

func (p Person) hasUnknowns() bool {
	if p.Unknown.hasUnknowns() {
		return true
	}
	for _, n := range p.Names {
		if n.hasUnknowns() {
			return true
		}
	}
	return false
}

type Place struct {
	PrimaryRecord
	Type          string        `xml:"type,attr"`
	Name          PlaceName     `xml:"pname"`
	Coordinates   Coordinates   `xml:"coord"`
	EncompassedBy []PlaceRef    `xml:"placeref"`
	Citations     []CitationRef `xml:"citationref"`
	Media         []MediaRef    `xml:"objref"`
	Notes         []NoteRef     `xml:"noteref"`
	URLs          []URL         `xml:"url"`
}

type Tag struct {
	Record
	Name     string `xml:"name,attr"`
	Color    string `xml:"color,attr"`
	Priority uint   `xml:"priority,attr"`
}

type Repository struct {
	PrimaryRecord
	Name      string    `xml:"rname"`
	Type      string    `xml:"type"`
	Addresses []Address `xml:"address"`
	Notes     []NoteRef `xml:"noteref"`
	URLs      []URL     `xml:"url"`
}

type Source struct {
	PrimaryRecord
	Title        string          `xml:"stitle"`
	Author       string          `xml:"sauthor"`
	PubInfo      string          `xml:"spubinfo"`
	Attributes   []Attribute     `xml:"attribute"`
	Media        []MediaRef      `xml:"objref"`
	Notes        []NoteRef       `xml:"noteref"`
	Repositories []RepositoryRef `xml:"reporef"`
}
