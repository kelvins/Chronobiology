
package chronobiology_test

import (
    "time"
    "testing"
    "github.com/kelvins/chronobiology"
)

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
    // Table tests 1
    var tTests1 = []struct {
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
    for _, pair := range tTests1 {
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
