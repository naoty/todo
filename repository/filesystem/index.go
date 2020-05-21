package filesystem

import (
	"fmt"
)

type index struct {
	Todos    map[string][]int `json:"todos"`
	Archived map[string][]int `json:"archived"`
	Metadata metadata         `json:"metadata"`
}

func (idx *index) removeID(id int) {
	idFound := false
	for key, ids := range idx.Todos {
		if idFound {
			break
		}

		var _ids []int
		for _, _id := range ids {
			if _id == id {
				idFound = true
				continue
			}

			_ids = append(_ids, _id)
		}

		if len(_ids) == 0 {
			delete(idx.Todos, key)
		} else {
			idx.Todos[key] = _ids
		}
	}

	key := fmt.Sprintf("%d", id)
	delete(idx.Todos, key)
}

func (idx *index) generateNextID() int {
	if len(idx.Metadata.MissingIDs) == 0 {
		id := idx.Metadata.LastID + 1
		idx.Metadata.LastID = id
		return id
	}

	id := idx.Metadata.MissingIDs[0]
	idx.Metadata.MissingIDs = idx.Metadata.MissingIDs[1:]
	return id
}
