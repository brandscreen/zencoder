package zencoder

import (
	"fmt"
)

// Get Output Details
func (z *Zencoder) GetOutputDetails(id string) (*OutputMediaFile, error) {
	var details OutputMediaFile

	if err := z.getBody(fmt.Sprintf("outputs/%s.json", id), &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// Output Progress
func (z *Zencoder) GetOutputProgress(id string) (*FileProgress, error) {
	var details FileProgress

	if err := z.getBody(fmt.Sprintf("outputs/%s/progress.json", id), &details); err != nil {
		return nil, err
	}

	return &details, nil
}
