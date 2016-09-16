
// Package that provides chronobiology functions to analyse time series data
package chronobiology

import (
    "time"
    "errors"
    "math"
)

// Used to truncate a float64 value
func round(value float64) float64 {
    return math.Floor(value + .5)
}

// Used to truncate a float64 value to a particular precision
func roundPlus(value float64, places int) (float64) {
    shift := math.Pow(10, float64(places))
    return round(value * shift) / shift;
}

// Function that finds the highest activity average of the followed X hours (defined by parameter)
func HigherActivity(hours int, dateTime []time.Time, data []float64) (higherActivity float64, onsetHigherActivity time.Time, err error) {

    if hours == 0 {
        err = errors.New("InvalidHours")
        return
    }
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }
    if dateTime[0].Add(time.Duration(hours) * time.Hour).After( dateTime[len(dateTime)-1] ) {
        err = errors.New("HoursHigher")
        return
    }

    for index := 0; index < len(dateTime); index++ {

        startDateTime := dateTime[index]
        finalDateTime := startDateTime.Add(time.Duration(hours) * time.Hour)
        tempDateTime  := startDateTime

        if finalDateTime.After( dateTime[len(dateTime)-1] ) {
            break
        }

        currentActivity := 0.0
        tempIndex := index
        count := 0

        for tempDateTime.Before(finalDateTime) {
            currentActivity += data[tempIndex]
            count += 1
            tempIndex += 1

            if tempIndex >= len(dateTime) {
                break
            }

            tempDateTime = dateTime[tempIndex]
        }

        currentActivity /= float64(count)

        if currentActivity > higherActivity || higherActivity == 0.0 {
            higherActivity = roundPlus(currentActivity, 4)
            onsetHigherActivity = startDateTime
        }
    }

    return
}

// Function that finds the lowest activity average of the followed X hours (defined by parameter)
func LowerActivity(hours int, dateTime []time.Time, data []float64) (lowerActivity float64, onsetLowerActivity time.Time, err error) {

      if hours == 0 {
          err = errors.New("InvalidHours")
          return
      }
      if len(dateTime) == 0 || len(data) == 0 {
          err = errors.New("Empty")
          return
      }
      if len(dateTime) != len(data) {
          err = errors.New("DifferentSize")
          return
      }
      if dateTime[0].Add(time.Duration(hours) * time.Hour).After( dateTime[len(dateTime)-1] ) {
          err = errors.New("HoursHigher")
          return
      }

      firstTime := true

      for index := 0; index < len(dateTime); index++ {

          startDateTime := dateTime[index]
          finalDateTime := startDateTime.Add(time.Duration(hours) * time.Hour)
          tempDateTime  := startDateTime

          if finalDateTime.After( dateTime[len(dateTime)-1] ) {
              break
          }

          currentActivity := 0.0
          tempIndex := index
          count := 0

          for tempDateTime.Before(finalDateTime) {
              currentActivity += data[tempIndex]
              count += 1
              tempIndex += 1

              if tempIndex >= len(dateTime) {
                  break
              }

              tempDateTime = dateTime[tempIndex]
          }

          currentActivity /= float64(count)

          if currentActivity < lowerActivity || firstTime == true {
              lowerActivity = roundPlus(currentActivity, 4)
              onsetLowerActivity = startDateTime
              firstTime = false
          }
      }

      return
}

// Function that finds the highest activity average of the followed 10 hours
func M10(dateTime []time.Time, data []float64) (higherActivity float64, onsetHigherActivity time.Time, err error) {
    higherActivity, onsetHigherActivity, err = HigherActivity(10, dateTime, data)
    return
}

// Function that finds the lowest activity average of the following 5 hours
func L5(dateTime []time.Time, data []float64) (lowerActivity float64, onsetLowerActivity time.Time, err error) {
    lowerActivity, onsetLowerActivity, err = LowerActivity(5, dateTime, data)
    return
}

// Function that calculates the relative amplitude based on the formula (M10-L5)/(M10+L5)
func RelativeAmplitude(highestAverage float64, lowestAverage float64) (RA float64, err error) {
    if( highestAverage == 0.0 && lowestAverage == 0.0 ) {
        err = errors.New("NullValues")
        return
    }

    RA = (highestAverage-lowestAverage) / (highestAverage+lowestAverage)
    RA = roundPlus(RA, 4)
    return
}
