package jarvis

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

// CreateCpmcPunishFeedBack invokes the jarvis.CreateCpmcPunishFeedBack API synchronously
// api document: https://help.aliyun.com/api/jarvis/createcpmcpunishfeedback.html
func (client *Client) CreateCpmcPunishFeedBack(request *CreateCpmcPunishFeedBackRequest) (response *CreateCpmcPunishFeedBackResponse, err error) {
	response = CreateCreateCpmcPunishFeedBackResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCpmcPunishFeedBackWithChan invokes the jarvis.CreateCpmcPunishFeedBack API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createcpmcpunishfeedback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCpmcPunishFeedBackWithChan(request *CreateCpmcPunishFeedBackRequest) (<-chan *CreateCpmcPunishFeedBackResponse, <-chan error) {
	responseChan := make(chan *CreateCpmcPunishFeedBackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCpmcPunishFeedBack(request)
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

// CreateCpmcPunishFeedBackWithCallback invokes the jarvis.CreateCpmcPunishFeedBack API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createcpmcpunishfeedback.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCpmcPunishFeedBackWithCallback(request *CreateCpmcPunishFeedBackRequest, callback func(response *CreateCpmcPunishFeedBackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCpmcPunishFeedBackResponse
		var err error
		defer close(result)
		response, err = client.CreateCpmcPunishFeedBack(request)
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

// CreateCpmcPunishFeedBackRequest is the request struct for api CreateCpmcPunishFeedBack
type CreateCpmcPunishFeedBackRequest struct {
	*requests.RpcRequest
	FeedBack     requests.Integer `position:"Query" name:"FeedBack"`
	SrcIP        string           `position:"Query" name:"SrcIP"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	DstPort      requests.Integer `position:"Query" name:"DstPort"`
	ProtocolName string           `position:"Query" name:"ProtocolName"`
	SrcPort      requests.Integer `position:"Query" name:"SrcPort"`
	PunishType   string           `position:"Query" name:"PunishType"`
	GmtCreate    string           `position:"Query" name:"GmtCreate"`
	DstIP        string           `position:"Query" name:"DstIP"`
	Lang         string           `position:"Query" name:"Lang"`
	SourceCode   string           `position:"Query" name:"SourceCode"`
}

// CreateCpmcPunishFeedBackResponse is the response struct for api CreateCpmcPunishFeedBack
type CreateCpmcPunishFeedBackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateCreateCpmcPunishFeedBackRequest creates a request to invoke CreateCpmcPunishFeedBack API
func CreateCreateCpmcPunishFeedBackRequest() (request *CreateCpmcPunishFeedBackRequest) {
	request = &CreateCpmcPunishFeedBackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "CreateCpmcPunishFeedBack", "jarvis", "openAPI")
	return
}

// CreateCreateCpmcPunishFeedBackResponse creates a response to parse from CreateCpmcPunishFeedBack response
func CreateCreateCpmcPunishFeedBackResponse() (response *CreateCpmcPunishFeedBackResponse) {
	response = &CreateCpmcPunishFeedBackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
