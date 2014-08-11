// Copyright 2013 Sevki Hasirci .  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fathertime converts time.Time in to human readable string
package fathertime

import (
	"fmt"
	"time"
)

const lssthnd = "less than %d %s"
const lssthns = "less than a %s"
const aboutnd = "about %d %s"
const day time.Duration = 86400000000000
const month time.Duration = 2628000000001209
const year time.Duration = 31535999999964780

/*
TimeInWords returns time in words.
Conversion follows the format.
0-4   secs                                                                # => less than 5 seconds
5-9   secs                                                                # => less than 10 seconds
10-19 secs                                                                # => less than 20 seconds
20-39 secs                                                                # => half a minute
40-59 secs                                                                # => less than a minute
60-89 secs                                                                # => 1 minute
1 min, 30 secs <-> 44 mins, 29 secs                                       # => [2..44] minutes
44 mins, 30 secs <-> 89 mins, 29 secs                                     # => about 1 hour
89 mins, 30 secs <-> 23 hrs, 59 mins, 29 secs                             # => about [2..24] hours
23 hrs, 59 mins, 30 secs <-> 41 hrs, 59 mins, 29 secs                     # => 1 day
41 hrs, 59 mins, 30 secs  <-> 29 days, 23 hrs, 59 mins, 29 secs           # => [2..29] days
29 days, 23 hrs, 59 mins, 30 secs <-> 59 days, 23 hrs, 59 mins, 29 secs   # => about 1 month
59 days, 23 hrs, 59 mins, 30 secs <-> 1 yr minus 1 sec                    # => [2..12] months
1 yr <-> 1 yr, 3 months                                                   # => about 1 year
1 yr, 3 months <-> 1 yr, 9 months                                         # => over 1 year
1 yr, 9 months <-> 2 yr minus 1 sec                                       # => almost 2 years
2 yrs <-> max time or date                                                # => (same rules as 1 yr)
*/
func TimeInWords(t time.Time) string {

	now := time.Now()
	d := now.Sub(t)
	if d >= time.Second && d <= (time.Second*4) {
		return fmt.Sprintf(lssthnd, 5, "seconds")
	} else if d >= (time.Second*5) && d < (time.Second*10) {
		return fmt.Sprintf(lssthnd, 10, "seconds")
	} else if d >= (time.Second*10) && d < (time.Second*20) {
		return fmt.Sprintf(lssthnd, 20, "seconds")
	} else if d >= (time.Second*20) && d < (time.Second*40) {
		return "half a minute"
	} else if d >= (time.Second*40) && d < (time.Second*60) {
		return fmt.Sprintf(lssthns, "minute")
	} else if d >= (time.Second*60) && d < time.Minute+(time.Second*30) {
		return "1 minute"
	} else if d >= time.Minute+(time.Second*30) && d < (time.Minute*44)+(time.Second*30) {
		return fmt.Sprintf("%d minutes", (d / time.Minute))
	} else if d >= (time.Minute*44)+(time.Second*30) && d < (time.Minute*89)+(time.Second*30) {
		return fmt.Sprintf(aboutnd, d/time.Hour, "hour")
	} else if d >= (time.Minute*89)+(time.Second*30) && d < (time.Hour*29)+(time.Minute*59)+(time.Second*30) {
		return fmt.Sprintf(aboutnd, (d / time.Hour), "hours")
	} else if d >= (time.Hour*23)+(time.Minute*59)+(time.Second*30) && d < (time.Hour*41)+(time.Minute*59)+(time.Second*30) {
		return "1 day"
	} else if d >= (time.Hour*41)+(time.Minute*59)+(time.Second*30) && d < (day*29)+(time.Hour*23)+(time.Minute*59)+(time.Second*30) {
		return fmt.Sprintf("%d days", d/(time.Hour*24))
	} else if d >= (day*29)+(time.Hour*23)+(time.Minute*59)+(time.Second*30) && d < (day*59)+(time.Hour*23)+(time.Minute*59)+(time.Second*30) {
		return fmt.Sprintf(aboutnd, 1, "month")
	} else if d >= (day*59)+(time.Hour*23)+(time.Minute*59)+(time.Second*30) && d < (year) {
		return fmt.Sprintf(aboutnd, d/month+1, "months")
	} else if d >= year && d < year+(3*month) {
		return fmt.Sprintf(aboutnd, 1, "year")
	} else if d >= year+(3*month) && d < year+(9*month) {
		return "over 1 year"
	} else if d >= year+(9*month) && d < (year*2) {
		return "almost 2 years"
	} else {
		return fmt.Sprintf(aboutnd, d/year, "years")
	}

}
