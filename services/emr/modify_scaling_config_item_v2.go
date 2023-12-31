package emr

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

// ModifyScalingConfigItemV2 invokes the emr.ModifyScalingConfigItemV2 API synchronously
func (client *Client) ModifyScalingConfigItemV2(request *ModifyScalingConfigItemV2Request) (response *ModifyScalingConfigItemV2Response, err error) {
	response = CreateModifyScalingConfigItemV2Response()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingConfigItemV2WithChan invokes the emr.ModifyScalingConfigItemV2 API asynchronously
func (client *Client) ModifyScalingConfigItemV2WithChan(request *ModifyScalingConfigItemV2Request) (<-chan *ModifyScalingConfigItemV2Response, <-chan error) {
	responseChan := make(chan *ModifyScalingConfigItemV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingConfigItemV2(request)
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

// ModifyScalingConfigItemV2WithCallback invokes the emr.ModifyScalingConfigItemV2 API asynchronously
func (client *Client) ModifyScalingConfigItemV2WithCallback(request *ModifyScalingConfigItemV2Request, callback func(response *ModifyScalingConfigItemV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingConfigItemV2Response
		var err error
		defer close(result)
		response, err = client.ModifyScalingConfigItemV2(request)
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

// ModifyScalingConfigItemV2Request is the request struct for api ModifyScalingConfigItemV2
type ModifyScalingConfigItemV2Request struct {
	*requests.RpcRequest
	ConfigItemBizId       string           `position:"Query" name:"ConfigItemBizId"`
	ConfigItemType        string           `position:"Query" name:"ConfigItemType"`
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupBizId     string           `position:"Query" name:"ScalingGroupBizId"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	ConfigItemInformation string           `position:"Query" name:"ConfigItemInformation"`
}

// ModifyScalingConfigItemV2Response is the response struct for api ModifyScalingConfigItemV2
type ModifyScalingConfigItemV2Response struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateModifyScalingConfigItemV2Request creates a request to invoke ModifyScalingConfigItemV2 API
func CreateModifyScalingConfigItemV2Request() (request *ModifyScalingConfigItemV2Request) {
	request = &ModifyScalingConfigItemV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyScalingConfigItemV2", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyScalingConfigItemV2Response creates a response to parse from ModifyScalingConfigItemV2 response
func CreateModifyScalingConfigItemV2Response() (response *ModifyScalingConfigItemV2Response) {
	response = &ModifyScalingConfigItemV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
