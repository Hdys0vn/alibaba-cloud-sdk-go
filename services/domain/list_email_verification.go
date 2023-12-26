package domain

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

// ListEmailVerification invokes the domain.ListEmailVerification API synchronously
func (client *Client) ListEmailVerification(request *ListEmailVerificationRequest) (response *ListEmailVerificationResponse, err error) {
	response = CreateListEmailVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// ListEmailVerificationWithChan invokes the domain.ListEmailVerification API asynchronously
func (client *Client) ListEmailVerificationWithChan(request *ListEmailVerificationRequest) (<-chan *ListEmailVerificationResponse, <-chan error) {
	responseChan := make(chan *ListEmailVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEmailVerification(request)
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

// ListEmailVerificationWithCallback invokes the domain.ListEmailVerification API asynchronously
func (client *Client) ListEmailVerificationWithCallback(request *ListEmailVerificationRequest, callback func(response *ListEmailVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEmailVerificationResponse
		var err error
		defer close(result)
		response, err = client.ListEmailVerification(request)
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

// ListEmailVerificationRequest is the request struct for api ListEmailVerification
type ListEmailVerificationRequest struct {
	*requests.RpcRequest
	EndCreateTime      requests.Integer `position:"Query" name:"EndCreateTime"`
	PageNum            requests.Integer `position:"Query" name:"PageNum"`
	VerificationStatus requests.Integer `position:"Query" name:"VerificationStatus"`
	BeginCreateTime    requests.Integer `position:"Query" name:"BeginCreateTime"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	UserClientIp       string           `position:"Query" name:"UserClientIp"`
	Lang               string           `position:"Query" name:"Lang"`
	Email              string           `position:"Query" name:"Email"`
}

// ListEmailVerificationResponse is the response struct for api ListEmailVerification
type ListEmailVerificationResponse struct {
	*responses.BaseResponse
	PrePage        bool                `json:"PrePage" xml:"PrePage"`
	CurrentPageNum int                 `json:"CurrentPageNum" xml:"CurrentPageNum"`
	RequestId      string              `json:"RequestId" xml:"RequestId"`
	PageSize       int                 `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int                 `json:"TotalPageNum" xml:"TotalPageNum"`
	TotalItemNum   int                 `json:"TotalItemNum" xml:"TotalItemNum"`
	NextPage       bool                `json:"NextPage" xml:"NextPage"`
	Data           []EmailVerification `json:"Data" xml:"Data"`
}

// CreateListEmailVerificationRequest creates a request to invoke ListEmailVerification API
func CreateListEmailVerificationRequest() (request *ListEmailVerificationRequest) {
	request = &ListEmailVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "ListEmailVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateListEmailVerificationResponse creates a response to parse from ListEmailVerification response
func CreateListEmailVerificationResponse() (response *ListEmailVerificationResponse) {
	response = &ListEmailVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}