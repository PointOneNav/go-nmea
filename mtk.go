package nmea

const (
	// TypeMTK type for PMTK sentences
	TypeMTK = "PMTK"
)

// MTK is sentence for NMEA embedded command packet protocol.
// https://www.rhydolabz.com/documents/25/PMTK_A11.pdf
// https://www.sparkfun.com/datasheets/GPS/Modules/PMTK_Protocol.pdf
//
// The maximum length of each packet is restricted to 255 bytes which is longer than NMEA0183 82 bytes.
//
// Format: $PMTKxxx,c-c*hh<CR><LF>
// Example: $PMTK000*32<CR><LF>
//			$PMTK001,101,0*33<CR><LF>
type MTK struct {
	BaseSentence
	Cmd, // Three bytes character string. From "000" to "999". An identifier used to tell the decoder how to decode the packet
	// Flag is flag field in PMTK001 packet.
	// Note: this field on only relevant for `PMTK001,Cmd,Flag` sentence.
	// Actual MTK protocol has variable amount of fields (whole sentence can be up to 255 bytes)
	//
	// Actual docs say:
	// DataField: The DataField has variable length depending on the packet type. A comma symbol ‘,’ must be inserted
	// ahead each data filed to help the decoder process the DataField.
	Flag int64
}

// newMTK constructor
func newMTK(s BaseSentence) (MTK, error) {
	p := NewParser(s)
	cmd := p.Int64(0, "command")
	flag := p.Int64(1, "flag")
	return MTK{
		BaseSentence: s,
		Cmd:          cmd,
		Flag:         flag,
	}, p.Err()
}
