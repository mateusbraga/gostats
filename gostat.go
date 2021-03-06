package gostat

import (
    "math"
    "sort"
)

func TakeExtremes(s []int64) []int64 {
    sort.Sort(int64slice(s))

    lowerBound := float64(len(s)) * 0.05
    upperBound := float64(len(s)) * 0.95

    return s[int(lowerBound):int(upperBound)]
}

func Mean(s []int64) (float64) {
    var mean float64

    switch l := len(s); {
    case l == 0:
        mean = 0
    case l == 1:
        mean = 1
    case l >= 2:
        mean = (float64(s[0]) + float64(s[1]))/2.0

        i := 3
        for _, n := range s[2:] {
            mean = (mean + (float64(n)/float64(i-1))) * float64(i-1)/float64(i)
            i++
        }
    }

    return mean
}

func StandardDeviation(s []int64) (float64) {
    var temp float64
    mean := Mean(s)

    switch l := len(s); {
    case l == 0:
        temp = 0
    case l == 1:
        temp = 0
    case l >= 2:
        temp = (squaredDiff(float64(s[0]), mean) + squaredDiff(float64(s[1]), mean))/2.0

        i := 3
        for _, n := range s[2:] {
            squaredDiff := squaredDiff(float64(n), mean)

            temp = (temp + (float64(squaredDiff)/float64(i-1))) * float64(i-1)/float64(i)
            i++
        }
    }

    return math.Sqrt(temp)
}

func squaredDiff(a, b float64) float64 {
    temp := a - b
    return temp * temp
}

type int64slice []int64
func (s int64slice) Len() int { return len(s) }
func (s int64slice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s int64slice) Less(i, j int) bool { return s[i] < s[j] }
