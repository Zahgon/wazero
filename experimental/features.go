package experimental

import "github.com/tetratelabs/wazero/api"

const CoreFeaturesThreads = api.CoreFeatureSIMD << 1

const CoreFeaturesTailCall = api.CoreFeatureSIMD << 2

const CoreFeaturesExtendedConst = api.CoreFeatureSIMD << 3

const CoreFeaturesExceptionHandling = api.CoreFeatureSIMD << 4

const CoreFeaturesTypedFunctionReferences = api.CoreFeatureSIMD << 5
