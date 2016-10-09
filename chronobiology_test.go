
package chronobiology_test

import (
    "time"
    "testing"
    "reflect"
    "github.com/kelvins/chronobiology"
)

/* Internal Functions */

func floatEquals(a, b float64) bool {
    var epsilon float64 = 0.00000001
  	if ((a - b) < epsilon && (b - a) < epsilon) {
  		  return true
  	}
  	return false
}

func sliceTimeEquals(slice1 []time.Time, slice2 []time.Time) (bool) {
    if len(slice1) != len(slice2) {
        return false
    }

    for index := 0; index < len(slice1); index++ {
        if !slice1[index].Equal(slice2[index]) {
            return false
        }
    }

    return true
}

func sliceFloatEquals(slice1 []float64, slice2 []float64) (bool) {
    if len(slice1) != len(slice2) {
        return false
    }

    for index := 0; index < len(slice1); index++ {
        if !floatEquals(slice1[index], slice2[index]) {
            return false
        }
    }

    return true
}

/* ################# */

func TestInvalidParametersHigherActivity(t *testing.T) {
    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    // Call the function with empty slices
    _, _, err := chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 8; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Call the function with myData empty
    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    myData = append(myData, 450.0) // 01
    myData = append(myData, 050.0) // 02
    myData = append(myData, 025.0) // 03
    myData = append(myData, 020.0) // 04
    myData = append(myData, 100.0) // 05
    myData = append(myData, 500.0) // 06
    myData = append(myData, 250.0) // 07

    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: DifferentSize")
    }

    myData = append(myData, 050.0) // 08

    _, _, err = chronobiology.HigherActivity(0, myDateTime, myData)

    if err == nil {
        t.Error("Expect: InvalidHours")
    }

    _, _, err = chronobiology.HigherActivity(20, myDateTime, myData)

    if err == nil {
        t.Error("Expect: HoursHigher")
    }
}

func TestHigherActivity(t *testing.T) {
    // Table tests
    var tTests = []struct {
        hours int
        higherActivity float64
        onsetHigherActivity string
    }{
        { 01, 990.0000, "01/01/2015 08:00:00" },
        { 02, 620.0000, "01/01/2015 07:00:00" },
        { 05, 482.0000, "01/01/2015 06:00:00" },
        { 06, 418.3333, "01/01/2015 05:00:00" },
        { 07, 364.2857, "01/01/2015 05:00:00" },
        { 10, 305.5000, "01/01/2015 01:00:00" },
    }

    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 12; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Creates the data slice
    myData = append(myData, 450.0) // 01
    myData = append(myData, 050.0) // 02
    myData = append(myData, 025.0) // 03
    myData = append(myData, 020.0) // 04
    myData = append(myData, 100.0) // 05
    myData = append(myData, 500.0) // 06
    myData = append(myData, 250.0) // 07
    myData = append(myData, 990.0) // 08
    myData = append(myData, 130.0) // 09
    myData = append(myData, 540.0) // 10
    myData = append(myData, 040.0) // 11
    myData = append(myData, 050.0) // 12

    // Test with all values in the table
    for _, pair := range tTests {
        higherActivity, onsetHigherActivity, err := chronobiology.HigherActivity(pair.hours, myDateTime, myData)
        if err != nil {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expect: error not nil",
            )
        }
        if !floatEquals(higherActivity, pair.higherActivity) {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.higherActivity,
                "received: ", higherActivity,
            )
        }
        if onsetHigherActivity.Format("02/01/2006 15:04:05") != pair.onsetHigherActivity {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.onsetHigherActivity,
                "received: ", onsetHigherActivity.Format("02/01/2006 15:04:05"),
            )
        }
    }
}

func TestInvalidParametersLowerActivity(t *testing.T) {
    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    // Call the function with empty slices
    _, _, err := chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 8; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Call the function with myData empty
    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: Empty")
    }

    myData = append(myData, 450.0) // 01
    myData = append(myData, 050.0) // 02
    myData = append(myData, 025.0) // 03
    myData = append(myData, 020.0) // 04
    myData = append(myData, 100.0) // 05
    myData = append(myData, 500.0) // 06
    myData = append(myData, 250.0) // 07

    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: DifferentSize")
    }

    myData = append(myData, 050.0) // 08

    _, _, err = chronobiology.HigherActivity(0, myDateTime, myData)

    if err == nil {
        t.Error("Expect: InvalidHours")
    }

    _, _, err = chronobiology.HigherActivity(20, myDateTime, myData)

    if err == nil {
        t.Error("Expect: HoursHigher")
    }
}

func TestLowerActivity(t *testing.T) {
    // Table tests
    var tTests = []struct {
        hours int
        lowerActivity float64
        onsetLowerActivity string
    }{
        { 01, 031.5, "01/01/2016 11:00:00" },
        { 02, 061.5, "01/01/2016 10:00:00" },
        { 04, 121.5, "01/01/2016 08:00:00" },
        { 06, 181.5, "01/01/2016 06:00:00" },
        { 07, 211.5, "01/01/2016 05:00:00" },
        { 10, 301.5, "01/01/2016 02:00:00" },
    }

    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the slices
    var myDateTime []time.Time
    var myData []float64

    tempDateTime := time.Date(2016,1,1,0,0,0,0,utc)

    // Fill the myDateTime (12 hours * 60 minutes) time.Minute
    for index := 0; index < (12*60); index++ {
        tempDateTime = tempDateTime.Add(1 * time.Minute)
        myDateTime = append(myDateTime, tempDateTime)
        myData = append(myData, float64((12*60)-index))
    }

    // Test with all values in the table
    for _, pair := range tTests {
        lowerActivity, onsetLowerActivity, err := chronobiology.LowerActivity(pair.hours, myDateTime, myData)
        if err != nil {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expect: error not nil",
            )
        }
        if !floatEquals(lowerActivity, pair.lowerActivity) {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.lowerActivity,
                "received: ", lowerActivity,
            )
        }
        if onsetLowerActivity.Format("02/01/2006 15:04:05") != pair.onsetLowerActivity {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expected: ", pair.onsetLowerActivity,
                "received: ", onsetLowerActivity.Format("02/01/2006 15:04:05"),
            )
        }
    }
}

func TestRelativeAmplitude(t *testing.T) {
    // Expect an error
    _, err := chronobiology.RelativeAmplitude(0.0, 0.0)
    if err == nil {
        t.Error(
            "Error: ", err,
        )
    }

    // Table tests
    var tTests = []struct {
        highestAverage float64
        lowestAverage float64
        relativeAmplitude float64
    }{
        { 180.0, 050.0, 0.5652 },
        { 550.0, 125.0, 0.6296 },
        { 101.0, 100.5, 0.0025 },
        { 898.0, 315.0, 0.4806 },
        { 211.5, 075.5, 0.4739 },
        { 620.0, 020.0, 0.9375 },
        { 780.0, 010.0, 0.9747 },
    }

    // Test with all values in the table
    for _, pair := range tTests {
        relativeAmplitude, err := chronobiology.RelativeAmplitude(pair.highestAverage, pair.lowestAverage)
        if err != nil {
            t.Error(
                "Error: ", err,
            )
        }
        if !floatEquals(relativeAmplitude, pair.relativeAmplitude) {
            t.Error(
                "Expected: ", pair.relativeAmplitude,
                "Received: ", relativeAmplitude,
            )
        }
    }
}

func TestFindEpoch(t *testing.T) {

    utc, _ := time.LoadLocation("UTC")
    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    var dateTimeEmpty []time.Time

    var dateTime60sec []time.Time
    for index := 0; index < 420; index++ {
        tempDateTime  = tempDateTime.Add(1 * time.Minute)
        dateTime60sec = append(dateTime60sec, tempDateTime)
    }

    var dateTime30sec []time.Time
    for index := 0; index < 120; index++ {
        tempDateTime  = tempDateTime.Add(30 * time.Second)
        dateTime30sec = append(dateTime30sec, tempDateTime)
    }

    var dateTime5sec []time.Time
    for index := 0; index < 10; index++ {
        tempDateTime = tempDateTime.Add(5 * time.Second)
        dateTime5sec = append(dateTime5sec, tempDateTime)
    }

    var dateTime180sec []time.Time
    for index := 0; index < 820; index++ {
        tempDateTime   = tempDateTime.Add(3 * time.Minute)
        dateTime180sec = append(dateTime180sec, tempDateTime)
    }

    var dateTime120sec []time.Time
    for index := 0; index < 100; index++ {
        tempDateTime   = tempDateTime.Add(2 * time.Minute)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }
    for index := 0; index < 30; index++ {
        tempDateTime   = tempDateTime.Add(1 * time.Minute)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }
    for index := 0; index < 15; index++ {
        tempDateTime   = tempDateTime.Add(30 * time.Second)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }
    for index := 0; index < 100; index++ {
        tempDateTime   = tempDateTime.Add(2 * time.Minute)
        dateTime120sec = append(dateTime120sec, tempDateTime)
    }

    var dateTime360sec []time.Time
    for index := 0; index < 50; index++ {
        tempDateTime   = tempDateTime.Add(6 * time.Minute)
        dateTime360sec = append(dateTime360sec, tempDateTime)
    }
    for index := 0; index < 49; index++ {
        tempDateTime   = tempDateTime.Add(4 * time.Minute)
        dateTime360sec = append(dateTime360sec, tempDateTime)
    }
    for index := 0; index < 49; index++ {
        tempDateTime   = tempDateTime.Add(2 * time.Minute)
        dateTime360sec = append(dateTime360sec, tempDateTime)
    }

    var dateTimeInvalid []time.Time
    for index := 0; index < 250; index++ {
        dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
    }

    // Table tests
    var tTests = []struct {
        dateTime []time.Time
        epoch int
    }{
        { dateTimeEmpty,    0 },
        { dateTime60sec,   60 },
        { dateTime30sec,   30 },
        { dateTime5sec,     5 },
        { dateTime180sec, 180 },
        { dateTime120sec, 120 },
        { dateTime360sec, 360 },
        { dateTimeInvalid,  0 },
    }

    // Test with all values in the table
    for _, pair := range tTests {
        epoch := chronobiology.FindEpoch(pair.dateTime)
        if epoch != pair.epoch {
            t.Error(
                "Expected: ", pair.epoch,
                "Received: ", epoch,
            )
        }
    }
}

func TestConvertDataBasedOnEpoch(t *testing.T) {

    utc, _ := time.LoadLocation("UTC")
    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    var dateTimeEmpty []time.Time
    var dataEmpty []float64

    _, _, err := chronobiology.ConvertDataBasedOnEpoch(dateTimeEmpty, dataEmpty, 120)

    if err == nil {
        t.Error("Expect error Empty")
    }

    var dateTimeInvalid []time.Time
    var dataInvalid []float64

    tempDateTime    = tempDateTime.Add(1 * time.Minute)
    dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
    tempDateTime    = tempDateTime.Add(1 * time.Minute)
    dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
    dataInvalid     = append(dataInvalid, 123.5)

    _, _, err = chronobiology.ConvertDataBasedOnEpoch(dateTimeInvalid, dataInvalid, 120)

    if err == nil {
        t.Error("Expect error DifferentSize")
    }

    var dateTime60secs []time.Time
    var data60secs []float64

    tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
    for index := 0; index < 40; index++ {
        dateTime60secs = append(dateTime60secs, tempDateTime)
        data60secs     = append(data60secs, 250.0)
        tempDateTime   = tempDateTime.Add(60 * time.Second)
    }

    var newDateTime30secs []time.Time
    var newData30secs []float64

    tempDateTime = dateTime60secs[0]
    tempDateTime = tempDateTime.Add(-(60 * time.Second))
    for index := 0; index < 80; index++ {
        tempDateTime      = tempDateTime.Add(30 * time.Second)
        newDateTime30secs = append(newDateTime30secs, tempDateTime)
        newData30secs     = append(newData30secs, 250.0)
    }

    var newDateTime120secs []time.Time
    var newData120secs []float64

    tempDateTime = dateTime60secs[0]
    tempDateTime = tempDateTime.Add(-(60 * time.Second))
    for index := 0; index < 20; index++ {
        tempDateTime       = tempDateTime.Add(120 * time.Second)
        newDateTime120secs = append(newDateTime120secs, tempDateTime)
        newData120secs     = append(newData120secs, 250.0)
    }

    var newDateTime90secs []time.Time
    var newData90secs []float64

    tempDateTime = dateTime60secs[0]
    tempDateTime = tempDateTime.Add(-(60 * time.Second))
    for index := 0; index < 26; index++ {
        tempDateTime      = tempDateTime.Add(90 * time.Second)
        newDateTime90secs = append(newDateTime90secs, tempDateTime)
        newData90secs     = append(newData90secs, 250.0)
    }

    var newDateTime15secs []time.Time
    var newData15secs []float64

    tempDateTime = dateTime60secs[0]
    tempDateTime = tempDateTime.Add(-(60 * time.Second))
    for index := 0; index < 160; index++ {
        tempDateTime      = tempDateTime.Add(15 * time.Second)
        newDateTime15secs = append(newDateTime15secs, tempDateTime)
        newData15secs     = append(newData15secs, 250.0)
    }

    var newDateTime240secs []time.Time
    var newData240secs []float64

    tempDateTime = dateTime60secs[0]
    tempDateTime = tempDateTime.Add(-(60 * time.Second))
    for index := 0; index < 10; index++ {
        tempDateTime       = tempDateTime.Add(240 * time.Second)
        newDateTime240secs = append(newDateTime240secs, tempDateTime)
        newData240secs     = append(newData240secs, 250.0)
    }

    // Table tests
    var tTests = []struct {
        dateTime []time.Time
        data []float64
        newEpoch int
        newDateTime []time.Time
        newData []float64
    }{
        { dateTime60secs, data60secs,  30,  newDateTime30secs,  newData30secs },
        { dateTime60secs, data60secs, 120, newDateTime120secs, newData120secs },
        { dateTime60secs, data60secs,  90,  newDateTime90secs,  newData90secs },
        { dateTime60secs, data60secs,  15,  newDateTime15secs,  newData15secs },
        { dateTime60secs, data60secs, 240, newDateTime240secs, newData240secs },
    }

    // Test with all values in the table
    for _, table := range tTests {
        newDateTime, newData, err := chronobiology.ConvertDataBasedOnEpoch(table.dateTime, table.data, table.newEpoch)

        if err != nil {
            t.Error("Expected error = nil.")
        }
        if !reflect.DeepEqual(newDateTime, table.newDateTime) {
            t.Error("Different dateTime slices. NewEpoch : ", table.newEpoch,)
        }
        if !reflect.DeepEqual(newData, table.newData) {
            t.Error("Different data slices. NewEpoch : ", table.newEpoch,)
        }
    }
}

func TestIntradailyVariability(t *testing.T) {

    /* TEST WITH INVALID PARAMETERS */

    utc, _ := time.LoadLocation("UTC")
    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    var dateTimeEmpty []time.Time
    var dataEmpty []float64

    _, err := chronobiology.IntradailyVariability(dateTimeEmpty, dataEmpty)

    if err == nil {
        t.Error("Expected error : Empty")
    }

    var dateTimeDifferentSize []time.Time
    var dataDifferentSize []float64

    dateTimeDifferentSize = append(dateTimeDifferentSize, tempDateTime)
    tempDateTime          = tempDateTime.Add(60 * time.Second)
    dateTimeDifferentSize = append(dateTimeDifferentSize, tempDateTime)
    tempDateTime          = tempDateTime.Add(60 * time.Second)

    dataDifferentSize = append(dataDifferentSize, 250.0)
    dataDifferentSize = append(dataDifferentSize, 250.0)
    dataDifferentSize = append(dataDifferentSize, 250.0)

    _, err = chronobiology.IntradailyVariability(dateTimeEmpty, dataEmpty)

    if err == nil {
        t.Error("Expected error : DifferentSize")
    }

    var dateTimeLess2Hours []time.Time
    var dataLess2Hours []float64

    dateTimeLess2Hours = append(dateTimeLess2Hours, tempDateTime)
    tempDateTime       = tempDateTime.Add(60 * time.Second)
    dateTimeLess2Hours = append(dateTimeLess2Hours, tempDateTime)
    tempDateTime       = tempDateTime.Add(60 * time.Second)
    dateTimeLess2Hours = append(dateTimeLess2Hours, tempDateTime)
    tempDateTime       = tempDateTime.Add(60 * time.Second)

    dataLess2Hours = append(dataLess2Hours, 250.0)
    dataLess2Hours = append(dataLess2Hours, 250.0)
    dataLess2Hours = append(dataLess2Hours, 250.0)

    _, err = chronobiology.IntradailyVariability(dateTimeEmpty, dataEmpty)

    if err == nil {
        t.Error("Expected error : LessThan2Hours")
    }

    var dateTime60secs []time.Time
    var data60secs []float64

    tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
    for index := 0; index < 240; index++ {
        dateTime60secs = append(dateTime60secs, tempDateTime)

        var value float64
        if index < 50 {
            value = 100.0
        } else if index < 100 {
            value = 150.0
        } else if index < 150 {
            value = 225.0
        } else if index < 200 {
            value = 250.0
        } else {
            value = 300.0
        }

        data60secs     = append(data60secs, value)
        tempDateTime   = tempDateTime.Add(60 * time.Second)
    }

    dateTime30secs, data30secs, err := chronobiology.ConvertDataBasedOnEpoch(dateTime60secs, data60secs, 30)

    // Table tests
    var tTests = []struct {
        dateTime []time.Time
        data []float64
        iv1 float64
        iv2 float64
    }{
        { dateTime60secs, data60secs, 0.0096, 0.0192 },
        { dateTime30secs, data30secs, 0.0096, 0.0192 },
    }

    // Test with all values in the table
    for _, table := range tTests {
        iv, err := chronobiology.IntradailyVariability(table.dateTime, table.data)

        if err != nil {
            t.Error("Expected error = nil.")
        } else {
            if !floatEquals(iv[1], table.iv1) {
                t.Error(
                    "Expected: ", table.iv1,
                    "Received: ", iv[1],
                )
            }
            if !floatEquals(iv[2], table.iv2) {
                t.Error(
                    "Expected: ", table.iv2,
                    "Received: ", iv[2],
                )
            }
        }
    }
}

func TestAverageDay(t *testing.T) {

      utc, _ := time.LoadLocation("UTC")
      tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

      /* Test with invalid parameters */

      var dateTimeInvalid []time.Time
      var dataInvalid []float64

      _, _, err := chronobiology.AverageDay(dateTimeInvalid, dataInvalid)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      dataInvalid = append(dataInvalid, 35.50)

      _, _, err = chronobiology.AverageDay(dateTimeInvalid, dataInvalid)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      dataInvalid = nil

      for index := 0; index < 20; index++ {
          dateTimeInvalid = append(dateTimeInvalid, tempDateTime)
          tempDateTime = tempDateTime.Add(1 * time.Hour)
          dataInvalid = append(dataInvalid, 35.50)
      }

      _, _, err = chronobiology.AverageDay(dateTimeInvalid, dataInvalid)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      /* Test with valid parameters */

      /* Test 1 */

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)

      var dateTime1 []time.Time
      var data1 []float64

      for index := 0; index < 72; index++ {
          dateTime1 = append(dateTime1, tempDateTime)
          tempDateTime = tempDateTime.Add(1 * time.Hour)

          if index < 24 {
              data1 = append(data1, 45.50)
          } else if index < 48 {
              data1 = append(data1, 102.50)
          } else {
              data1 = append(data1, 86.50)
          }
      }

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)

      var newDateTime1 []time.Time
      var newData1 []float64

      for index := 0; index < 24; index++ {
            newDateTime1 = append(newDateTime1, tempDateTime)
            tempDateTime = tempDateTime.Add(1 * time.Hour)
            newData1 = append(newData1, 78.1667)
      }

      /* Test 2 */

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)

      var dateTime2 []time.Time
      var data2 []float64

      for index := 0; index < 60; index++ {
          dateTime2 = append(dateTime2, tempDateTime)
          tempDateTime = tempDateTime.Add(1 * time.Hour)

          if index < 24 {
              data2 = append(data2, 50.00)
          } else if index < 48 {
              data2 = append(data2, 150.00)
          } else {
              data2 = append(data2, 100.00)
          }
      }

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)

      var newDateTime2 []time.Time
      var newData2 []float64

      for index := 0; index < 24; index++ {
            newDateTime2 = append(newDateTime2, tempDateTime)
            tempDateTime = tempDateTime.Add(1 * time.Hour)
            newData2 = append(newData2, 100.00)
      }

      // Table tests
      var tTests = []struct {
          dateTime []time.Time
          data []float64
          expectedDateTime []time.Time
          expectedData []float64
      }{
          { dateTime1, data1, newDateTime1, newData1 },
          { dateTime2, data2, newDateTime2, newData2 },
      }

      // Test with all values in the table
      for _, table := range tTests {
          newDateTime, newData, err := chronobiology.AverageDay(table.dateTime, table.data)

          if err != nil {
              t.Error("Expected error = nil.")
          } else {
              if !sliceTimeEquals(newDateTime, table.expectedDateTime) {
                  t.Error(
                      "Different DateTime Slices.",
                  )
              }
              if !sliceFloatEquals(newData, table.expectedData) {
                  t.Error(
                      "Different Data Slices.",
                  )
              }
          }
      }
}

func TestFilterDataByDateTime(t *testing.T) {

      utc, _ := time.LoadLocation("UTC")
      tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

      /* Test with invalid parameters */

      var dateTimeInvalid []time.Time
      var dataInvalid []float64
      startTimeInvalid := time.Date(2015,1,1,0,0,0,0,utc)
      endTimeInvalid   := time.Date(2015,1,2,0,0,0,0,utc)

      // Empty slices
      _, _, err := chronobiology.FilterDataByDateTime(dateTimeInvalid, dataInvalid, startTimeInvalid, endTimeInvalid)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      dataInvalid = append(dataInvalid, 35.50)

      // Different sizes
      _, _, err = chronobiology.FilterDataByDateTime(dateTimeInvalid, dataInvalid, startTimeInvalid, endTimeInvalid)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      dateTimeInvalid  = append(dateTimeInvalid, startTimeInvalid)
      startTimeInvalid = time.Date(2015,1,1,0,0,0,0,utc)
      endTimeInvalid   = time.Date(2014,1,1,0,0,0,0,utc)

      // Invalid time range
      _, _, err = chronobiology.FilterDataByDateTime(dateTimeInvalid, dataInvalid, startTimeInvalid, endTimeInvalid)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      /* Test with valid parameters */

      /* The base data */

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)

      var dateTime []time.Time
      var data []float64

      // 01/01/2015 - 00:00:00 <-> 04/01/2015 - 23:00:00
      for index := 0; index < 96; index++ {
          dateTime = append(dateTime, tempDateTime)
          tempDateTime = tempDateTime.Add(1 * time.Hour)
          data = append(data, 100.00)
      }

      /* Test 1 */

      var newDateTime1 []time.Time
      var newData1 []float64
      startTime1 := time.Date(2015,1,3,0,0,0,0,utc)
      endTime1   := time.Date(2015,1,9,0,0,0,0,utc)

      tempDateTime = time.Date(2015,1,3,0,0,0,0,utc)
      // 03/01/2015 - 00:00:00 <-> 04/01/2015 - 23:00:00
      for index := 0; index < 48; index++ {
          newDateTime1 = append(newDateTime1, tempDateTime)
          tempDateTime = tempDateTime.Add(1 * time.Hour)
          newData1     = append(newData1, 100.00)
      }

     /* Test 2 */

      var newDateTime2 []time.Time
      var newData2 []float64
      startTime2 := time.Date(2015,1,2,2,0,0,0,utc)
      endTime2   := time.Date(2015,1,3,9,0,0,0,utc)

      tempDateTime = time.Date(2015,1,2,2,0,0,0,utc)
      // 02/01/2015 - 02:00:00 <-> 03/01/2015 - 09:00:00
      for index := 0; index < 32; index++ {
          newDateTime2 = append(newDateTime2, tempDateTime)
          tempDateTime = tempDateTime.Add(1 * time.Hour)
          newData2     = append(newData2, 100.00)
      }

      /* Test 3 */

      startTime3 := time.Date(2014,12,20,0,0,0,0,utc)
      endTime3   := time.Date(2015,2,10,0,0,0,0,utc)

      // Table tests
      var tTests = []struct {
          dateTime []time.Time
          data []float64
          startTime time.Time
          endTime time.Time
          newDateTime []time.Time
          newData []float64
      }{
          { dateTime, data, startTime1, endTime1, newDateTime1, newData1 },
          { dateTime, data, startTime2, endTime2, newDateTime2, newData2 },
          { dateTime, data, startTime3, endTime3, dateTime, data },
      }

      // Test with all values in the table
      for _, table := range tTests {
          newDateTime, newData, err := chronobiology.FilterDataByDateTime(table.dateTime, table.data, table.startTime, table.endTime)

          if err != nil {
              t.Error("Expected error = nil.")
          } else {
              if !sliceTimeEquals(newDateTime, table.newDateTime) {
                  t.Error(
                      "Different DateTime Slices.",
                  )
              }
              if !sliceFloatEquals(newData, table.newData) {
                  t.Error(
                      "Different Data Slices.",
                  )
              }
          }
      }
}

func TestFillGapsInData(t *testing.T) {

      utc, _ := time.LoadLocation("UTC")
      tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

      /* Test with invalid parameters */

      var dateTimeInvalid []time.Time
      var dataInvalid []float64

      // Empty slices
      _, _, err := chronobiology.FillGapsInData(dateTimeInvalid, dataInvalid, 0.0)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      dataInvalid = append(dataInvalid, 35.50)

      // Different sizes
      _, _, err = chronobiology.FillGapsInData(dateTimeInvalid, dataInvalid, 0.0)

      if err == nil {
          t.Error("Expected error != nil.")
      }

      /* Test with valid parameters */

      /* Test 1 - Without any gap */

      var dateTime1 []time.Time
      var data1 []float64

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
      for index := 0; index < 2880; index++ {
          dateTime1    = append(dateTime1, tempDateTime)
          tempDateTime = tempDateTime.Add(60 * time.Second)
          data1        = append(data1, 100.00)
      }

      /* Test 2 */

      var dateTime2 []time.Time
      var data2 []float64

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
      for index := 0; index < 2880; index++ {
        if index < 2000 || index > 2100 {
            dateTime2 = append(dateTime2, tempDateTime)
            data2     = append(data2, 100.00)
        }
        tempDateTime = tempDateTime.Add(60 * time.Second)
      }

      var newDateTime2 []time.Time
      var newData2 []float64

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
      for index := 0; index < 2880; index++ {
          newDateTime2 = append(newDateTime2, tempDateTime)
          tempDateTime = tempDateTime.Add(60 * time.Second)
          if index < 2000 || index > 2100 {
              newData2 = append(newData2, 100.00)
          } else {
              newData2 = append(newData2, 0.0)
          }
      }

      /* Test 3 */

      var dateTime3 []time.Time
      var data3 []float64

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
      for index := 0; index < 8640; index++ {
        if index < 3000 || index > 5000 {
            dateTime3 = append(dateTime3, tempDateTime)
            data3     = append(data3, 100.00)
        }
        tempDateTime = tempDateTime.Add(30 * time.Second)
      }

      var newDateTime3 []time.Time
      var newData3 []float64

      tempDateTime = time.Date(2015,1,1,0,0,0,0,utc)
      for index := 0; index < 8640; index++ {
          newDateTime3 = append(newDateTime3, tempDateTime)
          newData3     = append(newData3, 100.00)
          tempDateTime = tempDateTime.Add(30 * time.Second)
      }

      // Table tests
      var tTests = []struct {
          dateTime []time.Time
          data []float64
          value float64
          newDateTime []time.Time
          newData []float64
      }{
          { dateTime1, data1,    0.0,    dateTime1,    data1 },
          { dateTime2, data2,    0.0, newDateTime2, newData2 },
          { dateTime3, data3, 100.00, newDateTime3, newData3 },
      }

      // Test with all values in the table
      for _, table := range tTests {
          newDateTime, newData, err := chronobiology.FillGapsInData(table.dateTime, table.data, table.value)

          if err != nil {
              t.Error("Expected error = nil.")
          } else {
              if !sliceTimeEquals(newDateTime, table.newDateTime) {
                  t.Error(
                      "Different DateTime Slices.",
                  )
              }
              if !sliceFloatEquals(newData, table.newData) {
                  t.Error(
                      "Different Data Slices.",
                  )
              }
          }
      }
}
