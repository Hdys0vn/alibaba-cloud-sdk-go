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

// GetAddressGeocode invokes the address_purification.GetAddressGeocode API synchronously
func (client *Client) GetAddressGeocode(request *GetAddressGeocodeRequest) (response *GetAddressGeocodeResponse, err error) {
	response = CreateGetAddressGeocodeResponse()
	err = client.DoAction(request, response)
	return
}

// GetAddressGeocodeWithChan invokes the address_purification.GetAddressGeocode API asynchronously
func (client *Client) GetAddressGeocodeWithChan(request *GetAddressGeocodeRequest) (<-chan *GetAddressGeocodeResponse, <-chan error) {
	responseChan := make(chan *GetAddressGeocodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAddressGeocode(request)
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

// GetAddressGeocodeWithCallback invokes the address_purification.GetAddressGeocode API asynchronously
func (client *Client) GetAddressGeocodeWithCallback(request *GetAddressGeocodeRequest, callback func(response *GetAddressGeocodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAddressGeocodeResponse
		var err error
		defer close(result)
		response, err = client.GetAddressGeocode(request)
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

// GetAddressGeocodeRequest is the request struct for api GetAddressGeocode
type GetAddressGeocodeRequest struct {
	*requests.RpcRequest
	DefaultProvince string `position:"Body" name:"DefaultProvince"`
	ServiceCode     string `position:"Body" name:"ServiceCode"`
	DefaultCity     string `position:"Body" name:"DefaultCity"`
	DefaultDistrict string `position:"Body" name:"DefaultDistrict"`
	AppKey          string `position:"Body" name:"AppKey"`
	Text            string `position:"Body" name:"Text"`
}

// GetAddressGeocodeResponse is the response struct for api GetAddressGeocode
type GetAddressGeocodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetAddressGeocodeRequest creates a request to invoke GetAddressGeocode API
func CreateGetAddressGeocodeRequest() (request *GetAddressGeocodeRequest) {
	request = &GetAddressGeocodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("address-purification", "2019-11-18", "GetAddressGeocode", "addrp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAddressGeocodeResponse creates a response to parse from GetAddressGeocode response
func CreateGetAddressGeocodeResponse() (response *GetAddressGeocodeResponse) {
	response = &GetAddressGeocodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}