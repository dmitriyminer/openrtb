package response

import "github.com/bsm/openrtb"

type ImageTypeID int

const (
	ImageTypeIcon ImageTypeID = 1 // Icon image
	ImageTypeLogo ImageTypeID = 2 // Logo image for the brand/app
	ImageTypeMain ImageTypeID = 3 // Large image preview for the ad
)

// Image object contains response image.
type Image struct {
	TypeID ImageTypeID       `json:"type,omitempty"` // Type ID of the image element supported by the publisher
	URL    string            `json:"url,omitempty"`  // URL of the image asset
	Width  int               `json:"w,omitempty"`    // Width of the image in pixels
	Height int               `json:"h,omitempty"`    // Height of the image in pixels
	Ext    openrtb.Extension `json:"ext,omitempty"`
}
