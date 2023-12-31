package yundun_bastionhost

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

// DescribeInstanceLoginTicket invokes the yundun_bastionhost.DescribeInstanceLoginTicket API synchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstanceloginticket.html
func (client *Client) DescribeInstanceLoginTicket(request *DescribeInstanceLoginTicketRequest) (response *DescribeInstanceLoginTicketResponse, err error) {
	response = CreateDescribeInstanceLoginTicketResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceLoginTicketWithChan invokes the yundun_bastionhost.DescribeInstanceLoginTicket API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstanceloginticket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceLoginTicketWithChan(request *DescribeInstanceLoginTicketRequest) (<-chan *DescribeInstanceLoginTicketResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceLoginTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceLoginTicket(request)
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

// DescribeInstanceLoginTicketWithCallback invokes the yundun_bastionhost.DescribeInstanceLoginTicket API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstanceloginticket.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceLoginTicketWithCallback(request *DescribeInstanceLoginTicketRequest, callback func(response *DescribeInstanceLoginTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceLoginTicketResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceLoginTicket(request)
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

// DescribeInstanceLoginTicketRequest is the request struct for api DescribeInstanceLoginTicket
type DescribeInstanceLoginTicketRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeInstanceLoginTicketResponse is the response struct for api DescribeInstanceLoginTicket
type DescribeInstanceLoginTicketResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	LoginTicket LoginTicket `json:"LoginTicket" xml:"LoginTicket"`
}

// CreateDescribeInstanceLoginTicketRequest creates a request to invoke DescribeInstanceLoginTicket API
func CreateDescribeInstanceLoginTicketRequest() (request *DescribeInstanceLoginTicketRequest) {
	request = &DescribeInstanceLoginTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-bastionhost", "2018-10-10", "DescribeInstanceLoginTicket", "bastionhost", "openAPI")
	return
}

// CreateDescribeInstanceLoginTicketResponse creates a response to parse from DescribeInstanceLoginTicket response
func CreateDescribeInstanceLoginTicketResponse() (response *DescribeInstanceLoginTicketResponse) {
	response = &DescribeInstanceLoginTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
