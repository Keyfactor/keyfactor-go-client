// Copyright 2024 Keyfactor
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

//
//import (
//	"reflect"
//	"testing"
//)
//
//func TestClient_DeployPFXCertificate(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		args *DeployPFXArgs
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *DeployPFXResp
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			got, err := c.DeployPFXCertificate(tt.args.args)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("DeployPFXCertificate() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("DeployPFXCertificate() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestClient_DownloadCertificate(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		certId       int
//		thumbprint   string
//		serialNumber string
//		issuerDn     string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *x509.Certificate
//		want1   []*x509.Certificate
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			got, got1, err := c.DownloadCertificate(tt.args.certId, tt.args.thumbprint, tt.args.serialNumber, tt.args.issuerDn)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("DownloadCertificate() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("DownloadCertificate() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("DownloadCertificate() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func TestClient_EnrollCSR(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		ea *EnrollCSRFctArgs
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *EnrollResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			got, err := c.EnrollCSR(tt.args.ea)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("EnrollCSR() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("EnrollCSR() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestClient_EnrollPFX(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		ea *EnrollPFXFctArgs
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *EnrollResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			got, err := c.EnrollPFX(tt.args.ea)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("EnrollPFX() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("EnrollPFX() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestClient_GetCertificateContext(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		gca *GetCertificateContextArgs
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *GetCertificateResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			got, err := c.GetCertificateContext(tt.args.gca)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetCertificateContext() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetCertificateContext() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestClient_RecoverCertificate(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		certId       int
//		thumbprint   string
//		serialNumber string
//		issuerDn     string
//		password     string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    interface{}
//		want1   *x509.Certificate
//		want2   []*x509.Certificate
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			got, got1, got2, err := c.RecoverCertificate(tt.args.certId, tt.args.thumbprint, tt.args.serialNumber, tt.args.issuerDn, tt.args.password)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("RecoverCertificate() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("RecoverCertificate() got = %v, want %v", got, tt.want)
//			}
//			if !reflect.DeepEqual(got1, tt.want1) {
//				t.Errorf("RecoverCertificate() got1 = %v, want %v", got1, tt.want1)
//			}
//			if !reflect.DeepEqual(got2, tt.want2) {
//				t.Errorf("RecoverCertificate() got2 = %v, want %v", got2, tt.want2)
//			}
//		})
//	}
//}
//
//func TestClient_RevokeCert(t *testing.T) {
//	type fields struct {
//		Hostname        string
//		httpClient      *http.Client
//		basicAuthString string
//	}
//	type args struct {
//		ra *RevokeCertArgs
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &Client{
//				Hostname:        tt.fields.Hostname,
//				httpClient:      tt.fields.httpClient,
//				basicAuthString: tt.fields.basicAuthString,
//			}
//			if err := c.RevokeCert(tt.args.ra); (err != nil) != tt.wantErr {
//				t.Errorf("RevokeCert() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_createSubject(t *testing.T) {
//	type args struct {
//		cs CertificateSubject
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := createSubject(tt.args.cs)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("createSubject() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("createSubject() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_decodePKCS12Blob(t *testing.T) {
//	type args struct {
//		resp *EnrollResponse
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := decodePKCS12Blob(tt.args.resp); (err != nil) != tt.wantErr {
//				t.Errorf("decodePKCS12Blob() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_mapTupleArrayToInterface(t *testing.T) {
//	type args struct {
//		i []StringTuple
//	}
//	tests := []struct {
//		name string
//		args args
//		want map[string]interface{}
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := mapTupleArrayToInterface(tt.args.i); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("mapTupleArrayToInterface() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_validateDeployPFXArgs(t *testing.T) {
//	type args struct {
//		dpfxa *DeployPFXArgs
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := validateDeployPFXArgs(tt.args.dpfxa); (err != nil) != tt.wantErr {
//				t.Errorf("validateDeployPFXArgs() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
