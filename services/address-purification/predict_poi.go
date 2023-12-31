package address_purification

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

// PredictPOI invokes the address_purification.PredictPOI API synchronously
func (client *Client) PredictPOI(request *PredictPOIRequest) (response *PredictPOIResponse, err error) {
	response = CreatePredictPOIResponse()
	err = client.DoAction(request, response)
	return
}

// PredictPOIWithChan invokes the address_purification.PredictPOI API asynchronously
func (client *Client) PredictPOIWithChan(request *PredictPOIRequest) (<-chan *PredictPOIResponse, <-chan error) {
	responseChan := make(chan *PredictPOIResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PredictPOI(request)
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

// PredictPOIWithCallback invokes the address_purification.PredictPOI API asynchronously
func (client *Client) PredictPOIWithCallback(request *PredictPOIRequest, callback func(response *PredictPOIResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PredictPOIResponse
		var err error
		defer close(result)
		response, err = client.PredictPOI(request)
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

// PredictPOIRequest is the request struct for api PredictPOI
type PredictPOIRequest struct {
	*requests.RpcRequest
	DefaultProvince string `position:"Body" name:"DefaultProvince"`
	ServiceCode     string `position:"Body" name:"ServiceCode"`
	DefaultCity     string `position:"Body" name:"DefaultCity"`
	DefaultDistrict string `position:"Body" name:"DefaultDistrict"`
	AppKey          string `position:"Body" name:"AppKey"`
	Text            string `position:"Body" name:"Text"`
}

// PredictPOIResponse is the response struct for api PredictPOI
type PredictPOIResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreatePredictPOIRequest creates a request to invoke PredictPOI API
func CreatePredictPOIRequest() (request *PredictPOIRequest) {
	request = &PredictPOIRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("address-purification", "2019-11-18", "PredictPOI", "addrp", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePredictPOIResponse creates a response to parse from PredictPOI response
func CreatePredictPOIResponse() (response *PredictPOIResponse) {
	response = &PredictPOIResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
