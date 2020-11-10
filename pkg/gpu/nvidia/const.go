package nvidia

import (
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

// MemoryUnit describes GPU Memory, now only supports Gi, Mi
type MemoryUnit string

const (
	resourceName  = "ucloud.cn/gpu-mem"
	resourceCount = "ucloud.cn/gpu-count"
	serverSock    = pluginapi.DevicePluginPath + "ucloudshare.sock"

	OptimisticLockErrorMsg = "the object has been modified; please apply your changes to the latest version and try again"

	envNVGPU                   = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex           = "UCLOUD_CN_GPU_MEM_IDX"
	EnvResourceByPod           = "UCLOUD_CN_GPU_MEM_POD"
	EnvResourceByContainer     = "UCLOUD_CN_GPU_MEM_CONTAINER"
	EnvResourceByDev           = "UCLOUD_CN_GPU_MEM_DEV"
	EnvAssignedFlag            = "UCLOUD_CN_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime      = "UCLOUD_CN_GPU_MEM_ASSUME_TIME"
	EnvResourceAssignTime      = "UCLOUD_CN_GPU_MEM_ASSIGN_TIME"
	EnvNodeLabelForDisableCGPU = "cgpu.disable.isolation"

	GiBPrefix = MemoryUnit("GiB")
	MiBPrefix = MemoryUnit("MiB")
)
