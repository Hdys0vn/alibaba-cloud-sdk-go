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

// CreateTopicRouteTable invokes the iot.CreateTopicRouteTable API synchronously
func (client *Client) CreateTopicRouteTable(request *CreateTopicRouteTableRequest) (response *CreateTopicRouteTableResponse, err error) {
	response = CreateCreateTopicRouteTableResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTopicRouteTableWithChan invokes the iot.CreateTopicRouteTable API asynchronously
func (client *Client) CreateTopicRouteTableWithChan(request *CreateTopicRouteTableRequest) (<-chan *CreateTopicRouteTableResponse, <-chan error) {
	responseChan := make(chan *CreateTopicRouteTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTopicRouteTable(request)
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

// CreateTopicRouteTableWithCallback invokes the iot.CreateTopicRouteTable API asynchronously
func (client *Client) CreateTopicRouteTableWithCallback(request *CreateTopicRouteTableRequest, callback func(response *CreateTopicRouteTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTopicRouteTableResponse
		var err error
		defer close(result)
		response, err = client.CreateTopicRouteTable(request)
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

// CreateTopicRouteTableRequest is the request struct for api CreateTopicRouteTable
type CreateTopicRouteTableRequest struct {
	*requests.RpcRequest
	RealTenantId      string    `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string    `position:"Query" name:"RealTripartiteKey"`
	IotInstanceId     string    `position:"Query" name:"IotInstanceId"`
	DstTopic          *[]string `position:"Query" name:"DstTopic"  type:"Repeated"`
	ApiProduct        string    `position:"Body" name:"ApiProduct"`
	ApiRevision       string    `position:"Body" name:"ApiRevision"`
	SrcTopic          string    `position:"Query" name:"SrcTopic"`
}

// CreateTopicRouteTableResponse is the response struct for api CreateTopicRouteTable
type CreateTopicRouteTableResponse struct {
	*responses.BaseResponse
	RequestId     string                               `json:"RequestId" xml:"RequestId"`
	Success       bool                                 `json:"Success" xml:"Success"`
	Code          string                               `json:"Code" xml:"Code"`
	IsAllSucceed  bool                                 `json:"IsAllSucceed" xml:"IsAllSucceed"`
	ErrorMessage  string                               `json:"ErrorMessage" xml:"ErrorMessage"`
	FailureTopics FailureTopicsInCreateTopicRouteTable `json:"FailureTopics" xml:"FailureTopics"`
}

// CreateCreateTopicRouteTableRequest creates a request to invoke CreateTopicRouteTable API
func CreateCreateTopicRouteTableRequest() (request *CreateTopicRouteTableRequest) {
	request = &CreateTopicRouteTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateTopicRouteTable", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTopicRouteTableResponse creates a response to parse from CreateTopicRouteTable response
func CreateCreateTopicRouteTableResponse() (response *CreateTopicRouteTableResponse) {
	response = &CreateTopicRouteTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
