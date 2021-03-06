//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type MpaaRating string

const (
	MpaaRating_G    MpaaRating = "G"
	MpaaRating_Pg   MpaaRating = "PG"
	MpaaRating_Pg13 MpaaRating = "PG-13"
	MpaaRating_R    MpaaRating = "R"
	MpaaRating_Nc17 MpaaRating = "NC-17"
)

func (e *MpaaRating) Scan(value interface{}) error {
	if v, ok := value.(string); !ok {
		return errors.New("jet: Invalid data for MpaaRating enum")
	} else {
		switch string(v) {
		case "G":
			*e = MpaaRating_G
		case "PG":
			*e = MpaaRating_Pg
		case "PG-13":
			*e = MpaaRating_Pg13
		case "R":
			*e = MpaaRating_R
		case "NC-17":
			*e = MpaaRating_Nc17
		default:
			return errors.New("jet: Inavlid data " + string(v) + "for MpaaRating enum")
		}

		return nil
	}
}

func (e MpaaRating) String() string {
	return string(e)
}
