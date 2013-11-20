fathertime
==========

Prints out time in words.

follows this scheme
```text
0-4   secs                                                                => less than 5 seconds
5-9   secs                                                                => less than 10 seconds
10-19 secs                                                                => less than 20 seconds
20-39 secs                                                                => half a minute
40-59 secs                                                                => less than a minute
60-89 secs                                                                => 1 minute
1 min, 30 secs <-> 44 mins, 29 secs                                       => [2..44] minutes
44 mins, 30 secs <-> 89 mins, 29 secs                                     => about 1 hour
89 mins, 30 secs <-> 23 hrs, 59 mins, 29 secs                             => about [2..24] hours
23 hrs, 59 mins, 30 secs <-> 41 hrs, 59 mins, 29 secs                     => 1 day
41 hrs, 59 mins, 30 secs  <-> 29 days, 23 hrs, 59 mins, 29 secs           => [2..29] days
29 days, 23 hrs, 59 mins, 30 secs <-> 59 days, 23 hrs, 59 mins, 29 secs   => about 1 month
59 days, 23 hrs, 59 mins, 30 secs <-> 1 yr minus 1 sec                    => [2..12] months
1 yr <-> 1 yr, 3 months                                                   => about 1 year
1 yr, 3 months <-> 1 yr, 9 months                                         => over 1 year
1 yr, 9 months <-> 2 yr minus 1 sec                                       => almost 2 years
2 yrs <-> max time or date                                                => (same rules as 1 yr)
```
which I lovingly borrowed from a similar rails function called [distance_of_time_in_words](http://apidock.com/rails/ActionView/Helpers/DateHelper/distance_of_time_in_words) 

---

Copyright (c) 2013, Sevki Hasirci
