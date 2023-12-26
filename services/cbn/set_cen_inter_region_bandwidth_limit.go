package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// SetCenInterRegionBandwidthLimit invokes the cbn.SetCenInterRegionBandwidthLimit API synchronously
func (client *Client) SetCenInterRegionBandwidthLimit(request *SetCenInterRegionBandwidthLimitRequest) (response *SetCenInterRegionBandwidthLimitResponse, err error) {
	response = CreateSetCenInterRegionBandwidthLimitResponse()
	err = client.DoAction(request, response)
	return
}

// SetCenInterRegionBandwidthLimitWithChan invokes the cbn.SetCenInterRegionBandwidthLimit API asynchronously
func (client *Client) SetCenInterRegionBandwidthLimitWithChan(request *SetCenInterRegionBandwidthLimitRequest) (<-chan *SetCenInterRegionBandwidthLimitResponse, <-chan error) {
	responseChan := make(chan *SetCenInterRegionBandwidthLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCenInterRegionBandwidthLimit(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetCenInterRegionBandwidthLimitWithCallback invokes the cbn.SetCenInterRegionBandwidthLimit API asynchronously
func (client *Client) SetCenInterRegionBandwidthLimitWithCallback(request *SetCenInterRegionBandwidthLimitRequest, callback func(response *SetCenInterRegionBandwidthLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCenInterRegionBandwidthLimitResponse
		var err error
		defer close(result)
		response, err = client.SetCenInterRegionBandwidthLimit(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetCenInterRegionBandwidthLimitRequest is the request struct for api SetCenInterRegionBandwidthLimit
type SetCenInterRegionBandwidthLimitRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OppositeRegionId     string           `position:"Query" name:"OppositeRegionId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	LocalRegionId        string           `position:"Query" name:"LocalRegionId"`
	BandwidthLimit       requests.Integer `position:"Query" name:"BandwidthLimit"`
}

// SetCenInterRegionBandwidthLimitResponse is the response struct for api SetCenInterRegionBandwidthLimit
type SetCenInterRegionBandwidthLimitResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCenInterRegionBandwidthLimitRequest creates a request to invoke SetCenInterRegionBandwidthLimit API
func CreateSetCenInterRegionBandwidthLimitRequest() (request *SetCenInterRegionBandwidthLimitRequest) {
	request = &SetCenInterRegionBandwidthLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "SetCenInterRegionBandwidthLimit", "", "")
	request.Method = requests.POST
	return
}

// CreateSetCenInterRegionBandwidthLimitResponse creates a response to parse from SetCenInterRegionBandwidthLimit response
func CreateSetCenInterRegionBandwidthLimitResponse() (response *SetCenInterRegionBandwidthLimitResponse) {
	response = &SetCenInterRegionBandwidthLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
