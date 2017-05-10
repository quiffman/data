package list

import (
	"path/filepath"
)

const (
	CMTsFile = "data.csv"
)

func ReadCMTs(path string) ([]CMT, error) {
	var raw CMTs
	if err := ReadFile(filepath.Join(path, CMTsFile), &raw); err != nil {
		return nil, err
	}
	return raw, nil
}

func ReadCMTsMap(path string) (map[string]CMT, error) {
	cmts := make(map[string]CMT)

	var raw CMTs
	if err := ReadFile(filepath.Join(path, CMTsFile), &raw); err != nil {
		return nil, err
	}
	for _, l := range raw {
		cmts[l.EVENT_ID] = l
	}

	return cmts, nil
}
