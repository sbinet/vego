package vego

// Vis is a lyra visualization
type Vis struct {
	Width   int     `json:"width"`
	Height  int     `json:"height"`
	Padding Padding `json:"padding"`
	Data    []Data  `json:"data"`
	Scales  []Scale `json:"scales"`
	Axes    []Axis  `json:"axes"`
	Marks   []Mark  `json:"marks"`
}

type Padding struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Bottom int `json:"bottom"`
	Rigth  int `json:"right"`
}

type Item2 struct {
	Name  string
	Value Value
}

type Item map[string]Value

type Value interface{}

type Data struct {
	Name      string  `json:"name"`      // a unique name for the data set
	Format    Value   `json:"format"`    // specifies the format of the data file, if loaded from a URL (json|csv|tsv|topojson|treejson)
	Values    []Item  `json:"values"`    // actual data set to use
	Source    Item    `json:"source"`    // name of another data set to use as the source for this data set
	Url       string  `json:"url"`       // URL from which to load the data set
	Transform []Value `json:"transform"` // array of transforms to perform on the data
}

type Scale struct {
	Name      string    `json:"name"`      // unique name for the scale
	Type      ScaleType `json:"type"`      // type of scale
	Domain    DataRef   `json:"domain"`    // domain of the scale
	DomainMin DataRef   `json:"domainMin"` // sets the minimum value in the scale domain
	DomainMax DataRef   `json:"domainMax"` // sets the maximum value in the scale domain
	Range     string    `json:"range"`     // range of the scale, representing the set of visual values
	RangeMin  Value     `json:"rangeMin"`  // sets the minimum value in the scale range
	RangeMax  Value     `json:"rangeMax"`  // sets the maximum value in the scale range
	Reverse   bool      `json:"reverse"`   // flips the scale range
	Round     bool      `json:"round"`     // rounds numeric output values to integers
	Nice      bool      `json:"nice"`
}

type ScaleType string

// List of allowed ScaleTypes ("linear"=default|"ordinal"|"time"|"utc"|)
const (
	LinearScale    ScaleType = "linear"
	OrdinalScale             = "ordinal"
	TimeScale                = "time"
	UTCScale                 = "utc"
	LogScale                 = "log"
	PowScale                 = "pow"
	SqrtScale                = "sqrt"
	QuantileScale            = "quantile"
	QuantizeScale            = "quantize"
	ThresholdScale           = "threshold"
)

type Range interface {
	isRange()
}

type RangeString string

func (RangeString) isRange() {}

type RangeRef DataRef

func (RangeRef) isRange() {}

type DataRef struct {
	Data  string `json:"data"`
	Field string `json:"field"`
}

type Domain interface {
	isDomain()
}

type DomainRef DataRef

func (DomainRef) isDomain() {}

type Axis struct {
	Type          string  `json:"type"`          // the type of axis ("x"|"y")
	Scale         string  `json:"scale"`         // the name of the scale backing the axis component
	Orient        string  `json:"orient"`        // the orientation of the axis ("top"|"bottom"|"left"|"right")
	Title         string  `json:"title"`         // title for the axis
	TitleOffset   int     `json:"titleOffset"`   // offset in pixels from the axis at which to place the title
	Format        string  `json:"format"`        // formatting pattern for axis labels (D3's format pattern)
	Ticks         int     `json:"ticks"`         // desired number of ticks
	Values        []Value `json:"values"`        // explicitly set the visible axis tick values
	Subdivide     int     `json:"subdivide"`     // number of minor ticks between major ticks
	TickPadding   int     `json:"tickPadding"`   // padding in pixels between ticks and text labels
	TickSize      int     `json:"tickSize"`      // size in pixels of major, minor and text labels
	TickSizeMajor int     `json:"tickSizeMajor"` // size in pixels of major ticks
	TickSizeMinor int     `json:"tickSizeMinor"` // size in pixels of minor ticks
	TickSizeEnd   int     `json:"tickSizeEnd"`   // size in pixels of end ticks
	Offset        Value   `json:"offset"`        // offset in pixels by which to displace the axis from the edge of the enclosing group or data rectangle
	Layer         string  `json:"layer"`         // indicates whether the axis should be placed above or below data marks ("front"=default|"back")
	Grid          bool    `json:"grid"`          // flag to indicate if gridlines should be created in addition to ticks
	Properties    Value   `json:"properties"`    // optional mark property definitions for custom axis styling
}

type Mark struct {
	Type       string     `json:"type"`
	From       Provenance `json:"from"`
	Properties Properties `json:"properties"`
}

type Provenance struct {
	Data string `json:"data"`
}

type Properties struct {
	Enter  EnterProp  `json:"enter"`
	Update UpdateProp `json:"update"`
	Hover  HoverProp  `json:"hover"`
}

type EnterProp struct {
	X     Item `json:"x"`
	Width Item `json:"width"`
	Y     Item `json:"y"`
	Y2    Item `json:"y2"`
}

type HoverProp struct {
	Fill Item `json:"fill"`
}

type UpdateProp struct {
	Fill Item `json:"fill"`
}
