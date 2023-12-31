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

// DeleteFeeRecordByBizId invokes the cloudcallcenter.DeleteFeeRecordByBizId API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletefeerecordbybizid.html
func (client *Client) DeleteFeeRecordByBizId(request *DeleteFeeRecordByBizIdRequest) (response *DeleteFeeRecordByBizIdResponse, err error) {
	response = CreateDeleteFeeRecordByBizIdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFeeRecordByBizIdWithChan invokes the cloudcallcenter.DeleteFeeRecordByBizId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletefeerecordbybizid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFeeRecordByBizIdWithChan(request *DeleteFeeRecordByBizIdRequest) (<-chan *DeleteFeeRecordByBizIdResponse, <-chan error) {
	responseChan := make(chan *DeleteFeeRecordByBizIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFeeRecordByBizId(request)
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

// DeleteFeeRecordByBizIdWithCallback invokes the cloudcallcenter.DeleteFeeRecordByBizId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletefeerecordbybizid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFeeRecordByBizIdWithCallback(request *DeleteFeeRecordByBizIdRequest, callback func(response *DeleteFeeRecordByBizIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFeeRecordByBizIdResponse
		var err error
		defer close(result)
		response, err = client.DeleteFeeRecordByBizId(request)
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

// DeleteFeeRecordByBizIdRequest is the request struct for api DeleteFeeRecordByBizId
type DeleteFeeRecordByBizIdRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"bizId"`
}

// DeleteFeeRecordByBizIdResponse is the response struct for api DeleteFeeRecordByBizId
type DeleteFeeRecordByBizIdResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Count          int    `json:"Count" xml:"Count"`
}

// CreateDeleteFeeRecordByBizIdRequest creates a request to invoke DeleteFeeRecordByBizId API
func CreateDeleteFeeRecordByBizIdRequest() (request *DeleteFeeRecordByBizIdRequest) {
	request = &DeleteFeeRecordByBizIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteFeeRecordByBizId", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteFeeRecordByBizIdResponse creates a response to parse from DeleteFeeRecordByBizId response
func CreateDeleteFeeRecordByBizIdResponse() (response *DeleteFeeRecordByBizIdResponse) {
	response = &DeleteFeeRecordByBizIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
