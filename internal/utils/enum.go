package utils

type ScanType string

const (
	Full  ScanType = "full"
	Quick ScanType = "quick"
)

func (ScanType) Values() (kinds []string) {
	for _, s := range []ScanType{Full, Quick} {
		kinds = append(kinds, string(s))
	}
	return
}

type AsynqQueue string

const (
	ScannerQueue            AsynqQueue = "scanner"
	ThumbnailGeneratorQueue AsynqQueue = "thumbnail_generator"
)

func (AsynqQueue) Values() (kinds []string) {
	for _, s := range []AsynqQueue{ScannerQueue} {
		kinds = append(kinds, string(s))
	}
	return
}
