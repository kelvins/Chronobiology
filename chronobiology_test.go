
package chronobiology_test

import (
    "time"
    "testing"
    "github.com/kelvins/chronobiology"
)

type testHigherActivityPair struct {
  hours int
  higherActivity float64
  onsetHigherActivity string
}

var higherActivityTests = []testHigherActivityPair{
  { 05, 482.0000, "01/01/2015 06:00:00" },
  { 06, 418.3333, "01/01/2015 05:00:00" },
  { 07, 364.2857, "01/01/2015 05:00:00" },
  { 10, 305.5000, "01/01/2015 01:00:00" },
}

func TestHigherActivity(t *testing.T) {

    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the vectors
    var myDateTime []time.Time
    var myData []float64

    // Call the function with empty vectors
    _, _, err := chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: dateTime is empty")
    }

    tempDateTime := time.Date(2015,1,1,0,0,0,0,utc)

    // Fill the myDateTime with 1 - 12 hours
    for index := 0; index < 12; index++ {
        tempDateTime = tempDateTime.Add(1 * time.Hour)
        myDateTime = append(myDateTime, tempDateTime)
    }

    // Call the function with myData empty
    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: data is empty")
    }

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

    _, _, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("Expect: dateTime and data has different sizes")
    }

    myData = append(myData, 050) // 12

    _, _, err = chronobiology.HigherActivity(0, myDateTime, myData)

    if err == nil {
        t.Error("Expect: invalid hours")
    }

    _, _, err = chronobiology.HigherActivity(20, myDateTime, myData)

    if err == nil {
        t.Error("Expect: time range lower than the hours passed as parameter")
    }

    for _, pair := range higherActivityTests {
        higherActivity, onsetHigherActivity, err := chronobiology.HigherActivity(pair.hours, myDateTime, myData)
        if err != nil {
            t.Error(
                "For: ", pair.hours, " hours - ",
                "expect: error not nil",
            )
        }
        if higherActivity != pair.higherActivity {
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
