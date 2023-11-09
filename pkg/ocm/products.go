package ocm

import (
	"time"

	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

const RosaProductId = "rosa"

func (c *Client) GetTechnologyPreview(id string) (*cmv1.ProductTechnologyPreview, bool, error) {
	response, err := c.ocm.ClustersMgmt().V1().
		Products().Product(RosaProductId).
		TechnologyPreviews().TechnologyPreview(id).
		Get().
		Send()
	if response.Status() == 404 {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, handleErr(response.Error(), err)
	}
	return response.Body(), true, nil
}

func (c *Client) GetTechnologyPreviewMessage(id string, forTime time.Time) (string, error) {
	techPreview, exists, err := c.GetTechnologyPreview(id)
	if err != nil {
		return "", err
	}
	// It assumes that the technology_preview endpoint will be available in CS before the ROSA CLI 1.2.30 is out
	// If no technology preview found, a feature should be considered GA
	// If there is a technology preview for HCP and it's active we'll show the message
	if exists && techPreview.EndDate().After(forTime) {
		return techPreview.AdditionalText(), nil
	}
	return "", nil
}
