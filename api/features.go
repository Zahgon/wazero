package api

type CoreFeatures uint64

const CoreFeaturesV1 = CoreFeatureMutableGlobal

const CoreFeaturesV2 = CoreFeaturesV1 |
	CoreFeatureBulkMemoryOperations |
	CoreFeatureMultiValue |
	CoreFeatureNonTrappingFloatToIntConversion |
	CoreFeatureReferenceTypes |
	CoreFeatureSignExtensionOps |
	CoreFeatureSIMD

const (
	CoreFeatureBulkMemoryOperations CoreFeatures = 1 << iota

	CoreFeatureMultiValue

	CoreFeatureMutableGlobal

	CoreFeatureNonTrappingFloatToIntConversion

	CoreFeatureReferenceTypes

	CoreFeatureSignExtensionOps

	CoreFeatureSIMD
)

func (f CoreFeatures) SetEnabled(feature CoreFeatures, val bool) CoreFeatures {
	_ = "STUB: not implemented"
	return *new(CoreFeatures)
}

func (f CoreFeatures) IsEnabled(feature CoreFeatures) bool { _ = "STUB: not implemented"; return false }

func (f CoreFeatures) RequireEnabled(feature CoreFeatures) error {
	_ = "STUB: not implemented"
	return nil
}

func (f CoreFeatures) String() string { _ = "STUB: not implemented"; return "" }

func featureName(f CoreFeatures) string { _ = "STUB: not implemented"; return "" }
