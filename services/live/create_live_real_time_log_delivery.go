package live

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

// CreateLiveRealTimeLogDelivery invokes the live.CreateLiveRealTimeLogDelivery API synchronously
func (client *Client) CreateLiveRealTimeLogDelivery(request *CreateLiveRealTimeLogDeliveryRequest) (response *CreateLiveRealTimeLogDeliveryResponse, err error) {
	response = CreateCreateLiveRealTimeLogDeliveryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLiveRealTimeLogDeliveryWithChan invokes the live.CreateLiveRealTimeLogDelivery API asynchronously
func (client *Client) CreateLiveRealTimeLogDeliveryWithChan(request *CreateLiveRealTimeLogDeliveryRequest) (<-chan *CreateLiveRealTimeLogDeliveryResponse, <-chan error) {
	responseChan := make(chan *CreateLiveRealTimeLogDeliveryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLiveRealTimeLogDelivery(request)
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

// CreateLiveRealTimeLogDeliveryWithCallback invokes the live.CreateLiveRealTimeLogDelivery API asynchronously
func (client *Client) CreateLiveRealTimeLogDeliveryWithCallback(request *CreateLiveRealTimeLogDeliveryRequest, callback func(response *CreateLiveRealTimeLogDeliveryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLiveRealTimeLogDeliveryResponse
		var err error
		defer close(result)
		response, err = client.CreateLiveRealTimeLogDelivery(request)
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

// CreateLiveRealTimeLogDeliveryRequest is the request struct for api CreateLiveRealTimeLogDelivery
type CreateLiveRealTimeLogDeliveryRequest struct {
	*requests.RpcRequest
	Project    string           `position:"Query" name:"Project"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Region     string           `position:"Query" name:"Region"`
	Logstore   string           `position:"Query" name:"Logstore"`
}

// CreateLiveRealTimeLogDeliveryResponse is the response struct for api CreateLiveRealTimeLogDelivery
type CreateLiveRealTimeLogDeliveryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLiveRealTimeLogDeliveryRequest creates a request to invoke CreateLiveRealTimeLogDelivery API
func CreateCreateLiveRealTimeLogDeliveryRequest() (request *CreateLiveRealTimeLogDeliveryRequest) {
	request = &CreateLiveRealTimeLogDeliveryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateLiveRealTimeLogDelivery", "live", "openAPI")
	request.Method = requests.GET
	return
}

// CreateCreateLiveRealTimeLogDeliveryResponse creates a response to parse from CreateLiveRealTimeLogDelivery response
func CreateCreateLiveRealTimeLogDeliveryResponse() (response *CreateLiveRealTimeLogDeliveryResponse) {
	response = &CreateLiveRealTimeLogDeliveryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
