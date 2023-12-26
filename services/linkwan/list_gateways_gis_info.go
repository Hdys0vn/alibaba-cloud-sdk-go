package linkwan

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

// ListGatewaysGisInfo invokes the linkwan.ListGatewaysGisInfo API synchronously
func (client *Client) ListGatewaysGisInfo(request *ListGatewaysGisInfoRequest) (response *ListGatewaysGisInfoResponse, err error) {
	response = CreateListGatewaysGisInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ListGatewaysGisInfoWithChan invokes the linkwan.ListGatewaysGisInfo API asynchronously
func (client *Client) ListGatewaysGisInfoWithChan(request *ListGatewaysGisInfoRequest) (<-chan *ListGatewaysGisInfoResponse, <-chan error) {
	responseChan := make(chan *ListGatewaysGisInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGatewaysGisInfo(request)
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

// ListGatewaysGisInfoWithCallback invokes the linkwan.ListGatewaysGisInfo API asynchronously
func (client *Client) ListGatewaysGisInfoWithCallback(request *ListGatewaysGisInfoRequest, callback func(response *ListGatewaysGisInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGatewaysGisInfoResponse
		var err error
		defer close(result)
		response, err = client.ListGatewaysGisInfo(request)
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

// ListGatewaysGisInfoRequest is the request struct for api ListGatewaysGisInfo
type ListGatewaysGisInfoRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// ListGatewaysGisInfoResponse is the response struct for api ListGatewaysGisInfo
type ListGatewaysGisInfoResponse struct {
	*responses.BaseResponse
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Success   bool             `json:"Success" xml:"Success"`
	Data      []GatewayGisInfo `json:"Data" xml:"Data"`
}

// CreateListGatewaysGisInfoRequest creates a request to invoke ListGatewaysGisInfo API
func CreateListGatewaysGisInfoRequest() (request *ListGatewaysGisInfoRequest) {
	request = &ListGatewaysGisInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListGatewaysGisInfo", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGatewaysGisInfoResponse creates a response to parse from ListGatewaysGisInfo response
func CreateListGatewaysGisInfoResponse() (response *ListGatewaysGisInfoResponse) {
	response = &ListGatewaysGisInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
