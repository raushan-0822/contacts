package handler

// import (
// 	"appapi/types"
// 	"context"
// )

// //Throttler defines interface for throttling
// type Throttler interface {
// 	Throttle(context.Context) bool
// }

// //basicThrottler type for mocking Throttler
// type basicThrottler struct {
// 	Unit  string
// 	Limit int
// 	Type  string
// }

// //BasicThrottler creates basic throtlle
// func BasicThrottler(limit int, unit string) Throttler {
// 	return &basicThrottler{Unit: unit, Limit: limit, Type: "basicthrottle"}
// }

// //Throttle implementing throttle interface
// func (bt basicThrottler) Throttle(ctx types.Context) bool {
// 	path, _ := ctx.Get("path").(string)
// 	accountSid, _ := ctx.Get("accountsid").(string)
// 	action, _ := ctx.Get("action").(string)
// 	// Default throttle limits from resource contract
// 	threshold := bt.Limit
// 	interval := bt.Unit
// 	throttleType := bt.Type
// 	// Get updated throttle settings if any for a given tenant from settings table
// 	//mysql.UpdateThrottleFromSettings(ctx, accountSid, path, action, throttleType, &threshold, &interval)
// 	// Decide whether or not to throttle this request
// 	//return mysql.ShouldThrottle(accountSid, path, action, throttleType, threshold, interval)
// 	return false
// }
