package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type rocket struct {
	id, start int
	data      time.Time
	status    string
	durations int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			panic("cant close buffer")
		}
	}(out)

	var amountRecords int
	_, err := fmt.Fscan(in, &amountRecords)
	if err != nil {
		panic("cant read size")
	}

	rocketSlice := make([]*rocket, 0)
	rocketSet := make(map[int]*rocket)

	// считаем логи
	for N := 0; N < amountRecords; N++ {
		var day, hour, min, id int
		var status string
		_, err := fmt.Fscan(in, &day, &hour, &min, &id, &status)
		if err != nil {
			panic("cant read record")
		}

		rocketSlice = append(rocketSlice, &rocket{
			id:     id,
			data:   time.Date(0, 0, day, hour, min, 0, 0, time.UTC),
			status: status,
			start:  day*1440 + hour*60 + min,
		})
	}

	sort.Slice(rocketSlice, func(i, j int) bool {
		return rocketSlice[i].data.Before(rocketSlice[j].data)
	})

	for _, Rocket := range rocketSlice {
		if _, ok := rocketSet[Rocket.id]; !ok {

			rocketSet[Rocket.id] = &rocket{
				id:    Rocket.id,
				start: Rocket.start,
			}
		} else {
			if Rocket.status == "C" || Rocket.status == "S" {
				rocketSet[Rocket.id].durations += Rocket.start - rocketSet[Rocket.id].start
			} else if Rocket.status == "A" {
				rocketSet[Rocket.id].start = Rocket.start
			}
		}
	}

	sliceId := make([]int, len(rocketSet))
	j := 0
	for id := range rocketSet {
		sliceId[j] = id
		j++
	}

	sort.Slice(sliceId, func(i, j int) bool { return sliceId[i] < sliceId[j] })

	for _, id := range sliceId {
		fmt.Fprintf(out, "%d ", rocketSet[id].durations)
	}

}

// 8
//50 7 25 3632 A
//14 23 52 212372 S
//15 0 5 3632 C
//14 21 30 212372 A
//50 7 26 3632 C
//14 21 30 3632 A
//14 21 40 212372 B
//14 23 52 3632 B
