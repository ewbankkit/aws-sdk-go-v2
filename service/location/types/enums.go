// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BatchItemErrorCode string

// Values returns all known values for BatchItemErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BatchItemErrorCode) Values() []BatchItemErrorCode {
	return []BatchItemErrorCode{
		"AccessDeniedError",
		"ConflictError",
		"InternalServerError",
		"ResourceNotFoundError",
		"ThrottlingError",
		"ValidationError",
	}
}

type DimensionUnit string

// Values returns all known values for DimensionUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DimensionUnit) Values() []DimensionUnit {
	return []DimensionUnit{
		"Meters",
		"Feet",
	}
}

type DistanceUnit string

// Values returns all known values for DistanceUnit. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DistanceUnit) Values() []DistanceUnit {
	return []DistanceUnit{
		"Kilometers",
		"Miles",
	}
}

type IntendedUse string

// Values returns all known values for IntendedUse. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IntendedUse) Values() []IntendedUse {
	return []IntendedUse{
		"SingleUse",
		"Storage",
	}
}

type PositionFiltering string

// Values returns all known values for PositionFiltering. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PositionFiltering) Values() []PositionFiltering {
	return []PositionFiltering{
		"TimeBased",
		"DistanceBased",
		"AccuracyBased",
	}
}

type PricingPlan string

// Values returns all known values for PricingPlan. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PricingPlan) Values() []PricingPlan {
	return []PricingPlan{
		"RequestBasedUsage",
		"MobileAssetTracking",
		"MobileAssetManagement",
	}
}

type RouteMatrixErrorCode string

// Values returns all known values for RouteMatrixErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RouteMatrixErrorCode) Values() []RouteMatrixErrorCode {
	return []RouteMatrixErrorCode{
		"RouteNotFound",
		"RouteTooLong",
		"PositionsNotFound",
		"DestinationPositionNotFound",
		"DeparturePositionNotFound",
		"OtherValidationError",
	}
}

type Status string

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"Active",
		"Expired",
	}
}

type TravelMode string

// Values returns all known values for TravelMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TravelMode) Values() []TravelMode {
	return []TravelMode{
		"Car",
		"Truck",
		"Walking",
		"Bicycle",
		"Motorcycle",
	}
}

type ValidationExceptionReason string

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UnknownOperation",
		"Missing",
		"CannotParse",
		"FieldValidationFailed",
		"Other",
	}
}

type VehicleWeightUnit string

// Values returns all known values for VehicleWeightUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VehicleWeightUnit) Values() []VehicleWeightUnit {
	return []VehicleWeightUnit{
		"Kilograms",
		"Pounds",
	}
}
