package airec

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

// OfflineFilteringAlgorithm invokes the airec.OfflineFilteringAlgorithm API synchronously
func (client *Client) OfflineFilteringAlgorithm(request *OfflineFilteringAlgorithmRequest) (response *OfflineFilteringAlgorithmResponse, err error) {
	response = CreateOfflineFilteringAlgorithmResponse()
	err = client.DoAction(request, response)
	return
}

// OfflineFilteringAlgorithmWithChan invokes the airec.OfflineFilteringAlgorithm API asynchronously
func (client *Client) OfflineFilteringAlgorithmWithChan(request *OfflineFilteringAlgorithmRequest) (<-chan *OfflineFilteringAlgorithmResponse, <-chan error) {
	responseChan := make(chan *OfflineFilteringAlgorithmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OfflineFilteringAlgorithm(request)
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

// OfflineFilteringAlgorithmWithCallback invokes the airec.OfflineFilteringAlgorithm API asynchronously
func (client *Client) OfflineFilteringAlgorithmWithCallback(request *OfflineFilteringAlgorithmRequest, callback func(response *OfflineFilteringAlgorithmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OfflineFilteringAlgorithmResponse
		var err error
		defer close(result)
		response, err = client.OfflineFilteringAlgorithm(request)
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

// OfflineFilteringAlgorithmRequest is the request struct for api OfflineFilteringAlgorithm
type OfflineFilteringAlgorithmRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"instanceId"`
	AlgorithmId string `position:"Path" name:"algorithmId"`
}

// OfflineFilteringAlgorithmResponse is the response struct for api OfflineFilteringAlgorithm
type OfflineFilteringAlgorithmResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"requestId" xml:"requestId"`
	Result    ResultInOfflineFilteringAlgorithm `json:"result" xml:"result"`
}

// CreateOfflineFilteringAlgorithmRequest creates a request to invoke OfflineFilteringAlgorithm API
func CreateOfflineFilteringAlgorithmRequest() (request *OfflineFilteringAlgorithmRequest) {
	request = &OfflineFilteringAlgorithmRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "OfflineFilteringAlgorithm", "/v2/openapi/instances/[instanceId]/filtering-algorithms/[algorithmId]/actions/offline", "airec", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOfflineFilteringAlgorithmResponse creates a response to parse from OfflineFilteringAlgorithm response
func CreateOfflineFilteringAlgorithmResponse() (response *OfflineFilteringAlgorithmResponse) {
	response = &OfflineFilteringAlgorithmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
