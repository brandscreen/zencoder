package zencoder

import (
	"fmt"
)

// Get Input Details
func (z *Zencoder) GetInputDetails(id string) (*InputMediaFile, error) {
	var details InputMediaFile

	if err := z.getBody(fmt.Sprintf("inputs/%s.json", id), &details); err != nil {
		return nil, err
	}

	return &details, nil
}

// Input Progress
func (z *Zencoder) GetInputProgress(id string) (*FileProgress, error) {
	var details FileProgress

	if err := z.getBody(fmt.Sprintf("inputs/%s/progress.json", id), &details); err != nil {
		return nil, err
	}

	return &details, nil
}
