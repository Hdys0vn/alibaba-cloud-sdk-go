package imm

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

// GetWebofficeURL invokes the imm.GetWebofficeURL API synchronously
func (client *Client) GetWebofficeURL(request *GetWebofficeURLRequest) (response *GetWebofficeURLResponse, err error) {
	response = CreateGetWebofficeURLResponse()
	err = client.DoAction(request, response)
	return
}

// GetWebofficeURLWithChan invokes the imm.GetWebofficeURL API asynchronously
func (client *Client) GetWebofficeURLWithChan(request *GetWebofficeURLRequest) (<-chan *GetWebofficeURLResponse, <-chan error) {
	responseChan := make(chan *GetWebofficeURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWebofficeURL(request)
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

// GetWebofficeURLWithCallback invokes the imm.GetWebofficeURL API asynchronously
func (client *Client) GetWebofficeURLWithCallback(request *GetWebofficeURLRequest, callback func(response *GetWebofficeURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWebofficeURLResponse
		var err error
		defer close(result)
		response, err = client.GetWebofficeURL(request)
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

// GetWebofficeURLRequest is the request struct for api GetWebofficeURL
type GetWebofficeURLRequest struct {
	*requests.RpcRequest
	SrcType         string           `position:"Query" name:"SrcType"`
	Project         string           `position:"Query" name:"Project"`
	File            string           `position:"Query" name:"File"`
	Hidecmb         requests.Boolean `position:"Query" name:"Hidecmb"`
	NotifyEndpoint  string           `position:"Query" name:"NotifyEndpoint"`
	FileID          string           `position:"Query" name:"FileID"`
	Watermark       string           `position:"Query" name:"Watermark"`
	NotifyTopicName string           `position:"Query" name:"NotifyTopicName"`
	Permission      string           `position:"Query" name:"Permission"`
	User            string           `position:"Query" name:"User"`
}

// GetWebofficeURLResponse is the response struct for api GetWebofficeURL
type GetWebofficeURLResponse struct {
	*responses.BaseResponse
	RefreshToken            string `json:"RefreshToken" xml:"RefreshToken"`
	RequestId               string `json:"RequestId" xml:"RequestId"`
	AccessToken             string `json:"AccessToken" xml:"AccessToken"`
	RefreshTokenExpiredTime string `json:"RefreshTokenExpiredTime" xml:"RefreshTokenExpiredTime"`
	WebofficeURL            string `json:"WebofficeURL" xml:"WebofficeURL"`
	AccessTokenExpiredTime  string `json:"AccessTokenExpiredTime" xml:"AccessTokenExpiredTime"`
}

// CreateGetWebofficeURLRequest creates a request to invoke GetWebofficeURL API
func CreateGetWebofficeURLRequest() (request *GetWebofficeURLRequest) {
	request = &GetWebofficeURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetWebofficeURL", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetWebofficeURLResponse creates a response to parse from GetWebofficeURL response
func CreateGetWebofficeURLResponse() (response *GetWebofficeURLResponse) {
	response = &GetWebofficeURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}