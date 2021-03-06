package host

import (
	"context"
	"testing"
	"time"

	"github.com/filanov/bm-inventory/internal/connectivity"

	"github.com/go-openapi/swag"

	"github.com/filanov/bm-inventory/internal/hardware"
	"github.com/filanov/bm-inventory/models"
	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("disconnected_state", func() {
	var (
		ctx                       = context.Background()
		state                     API
		db                        *gorm.DB
		currentState              = HostStatusDisconnected
		host                      models.Host
		id, clusterId             strfmt.UUID
		updateReply               *UpdateReply
		updateErr                 error
		expectedReply             *expect
		ctrl                      *gomock.Controller
		mockHWValidator           *hardware.MockValidator
		mockConnectivityValidator *connectivity.MockValidator
	)

	BeforeEach(func() {
		db = prepareDB()
		ctrl = gomock.NewController(GinkgoT())
		mockHWValidator = hardware.NewMockValidator(ctrl)
		mockConnectivityValidator = connectivity.NewMockValidator(ctrl)
		state = &Manager{disconnected: NewDisconnectedState(getTestLog(), db, mockHWValidator)}

		id = strfmt.UUID(uuid.New().String())
		clusterId = strfmt.UUID(uuid.New().String())
		host = getTestHost(id, clusterId, currentState)
		host.CheckedInAt = strfmt.DateTime(time.Now().Add(-time.Hour))
		Expect(db.Create(&host).Error).ShouldNot(HaveOccurred())
		expectedReply = &expect{expectedState: currentState}
		addTestCluster(clusterId, "1.2.3.5", "1.2.3.6", "1.2.3.0/24", db)
	})

	Context("update hw info", func() {
		It("update", func() {
			expectedStatusInfo := mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, false, true, true)
			updateReply, updateErr = state.UpdateHwInfo(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal(defaultInventory()))
				Expect(h.HardwareInfo).Should(Equal("some hw info"))
				Expect(swag.StringValue(h.StatusInfo)).Should(Equal(expectedStatusInfo))
			}
		})
	})

	Context("update inventory", func() {
		It("sufficient_hw", func() {
			expectedStatusInfo := mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, false, true, true)
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.expectedState = HostStatusDisconnected
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal("some hw info"))
				Expect(swag.StringValue(h.StatusInfo)).Should(Equal(expectedStatusInfo))
			}
		})
		It("insufficient_hw", func() {
			mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, false, false, true)
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.expectedState = HostStatusDisconnected
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal("some hw info"))
				Expect(swag.StringValue(h.StatusInfo)).Should(Equal(""))
			}
		})
		It("hw_validation_error", func() {
			mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, true, false, true)
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.expectedState = HostStatusDisconnected
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal(defaultInventory()))
				Expect(swag.StringValue(h.StatusInfo)).Should(Equal(""))
			}
		})
		It("sufficient_hw_insufficient_connectivity", func() {
			host.Role = ""
			mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, false, true, false)
			updateReply, updateErr = state.UpdateInventory(ctx, &host, "some hw info")
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.expectedState = HostStatusDisconnected
			expectedReply.postCheck = func() {
				h := getHost(id, clusterId, db)
				Expect(h.Inventory).Should(Equal("some hw info"))
				Expect(swag.StringValue(h.StatusInfo)).Should(Equal(""))
			}
		})
	})

	Context("refresh_status", func() {
		It("keep_alive", func() {
			host.CheckedInAt = strfmt.DateTime(time.Now().Add(-time.Minute))
			host.Inventory = ""
			mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, false, true, true)
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.expectedState = HostStatusDiscovering
		})
		It("keep_alive_timeout", func() {
			host.CheckedInAt = strfmt.DateTime(time.Now().Add(-time.Hour))
			mockConnectivityAndHwValidators(&host, mockHWValidator, mockConnectivityValidator, false, true, true)
			updateReply, updateErr = state.RefreshStatus(ctx, &host, db)
			expectedReply.expectedState = HostStatusDisconnected
		})
	})

	It("install", func() {
		updateReply, updateErr = state.Install(ctx, &host, nil)
		expectedReply.expectError = true
	})

	It("enable_host", func() {
		updateReply, updateErr = state.EnableHost(ctx, &host)
	})

	It("disable_host", func() {
		updateReply, updateErr = state.DisableHost(ctx, &host)
		expectedReply.expectedState = HostStatusDisabled
	})

	AfterEach(func() {
		ctrl.Finish()
		postValidation(expectedReply, currentState, db, id, clusterId, updateReply, updateErr)
		// cleanup
		db.Close()
		expectedReply = nil
		updateReply = nil
		updateErr = nil
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "disconnected host state tests")
}
