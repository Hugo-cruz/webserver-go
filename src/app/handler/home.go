package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"webserver/src/app/domain"
	"webserver/src/app/port"
)

type devicesHandler struct {
	deviceUseCase port.UseCase
}

func NewDevicesHandler(deviceUseCase port.UseCase) *devicesHandler {
	return &devicesHandler{
		deviceUseCase: deviceUseCase,
	}
}

func HomeHandler(context *gin.Context) {
	fmt.Println("aaaaa")
}

func (h *devicesHandler) InitializeDB(ctx *gin.Context) {
	err := h.deviceUseCase.InitializeRepository()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "DB Initialized"})
	}
}

func (h *devicesHandler) GetDeviceByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	device, err := h.deviceUseCase.GetDeviceByID(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Device with ID %d", ID),
			"data": device})
	}
}

func (h *devicesHandler) ListDevices(ctx *gin.Context) {
	devices, err := h.deviceUseCase.ListDevices(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Devices %d", len(devices)),
			"data": devices})
	}
}

func (h *devicesHandler) CreateDevice(ctx *gin.Context) {
	device := domain.Device{}
	err := ctx.ShouldBind(&device)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = h.deviceUseCase.AddDevice(ctx, device)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Device created"})
	}
}

func (h *devicesHandler) UpdateDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	device := domain.Device{}
	err = ctx.ShouldBind(&device)
	device.ID = ID
	log.Println(device)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = h.deviceUseCase.UpdateDevice(ctx, ID, device)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Device updated"})
	}

}

func (h *devicesHandler) DeleteDevice(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = h.deviceUseCase.DeleteDevice(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Device deleted"})
	}
}

func (h *devicesHandler) SearchByBrand(ctx *gin.Context) {
	brand := ctx.Query("brand")
	devices, err := h.deviceUseCase.SearchDevicesByBrand(ctx, brand)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Devices %d", len(devices)),
		"data": devices,
	})
}
