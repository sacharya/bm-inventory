// Code generated by mockery v1.0.0. DO NOT EDIT.

package installer

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockAPI is an autogenerated mock type for the API type
type MockAPI struct {
	mock.Mock
}

// DeregisterCluster provides a mock function with given fields: ctx, params
func (_m *MockAPI) DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *DeregisterClusterNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *DeregisterClusterParams) *DeregisterClusterNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeregisterClusterNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeregisterClusterParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeregisterHost provides a mock function with given fields: ctx, params
func (_m *MockAPI) DeregisterHost(ctx context.Context, params *DeregisterHostParams) (*DeregisterHostNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *DeregisterHostNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *DeregisterHostParams) *DeregisterHostNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeregisterHostNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeregisterHostParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableHost provides a mock function with given fields: ctx, params
func (_m *MockAPI) DisableHost(ctx context.Context, params *DisableHostParams) (*DisableHostNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *DisableHostNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *DisableHostParams) *DisableHostNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DisableHostNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DisableHostParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadClusterFiles provides a mock function with given fields: ctx, params, writer
func (_m *MockAPI) DownloadClusterFiles(ctx context.Context, params *DownloadClusterFilesParams, writer io.Writer) (*DownloadClusterFilesOK, error) {
	ret := _m.Called(ctx, params, writer)

	var r0 *DownloadClusterFilesOK
	if rf, ok := ret.Get(0).(func(context.Context, *DownloadClusterFilesParams, io.Writer) *DownloadClusterFilesOK); ok {
		r0 = rf(ctx, params, writer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DownloadClusterFilesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DownloadClusterFilesParams, io.Writer) error); ok {
		r1 = rf(ctx, params, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadClusterISO provides a mock function with given fields: ctx, params, writer
func (_m *MockAPI) DownloadClusterISO(ctx context.Context, params *DownloadClusterISOParams, writer io.Writer) (*DownloadClusterISOOK, error) {
	ret := _m.Called(ctx, params, writer)

	var r0 *DownloadClusterISOOK
	if rf, ok := ret.Get(0).(func(context.Context, *DownloadClusterISOParams, io.Writer) *DownloadClusterISOOK); ok {
		r0 = rf(ctx, params, writer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DownloadClusterISOOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DownloadClusterISOParams, io.Writer) error); ok {
		r1 = rf(ctx, params, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadClusterKubeconfig provides a mock function with given fields: ctx, params, writer
func (_m *MockAPI) DownloadClusterKubeconfig(ctx context.Context, params *DownloadClusterKubeconfigParams, writer io.Writer) (*DownloadClusterKubeconfigOK, error) {
	ret := _m.Called(ctx, params, writer)

	var r0 *DownloadClusterKubeconfigOK
	if rf, ok := ret.Get(0).(func(context.Context, *DownloadClusterKubeconfigParams, io.Writer) *DownloadClusterKubeconfigOK); ok {
		r0 = rf(ctx, params, writer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DownloadClusterKubeconfigOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DownloadClusterKubeconfigParams, io.Writer) error); ok {
		r1 = rf(ctx, params, writer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnableHost provides a mock function with given fields: ctx, params
func (_m *MockAPI) EnableHost(ctx context.Context, params *EnableHostParams) (*EnableHostNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *EnableHostNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *EnableHostParams) *EnableHostNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EnableHostNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EnableHostParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateClusterISO provides a mock function with given fields: ctx, params
func (_m *MockAPI) GenerateClusterISO(ctx context.Context, params *GenerateClusterISOParams) (*GenerateClusterISOCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *GenerateClusterISOCreated
	if rf, ok := ret.Get(0).(func(context.Context, *GenerateClusterISOParams) *GenerateClusterISOCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GenerateClusterISOCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GenerateClusterISOParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: ctx, params
func (_m *MockAPI) GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *GetClusterOK
	if rf, ok := ret.Get(0).(func(context.Context, *GetClusterParams) *GetClusterOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetClusterOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetClusterParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCredentials provides a mock function with given fields: ctx, params
func (_m *MockAPI) GetCredentials(ctx context.Context, params *GetCredentialsParams) (*GetCredentialsOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *GetCredentialsOK
	if rf, ok := ret.Get(0).(func(context.Context, *GetCredentialsParams) *GetCredentialsOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetCredentialsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetCredentialsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHost provides a mock function with given fields: ctx, params
func (_m *MockAPI) GetHost(ctx context.Context, params *GetHostParams) (*GetHostOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *GetHostOK
	if rf, ok := ret.Get(0).(func(context.Context, *GetHostParams) *GetHostOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetHostOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetHostParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNextSteps provides a mock function with given fields: ctx, params
func (_m *MockAPI) GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *GetNextStepsOK
	if rf, ok := ret.Get(0).(func(context.Context, *GetNextStepsParams) *GetNextStepsOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetNextStepsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetNextStepsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InstallCluster provides a mock function with given fields: ctx, params
func (_m *MockAPI) InstallCluster(ctx context.Context, params *InstallClusterParams) (*InstallClusterAccepted, error) {
	ret := _m.Called(ctx, params)

	var r0 *InstallClusterAccepted
	if rf, ok := ret.Get(0).(func(context.Context, *InstallClusterParams) *InstallClusterAccepted); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*InstallClusterAccepted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *InstallClusterParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields: ctx, params
func (_m *MockAPI) ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *ListClustersOK
	if rf, ok := ret.Get(0).(func(context.Context, *ListClustersParams) *ListClustersOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListClustersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListClustersParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListHosts provides a mock function with given fields: ctx, params
func (_m *MockAPI) ListHosts(ctx context.Context, params *ListHostsParams) (*ListHostsOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *ListHostsOK
	if rf, ok := ret.Get(0).(func(context.Context, *ListHostsParams) *ListHostsOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListHostsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListHostsParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostStepReply provides a mock function with given fields: ctx, params
func (_m *MockAPI) PostStepReply(ctx context.Context, params *PostStepReplyParams) (*PostStepReplyNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *PostStepReplyNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *PostStepReplyParams) *PostStepReplyNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PostStepReplyNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PostStepReplyParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCluster provides a mock function with given fields: ctx, params
func (_m *MockAPI) RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *RegisterClusterCreated
	if rf, ok := ret.Get(0).(func(context.Context, *RegisterClusterParams) *RegisterClusterCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RegisterClusterCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *RegisterClusterParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterHost provides a mock function with given fields: ctx, params
func (_m *MockAPI) RegisterHost(ctx context.Context, params *RegisterHostParams) (*RegisterHostCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *RegisterHostCreated
	if rf, ok := ret.Get(0).(func(context.Context, *RegisterHostParams) *RegisterHostCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RegisterHostCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *RegisterHostParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDebugStep provides a mock function with given fields: ctx, params
func (_m *MockAPI) SetDebugStep(ctx context.Context, params *SetDebugStepParams) (*SetDebugStepNoContent, error) {
	ret := _m.Called(ctx, params)

	var r0 *SetDebugStepNoContent
	if rf, ok := ret.Get(0).(func(context.Context, *SetDebugStepParams) *SetDebugStepNoContent); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SetDebugStepNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SetDebugStepParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCluster provides a mock function with given fields: ctx, params
func (_m *MockAPI) UpdateCluster(ctx context.Context, params *UpdateClusterParams) (*UpdateClusterCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *UpdateClusterCreated
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateClusterParams) *UpdateClusterCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateClusterCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateClusterParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHostInstallProgress provides a mock function with given fields: ctx, params
func (_m *MockAPI) UpdateHostInstallProgress(ctx context.Context, params *UpdateHostInstallProgressParams) (*UpdateHostInstallProgressOK, error) {
	ret := _m.Called(ctx, params)

	var r0 *UpdateHostInstallProgressOK
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateHostInstallProgressParams) *UpdateHostInstallProgressOK); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateHostInstallProgressOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateHostInstallProgressParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UploadClusterIngressCert provides a mock function with given fields: ctx, params
func (_m *MockAPI) UploadClusterIngressCert(ctx context.Context, params *UploadClusterIngressCertParams) (*UploadClusterIngressCertCreated, error) {
	ret := _m.Called(ctx, params)

	var r0 *UploadClusterIngressCertCreated
	if rf, ok := ret.Get(0).(func(context.Context, *UploadClusterIngressCertParams) *UploadClusterIngressCertCreated); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UploadClusterIngressCertCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UploadClusterIngressCertParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
