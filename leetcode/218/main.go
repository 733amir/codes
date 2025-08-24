package main

func getSkyline(buildings []Building) []Skyline {
	merged := []Building{buildings[0]}

	for _, b := range buildings[1:] {
		oldMerged := merged
		merged = make([]Building, 1, len(merged)+2)
		merged[0] = b

		for i, m := range oldMerged {
			got, c := merger(m, merged[len(merged)-1])
			merged = append(merged[:len(merged)-1], got...)

			if !c {
				merged = append(merged, oldMerged[i+1:]...)
				break
			}
		}
	}

	sl := []Skyline{{start: merged[0].start, height: merged[0].height}}
	for i := 1; i < len(merged); i++ {
		if merged[i-1].end != merged[i].start {
			sl = append(sl, Skyline{start: merged[i-1].end, height: 0})
		}
		if merged[i].height != sl[len(sl)-1].height {
			sl = append(sl, Skyline{start: merged[i].start, height: merged[i].height})
		}
	}
	sl = append(sl, Skyline{start: merged[len(merged)-1].end, height: 0})
	return sl
}

func merger(a, b Building) (merged []Building, cont bool) {
	if !collision(a, b) {
		return []Building{a, b}, a.end <= b.start
	}

	if a.start == b.start && a.end == b.end {
		return []Building{
			{start: a.start, end: a.end, height: max(a.height, b.height)},
		}, false
	}

	if a.start == b.start {
		h := a.height
		if b.end == max(a.end, b.end) {
			h = b.height
		}
		return []Building{
			{start: a.start, end: min(a.end, b.end), height: max(a.height, b.height)},
			{start: min(a.end, b.end), end: max(a.end, b.end), height: h},
		}, b.end == max(a.end, b.end)
	}

	if a.end == b.end {
		h := a.height
		if b.start == min(a.start, b.start) {
			h = b.height
		}
		return []Building{
			{start: min(a.start, b.start), end: max(a.start, b.start), height: h},
			{start: max(a.start, b.start), end: a.end, height: max(a.height, b.height)},
		}, false
	}

	h1 := a.height
	if b.start < a.start {
		h1 = b.height
	}
	h2 := a.height
	if a.end < b.end {
		h2 = b.height
	}
	return []Building{
		{start: min(a.start, b.start), end: max(a.start, b.start), height: h1},
		{start: max(a.start, b.start), end: min(a.end, b.end), height: max(a.height, b.height)},
		{start: min(a.end, b.end), end: max(a.end, b.end), height: h2},
	}, b.end == max(a.end, b.end)
}

func collision(a, b Building) bool {
	return max(a.start, b.start) < min(a.end, b.end)
}

type Building struct {
	start  int
	end    int
	height int
}

type Skyline struct {
	start  int
	height int
}
