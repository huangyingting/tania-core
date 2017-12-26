package reservoir

const (
	ReservoirErrorNameEmptyCode = iota
	ReservoirErrorNameNotEnoughCharacterCode
	ReservoirErrorNameExceedMaximunCharacterCode
	ReservoirErrorNameAlphanumericOnlyCode

	ReservoirErrorPHInvalidCode

	ReservoirErrorECInvalidCode

	ReservoirErrorBucketCapacityInvalidCode
	ReservoirErrorBucketVolumeInvalidCode

	ReservoirErrorWaterSourceAlreadyAttachedCode
)

// ReservoirError is a custom error from Go built-in error
type ReservoirError struct {
	Code int
}

func (e ReservoirError) Error() string {
	switch e.Code {
	case ReservoirErrorNameEmptyCode:
		return "Reservoir name is required."
	case ReservoirErrorNameNotEnoughCharacterCode:
		return "Not enough character on Reservoir Name"
	case ReservoirErrorNameExceedMaximunCharacterCode:
		return "Reservoir name cannot more than 100 characters"
	case ReservoirErrorNameAlphanumericOnlyCode:
		return "Reservoir name should be alphanumeric"
	case ReservoirErrorPHInvalidCode:
		return "Reservoir pH value is invalid."
	case ReservoirErrorECInvalidCode:
		return "Reservoir EC value is invalid."
	case ReservoirErrorBucketCapacityInvalidCode:
		return "Reservoir bucket capacity is invalid."
	case ReservoirErrorWaterSourceAlreadyAttachedCode:
		return "Reservoir water source is already attached."
	case ReservoirErrorBucketVolumeInvalidCode:
		return "Reservoir bucket volume is invalid."
	default:
		return "Unrecognized Reservoir Error Code"
	}
}
