package elasticsearch

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

// RestartCollector invokes the elasticsearch.RestartCollector API synchronously
func (client *Client) RestartCollector(request *RestartCollectorRequest) (response *RestartCollectorResponse, err error) {
	response = CreateRestartCollectorResponse()
	err = client.DoAction(request, response)
	return
}

// RestartCollectorWithChan invokes the elasticsearch.RestartCollector API asynchronously
func (client *Client) RestartCollectorWithChan(request *RestartCollectorRequest) (<-chan *RestartCollectorResponse, <-chan error) {
	responseChan := make(chan *RestartCollectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartCollector(request)
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

// RestartCollectorWithCallback invokes the elasticsearch.RestartCollector API asynchronously
func (client *Client) RestartCollectorWithCallback(request *RestartCollectorRequest, callback func(response *RestartCollectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartCollectorResponse
		var err error
		defer close(result)
		response, err = client.RestartCollector(request)
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

// RestartCollectorRequest is the request struct for api RestartCollector
type RestartCollectorRequest struct {
	*requests.RoaRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	ResId       string `position:"Path" name:"ResId"`
}

// RestartCollectorResponse is the response struct for api RestartCollector
type RestartCollectorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateRestartCollectorRequest creates a request to invoke RestartCollector API
func CreateRestartCollectorRequest() (request *RestartCollectorRequest) {
	request = &RestartCollectorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "RestartCollector", "/openapi/collectors/[ResId]/actions/restart", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestartCollectorResponse creates a response to parse from RestartCollector response
func CreateRestartCollectorResponse() (response *RestartCollectorResponse) {
	response = &RestartCollectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
