package ga

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

// ListAvailableAccelerateAreas invokes the ga.ListAvailableAccelerateAreas API synchronously
func (client *Client) ListAvailableAccelerateAreas(request *ListAvailableAccelerateAreasRequest) (response *ListAvailableAccelerateAreasResponse, err error) {
	response = CreateListAvailableAccelerateAreasResponse()
	err = client.DoAction(request, response)
	return
}

// ListAvailableAccelerateAreasWithChan invokes the ga.ListAvailableAccelerateAreas API asynchronously
func (client *Client) ListAvailableAccelerateAreasWithChan(request *ListAvailableAccelerateAreasRequest) (<-chan *ListAvailableAccelerateAreasResponse, <-chan error) {
	responseChan := make(chan *ListAvailableAccelerateAreasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAvailableAccelerateAreas(request)
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

// ListAvailableAccelerateAreasWithCallback invokes the ga.ListAvailableAccelerateAreas API asynchronously
func (client *Client) ListAvailableAccelerateAreasWithCallback(request *ListAvailableAccelerateAreasRequest, callback func(response *ListAvailableAccelerateAreasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAvailableAccelerateAreasResponse
		var err error
		defer close(result)
		response, err = client.ListAvailableAccelerateAreas(request)
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

// ListAvailableAccelerateAreasRequest is the request struct for api ListAvailableAccelerateAreas
type ListAvailableAccelerateAreasRequest struct {
	*requests.RpcRequest
	AcceleratorId string `position:"Query" name:"AcceleratorId"`
}

// ListAvailableAccelerateAreasResponse is the response struct for api ListAvailableAccelerateAreas
type ListAvailableAccelerateAreasResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Areas     []AreasItem `json:"Areas" xml:"Areas"`
}

// CreateListAvailableAccelerateAreasRequest creates a request to invoke ListAvailableAccelerateAreas API
func CreateListAvailableAccelerateAreasRequest() (request *ListAvailableAccelerateAreasRequest) {
	request = &ListAvailableAccelerateAreasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ListAvailableAccelerateAreas", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAvailableAccelerateAreasResponse creates a response to parse from ListAvailableAccelerateAreas response
func CreateListAvailableAccelerateAreasResponse() (response *ListAvailableAccelerateAreasResponse) {
	response = &ListAvailableAccelerateAreasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
