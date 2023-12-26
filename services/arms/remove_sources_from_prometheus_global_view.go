package arms

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

// RemoveSourcesFromPrometheusGlobalView invokes the arms.RemoveSourcesFromPrometheusGlobalView API synchronously
func (client *Client) RemoveSourcesFromPrometheusGlobalView(request *RemoveSourcesFromPrometheusGlobalViewRequest) (response *RemoveSourcesFromPrometheusGlobalViewResponse, err error) {
	response = CreateRemoveSourcesFromPrometheusGlobalViewResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveSourcesFromPrometheusGlobalViewWithChan invokes the arms.RemoveSourcesFromPrometheusGlobalView API asynchronously
func (client *Client) RemoveSourcesFromPrometheusGlobalViewWithChan(request *RemoveSourcesFromPrometheusGlobalViewRequest) (<-chan *RemoveSourcesFromPrometheusGlobalViewResponse, <-chan error) {
	responseChan := make(chan *RemoveSourcesFromPrometheusGlobalViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveSourcesFromPrometheusGlobalView(request)
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

// RemoveSourcesFromPrometheusGlobalViewWithCallback invokes the arms.RemoveSourcesFromPrometheusGlobalView API asynchronously
func (client *Client) RemoveSourcesFromPrometheusGlobalViewWithCallback(request *RemoveSourcesFromPrometheusGlobalViewRequest, callback func(response *RemoveSourcesFromPrometheusGlobalViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveSourcesFromPrometheusGlobalViewResponse
		var err error
		defer close(result)
		response, err = client.RemoveSourcesFromPrometheusGlobalView(request)
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

// RemoveSourcesFromPrometheusGlobalViewRequest is the request struct for api RemoveSourcesFromPrometheusGlobalView
type RemoveSourcesFromPrometheusGlobalViewRequest struct {
	*requests.RpcRequest
	SourceNames         string `position:"Query" name:"SourceNames"`
	GlobalViewClusterId string `position:"Query" name:"GlobalViewClusterId"`
	GroupName           string `position:"Query" name:"GroupName"`
}

// RemoveSourcesFromPrometheusGlobalViewResponse is the response struct for api RemoveSourcesFromPrometheusGlobalView
type RemoveSourcesFromPrometheusGlobalViewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      int    `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRemoveSourcesFromPrometheusGlobalViewRequest creates a request to invoke RemoveSourcesFromPrometheusGlobalView API
func CreateRemoveSourcesFromPrometheusGlobalViewRequest() (request *RemoveSourcesFromPrometheusGlobalViewRequest) {
	request = &RemoveSourcesFromPrometheusGlobalViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "RemoveSourcesFromPrometheusGlobalView", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveSourcesFromPrometheusGlobalViewResponse creates a response to parse from RemoveSourcesFromPrometheusGlobalView response
func CreateRemoveSourcesFromPrometheusGlobalViewResponse() (response *RemoveSourcesFromPrometheusGlobalViewResponse) {
	response = &RemoveSourcesFromPrometheusGlobalViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
