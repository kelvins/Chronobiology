
package chronobiology_test

import (
    "time"
    "testing"
    "github.com/kelvins/chronobiology"
)

func TestHigherActivity(t *testing.T) {

    // Get UTC
    utc, _ := time.LoadLocation("UTC")

    // Create the vectors
    var myDateTime []time.Time
    var myData []float64

    // Call the function with empty vectors
    _, _, err := chronobiology.HigherActivity(5, myDateTime, myData)

    if err == nil {
        t.Error("1 - expect: dateTime is empty")
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
        t.Error("2 - expect: data is empty")
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
        t.Error("3 - dateTime and data has different sizes")
    }

    myData = append(myData, 050) // 12

    higherActivity, onsetHigherActivity, err := chronobiology.HigherActivity(0, myDateTime, myData)

    if err == nil {
        t.Error("4 - expect: invalid hours")
    }

    higherActivity, onsetHigherActivity, err = chronobiology.HigherActivity(5, myDateTime, myData)

    if err != nil {
        t.Error("5 - error not nil")
    }
    if higherActivity != 482.0 {
        t.Error("6 - expect: 482.0 - received: ", higherActivity)
    }
    if onsetHigherActivity.Format("02/01/2006 15:04:05") != "01/01/2015 06:00:00" {
        t.Error("7 - expect: 01/01/2015 06:00:00 - received: ", onsetHigherActivity.Format("02/01/2006 15:04:05"))
    }

    higherActivity, onsetHigherActivity, err = chronobiology.HigherActivity(6, myDateTime, myData)

    if err != nil {
        t.Error("8 - error not nil")
    }
    if higherActivity != 418.3333 {
        t.Error("9 - expect: 418.3333 - received: ", higherActivity)
    }
    if onsetHigherActivity.Format("02/01/2006 15:04:05") != "01/01/2015 05:00:00" {
        t.Error("10 - expect: 01/01/2015 05:00:00 - received: ", onsetHigherActivity.Format("02/01/2006 15:04:05"))
    }

    higherActivity, onsetHigherActivity, err = chronobiology.HigherActivity(7, myDateTime, myData)

    if err != nil {
        t.Error("11 - error not nil")
    }
    if higherActivity != 364.2857 {
        t.Error("12 - expect: 364.2857 - received: ", higherActivity)
    }
    if onsetHigherActivity.Format("02/01/2006 15:04:05") != "01/01/2015 05:00:00" {
        t.Error("13 - expect: 01/01/2015 05:00:00 - received: ", onsetHigherActivity.Format("02/01/2006 15:04:05"))
    }

    higherActivity, onsetHigherActivity, err = chronobiology.HigherActivity(10, myDateTime, myData)

    if err != nil {
        t.Error("14 - error not nil")
    }
    if higherActivity != 305.5000 {
        t.Error("15 - expect: 305.5000 - received: ", higherActivity)
    }
    if onsetHigherActivity.Format("02/01/2006 15:04:05") != "01/01/2015 01:00:00" {
        t.Error("16 - expect: 01/01/2015 01:00:00 - received: ", onsetHigherActivity.Format("02/01/2006 15:04:05"))
    }
}
