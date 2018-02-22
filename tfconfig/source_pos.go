package tfconfig

import (
	legacyhcltoken "github.com/hashicorp/hcl/hcl/token"
	"github.com/hashicorp/hcl2/hcl"
)

// SourcePos is a pointer to a particular location in a source file.
//
// This type is embedded into other structs to allow callers to locate the
// definition of each described module element. The SourcePos of an element
// is usually the first line of its definition, although the definition can
// be a little "fuzzy" with JSON-based config files.
type SourcePos struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
}

func sourcePos(filename string, line int) SourcePos {
	return SourcePos{
		Filename: filename,
		Line:     line,
	}
}

func sourcePosHCL(rng hcl.Range) SourcePos {
	return SourcePos{
		Filename: rng.Filename,
		Line:     rng.Start.Line,
	}
}

func sourcePosLegacyHCL(pos legacyhcltoken.Pos) SourcePos {
	return SourcePos{
		Filename: pos.Filename,
		Line:     pos.Line,
	}
}
