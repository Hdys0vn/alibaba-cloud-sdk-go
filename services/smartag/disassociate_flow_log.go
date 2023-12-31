package smartag

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

// DisassociateFlowLog invokes the smartag.DisassociateFlowLog API synchronously
func (client *Client) DisassociateFlowLog(request *DisassociateFlowLogRequest) (response *DisassociateFlowLogResponse, err error) {
	response = CreateDisassociateFlowLogResponse()
	err = client.DoAction(request, response)
	return
}

// DisassociateFlowLogWithChan invokes the smartag.DisassociateFlowLog API asynchronously
func (client *Client) DisassociateFlowLogWithChan(request *DisassociateFlowLogRequest) (<-chan *DisassociateFlowLogResponse, <-chan error) {
	responseChan := make(chan *DisassociateFlowLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisassociateFlowLog(request)
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

// DisassociateFlowLogWithCallback invokes the smartag.DisassociateFlowLog API asynchronously
func (client *Client) DisassociateFlowLogWithCallback(request *DisassociateFlowLogRequest, callback func(response *DisassociateFlowLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisassociateFlowLogResponse
		var err error
		defer close(result)
		response, err = client.DisassociateFlowLog(request)
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

// DisassociateFlowLogRequest is the request struct for api DisassociateFlowLog
type DisassociateFlowLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
}

// DisassociateFlowLogResponse is the response struct for api DisassociateFlowLog
type DisassociateFlowLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisassociateFlowLogRequest creates a request to invoke DisassociateFlowLog API
func CreateDisassociateFlowLogRequest() (request *DisassociateFlowLogRequest) {
	request = &DisassociateFlowLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DisassociateFlowLog", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisassociateFlowLogResponse creates a response to parse from DisassociateFlowLog response
func CreateDisassociateFlowLogResponse() (response *DisassociateFlowLogResponse) {
	response = &DisassociateFlowLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
