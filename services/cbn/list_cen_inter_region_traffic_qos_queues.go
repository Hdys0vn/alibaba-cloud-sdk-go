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

// ListCenInterRegionTrafficQosQueues invokes the cbn.ListCenInterRegionTrafficQosQueues API synchronously
func (client *Client) ListCenInterRegionTrafficQosQueues(request *ListCenInterRegionTrafficQosQueuesRequest) (response *ListCenInterRegionTrafficQosQueuesResponse, err error) {
	response = CreateListCenInterRegionTrafficQosQueuesResponse()
	err = client.DoAction(request, response)
	return
}

// ListCenInterRegionTrafficQosQueuesWithChan invokes the cbn.ListCenInterRegionTrafficQosQueues API asynchronously
func (client *Client) ListCenInterRegionTrafficQosQueuesWithChan(request *ListCenInterRegionTrafficQosQueuesRequest) (<-chan *ListCenInterRegionTrafficQosQueuesResponse, <-chan error) {
	responseChan := make(chan *ListCenInterRegionTrafficQosQueuesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCenInterRegionTrafficQosQueues(request)
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

// ListCenInterRegionTrafficQosQueuesWithCallback invokes the cbn.ListCenInterRegionTrafficQosQueues API asynchronously
func (client *Client) ListCenInterRegionTrafficQosQueuesWithCallback(request *ListCenInterRegionTrafficQosQueuesRequest, callback func(response *ListCenInterRegionTrafficQosQueuesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCenInterRegionTrafficQosQueuesResponse
		var err error
		defer close(result)
		response, err = client.ListCenInterRegionTrafficQosQueues(request)
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

// ListCenInterRegionTrafficQosQueuesRequest is the request struct for api ListCenInterRegionTrafficQosQueues
type ListCenInterRegionTrafficQosQueuesRequest struct {
	*requests.RpcRequest
	TrafficQosQueueName        string           `position:"Query" name:"TrafficQosQueueName"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TrafficQosQueueId          string           `position:"Query" name:"TrafficQosQueueId"`
	TrafficQosQueueDescription string           `position:"Query" name:"TrafficQosQueueDescription"`
	NextToken                  string           `position:"Query" name:"NextToken"`
	TrafficQosPolicyId         string           `position:"Query" name:"TrafficQosPolicyId"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	TransitRouterId            string           `position:"Query" name:"TransitRouterId"`
	Version                    string           `position:"Query" name:"Version"`
	TransitRouterAttachmentId  string           `position:"Query" name:"TransitRouterAttachmentId"`
	MaxResults                 requests.Integer `position:"Query" name:"MaxResults"`
}

// ListCenInterRegionTrafficQosQueuesResponse is the response struct for api ListCenInterRegionTrafficQosQueues
type ListCenInterRegionTrafficQosQueuesResponse struct {
	*responses.BaseResponse
	NextToken        string            `json:"NextToken" xml:"NextToken"`
	RequestId        string            `json:"RequestId" xml:"RequestId"`
	TrafficQosQueues []TrafficQosQueue `json:"TrafficQosQueues" xml:"TrafficQosQueues"`
}

// CreateListCenInterRegionTrafficQosQueuesRequest creates a request to invoke ListCenInterRegionTrafficQosQueues API
func CreateListCenInterRegionTrafficQosQueuesRequest() (request *ListCenInterRegionTrafficQosQueuesRequest) {
	request = &ListCenInterRegionTrafficQosQueuesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListCenInterRegionTrafficQosQueues", "", "")
	request.Method = requests.POST
	return
}

// CreateListCenInterRegionTrafficQosQueuesResponse creates a response to parse from ListCenInterRegionTrafficQosQueues response
func CreateListCenInterRegionTrafficQosQueuesResponse() (response *ListCenInterRegionTrafficQosQueuesResponse) {
	response = &ListCenInterRegionTrafficQosQueuesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
