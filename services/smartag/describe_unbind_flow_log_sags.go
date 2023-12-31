package smartag

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

// DescribeUnbindFlowLogSags invokes the smartag.DescribeUnbindFlowLogSags API synchronously
func (client *Client) DescribeUnbindFlowLogSags(request *DescribeUnbindFlowLogSagsRequest) (response *DescribeUnbindFlowLogSagsResponse, err error) {
	response = CreateDescribeUnbindFlowLogSagsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUnbindFlowLogSagsWithChan invokes the smartag.DescribeUnbindFlowLogSags API asynchronously
func (client *Client) DescribeUnbindFlowLogSagsWithChan(request *DescribeUnbindFlowLogSagsRequest) (<-chan *DescribeUnbindFlowLogSagsResponse, <-chan error) {
	responseChan := make(chan *DescribeUnbindFlowLogSagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUnbindFlowLogSags(request)
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

// DescribeUnbindFlowLogSagsWithCallback invokes the smartag.DescribeUnbindFlowLogSags API asynchronously
func (client *Client) DescribeUnbindFlowLogSagsWithCallback(request *DescribeUnbindFlowLogSagsRequest, callback func(response *DescribeUnbindFlowLogSagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUnbindFlowLogSagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeUnbindFlowLogSags(request)
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

// DescribeUnbindFlowLogSagsRequest is the request struct for api DescribeUnbindFlowLogSags
type DescribeUnbindFlowLogSagsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUnbindFlowLogSagsResponse is the response struct for api DescribeUnbindFlowLogSags
type DescribeUnbindFlowLogSagsResponse struct {
	*responses.BaseResponse
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Count     int                             `json:"Count" xml:"Count"`
	Sags      SagsInDescribeUnbindFlowLogSags `json:"Sags" xml:"Sags"`
}

// CreateDescribeUnbindFlowLogSagsRequest creates a request to invoke DescribeUnbindFlowLogSags API
func CreateDescribeUnbindFlowLogSagsRequest() (request *DescribeUnbindFlowLogSagsRequest) {
	request = &DescribeUnbindFlowLogSagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeUnbindFlowLogSags", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUnbindFlowLogSagsResponse creates a response to parse from DescribeUnbindFlowLogSags response
func CreateDescribeUnbindFlowLogSagsResponse() (response *DescribeUnbindFlowLogSagsResponse) {
	response = &DescribeUnbindFlowLogSagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
