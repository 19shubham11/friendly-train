package clockface

import (
	"testing"
	"time"
	"math"
)


func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2020, time.November, 29, hours, minutes, seconds, 0, time.UTC)
}


func testName(t time.Time) string {
    return t.Format("15:04:05")
}


func roughlyEqualFloat(x, y float64) bool {
	const thereshold = 1e-7
	return math.Abs(x - y) < thereshold
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloat(a.X, b.X) && roughlyEqualFloat(a.Y, b.Y)
}

func TestSecondsHandInRadians(t *testing.T) {

	cases := []struct {
		time time.Time
		angle float64
	} {
		{simpleTime(0,0, 30), math.Pi},
		{simpleTime(0,0, 0), 0},
		{simpleTime(0,0, 15), math.Pi / 2},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
        {simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := secondsInRadians(c.time)
            if got != c.angle {
                t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
            }
        })
    }
}


func TestSecondsHandVector(t *testing.T) {
	cases := []struct {
		time time.Time
		point Point
	} {
		{simpleTime(0, 0, 30), Point{0, -1} },
		{simpleTime(0, 0, 15), Point{1, 0} },
		{simpleTime(0, 0, 45), Point{-1, 0} },
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsHandPoint(c.time)

			if ! roughlyEqualPoints(got, c.point) {
				t.Fatalf("got %v want %v", got, c.point)
			}
		})
	}

}

func TestMinuteHandsInRadians(t *testing.T) {
	cases := []struct {
		time time.Time
		angle float64
	} {
		{simpleTime(0,0,0), 0},
		{simpleTime(0,30,0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := minutesInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
    cases := []struct {
        time  time.Time
        point Point
    }{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            got := minuteHandPoint(c.time)
            if !roughlyEqualPoints(got, c.point) {
                t.Fatalf("Wanted %v Point, but got %v", c.point, got)
            }
        })
    }
}