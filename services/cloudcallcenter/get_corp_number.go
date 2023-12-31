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

// GetCorpNumber invokes the cloudcallcenter.GetCorpNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcorpnumber.html
func (client *Client) GetCorpNumber(request *GetCorpNumberRequest) (response *GetCorpNumberResponse, err error) {
	response = CreateGetCorpNumberResponse()
	err = client.DoAction(request, response)
	return
}

// GetCorpNumberWithChan invokes the cloudcallcenter.GetCorpNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcorpnumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCorpNumberWithChan(request *GetCorpNumberRequest) (<-chan *GetCorpNumberResponse, <-chan error) {
	responseChan := make(chan *GetCorpNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCorpNumber(request)
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

// GetCorpNumberWithCallback invokes the cloudcallcenter.GetCorpNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcorpnumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCorpNumberWithCallback(request *GetCorpNumberRequest, callback func(response *GetCorpNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCorpNumberResponse
		var err error
		defer close(result)
		response, err = client.GetCorpNumber(request)
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

// GetCorpNumberRequest is the request struct for api GetCorpNumber
type GetCorpNumberRequest struct {
	*requests.RpcRequest
	TaobaoUid requests.Integer `position:"Query" name:"TaobaoUid"`
	RamId     requests.Integer `position:"Query" name:"RamId"`
}

// GetCorpNumberResponse is the response struct for api GetCorpNumber
type GetCorpNumberResponse struct {
	*responses.BaseResponse
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	Success        bool                   `json:"Success" xml:"Success"`
	Code           string                 `json:"Code" xml:"Code"`
	Message        string                 `json:"Message" xml:"Message"`
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Numbers        NumbersInGetCorpNumber `json:"Numbers" xml:"Numbers"`
}

// CreateGetCorpNumberRequest creates a request to invoke GetCorpNumber API
func CreateGetCorpNumberRequest() (request *GetCorpNumberRequest) {
	request = &GetCorpNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetCorpNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCorpNumberResponse creates a response to parse from GetCorpNumber response
func CreateGetCorpNumberResponse() (response *GetCorpNumberResponse) {
	response = &GetCorpNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
