package internal

import "github.com/RecursionExcursion/go-toolkit/core"

const batchSize = 100

var BatchRunner = core.RunBatchSizeClosure(batchSize)

const BetBotDataId = "data"
