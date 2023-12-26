package cdn

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

// ModifyRealtimeLogDelivery invokes the cdn.ModifyRealtimeLogDelivery API synchronously
func (client *Client) ModifyRealtimeLogDelivery(request *ModifyRealtimeLogDeliveryRequest) (response *ModifyRealtimeLogDeliveryResponse, err error) {
	response = CreateModifyRealtimeLogDeliveryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRealtimeLogDeliveryWithChan invokes the cdn.ModifyRealtimeLogDelivery API asynchronously
func (client *Client) ModifyRealtimeLogDeliveryWithChan(request *ModifyRealtimeLogDeliveryRequest) (<-chan *ModifyRealtimeLogDeliveryResponse, <-chan error) {
	responseChan := make(chan *ModifyRealtimeLogDeliveryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRealtimeLogDelivery(request)
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

// ModifyRealtimeLogDeliveryWithCallback invokes the cdn.ModifyRealtimeLogDelivery API asynchronously
func (client *Client) ModifyRealtimeLogDeliveryWithCallback(request *ModifyRealtimeLogDeliveryRequest, callback func(response *ModifyRealtimeLogDeliveryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRealtimeLogDeliveryResponse
		var err error
		defer close(result)
		response, err = client.ModifyRealtimeLogDelivery(request)
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

// ModifyRealtimeLogDeliveryRequest is the request struct for api ModifyRealtimeLogDelivery
type ModifyRealtimeLogDeliveryRequest struct {
	*requests.RpcRequest
	Domain   string `position:"Query" name:"Domain"`
	Project  string `position:"Query" name:"Project"`
	Region   string `position:"Query" name:"Region"`
	Logstore string `position:"Query" name:"Logstore"`
}

// ModifyRealtimeLogDeliveryResponse is the response struct for api ModifyRealtimeLogDelivery
type ModifyRealtimeLogDeliveryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRealtimeLogDeliveryRequest creates a request to invoke ModifyRealtimeLogDelivery API
func CreateModifyRealtimeLogDeliveryRequest() (request *ModifyRealtimeLogDeliveryRequest) {
	request = &ModifyRealtimeLogDeliveryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "ModifyRealtimeLogDelivery", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyRealtimeLogDeliveryResponse creates a response to parse from ModifyRealtimeLogDelivery response
func CreateModifyRealtimeLogDeliveryResponse() (response *ModifyRealtimeLogDeliveryResponse) {
	response = &ModifyRealtimeLogDeliveryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
