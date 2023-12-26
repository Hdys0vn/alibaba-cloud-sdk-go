package cloudauth

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

// CreateVerifySetting invokes the cloudauth.CreateVerifySetting API synchronously
func (client *Client) CreateVerifySetting(request *CreateVerifySettingRequest) (response *CreateVerifySettingResponse, err error) {
	response = CreateCreateVerifySettingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVerifySettingWithChan invokes the cloudauth.CreateVerifySetting API asynchronously
func (client *Client) CreateVerifySettingWithChan(request *CreateVerifySettingRequest) (<-chan *CreateVerifySettingResponse, <-chan error) {
	responseChan := make(chan *CreateVerifySettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVerifySetting(request)
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

// CreateVerifySettingWithCallback invokes the cloudauth.CreateVerifySetting API asynchronously
func (client *Client) CreateVerifySettingWithCallback(request *CreateVerifySettingRequest, callback func(response *CreateVerifySettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVerifySettingResponse
		var err error
		defer close(result)
		response, err = client.CreateVerifySetting(request)
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

// CreateVerifySettingRequest is the request struct for api CreateVerifySetting
type CreateVerifySettingRequest struct {
	*requests.RpcRequest
	GuideStep   requests.Boolean `position:"Query" name:"GuideStep"`
	ResultStep  requests.Boolean `position:"Query" name:"ResultStep"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Solution    string           `position:"Query" name:"Solution"`
	BizName     string           `position:"Query" name:"BizName"`
	BizType     string           `position:"Query" name:"BizType"`
	PrivacyStep requests.Boolean `position:"Query" name:"PrivacyStep"`
}

// CreateVerifySettingResponse is the response struct for api CreateVerifySetting
type CreateVerifySettingResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	BizType   string   `json:"BizType" xml:"BizType"`
	BizName   string   `json:"BizName" xml:"BizName"`
	Solution  string   `json:"Solution" xml:"Solution"`
	StepList  []string `json:"StepList" xml:"StepList"`
}

// CreateCreateVerifySettingRequest creates a request to invoke CreateVerifySetting API
func CreateCreateVerifySettingRequest() (request *CreateVerifySettingRequest) {
	request = &CreateVerifySettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "CreateVerifySetting", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVerifySettingResponse creates a response to parse from CreateVerifySetting response
func CreateCreateVerifySettingResponse() (response *CreateVerifySettingResponse) {
	response = &CreateVerifySettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
