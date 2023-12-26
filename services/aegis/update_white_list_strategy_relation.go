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

// UpdateWhiteListStrategyRelation invokes the aegis.UpdateWhiteListStrategyRelation API synchronously
// api document: https://help.aliyun.com/api/aegis/updatewhiteliststrategyrelation.html
func (client *Client) UpdateWhiteListStrategyRelation(request *UpdateWhiteListStrategyRelationRequest) (response *UpdateWhiteListStrategyRelationResponse, err error) {
	response = CreateUpdateWhiteListStrategyRelationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWhiteListStrategyRelationWithChan invokes the aegis.UpdateWhiteListStrategyRelation API asynchronously
// api document: https://help.aliyun.com/api/aegis/updatewhiteliststrategyrelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWhiteListStrategyRelationWithChan(request *UpdateWhiteListStrategyRelationRequest) (<-chan *UpdateWhiteListStrategyRelationResponse, <-chan error) {
	responseChan := make(chan *UpdateWhiteListStrategyRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWhiteListStrategyRelation(request)
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

// UpdateWhiteListStrategyRelationWithCallback invokes the aegis.UpdateWhiteListStrategyRelation API asynchronously
// api document: https://help.aliyun.com/api/aegis/updatewhiteliststrategyrelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWhiteListStrategyRelationWithCallback(request *UpdateWhiteListStrategyRelationRequest, callback func(response *UpdateWhiteListStrategyRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWhiteListStrategyRelationResponse
		var err error
		defer close(result)
		response, err = client.UpdateWhiteListStrategyRelation(request)
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

// UpdateWhiteListStrategyRelationRequest is the request struct for api UpdateWhiteListStrategyRelation
type UpdateWhiteListStrategyRelationRequest struct {
	*requests.RpcRequest
	SourceIp      string           `position:"Query" name:"SourceIp"`
	ProcessMethod requests.Integer `position:"Query" name:"ProcessMethod"`
	StrategyId    requests.Integer `position:"Query" name:"StrategyId"`
	Lang          string           `position:"Query" name:"Lang"`
	Type          requests.Integer `position:"Query" name:"Type"`
	Uuid          string           `position:"Query" name:"Uuid"`
	Status        requests.Integer `position:"Query" name:"Status"`
}

// UpdateWhiteListStrategyRelationResponse is the response struct for api UpdateWhiteListStrategyRelation
type UpdateWhiteListStrategyRelationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateWhiteListStrategyRelationRequest creates a request to invoke UpdateWhiteListStrategyRelation API
func CreateUpdateWhiteListStrategyRelationRequest() (request *UpdateWhiteListStrategyRelationRequest) {
	request = &UpdateWhiteListStrategyRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "UpdateWhiteListStrategyRelation", "vipaegis", "openAPI")
	return
}

// CreateUpdateWhiteListStrategyRelationResponse creates a response to parse from UpdateWhiteListStrategyRelation response
func CreateUpdateWhiteListStrategyRelationResponse() (response *UpdateWhiteListStrategyRelationResponse) {
	response = &UpdateWhiteListStrategyRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}