package trademark

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

// InsertTradeIntentionUser invokes the trademark.InsertTradeIntentionUser API synchronously
// api document: https://help.aliyun.com/api/trademark/inserttradeintentionuser.html
func (client *Client) InsertTradeIntentionUser(request *InsertTradeIntentionUserRequest) (response *InsertTradeIntentionUserResponse, err error) {
	response = CreateInsertTradeIntentionUserResponse()
	err = client.DoAction(request, response)
	return
}

// InsertTradeIntentionUserWithChan invokes the trademark.InsertTradeIntentionUser API asynchronously
// api document: https://help.aliyun.com/api/trademark/inserttradeintentionuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertTradeIntentionUserWithChan(request *InsertTradeIntentionUserRequest) (<-chan *InsertTradeIntentionUserResponse, <-chan error) {
	responseChan := make(chan *InsertTradeIntentionUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertTradeIntentionUser(request)
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

// InsertTradeIntentionUserWithCallback invokes the trademark.InsertTradeIntentionUser API asynchronously
// api document: https://help.aliyun.com/api/trademark/inserttradeintentionuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertTradeIntentionUserWithCallback(request *InsertTradeIntentionUserRequest, callback func(response *InsertTradeIntentionUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertTradeIntentionUserResponse
		var err error
		defer close(result)
		response, err = client.InsertTradeIntentionUser(request)
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

// InsertTradeIntentionUserRequest is the request struct for api InsertTradeIntentionUser
type InsertTradeIntentionUserRequest struct {
	*requests.RpcRequest
	Mobile         string           `position:"Query" name:"Mobile"`
	PartnerCode    string           `position:"Query" name:"PartnerCode"`
	Classification string           `position:"Query" name:"Classification"`
	Type           requests.Integer `position:"Query" name:"Type"`
	RegisterNumber string           `position:"Query" name:"RegisterNumber"`
	Vcode          string           `position:"Query" name:"Vcode"`
	UserName       string           `position:"Query" name:"UserName"`
}

// InsertTradeIntentionUserResponse is the response struct for api InsertTradeIntentionUser
type InsertTradeIntentionUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateInsertTradeIntentionUserRequest creates a request to invoke InsertTradeIntentionUser API
func CreateInsertTradeIntentionUserRequest() (request *InsertTradeIntentionUserRequest) {
	request = &InsertTradeIntentionUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "InsertTradeIntentionUser", "trademark", "openAPI")
	return
}

// CreateInsertTradeIntentionUserResponse creates a response to parse from InsertTradeIntentionUser response
func CreateInsertTradeIntentionUserResponse() (response *InsertTradeIntentionUserResponse) {
	response = &InsertTradeIntentionUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
