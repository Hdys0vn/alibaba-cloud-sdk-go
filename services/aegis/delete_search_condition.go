package aegis

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

// DeleteSearchCondition invokes the aegis.DeleteSearchCondition API synchronously
// api document: https://help.aliyun.com/api/aegis/deletesearchcondition.html
func (client *Client) DeleteSearchCondition(request *DeleteSearchConditionRequest) (response *DeleteSearchConditionResponse, err error) {
	response = CreateDeleteSearchConditionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSearchConditionWithChan invokes the aegis.DeleteSearchCondition API asynchronously
// api document: https://help.aliyun.com/api/aegis/deletesearchcondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSearchConditionWithChan(request *DeleteSearchConditionRequest) (<-chan *DeleteSearchConditionResponse, <-chan error) {
	responseChan := make(chan *DeleteSearchConditionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSearchCondition(request)
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

// DeleteSearchConditionWithCallback invokes the aegis.DeleteSearchCondition API asynchronously
// api document: https://help.aliyun.com/api/aegis/deletesearchcondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSearchConditionWithCallback(request *DeleteSearchConditionRequest, callback func(response *DeleteSearchConditionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSearchConditionResponse
		var err error
		defer close(result)
		response, err = client.DeleteSearchCondition(request)
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

// DeleteSearchConditionRequest is the request struct for api DeleteSearchCondition
type DeleteSearchConditionRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Name     string `position:"Query" name:"Name"`
}

// DeleteSearchConditionResponse is the response struct for api DeleteSearchCondition
type DeleteSearchConditionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSearchConditionRequest creates a request to invoke DeleteSearchCondition API
func CreateDeleteSearchConditionRequest() (request *DeleteSearchConditionRequest) {
	request = &DeleteSearchConditionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DeleteSearchCondition", "vipaegis", "openAPI")
	return
}

// CreateDeleteSearchConditionResponse creates a response to parse from DeleteSearchCondition response
func CreateDeleteSearchConditionResponse() (response *DeleteSearchConditionResponse) {
	response = &DeleteSearchConditionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}