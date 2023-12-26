package vs

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

// DescribeDeviceChannels invokes the vs.DescribeDeviceChannels API synchronously
func (client *Client) DescribeDeviceChannels(request *DescribeDeviceChannelsRequest) (response *DescribeDeviceChannelsResponse, err error) {
	response = CreateDescribeDeviceChannelsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeviceChannelsWithChan invokes the vs.DescribeDeviceChannels API asynchronously
func (client *Client) DescribeDeviceChannelsWithChan(request *DescribeDeviceChannelsRequest) (<-chan *DescribeDeviceChannelsResponse, <-chan error) {
	responseChan := make(chan *DescribeDeviceChannelsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeviceChannels(request)
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

// DescribeDeviceChannelsWithCallback invokes the vs.DescribeDeviceChannels API asynchronously
func (client *Client) DescribeDeviceChannelsWithCallback(request *DescribeDeviceChannelsRequest, callback func(response *DescribeDeviceChannelsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeviceChannelsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeviceChannels(request)
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

// DescribeDeviceChannelsRequest is the request struct for api DescribeDeviceChannels
type DescribeDeviceChannelsRequest struct {
	*requests.RpcRequest
	PageNum  requests.Integer `position:"Query" name:"PageNum"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
	Id       string           `position:"Query" name:"Id"`
	ShowLog  string           `position:"Query" name:"ShowLog"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDeviceChannelsResponse is the response struct for api DescribeDeviceChannels
type DescribeDeviceChannelsResponse struct {
	*responses.BaseResponse
	PageNum    int64     `json:"PageNum" xml:"PageNum"`
	PageSize   int64     `json:"PageSize" xml:"PageSize"`
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int64     `json:"TotalCount" xml:"TotalCount"`
	PageCount  int64     `json:"PageCount" xml:"PageCount"`
	Channels   []Channel `json:"Channels" xml:"Channels"`
}

// CreateDescribeDeviceChannelsRequest creates a request to invoke DescribeDeviceChannels API
func CreateDescribeDeviceChannelsRequest() (request *DescribeDeviceChannelsRequest) {
	request = &DescribeDeviceChannelsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeDeviceChannels", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDeviceChannelsResponse creates a response to parse from DescribeDeviceChannels response
func CreateDescribeDeviceChannelsResponse() (response *DescribeDeviceChannelsResponse) {
	response = &DescribeDeviceChannelsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
