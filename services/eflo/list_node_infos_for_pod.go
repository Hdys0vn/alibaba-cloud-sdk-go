package eflo

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

// ListNodeInfosForPod invokes the eflo.ListNodeInfosForPod API synchronously
func (client *Client) ListNodeInfosForPod(request *ListNodeInfosForPodRequest) (response *ListNodeInfosForPodResponse, err error) {
	response = CreateListNodeInfosForPodResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeInfosForPodWithChan invokes the eflo.ListNodeInfosForPod API asynchronously
func (client *Client) ListNodeInfosForPodWithChan(request *ListNodeInfosForPodRequest) (<-chan *ListNodeInfosForPodResponse, <-chan error) {
	responseChan := make(chan *ListNodeInfosForPodResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeInfosForPod(request)
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

// ListNodeInfosForPodWithCallback invokes the eflo.ListNodeInfosForPod API asynchronously
func (client *Client) ListNodeInfosForPodWithCallback(request *ListNodeInfosForPodRequest, callback func(response *ListNodeInfosForPodResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeInfosForPodResponse
		var err error
		defer close(result)
		response, err = client.ListNodeInfosForPod(request)
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

// ListNodeInfosForPodRequest is the request struct for api ListNodeInfosForPod
type ListNodeInfosForPodRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Body" name:"ClusterId"`
	ZoneId    string `position:"Body" name:"ZoneId"`
	NodeId    string `position:"Body" name:"NodeId"`
}

// ListNodeInfosForPodResponse is the response struct for api ListNodeInfosForPod
type ListNodeInfosForPodResponse struct {
	*responses.BaseResponse
	Code      int           `json:"Code" xml:"Code"`
	Message   string        `json:"Message" xml:"Message"`
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Content   []ContentItem `json:"Content" xml:"Content"`
}

// CreateListNodeInfosForPodRequest creates a request to invoke ListNodeInfosForPod API
func CreateListNodeInfosForPodRequest() (request *ListNodeInfosForPodRequest) {
	request = &ListNodeInfosForPodRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "ListNodeInfosForPod", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodeInfosForPodResponse creates a response to parse from ListNodeInfosForPod response
func CreateListNodeInfosForPodResponse() (response *ListNodeInfosForPodResponse) {
	response = &ListNodeInfosForPodResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}