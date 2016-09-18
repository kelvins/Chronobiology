
// Package that provides chronobiology functions to analyse time series data
package chronobiology

import (
    "time"
    "errors"
    "math"
)

/* BEGIN INTERNAL FUNCTIONS */

// Used to truncate a float64 value
func round(value float64) float64 {
    return math.Floor(value + .5)
}

// Used to truncate a float64 value to a particular precision
func roundPlus(value float64, places int) (float64) {
    shift := math.Pow(10, float64(places))
    return round(value * shift) / shift;
}

// Calculates the average of a float64 slice
func average(data []float64) (float64) {
    var average float64
    if len(data) == 0 {
      return average
    }
    for index := 0; index < len(data); index++ {
        average += data[index]
    }
    return average / float64(len(data))
}

// Searches for a value in a slice and returns its position
func findPosition(value int, data []int) (int) {
    if len(data) == 0 {
        return -1
    }
    for index := 0; index < len(data); index++ {
        if data[index] == value {
            return index
        }
    }
    return -1
}

// Finds the max value in a slice and returns its position
func findMaxPosition(data []int) (int) {
    if len(data) == 0 {
        return -1
    }
    var position int
    max := data[0]

    for index := 0; index < len(data); index++ {
        if data[index] > max {
            max = data[index]
            position = index
        }
    }
    return position
}

// Calculates the difference between two time.Time in seconds
func secondsTo(date1 time.Time, date2 time.Time) (int) {
    if date1.Equal(date2) || date1.After(date2) {
        return 0
    }

    year := date2.Year() - date1.Year()
    month := date2.Month() - date1.Month()

    // Does not calculate the seconds if is higher than 1 month
    if year > 0 || month > 0 {
        return 0
    }

    day := date2.Day() - date1.Day()
    hour := date2.Hour() - date1.Hour()
    minute := date2.Minute() - date1.Minute()
    second := date2.Second() - date1.Second()

    seconds := day*24*60*60
    seconds += hour*60*60
    seconds += minute*60
    seconds += second

    return seconds
}

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
  	if ((a - b) < EPSILON && (b - a) < EPSILON) {
  		  return true
  	}
  	return false
}

/* END INTERNAL FUNCTIONS */

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

        if currentActivity > higherActivity || floatEquals(higherActivity, 0.0) {
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

// Function that calculates the intradaily variability
func IntradailyVariability(dateTime []time.Time, data []float64) (iv []float64, err error) {

    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }
    if secondsTo(dateTime[0], dateTime[len(dateTime)-1]) < (2*60*60) {
        err = errors.New("LessThan2Hours")
    }

    // The zero position is allocated to store the average value of the iv vector
    iv = append(iv, 0.0)

    for mainIndex := 1; mainIndex <= 60; mainIndex++ {

        _, tempData, err := ConvertDataBasedOnEpoch(dateTime, data, (mainIndex*60))

        if err != nil {
            err = errors.New("ConvertDataBasedOnEpoch error")
            iv = nil
            return iv, err
        }

        if len(tempData) > 0 {

            average := average(tempData)

            // Calculates the numerator
            var numerator float64
            for index := 1; index < len(tempData); index++ {
                tempValue := tempData[index] - tempData[index-1]
                numerator += math.Pow(tempValue, 2)
            }
            numerator = numerator * float64(len(tempData))

            // Calculates the denominator
            var denominator float64
            for index := 0; index < len(tempData); index++ {
                tempValue := average - tempData[index]
                denominator += math.Pow(tempValue, 2)
            }
            denominator = denominator * (float64(len(tempData)) - 1.0)

            result := roundPlus((numerator/denominator), 4)
            iv = append(iv, result)

        } else {
            iv = append(iv, 0.0)
        }
    }

    // Calculates the IV average
    var average float64
    for index := 1; index < len(iv); index++ {
        average += iv[index]
    }
    average = average / float64(len(iv)-1)
    iv[0] = average

    return
}

// Function that finds the epoch of a time series (seconds)
func FindEpoch(dateTime []time.Time) (epoch int) {

    if len(dateTime) == 0 {
        return
    }

    var count []int
    var epochs []int

    for index := 1; index < len(dateTime); index++ {

        seconds := secondsTo(dateTime[index-1], dateTime[index])

        position := findPosition(seconds, epochs)
        if position > -1 {
            count[position] += 1
        }else {
            epochs = append(epochs, seconds)
            count  = append(count, 1)
        }
    }

    maxPos := findMaxPosition(count)
    epoch = epochs[maxPos]

    return
}

// Convert the data and dateTime slices to the new epoch passed by parameter
func ConvertDataBasedOnEpoch(dateTime []time.Time, data []float64, newEpoch int) (newDateTime []time.Time, newData []float64, err error) {

    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }

    currentEpoch := FindEpoch(dateTime)
    startDateTime := dateTime[0]

    if newEpoch == currentEpoch {
        return dateTime, data, nil
    }

    // Increase
    if newEpoch > currentEpoch {
        var tempEpoch int
        var tempData float64
        for index1 := 0; index1 < len(dateTime); index1++ {
            tempEpoch += currentEpoch
            tempData  += data[index1]
            if tempEpoch >= newEpoch {
                newDateTime = append(newDateTime, startDateTime)
                tempData    = tempData / (float64(newEpoch)/float64(currentEpoch))
                tempData    = roundPlus(tempData, 4)
                newData     = append(newData, tempData)
                tempEpoch = 0
                tempData  = 0.0
                startDateTime = startDateTime.Add(time.Duration(newEpoch) * time.Second)
            }
        }
    } else {
        // Decrease
        for index1 := 0; index1 < len(dateTime); index1++ {
            for index2 := 0; index2 < currentEpoch/newEpoch; index2++ {
                newDateTime = append(newDateTime, startDateTime)
                newData     = append(newData, data[index1])
                startDateTime = startDateTime.Add(time.Duration(newEpoch) * time.Second)
            }
        }
    }

    return
}
