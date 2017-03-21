Chronobiology GoLang Package
==========================

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)
[![Build Status](https://travis-ci.org/kelvins/chronobiology.svg?branch=master)](https://travis-ci.org/kelvins/chronobiology)
[![Coverage Status](https://coveralls.io/repos/github/kelvins/chronobiology/badge.svg?branch=master)](https://coveralls.io/github/kelvins/chronobiology?branch=master)

This is an open source package developed to provide **chronobiology** functions in **GoLang**.

You can use go get command:

    go get github.com/kelvins/chronobiology

----------

GoLang
---------------------------------
**Go**, also known as **GoLang**, is an open source programming language that makes it easy to build simple, reliable, and efficient software.

The Go programming language was developed by **Google** and released in 2009. To know more about the **GoLang** please visit the official [web site][1].

![](http://i.imgur.com/0QgXKrO.png)

----------

Chronobiology
---------------------------------

Basically, **chronobiology** is a field of biology that studies the periodic phenomena in living organisms and their adaptation to solar and lunar rhythms. These cycles are known as **biological rhythms**.

Chronobiological studies include but are not limited to comparative anatomy, physiology, genetics, molecular biology and behavior of organisms within biological rhythms mechanics.

----------

The project
---------------------------------
The **objective** of this project is to provide a easy way to access **chronobiology** functions that can perform **data analysis** in **time series** using **GoLang**.

Some analyzes and features that we pretend to provide with the package:

- [X] Convert the data based on the epoch
- [X] Filter data by date/time
- [X] Calculate Average Day
- [X] Fill gaps in data
- [X] The followed hours of highest activity (e.g. M10)
- [X] The followed hours of lowest activity (e.g. L5)
- [X] Relative Amplitude (RA)
- [X] Intradaily Variability (IV)
- [ ] Interdaily Stability (IS)

Functions provided in the version 1.2:

``` go
// Convert the data and dateTime slices to the new epoch passed by parameter
func ConvertDataBasedOnEpoch(dateTime []time.Time, data []float64, newEpoch int) (newDateTime []time.Time, newData []float64, err error){}

// Function that finds the epoch of a time series (seconds)
func FindEpoch(dateTime []time.Time) (epoch int){}

// Function that finds the highest activity average of the followed X hours (defined by parameter)
func HigherActivity(hours int, dateTime []time.Time, data []float64) (higherActivity float64, onsetHigherActivity time.Time, err error){}

// Function that finds the lowest activity average of the followed X hours (defined by parameter)
func LowerActivity(hours int, dateTime []time.Time, data []float64) (lowerActivity float64, onsetLowerActivity time.Time, err error){}

// Function that finds the highest activity average of the followed 10 hours
func M10(dateTime []time.Time, data []float64) (higherActivity float64, onsetHigherActivity time.Time, err error){}

// Function that finds the lowest activity average of the following 5 hours
func L5(dateTime []time.Time, data []float64) (lowerActivity float64, onsetLowerActivity time.Time, err error){}

// Function that calculates the relative amplitude based on the formula (M10-L5)/(M10+L5)
func RelativeAmplitude(highestAverage float64, lowestAverage float64) (RA float64, err error){}

// Function that calculates the intradaily variability
func IntradailyVariability(dateTime []time.Time, data []float64) (iv []float64, err error){}

// Function that calculates the interdaily stability
func InterdailyStability(dateTime []time.Time, data []float64) (is []float64, err error){}

// Function created to filter the data based on the startTime and endTime passed as parameter
func FilterDataByDateTime(dateTime []time.Time, data []float64, startTime time.Time, endTime time.Time) (newDateTime []time.Time, newData []float64, err error){}

// Function that searches for gaps in the time series and fills it with a specific value passed as parameter (usually zero)
func FillGapsInData(dateTime []time.Time, data []float64, value float64) (newDateTime []time.Time, newData []float64, err error){}

// Creates an average day based on the time series
func AverageDay(dateTime []time.Time, data []float64) (newDateTime []time.Time, newData []float64, err error) {}
```

**Note**: The functions were developed to work with default epoch of 60 seconds (or 15, 30, 120 seconds). If the epoch is something like 17 or 33 seconds, the results can be inaccurate.

**References**: Witting W, Kwa IH, Eikelenboom P, Mirmiran M, Swaab DF. Alterations in the circadian rest-activity rhythm in aging and Alzheimer's disease. Biol Psychiatry 1990;27:563-72.

----------

This project was created under the MIT license. Feel free to contribute by commenting, suggesting, creating issues or sending pull requests.

If you want more information about this project please contact-me: kelvinpfw@hotmail.com

  [1]: https://golang.org
