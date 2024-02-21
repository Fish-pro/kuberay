//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppStatus) DeepCopyInto(out *AppStatus) {
	*out = *in
	if in.HealthLastUpdateTime != nil {
		in, out := &in.HealthLastUpdateTime, &out.HealthLastUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make(map[string]ServeDeploymentStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppStatus.
func (in *AppStatus) DeepCopy() *AppStatus {
	if in == nil {
		return nil
	}
	out := new(AppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoscalerOptions) DeepCopyInto(out *AutoscalerOptions) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.IdleTimeoutSeconds != nil {
		in, out := &in.IdleTimeoutSeconds, &out.IdleTimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.UpscalingMode != nil {
		in, out := &in.UpscalingMode, &out.UpscalingMode
		*out = new(UpscalingMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoscalerOptions.
func (in *AutoscalerOptions) DeepCopy() *AutoscalerOptions {
	if in == nil {
		return nil
	}
	out := new(AutoscalerOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadGroupSpec) DeepCopyInto(out *HeadGroupSpec) {
	*out = *in
	if in.HeadService != nil {
		in, out := &in.HeadService, &out.HeadService
		*out = new(corev1.Service)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableIngress != nil {
		in, out := &in.EnableIngress, &out.EnableIngress
		*out = new(bool)
		**out = **in
	}
	if in.RayStartParams != nil {
		in, out := &in.RayStartParams, &out.RayStartParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadGroupSpec.
func (in *HeadGroupSpec) DeepCopy() *HeadGroupSpec {
	if in == nil {
		return nil
	}
	out := new(HeadGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadInfo) DeepCopyInto(out *HeadInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadInfo.
func (in *HeadInfo) DeepCopy() *HeadInfo {
	if in == nil {
		return nil
	}
	out := new(HeadInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayCluster) DeepCopyInto(out *RayCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayCluster.
func (in *RayCluster) DeepCopy() *RayCluster {
	if in == nil {
		return nil
	}
	out := new(RayCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterList) DeepCopyInto(out *RayClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RayCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterList.
func (in *RayClusterList) DeepCopy() *RayClusterList {
	if in == nil {
		return nil
	}
	out := new(RayClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterSpec) DeepCopyInto(out *RayClusterSpec) {
	*out = *in
	in.HeadGroupSpec.DeepCopyInto(&out.HeadGroupSpec)
	if in.WorkerGroupSpecs != nil {
		in, out := &in.WorkerGroupSpecs, &out.WorkerGroupSpecs
		*out = make([]WorkerGroupSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableInTreeAutoscaling != nil {
		in, out := &in.EnableInTreeAutoscaling, &out.EnableInTreeAutoscaling
		*out = new(bool)
		**out = **in
	}
	if in.AutoscalerOptions != nil {
		in, out := &in.AutoscalerOptions, &out.AutoscalerOptions
		*out = new(AutoscalerOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.HeadServiceAnnotations != nil {
		in, out := &in.HeadServiceAnnotations, &out.HeadServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterSpec.
func (in *RayClusterSpec) DeepCopy() *RayClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RayClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterStatus) DeepCopyInto(out *RayClusterStatus) {
	*out = *in
	out.DesiredCPU = in.DesiredCPU.DeepCopy()
	out.DesiredMemory = in.DesiredMemory.DeepCopy()
	out.DesiredGPU = in.DesiredGPU.DeepCopy()
	out.DesiredTPU = in.DesiredTPU.DeepCopy()
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Head = in.Head
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterStatus.
func (in *RayClusterStatus) DeepCopy() *RayClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RayClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayJob) DeepCopyInto(out *RayJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayJob.
func (in *RayJob) DeepCopy() *RayJob {
	if in == nil {
		return nil
	}
	out := new(RayJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayJobList) DeepCopyInto(out *RayJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RayJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayJobList.
func (in *RayJobList) DeepCopy() *RayJobList {
	if in == nil {
		return nil
	}
	out := new(RayJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayJobSpec) DeepCopyInto(out *RayJobSpec) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
	if in.RayClusterSpec != nil {
		in, out := &in.RayClusterSpec, &out.RayClusterSpec
		*out = new(RayClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubmitterPodTemplate != nil {
		in, out := &in.SubmitterPodTemplate, &out.SubmitterPodTemplate
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayJobSpec.
func (in *RayJobSpec) DeepCopy() *RayJobSpec {
	if in == nil {
		return nil
	}
	out := new(RayJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayJobStatus) DeepCopyInto(out *RayJobStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	in.RayClusterStatus.DeepCopyInto(&out.RayClusterStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayJobStatus.
func (in *RayJobStatus) DeepCopy() *RayJobStatus {
	if in == nil {
		return nil
	}
	out := new(RayJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayService) DeepCopyInto(out *RayService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayService.
func (in *RayService) DeepCopy() *RayService {
	if in == nil {
		return nil
	}
	out := new(RayService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceList) DeepCopyInto(out *RayServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RayService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceList.
func (in *RayServiceList) DeepCopy() *RayServiceList {
	if in == nil {
		return nil
	}
	out := new(RayServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceSpec) DeepCopyInto(out *RayServiceSpec) {
	*out = *in
	in.RayClusterSpec.DeepCopyInto(&out.RayClusterSpec)
	if in.ServiceUnhealthySecondThreshold != nil {
		in, out := &in.ServiceUnhealthySecondThreshold, &out.ServiceUnhealthySecondThreshold
		*out = new(int32)
		**out = **in
	}
	if in.DeploymentUnhealthySecondThreshold != nil {
		in, out := &in.DeploymentUnhealthySecondThreshold, &out.DeploymentUnhealthySecondThreshold
		*out = new(int32)
		**out = **in
	}
	if in.ServeService != nil {
		in, out := &in.ServeService, &out.ServeService
		*out = new(corev1.Service)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceSpec.
func (in *RayServiceSpec) DeepCopy() *RayServiceSpec {
	if in == nil {
		return nil
	}
	out := new(RayServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceStatus) DeepCopyInto(out *RayServiceStatus) {
	*out = *in
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make(map[string]AppStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.RayClusterStatus.DeepCopyInto(&out.RayClusterStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceStatus.
func (in *RayServiceStatus) DeepCopy() *RayServiceStatus {
	if in == nil {
		return nil
	}
	out := new(RayServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayServiceStatuses) DeepCopyInto(out *RayServiceStatuses) {
	*out = *in
	in.ActiveServiceStatus.DeepCopyInto(&out.ActiveServiceStatus)
	in.PendingServiceStatus.DeepCopyInto(&out.PendingServiceStatus)
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayServiceStatuses.
func (in *RayServiceStatuses) DeepCopy() *RayServiceStatuses {
	if in == nil {
		return nil
	}
	out := new(RayServiceStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleStrategy) DeepCopyInto(out *ScaleStrategy) {
	*out = *in
	if in.WorkersToDelete != nil {
		in, out := &in.WorkersToDelete, &out.WorkersToDelete
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleStrategy.
func (in *ScaleStrategy) DeepCopy() *ScaleStrategy {
	if in == nil {
		return nil
	}
	out := new(ScaleStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServeDeploymentStatus) DeepCopyInto(out *ServeDeploymentStatus) {
	*out = *in
	if in.HealthLastUpdateTime != nil {
		in, out := &in.HealthLastUpdateTime, &out.HealthLastUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServeDeploymentStatus.
func (in *ServeDeploymentStatus) DeepCopy() *ServeDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ServeDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerGroupSpec) DeepCopyInto(out *WorkerGroupSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
	if in.RayStartParams != nil {
		in, out := &in.RayStartParams, &out.RayStartParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	in.ScaleStrategy.DeepCopyInto(&out.ScaleStrategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerGroupSpec.
func (in *WorkerGroupSpec) DeepCopy() *WorkerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerGroupSpec)
	in.DeepCopyInto(out)
	return out
}
