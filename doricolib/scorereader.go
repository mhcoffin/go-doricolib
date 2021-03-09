package doricolib

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func ReadDoricoLib(name string) ScoreLib {
	result := ScoreLib{}
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Errorf("failed to read file %s: %w", name, err))
	}
	xml.Unmarshal(bytes, &result)
	return result
}
