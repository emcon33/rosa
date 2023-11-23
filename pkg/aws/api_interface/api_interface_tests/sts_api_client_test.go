package aws_test

import (
	"github.com/aws/aws-sdk-go-v2/service/sts"
	. "github.com/onsi/ginkgo/v2"
	f "github.com/openshift/rosa/pkg/aws/api_interface"
	m "github.com/openshift/rosa/pkg/aws/mocks"
)

var _ = Describe("StsApiClient", func() {
	It("is implemented by AWS SDK STS Client", func() {
		awsStsClient := &sts.Client{}
		var _ f.StsApiClient = awsStsClient
	})

	It("is implemented by MockSTSApiClient", func() {
		mockStsApiClient := &m.MockStsApiClient{}
		var _ f.StsApiClient = mockStsApiClient
	})
})
