package cloudcallcenter

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

// CallOnlinePrivacyNumber invokes the cloudcallcenter.CallOnlinePrivacyNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/callonlineprivacynumber.html
func (client *Client) CallOnlinePrivacyNumber(request *CallOnlinePrivacyNumberRequest) (response *CallOnlinePrivacyNumberResponse, err error) {
	response = CreateCallOnlinePrivacyNumberResponse()
	err = client.DoAction(request, response)
	return
}

// CallOnlinePrivacyNumberWithChan invokes the cloudcallcenter.CallOnlinePrivacyNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/callonlineprivacynumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CallOnlinePrivacyNumberWithChan(request *CallOnlinePrivacyNumberRequest) (<-chan *CallOnlinePrivacyNumberResponse, <-chan error) {
	responseChan := make(chan *CallOnlinePrivacyNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CallOnlinePrivacyNumber(request)
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

// CallOnlinePrivacyNumberWithCallback invokes the cloudcallcenter.CallOnlinePrivacyNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/callonlineprivacynumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CallOnlinePrivacyNumberWithCallback(request *CallOnlinePrivacyNumberRequest, callback func(response *CallOnlinePrivacyNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CallOnlinePrivacyNumberResponse
		var err error
		defer close(result)
		response, err = client.CallOnlinePrivacyNumber(request)
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

// CallOnlinePrivacyNumberRequest is the request struct for api CallOnlinePrivacyNumber
type CallOnlinePrivacyNumberRequest struct {
	*requests.RpcRequest
	TelA       string `position:"Query" name:"TelA"`
	TelB       string `position:"Query" name:"TelB"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// CallOnlinePrivacyNumberResponse is the response struct for api CallOnlinePrivacyNumber
type CallOnlinePrivacyNumberResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateCallOnlinePrivacyNumberRequest creates a request to invoke CallOnlinePrivacyNumber API
func CreateCallOnlinePrivacyNumberRequest() (request *CallOnlinePrivacyNumberRequest) {
	request = &CallOnlinePrivacyNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CallOnlinePrivacyNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateCallOnlinePrivacyNumberResponse creates a response to parse from CallOnlinePrivacyNumber response
func CreateCallOnlinePrivacyNumberResponse() (response *CallOnlinePrivacyNumberResponse) {
	response = &CallOnlinePrivacyNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
