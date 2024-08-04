package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webserver/src/app/domain"
	"webserver/src/app/port/device"
	common "webserver/src/commom"
)

type devicesHandler struct {
	deviceUseCase device.UseCase
}

func NewDevicesHandler(deviceUseCase device.UseCase) *devicesHandler {
	return &devicesHandler{
		deviceUseCase: deviceUseCase,
	}
}

func (h *devicesHandler) InitializeDB(ctx *gin.Context) {
	err := h.deviceUseCase.InitializeRepository()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "DB Initialized"})

}

func (h *devicesHandler) GetDeviceByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deviceItem, err := h.deviceUseCase.GetDeviceByID(ctx, ID)
	if err != nil {
		if err.Error() == common.ErrDeviceNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Device with ID %d", ID), "data": deviceItem})
}

func (h *devicesHandler) ListDevices(ctx *gin.Context) {
	devices, err := h.deviceUseCase.ListDevices(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Devices %d", len(devices)), "data": devices})
}

func (h *devicesHandler) CreateDevice(ctx *gin.Context) {
	var deviceToCreate domain.CreateDevice
	err := ctx.ShouldBind(&deviceToCreate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedDevice := parseDevice(deviceToCreate)
	err = h.deviceUseCase.AddDevice(ctx, &parsedDevice)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Device created"})

}

func (h *devicesHandler) UpdateDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deviceToUpdate := domain.Device{}
	err = ctx.ShouldBind(&deviceToUpdate)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	deviceToUpdate.ID = ID
	err = h.deviceUseCase.UpdateDevice(ctx, ID, &deviceToUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Device updated"})
}

func (h *devicesHandler) DeleteDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.deviceUseCase.DeleteDevice(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Device deleted"})

}

func (h *devicesHandler) SearchByBrand(ctx *gin.Context) {
	brand := ctx.Query("brand")
	devices, err := h.deviceUseCase.SearchDevicesByBrand(ctx, brand)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Devices %d", len(devices)), "data": devices})
}

func parseDevice(deviceToCreate domain.CreateDevice) domain.Device {
	return domain.Device{
		Brand: deviceToCreate.Brand,
		Name:  deviceToCreate.Name,
	}
}
