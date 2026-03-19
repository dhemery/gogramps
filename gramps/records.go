package gramps

type Citation struct {
	PrimaryObject
	Page       string      `xml:"page"`
	Confidence uint8       `xml:"confidence"`
	Date       DateVal     `xml:"datestr"`
	Attributes []Attribute `xml:"attribute"`
	Media      []MediaRef  `xml:"objref"`
	Notes      []NoteRef   `xml:"noteref"`
	Sources    []SourceRef `xml:"sourceref"`
}

type Event struct {
	PrimaryObject
	Type        string        `xml:"type"`
	Date        DateVal       `xml:"dateval"`
	Place       PlaceRef      `xml:"place"`
	Description string        `xml:"description"`
	Attributes  []Attribute   `xml:"attribute"`
	Citations   []CitationRef `xml:"citationref"`
	Notes       []NoteRef     `xml:"noteref"`
	Media       []MediaRef    `xml:"objref"`
}

type Family struct {
	PrimaryObject
	Rel        FamilyType           `xml:"rel"`
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
	PrimaryObject
	File       MediaFile     `xml:"file"`
	Date       DateVal       `xml:"dateval"`
	Attributes []Attribute   `xml:"attribute"`
	Citations  []CitationRef `xml:"citationref"`
	Notes      []NoteRef     `xml:"noteref"`
}

type Note struct {
	PrimaryObject
	Type string `xml:"type,attr"`
	Text string `xml:"text"`
}

type Person struct {
	PrimaryObject
	Unknown
	Privacy
	Gender        string         `xml:"gender"`
	Name          PersonName     `xml:"name"`
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

type Place struct {
	PrimaryObject
	Type          string        `xml:"type,attr"`
	Name          PlaceName     `xml:"pname"`
	Coordinates   Coordinates   `xml:"coord"`
	EncompassedBy []PlaceRef    `xml:"placeref"`
	Citations     []CitationRef `xml:"citationref"`
	Media         []MediaRef    `xml:"objref"`
	Notes         []NoteRef     `xml:"noteref"`
	URLs          []URL         `xml:"url"`
}

type Repository struct {
	PrimaryObject
	Name      string    `xml:"rname"`
	Type      string    `xml:"type"`
	Addresses []Address `xml:"address"`
	Notes     []NoteRef `xml:"noteref"`
	URLs      []URL     `xml:"url"`
}

type Source struct {
	PrimaryObject
	Title        string          `xml:"stitle"`
	Author       string          `xml:"sauthor"`
	PubInfo      string          `xml:"spubinfo"`
	Attributes   []Attribute     `xml:"attribute"`
	Media        []MediaRef      `xml:"objref"`
	Notes        []NoteRef       `xml:"noteref"`
	Repositories []RepositoryRef `xml:"reporef"`
}

