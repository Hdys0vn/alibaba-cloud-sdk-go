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

// QueryDeviceTunnel invokes the iot.QueryDeviceTunnel API synchronously
func (client *Client) QueryDeviceTunnel(request *QueryDeviceTunnelRequest) (response *QueryDeviceTunnelResponse, err error) {
	response = CreateQueryDeviceTunnelResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceTunnelWithChan invokes the iot.QueryDeviceTunnel API asynchronously
func (client *Client) QueryDeviceTunnelWithChan(request *QueryDeviceTunnelRequest) (<-chan *QueryDeviceTunnelResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceTunnelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceTunnel(request)
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

// QueryDeviceTunnelWithCallback invokes the iot.QueryDeviceTunnel API asynchronously
func (client *Client) QueryDeviceTunnelWithCallback(request *QueryDeviceTunnelRequest, callback func(response *QueryDeviceTunnelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceTunnelResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceTunnel(request)
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

// QueryDeviceTunnelRequest is the request struct for api QueryDeviceTunnel
type QueryDeviceTunnelRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	TunnelId      string `position:"Query" name:"TunnelId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// QueryDeviceTunnelResponse is the response struct for api QueryDeviceTunnel
type QueryDeviceTunnelResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryDeviceTunnelRequest creates a request to invoke QueryDeviceTunnel API
func CreateQueryDeviceTunnelRequest() (request *QueryDeviceTunnelRequest) {
	request = &QueryDeviceTunnelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceTunnel", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceTunnelResponse creates a response to parse from QueryDeviceTunnel response
func CreateQueryDeviceTunnelResponse() (response *QueryDeviceTunnelResponse) {
	response = &QueryDeviceTunnelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
