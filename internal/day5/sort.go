package day5

import "sort"

func sortSeatPosition(spList []seatPosition) {
	sps := &seatPositionSorter{}
	sps.data = spList
	sort.Sort(sps)
}

type seatPositionSorter struct {
	data []seatPosition
}

func (sps *seatPositionSorter) Len() int {
	return len(sps.data)
}

func (sps *seatPositionSorter) Swap(i, j int) {
	sps.data[i], sps.data[j] = sps.data[j], sps.data[i]
}

func (sps *seatPositionSorter) Less(i, j int) bool {
	p, q := &sps.data[i], &sps.data[j]
	if p.row != q.row {
		return p.row < q.row
	}
	return p.col < q.col
}
