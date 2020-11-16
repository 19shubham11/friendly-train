package reflection

import (
	"testing"
	"reflect"
)


type Person struct {
    Name    string
    Profile Profile
}

type Profile struct {
    Age int
    City string
}
func TestWalk(t *testing.T) {

    cases := []struct{
        Name string
        Input interface{}
        ExpectedCalls []string
    } {
        {
            "Struct with one string field",
            struct {
                Name string
            }{ "Helloo"},
            []string{"Helloo"},
		},
		{
			"Struct with multiple text fields",
			struct {
				Name string
				City string
            }{ "Helloo", "Computer"},
            []string{"Helloo", "Computer"},
		},{
            "Struct with non string field",
            struct {
                Name string
                Age  int
            }{"Shubham", 26},
            []string{"Shubham"},
        },
        {
            "Nested fields",
            Person{
                "Shubham",
                Profile {26, "Berlin"},
            },
            []string{"Shubham", "Berlin"},
        },
        {
            "Pointers to things",
            &Person{
                "Shubham",
                Profile{26, "Berlin"},
            },
            []string{"Shubham", "Berlin"},
        },{
            "Slices",
            []Profile {
                {33, "London"},
                {34, "Reykjavík"},
            },
            []string{"London", "Reykjavík"},
        },{
            "Arrays",
            [2]Profile {
                {33, "London"},
                {34, "Reykjavík"},
            },
            []string{"London", "Reykjavík"},
        },{
            "Map",
            map[string]string{
                "Foo": "Bar",
                "Baz": "Boz",
            },
            []string{"Bar", "Boz"},
        },
        
    }

    for _ , test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })

            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
        })
    }

    t.Run("with maps", func(t *testing.T) {
        aMap := map[string]string{
            "Foo": "Bar",
            "Baz": "Boz",
        }

        var got []string
        walk(aMap, func(input string) {
            got = append(got, input)
        })

        assertContains(t, got, "Bar")
        assertContains(t, got, "Boz")
    })

    t.Run("With channels", func(t *testing.T){
        aChannel := make(chan Profile)

        go func(){
            aChannel <- Profile{33, "Berlin"}
            aChannel <- Profile{34, "Katowice"}
            close(aChannel)
        }()
        var got []string
        want := []string{"Berlin", "Katowice"}

        walk(aChannel, func(input string){
            got = append(got, input)
        })

        if !(reflect.DeepEqual(got, want)) {
            t.Errorf("got %q, want%q", got, want)
        }

    })
    
}

func assertContains(t *testing.T, haysack []string, needle string) {
    t.Helper()
    contains := false

    for _, x := range haysack {
        if x == needle {
            contains = true
        }
    }

    if !contains {
        t.Errorf("expected %+v to contain %q but it didn't", haysack, needle)
    }
}