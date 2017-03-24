
// Package that provides chronobiology functions to analyse time series data
package chronobiology

import (
    "time"
    "math"
    "errors"
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

    // Get the number of seconds elapsed since 01/01/1970
    seconds1 := date1.Unix()
    seconds2 := date2.Unix()

    // Calculate the difference in seconds
    seconds := seconds2-seconds1

    // Return the seconds as int instead of int64
    return int(seconds)
}

// Compares two float values using a predetermined epsilon
func floatEquals(a, b float64) bool {
    var epsilon float64 = 0.000000001
    if ((a - b) < epsilon && (b - a) < epsilon) {
        return true
    }
    return false
}

// Function used to decrease the epoch
func decrease(dateTime []time.Time, data []float64, currentEpoch int, newEpoch int) (newDateTime []time.Time, newData[]float64) {

    startDateTime := dateTime[0]
    // The start time must be the same start time of the current recorded data
    startDateTime = startDateTime.Add(-(time.Duration(currentEpoch) * time.Second))

    for index1 := 0; index1 < len(dateTime); index1++ {
        // To each data "row", split it to X new "rows"
        for index2 := 0; index2 < currentEpoch/newEpoch; index2++ {
            startDateTime = startDateTime.Add(time.Duration(newEpoch) * time.Second)
            newDateTime = append(newDateTime, startDateTime)
            newData     = append(newData, data[index1])
        }
    }

    return
}

// Function used to increase the epoch
func increase(dateTime []time.Time, data []float64, currentEpoch int, newEpoch int) (newDateTime []time.Time, newData[]float64) {

    var tempEpoch int
    var tempData float64

    startDateTime := dateTime[0]
    // The start time must be the same start time of the current recorded data
    startDateTime = startDateTime.Add(-(time.Duration(currentEpoch) * time.Second))

    for index1 := 0; index1 < len(dateTime); index1++ {
        tempEpoch += currentEpoch
        tempData  += data[index1]

        if tempEpoch >= newEpoch {
            startDateTime = startDateTime.Add(time.Duration(newEpoch) * time.Second)
            newDateTime = append(newDateTime, startDateTime)

            tempData    = tempData / (float64(newEpoch)/float64(currentEpoch))
            tempData    = roundPlus(tempData, 4)
            newData     = append(newData, tempData)

            tempEpoch = 0
            tempData  = 0.0
        }
    }

    return
}

// Function used in the IS analysis to normalize the data to a specific epoch passed as parameter
func normalizeDataIS(dateTime []time.Time, data []float64, minutes int)(temporaryDateTime []time.Time, temporaryData []float64, err error) {

    // Check the parameters
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }
    if minutes <= 0 {
        err = errors.New("MinutesInvalid")
        return
    }

    // If the minute is equal to 1, just return the original slices
    if minutes == 1 {
        temporaryDateTime = dateTime
        temporaryData     = data
        return
    }

    // Get the last valid position according to the minutes passed by parameter
    lastValidIndex := -1
    for index := len(dateTime); index > 0; index-- {
        if index % minutes == 0 {
            lastValidIndex = index
            break
        }
    }

    // Store the first DateTime
    currentDateTime := dateTime[0]

    // "Normalize" the data based on the minutes passed as parameter
    for index := 0; index < lastValidIndex; index += minutes {

        tempData := 0.0
        count := 0

        for tempIndex := index; tempIndex < index+minutes; tempIndex++ {
            tempData += data[tempIndex]
            count++
        }

        currentDateTime   = currentDateTime.Add(time.Duration(minutes) * time.Minute)
        temporaryDateTime = append(temporaryDateTime, currentDateTime)
        temporaryData     = append(temporaryData, (tempData/float64(count)))
    }

    return
}

/* END INTERNAL FUNCTIONS */

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

// Function that finds the highest activity average of the followed X hours (defined by parameter)
func HigherActivity(hours int, dateTime []time.Time, data []float64) (higherActivity float64, onsetHigherActivity time.Time, err error) {

    // Check the parameters
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

      // Check the parameters
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
            return nil, err
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

    // Check the parameters
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }
    if newEpoch == 0 {
        err = errors.New("InvalidEpoch")
        return
    }

    currentEpoch := FindEpoch(dateTime)

    // Could not find the epoch
    if currentEpoch == 0 {
        err = errors.New("InvalidEpoch")
        return
    }
    if newEpoch == currentEpoch {
        return dateTime, data, nil
    }

    // If the new Epoch is not divisible or multipliable by the currentEpoch
    // It needs to be decreased to 1 second to then increase to the newEpoch
    if (newEpoch > currentEpoch && newEpoch % currentEpoch != 0) ||
       (currentEpoch > newEpoch && currentEpoch % newEpoch != 0)  {

        // Decrease to 1 second
        dateTime, data = decrease(dateTime, data, currentEpoch, 1);

        // Increase to the newEpoch
        newDateTime, newData = increase(dateTime, data, 1, newEpoch);

    } else {
        // Increase
        if newEpoch > currentEpoch {
            newDateTime, newData = increase(dateTime, data, currentEpoch, newEpoch);

        // Decrease
        } else {
            newDateTime, newData = decrease(dateTime, data, currentEpoch, newEpoch);
        }
    }

    return
}

// Function created to filter the data based on the startTime and endTime passed as parameter
func FilterDataByDateTime(dateTime []time.Time, data []float64, startTime time.Time, endTime time.Time) (newDateTime []time.Time, newData []float64, err error) {

    // Check the parameters
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }
    if endTime.Before(startTime) {
        err = errors.New("InvalidTimeRange")
        return
    }

    // Filter the data based on the startTime and endTime
    for index := 0; index < len(dateTime); index++ {

        if (dateTime[index].After(startTime) || dateTime[index].Equal(startTime)) &&
           (dateTime[index].Before(endTime)  || dateTime[index].Equal(endTime)) {

            newDateTime = append(newDateTime, dateTime[index])
            newData     = append(newData, data[index])
        }
    }

    return
}

// Function that calculates the interdaily stability
func InterdailyStability(dateTime []time.Time, data []float64) (is []float64, err error) {

    // Check the parameters
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }
    if secondsTo(dateTime[0], dateTime[len(dateTime)-1]) < (48*60*60) {
        err = errors.New("LessThan2Days")
        return
    }

    currentEpoch := FindEpoch(dateTime)

    // Could not find the epoch
    if currentEpoch == 0 {
        err = errors.New("InvalidEpoch")
        return
    }

    if currentEpoch != 60 {
        newDateTime, newData, _ := ConvertDataBasedOnEpoch(dateTime, data, 60)

        dateTime = newDateTime
        data     = newData
    }

    // The data should be divisible by 1440 (entire day)
    for len(dateTime) % 1440 != 0 {
        // Remove the last data
        dateTime = dateTime[:len(dateTime)-1]
        data     = data[:len(data)-1]
    }

    // The zero position is allocated to store the average value of the IS vector
    is = append(is, 0.0)

    // Calculate all 60 IS values
    for isIndex := 1; isIndex <= 60; isIndex++ {

        if 1440 % isIndex == 0 {

            // Normalizes data to the new epoch (minutes)
            temporaryDateTime, temporaryData, _ := normalizeDataIS(dateTime, data, isIndex)

            // Calculate the average day
            _, averageDayData, _ := AverageDay(temporaryDateTime, temporaryData)

            // Get the new N (length)
            n := len(temporaryData)

            // Calculate the number of points per day
            p := len(averageDayData)
            //p := 1440 / isIndex

            // Calculate the new average (Xm)
            average := average(temporaryData)

            numerator   := 0.0
            denominator := 0.0

            // The "h" value represents the same "h" from the IS calculation formula
            for h := 0; h < p; h++ {
                numerator += math.Pow((averageDayData[h]-average), 2)
            }

            // The "i" value represents the same "i" from the IS calculation formula
            for i := 0; i < n; i++ {
                denominator += math.Pow((temporaryData[i]-average), 2)
            }

            numerator   = float64(n) * numerator
            denominator = float64(p) * denominator

            // Prevent NaN
            if denominator == 0 {
                is = append(is, -1.0)
            } else {
                is = append(is, (numerator/denominator))
            }
        } else {
            // Append -1 in the positions that will not be used
            is = append(is, -1.0)
        }
    }

    // Calculates the IS average of all "valid" values
    average := 0.0
    count   := 0

    for index := 1; index < len(is); index++ {
        if is[index] > -1.0 {
            average += is[index]
            count++
        }
    }

    is[0] = average/float64(count)

    return
}

// Function that searches for gaps in the time series and fills it with a specific value passed as parameter (usually zero)
func FillGapsInData(dateTime []time.Time, data []float64, value float64) (newDateTime []time.Time, newData []float64, err error) {

    // Check the parameters
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }

    currentEpoch := FindEpoch(dateTime)

    // Could not find the epoch
    if currentEpoch == 0 {
        err = errors.New("InvalidEpoch")
        return
    }

    for index := 0; index < len(dateTime)-1; index++ {

        newDateTime   = append(newDateTime, dateTime[index])
        newData       = append(newData, data[index])

        // If this condition is true, then this is a gap
        if secondsTo(dateTime[index], dateTime[index+1]) >= (currentEpoch*2) {

            tempDateTime := dateTime[index]
            count := (secondsTo(dateTime[index], dateTime[index+1]) / currentEpoch) - 1

            for tempIndex := 0; tempIndex < count; tempIndex++ {
                tempDateTime = tempDateTime.Add(time.Duration(currentEpoch) * time.Second)
                newDateTime  = append(newDateTime, tempDateTime)
                newData      = append(newData, value)
            }
        }
    }

    newDateTime = append(newDateTime, dateTime[len(dateTime)-1])
    newData     = append(newData, data[len(dateTime)-1])

    return
}

// Creates an average day based on the time series.
func AverageDay(dateTime []time.Time, data []float64) (newDateTime []time.Time, newData []float64, err error) {

    // Check the parameters
    if len(dateTime) == 0 || len(data) == 0 {
        err = errors.New("Empty")
        return
    }
    if len(dateTime) != len(data) {
        err = errors.New("DifferentSize")
        return
    }

    currentEpoch := FindEpoch(dateTime)

    // Could not find the epoch
    if currentEpoch == 0 {
        err = errors.New("InvalidEpoch")
        return
    }

    if secondsTo(dateTime[0], dateTime[len(dateTime)-1]) < (24*60*60) {
        err = errors.New("LessThan1Day")
        return
    }

    gapValue := -999.999
    dateTime, data, _ = FillGapsInData(dateTime, data, gapValue)

    pointsPerDay := (60*1440) / currentEpoch

    var countPoints []int

    for index := 0; index < pointsPerDay; index++ {
        newData     = append(newData, 0.0)
        countPoints = append(countPoints, 0)
    }

    pointIndex  := 0
    for index := 0; index < len(data); index++ {
        if pointIndex >= pointsPerDay {
            pointIndex = 0
        }

        if !floatEquals(data[index], gapValue) {
            newData[pointIndex] += data[index]
            countPoints[pointIndex] += 1
        }

        pointIndex++
    }

    tempDateTime := dateTime[0]
    for index := 0; index < len(newData); index++ {
        newDateTime  = append(newDateTime, tempDateTime)
        tempDateTime = tempDateTime.Add(time.Duration(currentEpoch) * time.Second)
        newData[index] = roundPlus((newData[index] / float64(countPoints[index])), 4)
    }

    return
}
