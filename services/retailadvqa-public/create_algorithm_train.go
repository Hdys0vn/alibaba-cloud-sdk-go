package retailadvqa_public

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

// CreateAlgorithmTrain invokes the retailadvqa_public.CreateAlgorithmTrain API synchronously
func (client *Client) CreateAlgorithmTrain(request *CreateAlgorithmTrainRequest) (response *CreateAlgorithmTrainResponse, err error) {
	response = CreateCreateAlgorithmTrainResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAlgorithmTrainWithChan invokes the retailadvqa_public.CreateAlgorithmTrain API asynchronously
func (client *Client) CreateAlgorithmTrainWithChan(request *CreateAlgorithmTrainRequest) (<-chan *CreateAlgorithmTrainResponse, <-chan error) {
	responseChan := make(chan *CreateAlgorithmTrainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlgorithmTrain(request)
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

// CreateAlgorithmTrainWithCallback invokes the retailadvqa_public.CreateAlgorithmTrain API asynchronously
func (client *Client) CreateAlgorithmTrainWithCallback(request *CreateAlgorithmTrainRequest, callback func(response *CreateAlgorithmTrainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlgorithmTrainResponse
		var err error
		defer close(result)
		response, err = client.CreateAlgorithmTrain(request)
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

// CreateAlgorithmTrainRequest is the request struct for api CreateAlgorithmTrain
type CreateAlgorithmTrainRequest struct {
	*requests.RpcRequest
	AccessId                string                                `position:"Query" name:"AccessId"`
	TaskType                string                                `position:"Body" name:"TaskType"`
	DsId                    string                                `position:"Body" name:"DsId"`
	BehaviorTypeValue       string                                `position:"Body" name:"BehaviorTypeValue"`
	TenantId                string                                `position:"Query" name:"TenantId"`
	BehaviorTableInfo       CreateAlgorithmTrainBehaviorTableInfo `position:"Body" name:"BehaviorTableInfo"  type:"Struct"`
	Periods                 requests.Integer                      `position:"Body" name:"Periods"`
	BehaviorObjectTypeValue string                                `position:"Body" name:"BehaviorObjectTypeValue"`
	WorkspaceId             string                                `position:"Query" name:"WorkspaceId"`
	GoodsTableInfo          CreateAlgorithmTrainGoodsTableInfo    `position:"Body" name:"GoodsTableInfo"  type:"Struct"`
}

// CreateAlgorithmTrainBehaviorTableInfo is a repeated param struct in CreateAlgorithmTrainRequest
type CreateAlgorithmTrainBehaviorTableInfo struct {
	BehaviorCountsField        string `name:"BehaviorCountsField"`
	BehaviorDateField          string `name:"BehaviorDateField"`
	BehaviorTableName          string `name:"BehaviorTableName"`
	BehaviorObjectTypeField    string `name:"BehaviorObjectTypeField"`
	BehaviorChannelField       string `name:"BehaviorChannelField"`
	BehaviorAmountsField       string `name:"BehaviorAmountsField"`
	BehaviorTypeField          string `name:"BehaviorTypeField"`
	BehaviorUserIdField        string `name:"BehaviorUserIdField"`
	BehaviorObjectValueField   string `name:"BehaviorObjectValueField"`
	BehaviorObjectValueIdField string `name:"BehaviorObjectValueIdField"`
}

// CreateAlgorithmTrainGoodsTableInfo is a repeated param struct in CreateAlgorithmTrainRequest
type CreateAlgorithmTrainGoodsTableInfo struct {
	GoodsIdField   string `name:"GoodsIdField"`
	GoodsNameField string `name:"GoodsNameField"`
	GoodsTableName string `name:"GoodsTableName"`
}

// CreateAlgorithmTrainResponse is the response struct for api CreateAlgorithmTrain
type CreateAlgorithmTrainResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	ErrorDesc string                     `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string                     `json:"TraceId" xml:"TraceId"`
	ErrorCode string                     `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool                       `json:"Success" xml:"Success"`
	Data      DataInCreateAlgorithmTrain `json:"Data" xml:"Data"`
}

// CreateCreateAlgorithmTrainRequest creates a request to invoke CreateAlgorithmTrain API
func CreateCreateAlgorithmTrainRequest() (request *CreateAlgorithmTrainRequest) {
	request = &CreateAlgorithmTrainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "CreateAlgorithmTrain", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAlgorithmTrainResponse creates a response to parse from CreateAlgorithmTrain response
func CreateCreateAlgorithmTrainResponse() (response *CreateAlgorithmTrainResponse) {
	response = &CreateAlgorithmTrainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
