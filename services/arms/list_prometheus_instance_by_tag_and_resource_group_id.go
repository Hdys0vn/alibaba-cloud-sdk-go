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

// ListPrometheusInstanceByTagAndResourceGroupId invokes the arms.ListPrometheusInstanceByTagAndResourceGroupId API synchronously
func (client *Client) ListPrometheusInstanceByTagAndResourceGroupId(request *ListPrometheusInstanceByTagAndResourceGroupIdRequest) (response *ListPrometheusInstanceByTagAndResourceGroupIdResponse, err error) {
	response = CreateListPrometheusInstanceByTagAndResourceGroupIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrometheusInstanceByTagAndResourceGroupIdWithChan invokes the arms.ListPrometheusInstanceByTagAndResourceGroupId API asynchronously
func (client *Client) ListPrometheusInstanceByTagAndResourceGroupIdWithChan(request *ListPrometheusInstanceByTagAndResourceGroupIdRequest) (<-chan *ListPrometheusInstanceByTagAndResourceGroupIdResponse, <-chan error) {
	responseChan := make(chan *ListPrometheusInstanceByTagAndResourceGroupIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrometheusInstanceByTagAndResourceGroupId(request)
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

// ListPrometheusInstanceByTagAndResourceGroupIdWithCallback invokes the arms.ListPrometheusInstanceByTagAndResourceGroupId API asynchronously
func (client *Client) ListPrometheusInstanceByTagAndResourceGroupIdWithCallback(request *ListPrometheusInstanceByTagAndResourceGroupIdRequest, callback func(response *ListPrometheusInstanceByTagAndResourceGroupIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrometheusInstanceByTagAndResourceGroupIdResponse
		var err error
		defer close(result)
		response, err = client.ListPrometheusInstanceByTagAndResourceGroupId(request)
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

// ListPrometheusInstanceByTagAndResourceGroupIdRequest is the request struct for api ListPrometheusInstanceByTagAndResourceGroupId
type ListPrometheusInstanceByTagAndResourceGroupIdRequest struct {
	*requests.RpcRequest
	ResourceGroupId string                                              `position:"Query" name:"ResourceGroupId"`
	ListAsConsole   requests.Boolean                                    `position:"Query" name:"ListAsConsole"`
	Tag             *[]ListPrometheusInstanceByTagAndResourceGroupIdTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// ListPrometheusInstanceByTagAndResourceGroupIdTag is a repeated param struct in ListPrometheusInstanceByTagAndResourceGroupIdRequest
type ListPrometheusInstanceByTagAndResourceGroupIdTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListPrometheusInstanceByTagAndResourceGroupIdResponse is the response struct for api ListPrometheusInstanceByTagAndResourceGroupId
type ListPrometheusInstanceByTagAndResourceGroupIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListPrometheusInstanceByTagAndResourceGroupIdRequest creates a request to invoke ListPrometheusInstanceByTagAndResourceGroupId API
func CreateListPrometheusInstanceByTagAndResourceGroupIdRequest() (request *ListPrometheusInstanceByTagAndResourceGroupIdRequest) {
	request = &ListPrometheusInstanceByTagAndResourceGroupIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListPrometheusInstanceByTagAndResourceGroupId", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPrometheusInstanceByTagAndResourceGroupIdResponse creates a response to parse from ListPrometheusInstanceByTagAndResourceGroupId response
func CreateListPrometheusInstanceByTagAndResourceGroupIdResponse() (response *ListPrometheusInstanceByTagAndResourceGroupIdResponse) {
	response = &ListPrometheusInstanceByTagAndResourceGroupIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
