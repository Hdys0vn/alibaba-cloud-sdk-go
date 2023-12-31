package emr

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

// SubmitFlow invokes the emr.SubmitFlow API synchronously
func (client *Client) SubmitFlow(request *SubmitFlowRequest) (response *SubmitFlowResponse, err error) {
	response = CreateSubmitFlowResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitFlowWithChan invokes the emr.SubmitFlow API asynchronously
func (client *Client) SubmitFlowWithChan(request *SubmitFlowRequest) (<-chan *SubmitFlowResponse, <-chan error) {
	responseChan := make(chan *SubmitFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitFlow(request)
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

// SubmitFlowWithCallback invokes the emr.SubmitFlow API asynchronously
func (client *Client) SubmitFlowWithCallback(request *SubmitFlowRequest, callback func(response *SubmitFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitFlowResponse
		var err error
		defer close(result)
		response, err = client.SubmitFlow(request)
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

// SubmitFlowRequest is the request struct for api SubmitFlow
type SubmitFlowRequest struct {
	*requests.RpcRequest
	Conf      string `position:"Query" name:"Conf"`
	ProjectId string `position:"Query" name:"ProjectId"`
	FlowId    string `position:"Query" name:"FlowId"`
}

// SubmitFlowResponse is the response struct for api SubmitFlow
type SubmitFlowResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	Id         string `json:"Id" xml:"Id"`
	Data       string `json:"Data" xml:"Data"`
}

// CreateSubmitFlowRequest creates a request to invoke SubmitFlow API
func CreateSubmitFlowRequest() (request *SubmitFlowRequest) {
	request = &SubmitFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "SubmitFlow", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitFlowResponse creates a response to parse from SubmitFlow response
func CreateSubmitFlowResponse() (response *SubmitFlowResponse) {
	response = &SubmitFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
