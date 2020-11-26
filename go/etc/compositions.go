package main

import "fmt"

type Part struct {
    Name        string
    Description string
    NeedsSpare  bool
}

type Parts []Part

type Bicycle struct {
    Size string
    Parts
}

var (
    RoadBikeParts = Parts{
        {"chain", "10-speed", true},
        {"tire_size", "23", true},
        {"tape_color", "red", true},
    }

    MountainBikeParts = Parts{
        {"chain", "10-speed", true},
        {"tire_size", "2.1", true},
        {"front_shock", "Manitou", false},
        {"rear_shock", "Fox", true},
    }

    RecumbentBikeParts = Parts{
        {"chain", "9-speed", true},
        {"tire_size", "28", true},
        {"flag", "tall and orange", true},
    }
)

func (parts Parts) Spares() (spares Parts) {
    for _, part := range parts {
        if part.NeedsSpare {
            spares = append(spares, part)
        }
    }
    return spares
}

/*

This function satisfied the Stringer interface, which the fmt package makes use of. It is defined as:
type Stringer interface {
    String() string
}
*/
func (part Part) String() string {
    return fmt.Sprintf("%s: %s", part.Name, part.Description)
}

func main() {
    roadBike := Bicycle{Size: "L", Parts: RoadBikeParts}
    mountainBike := Bicycle{Size: "L", Parts: MountainBikeParts}
    recumbentBike := Bicycle{Size: "L", Parts: RecumbentBikeParts}

    fmt.Println(roadBike.Spares())
    // output = [{chain 10-speed true} {tire_size 23 true} {tape_color red true}]
    // after have func String() = [chain: 10-speed tire_size: 23 tape_color: red]
    fmt.Println(mountainBike.Spares())
    // output = [{chain 10-speed true} {tire_size 2.1 true} {rear_shock Fox true}]
    // after have func String() = [chain: 10-speed tire_size: 2.1 rear_shock: Fox]

    fmt.Println(recumbentBike.Spares())
	// output = [{chain 9-speed true} {tire_size 28 true} {flag tall and orange true}]
	// after have func String() = [chain: 9-speed tire_size: 28 flag: tall and orange]

}

