package iot

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

// QueryEdgeInstanceMessageRouting invokes the iot.QueryEdgeInstanceMessageRouting API synchronously
func (client *Client) QueryEdgeInstanceMessageRouting(request *QueryEdgeInstanceMessageRoutingRequest) (response *QueryEdgeInstanceMessageRoutingResponse, err error) {
	response = CreateQueryEdgeInstanceMessageRoutingResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEdgeInstanceMessageRoutingWithChan invokes the iot.QueryEdgeInstanceMessageRouting API asynchronously
func (client *Client) QueryEdgeInstanceMessageRoutingWithChan(request *QueryEdgeInstanceMessageRoutingRequest) (<-chan *QueryEdgeInstanceMessageRoutingResponse, <-chan error) {
	responseChan := make(chan *QueryEdgeInstanceMessageRoutingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEdgeInstanceMessageRouting(request)
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

// QueryEdgeInstanceMessageRoutingWithCallback invokes the iot.QueryEdgeInstanceMessageRouting API asynchronously
func (client *Client) QueryEdgeInstanceMessageRoutingWithCallback(request *QueryEdgeInstanceMessageRoutingRequest, callback func(response *QueryEdgeInstanceMessageRoutingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEdgeInstanceMessageRoutingResponse
		var err error
		defer close(result)
		response, err = client.QueryEdgeInstanceMessageRouting(request)
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

// QueryEdgeInstanceMessageRoutingRequest is the request struct for api QueryEdgeInstanceMessageRouting
type QueryEdgeInstanceMessageRoutingRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryEdgeInstanceMessageRoutingResponse is the response struct for api QueryEdgeInstanceMessageRouting
type QueryEdgeInstanceMessageRoutingResponse struct {
	*responses.BaseResponse
	RequestId    string                                `json:"RequestId" xml:"RequestId"`
	Success      bool                                  `json:"Success" xml:"Success"`
	Code         string                                `json:"Code" xml:"Code"`
	ErrorMessage string                                `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryEdgeInstanceMessageRouting `json:"Data" xml:"Data"`
}

// CreateQueryEdgeInstanceMessageRoutingRequest creates a request to invoke QueryEdgeInstanceMessageRouting API
func CreateQueryEdgeInstanceMessageRoutingRequest() (request *QueryEdgeInstanceMessageRoutingRequest) {
	request = &QueryEdgeInstanceMessageRoutingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryEdgeInstanceMessageRouting", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryEdgeInstanceMessageRoutingResponse creates a response to parse from QueryEdgeInstanceMessageRouting response
func CreateQueryEdgeInstanceMessageRoutingResponse() (response *QueryEdgeInstanceMessageRoutingResponse) {
	response = &QueryEdgeInstanceMessageRoutingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
