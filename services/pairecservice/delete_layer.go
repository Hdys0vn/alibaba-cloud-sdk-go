package pairecservice

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

// DeleteLayer invokes the pairecservice.DeleteLayer API synchronously
func (client *Client) DeleteLayer(request *DeleteLayerRequest) (response *DeleteLayerResponse, err error) {
	response = CreateDeleteLayerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLayerWithChan invokes the pairecservice.DeleteLayer API asynchronously
func (client *Client) DeleteLayerWithChan(request *DeleteLayerRequest) (<-chan *DeleteLayerResponse, <-chan error) {
	responseChan := make(chan *DeleteLayerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLayer(request)
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

// DeleteLayerWithCallback invokes the pairecservice.DeleteLayer API asynchronously
func (client *Client) DeleteLayerWithCallback(request *DeleteLayerRequest, callback func(response *DeleteLayerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLayerResponse
		var err error
		defer close(result)
		response, err = client.DeleteLayer(request)
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

// DeleteLayerRequest is the request struct for api DeleteLayer
type DeleteLayerRequest struct {
	*requests.RoaRequest
	LayerId    string `position:"Path" name:"LayerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DeleteLayerResponse is the response struct for api DeleteLayer
type DeleteLayerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLayerRequest creates a request to invoke DeleteLayer API
func CreateDeleteLayerRequest() (request *DeleteLayerRequest) {
	request = &DeleteLayerRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "DeleteLayer", "/api/v1/layers/[LayerId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteLayerResponse creates a response to parse from DeleteLayer response
func CreateDeleteLayerResponse() (response *DeleteLayerResponse) {
	response = &DeleteLayerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}