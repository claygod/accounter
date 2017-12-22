package accounter

// Indicator
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

//"errors"
//"log"
//"runtime"
//"sync/atomic"

const trialLimitConst int = 2147483647
const trialStop int = 64
const closed int64 = -2147483647

var trialLimit int = trialLimitConst
var batchSize int = 4
