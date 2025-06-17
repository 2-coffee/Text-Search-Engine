package utils

type Index map[string][]int

func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add the same ID
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	res := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		// finding the intersection
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			res = append(res, a[i])
			i++
			j++
		}
	}
	return res
}

func (idx Index) Search(text string) []int {
	var res []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if res == nil {
				res = ids
			} else {
				res = Intersection(res, ids)
			}
		} else {
			// Token does not exist
			return nil
		}
	}

	return res
}
