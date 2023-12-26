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

// AddAliClusterIdsToPrometheusGlobalView invokes the arms.AddAliClusterIdsToPrometheusGlobalView API synchronously
func (client *Client) AddAliClusterIdsToPrometheusGlobalView(request *AddAliClusterIdsToPrometheusGlobalViewRequest) (response *AddAliClusterIdsToPrometheusGlobalViewResponse, err error) {
	response = CreateAddAliClusterIdsToPrometheusGlobalViewResponse()
	err = client.DoAction(request, response)
	return
}

// AddAliClusterIdsToPrometheusGlobalViewWithChan invokes the arms.AddAliClusterIdsToPrometheusGlobalView API asynchronously
func (client *Client) AddAliClusterIdsToPrometheusGlobalViewWithChan(request *AddAliClusterIdsToPrometheusGlobalViewRequest) (<-chan *AddAliClusterIdsToPrometheusGlobalViewResponse, <-chan error) {
	responseChan := make(chan *AddAliClusterIdsToPrometheusGlobalViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAliClusterIdsToPrometheusGlobalView(request)
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

// AddAliClusterIdsToPrometheusGlobalViewWithCallback invokes the arms.AddAliClusterIdsToPrometheusGlobalView API asynchronously
func (client *Client) AddAliClusterIdsToPrometheusGlobalViewWithCallback(request *AddAliClusterIdsToPrometheusGlobalViewRequest, callback func(response *AddAliClusterIdsToPrometheusGlobalViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAliClusterIdsToPrometheusGlobalViewResponse
		var err error
		defer close(result)
		response, err = client.AddAliClusterIdsToPrometheusGlobalView(request)
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

// AddAliClusterIdsToPrometheusGlobalViewRequest is the request struct for api AddAliClusterIdsToPrometheusGlobalView
type AddAliClusterIdsToPrometheusGlobalViewRequest struct {
	*requests.RpcRequest
	GlobalViewClusterId string `position:"Query" name:"GlobalViewClusterId"`
	ClusterIds          string `position:"Query" name:"ClusterIds"`
	GroupName           string `position:"Query" name:"GroupName"`
}

// AddAliClusterIdsToPrometheusGlobalViewResponse is the response struct for api AddAliClusterIdsToPrometheusGlobalView
type AddAliClusterIdsToPrometheusGlobalViewResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddAliClusterIdsToPrometheusGlobalViewRequest creates a request to invoke AddAliClusterIdsToPrometheusGlobalView API
func CreateAddAliClusterIdsToPrometheusGlobalViewRequest() (request *AddAliClusterIdsToPrometheusGlobalViewRequest) {
	request = &AddAliClusterIdsToPrometheusGlobalViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "AddAliClusterIdsToPrometheusGlobalView", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddAliClusterIdsToPrometheusGlobalViewResponse creates a response to parse from AddAliClusterIdsToPrometheusGlobalView response
func CreateAddAliClusterIdsToPrometheusGlobalViewResponse() (response *AddAliClusterIdsToPrometheusGlobalViewResponse) {
	response = &AddAliClusterIdsToPrometheusGlobalViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
