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

// DescribeSagTrafficTopN invokes the smartag.DescribeSagTrafficTopN API synchronously
func (client *Client) DescribeSagTrafficTopN(request *DescribeSagTrafficTopNRequest) (response *DescribeSagTrafficTopNResponse, err error) {
	response = CreateDescribeSagTrafficTopNResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagTrafficTopNWithChan invokes the smartag.DescribeSagTrafficTopN API asynchronously
func (client *Client) DescribeSagTrafficTopNWithChan(request *DescribeSagTrafficTopNRequest) (<-chan *DescribeSagTrafficTopNResponse, <-chan error) {
	responseChan := make(chan *DescribeSagTrafficTopNResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagTrafficTopN(request)
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

// DescribeSagTrafficTopNWithCallback invokes the smartag.DescribeSagTrafficTopN API asynchronously
func (client *Client) DescribeSagTrafficTopNWithCallback(request *DescribeSagTrafficTopNRequest, callback func(response *DescribeSagTrafficTopNResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagTrafficTopNResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagTrafficTopN(request)
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

// DescribeSagTrafficTopNRequest is the request struct for api DescribeSagTrafficTopN
type DescribeSagTrafficTopNRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Size                 requests.Integer `position:"Query" name:"Size"`
}

// DescribeSagTrafficTopNResponse is the response struct for api DescribeSagTrafficTopN
type DescribeSagTrafficTopNResponse struct {
	*responses.BaseResponse
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	TrafficTopN []TrafficTopNItem `json:"TrafficTopN" xml:"TrafficTopN"`
}

// CreateDescribeSagTrafficTopNRequest creates a request to invoke DescribeSagTrafficTopN API
func CreateDescribeSagTrafficTopNRequest() (request *DescribeSagTrafficTopNRequest) {
	request = &DescribeSagTrafficTopNRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagTrafficTopN", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagTrafficTopNResponse creates a response to parse from DescribeSagTrafficTopN response
func CreateDescribeSagTrafficTopNResponse() (response *DescribeSagTrafficTopNResponse) {
	response = &DescribeSagTrafficTopNResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
