package psmdb

import (
	"fmt"
	"strconv"
	"strings"

	api "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func MongosDeployment(cr *api.PerconaServerMongoDB) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-" + "mongos",
			Namespace: cr.Namespace,
		},
	}
}

func MongosDeploymentSpec(cr *api.PerconaServerMongoDB, operatorPod corev1.Pod) (appsv1.DeploymentSpec, error) {
	ls := map[string]string{
		"app.kubernetes.io/component": "mongos",
	}

	c, err := mongosContainer(cr)
	if err != nil {
		return appsv1.DeploymentSpec{}, fmt.Errorf("failed to create container %v", err)
	}

	initContainers := initContainers(cr, operatorPod)
	for i := range initContainers {
		initContainers[i].Resources.Limits = c.Resources.Limits
		initContainers[i].Resources.Requests = c.Resources.Requests
	}

	return appsv1.DeploymentSpec{
		Replicas: &cr.Spec.Mongos.Size,
		Selector: &metav1.LabelSelector{
			MatchLabels: ls,
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels:      ls,
				Annotations: cr.Spec.Mongos.MultiAZ.Annotations,
			},
			Spec: corev1.PodSpec{
				SecurityContext:   cr.Spec.Mongos.PodSecurityContext,
				Affinity:          PodAffinity(cr.Spec.Mongos.MultiAZ.Affinity, ls),
				NodeSelector:      cr.Spec.Mongos.MultiAZ.NodeSelector,
				Tolerations:       cr.Spec.Mongos.MultiAZ.Tolerations,
				PriorityClassName: cr.Spec.Mongos.MultiAZ.PriorityClassName,
				RestartPolicy:     corev1.RestartPolicyAlways,
				ImagePullSecrets:  cr.Spec.ImagePullSecrets,
				Containers:        []corev1.Container{c},
				InitContainers:    initContainers,
				Volumes:           volumes(cr),
				SchedulerName:     cr.Spec.SchedulerName,
			},
		},
	}, nil
}

func initContainers(cr *api.PerconaServerMongoDB, operatorPod corev1.Pod) []corev1.Container {
	inits := []corev1.Container{}
	if cr.CompareVersion("1.5.0") >= 0 {
		inits = append(inits, EntrypointInitContainer(operatorPod.Spec.Containers[0].Image))
	}

	return inits
}

func mongosContainer(cr *api.PerconaServerMongoDB) (corev1.Container, error) {
	fvar := false

	resources, err := CreateResources(cr.Spec.Mongos.ResourcesSpec)
	if err != nil {
		return corev1.Container{}, fmt.Errorf("resource creation: %v", err)
	}

	volumes := []corev1.VolumeMount{
		{
			Name:      MongodDataVolClaimName,
			MountPath: MongodContainerDataDir,
		},
		{
			Name:      InternalKey(cr),
			MountPath: mongodSecretsDir,
			ReadOnly:  true,
		},
		{
			Name:      "ssl",
			MountPath: sslDir,
			ReadOnly:  true,
		},
		{
			Name:      "ssl-internal",
			MountPath: sslInternalDir,
			ReadOnly:  true,
		},
	}

	if *cr.Spec.Mongod.Security.EnableEncryption {
		volumes = append(volumes,
			corev1.VolumeMount{
				Name:      cr.Spec.Mongod.Security.EncryptionKeySecret,
				MountPath: mongodRESTencryptDir,
				ReadOnly:  true,
			},
		)
	}

	mongosArgs, err := mongosContainerArgs(cr, resources)
	if err != nil {
		return corev1.Container{}, err
	}
	container := corev1.Container{
		Name:            "mongos",
		Image:           cr.Spec.Image,
		ImagePullPolicy: cr.Spec.ImagePullPolicy,
		Args:            mongosArgs,
		Ports: []corev1.ContainerPort{
			{
				Name:          mongosPortName,
				HostPort:      cr.Spec.Mongos.HostPort,
				ContainerPort: cr.Spec.Mongos.Port,
			},
		},
		Env: []corev1.EnvVar{
			{
				Name:  "SERVICE_NAME",
				Value: cr.Name,
			},
			{
				Name:  "NAMESPACE",
				Value: cr.Namespace,
			},
			{
				Name:  "MONGODB_PORT",
				Value: strconv.Itoa(int(cr.Spec.Mongos.Port)),
			},
		},
		EnvFrom: []corev1.EnvFromSource{
			{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: cr.Spec.Secrets.Users,
					},
					Optional: &fvar,
				},
			},
		},
		WorkingDir:      MongodContainerDataDir,
		LivenessProbe:   &cr.Spec.Mongos.LivenessProbe.Probe,
		ReadinessProbe:  cr.Spec.Mongos.ReadinessProbe,
		SecurityContext: cr.Spec.Mongos.ContainerSecurityContext,
		Resources:       resources,
		VolumeMounts:    volumes,
	}

	if cr.CompareVersion("1.5.0") >= 0 {
		container.EnvFrom = []corev1.EnvFromSource{
			{
				SecretRef: &corev1.SecretEnvSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "internal-" + cr.Name + "-users",
					},
					Optional: &fvar,
				},
			},
		}
		container.Command = []string{"/data/db/ps-entry.sh"}
	}

	return container, nil
}

func findCfgReplset(replsets []*api.ReplsetSpec) (*api.ReplsetSpec, error) {
	for _, rs := range replsets {
		if rs.ClusterRole == "configsvr" {
			return rs, nil
		}
	}

	return nil, errors.New("failed to find config server replset configuration")
}

func mongosContainerArgs(m *api.PerconaServerMongoDB, resources corev1.ResourceRequirements) ([]string, error) {
	mdSpec := m.Spec.Mongod
	msSpec := m.Spec.Mongos
	cfgRs, err := findCfgReplset(m.Spec.Replsets)
	if err != nil {
		return nil, err
	}

	cfgInstanses := make([]string, 0, cfgRs.Size)
	for i := 0; i < int(cfgRs.Size); i++ {
		cfgInstanses = append(cfgInstanses,
			fmt.Sprintf("%s-%s-%d.%s-%s.%s.svc.cluster.local:%d",
				m.Name, cfgRs.Name, i, m.Name, cfgRs.Name, m.Namespace, msSpec.Port))
	}

	configDB := fmt.Sprintf("cfg0/%s", strings.Join(cfgInstanses, ","))
	args := []string{
		"mongos",
		"--bind_ip_all",
		"--port=" + strconv.Itoa(int(msSpec.Port)),
		"--sslAllowInvalidCertificates",
		"--configdb",
		configDB,
	}

	if m.Spec.UnsafeConf {
		args = append(args,
			"--clusterAuthMode=keyFile",
			"--keyFile="+mongodSecretsDir+"/mongodb-key",
		)
	} else {
		args = append(args,
			"--sslMode=preferSSL",
			"--clusterAuthMode=x509",
		)
	}

	if mdSpec.Security != nil && mdSpec.Security.RedactClientLogData {
		args = append(args, "--redactClientLogData")
	}

	if msSpec.SetParameter != nil {
		if msSpec.SetParameter.CursorTimeoutMillis > 0 {
			args = append(args,
				"--setParameter",
				"cursorTimeoutMillis="+strconv.Itoa(msSpec.SetParameter.CursorTimeoutMillis),
			)
		}
	}

	if msSpec.AuditLog != nil && msSpec.AuditLog.Destination == api.AuditLogDestinationFile {
		if msSpec.AuditLog.Filter == "" {
			msSpec.AuditLog.Filter = "{}"
		}
		args = append(args,
			"--auditDestination=file",
			"--auditFilter="+msSpec.AuditLog.Filter,
			"--auditFormat="+string(msSpec.AuditLog.Format),
		)
		switch msSpec.AuditLog.Format {
		case api.AuditLogFormatBSON:
			args = append(args, "--auditPath="+MongodContainerDataDir+"/auditLog.bson")
		default:
			args = append(args, "--auditPath="+MongodContainerDataDir+"/auditLog.json")
		}
	}

	return args, nil
}

func volumes(cr *api.PerconaServerMongoDB) []corev1.Volume {
	fvar := false
	t := true
	volumes := []corev1.Volume{
		{
			Name: InternalKey(cr),
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					DefaultMode: &secretFileMode,
					SecretName:  InternalKey(cr),
					Optional:    &fvar,
				},
			},
		},
		{
			Name: "ssl",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName:  cr.Spec.Secrets.SSL,
					Optional:    &cr.Spec.UnsafeConf,
					DefaultMode: &secretFileMode,
				},
			},
		},
		{
			Name: "ssl-internal",
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName:  cr.Spec.Secrets.SSLInternal,
					Optional:    &t,
					DefaultMode: &secretFileMode,
				},
			},
		},
		{
			Name: MongodDataVolClaimName,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},
	}

	if *cr.Spec.Mongod.Security.EnableEncryption {
		volumes = append(volumes,
			corev1.Volume{
				Name: cr.Spec.Mongod.Security.EncryptionKeySecret,
				VolumeSource: corev1.VolumeSource{
					Secret: &corev1.SecretVolumeSource{
						DefaultMode: &secretFileMode,
						SecretName:  cr.Spec.Mongod.Security.EncryptionKeySecret,
						Optional:    &fvar,
					},
				},
			},
		)
	}

	return volumes
}

func MongosService(m *api.PerconaServerMongoDB) *corev1.Service {
	ls := map[string]string{
		"app.kubernetes.io/component": "mongos",
	}

	svc := corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        m.Name + "-" + "mongos",
			Namespace:   m.Namespace,
			Annotations: m.Spec.Mongos.ServiceAnnotations,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       mongosPortName,
					Port:       m.Spec.Mongos.Port,
					TargetPort: intstr.FromInt(int(m.Spec.Mongos.Port)),
				},
			},
			Selector:                 ls,
			LoadBalancerSourceRanges: m.Spec.Mongos.LoadBalancerSourceRanges,
		},
	}

	if !m.Spec.Mongos.Expose.Enabled {
		svc.Spec.ClusterIP = "None"
	} else {
		switch m.Spec.Mongos.Expose.ExposeType {
		case corev1.ServiceTypeNodePort:
			svc.Spec.Type = corev1.ServiceTypeNodePort
			svc.Spec.ExternalTrafficPolicy = "Local"
		case corev1.ServiceTypeLoadBalancer:
			svc.Spec.Type = corev1.ServiceTypeLoadBalancer
			svc.Spec.ExternalTrafficPolicy = "Cluster"
		default:
			svc.Spec.Type = corev1.ServiceTypeClusterIP
		}
	}

	return &svc
}
