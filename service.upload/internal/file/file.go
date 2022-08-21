package file

// Properties of a File
type File struct {
	// A unique assigned identifier
	ID string
	// File name/title
	Title string
	// Url to access file
	URL string
	// Number of pages (NB: reference for PDF's, defaults to 1)
	NumberOfPages int
}
