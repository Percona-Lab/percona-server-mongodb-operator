package stub

import (
	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/psmdb/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	defaultVersion                        = "latest"
	defaultRunUID                   int64 = 1001
	defaultKeySecretName                  = "percona-server-mongodb-key"
	defaultUsersSecretName                = "percona-server-mongodb-users"
	defaultMongodSize               int32 = 3
	defaultReplsetName                    = "rs"
	defaultStorageEngine                  = v1alpha1.StorageEngineWiredTiger
	defaultMongodPort               int32 = 27017
	defaultWiredTigerCacheSizeRatio       = 0.5
	defaultInMemorySizeRatio              = 0.9
	defaultOperationProfilingMode         = v1alpha1.OperationProfilingModeSlowOp
	mongodContainerDataDir                = "/data/db"
	mongodContainerName                   = "mongod"
	mongodDataVolClaimName                = "mongod-data"
	mongodToolsVolName                    = "mongodb-tools"
	mongodPortName                        = "mongodb"
	mongodbInitiatorUrl                   = "https://github.com/percona/mongodb-orchestration-tools/releases/download/0.4.1/k8s-mongodb-initiator"
	mongodbHealthcheckUrl                 = "https://github.com/percona/mongodb-orchestration-tools/releases/download/0.4.1/mongodb-healthcheck"
)

// addPSMDBSpecDefaults sets default values for unset config params
func (h *Handler) addPSMDBSpecDefaults(spec *v1alpha1.PerconaServerMongoDBSpec) {
	if spec.Version == "" {
		spec.Version = defaultVersion
	}
	if spec.Secrets == nil {
		spec.Secrets = &v1alpha1.SecretsSpec{}
	}
	if spec.Secrets.Key == "" {
		spec.Secrets.Key = defaultKeySecretName
	}
	if spec.Secrets.Users == "" {
		spec.Secrets.Users = defaultUsersSecretName
	}
	if spec.Mongod == nil {
		spec.Mongod = &v1alpha1.MongodSpec{}
	}
	if spec.Mongod.Net == nil {
		spec.Mongod.Net = &v1alpha1.MongodSpecNet{}
	}
	if spec.Mongod.Net.Port == 0 {
		spec.Mongod.Net.Port = defaultMongodPort
	}
	if spec.Mongod.Storage == nil {
		spec.Mongod.Storage = &v1alpha1.MongodSpecStorage{}
	}
	if spec.Mongod.Storage.Engine == "" {
		spec.Mongod.Storage.Engine = defaultStorageEngine
	}

	switch spec.Mongod.Storage.Engine {
	case v1alpha1.StorageEngineInMemory:
		if spec.Mongod.Storage.InMemory == nil {
			spec.Mongod.Storage.InMemory = &v1alpha1.MongodSpecInMemory{}
		}
		if spec.Mongod.Storage.InMemory.EngineConfig == nil {
			spec.Mongod.Storage.InMemory.EngineConfig = &v1alpha1.MongodSpecInMemoryEngineConfig{}
		}
		if spec.Mongod.Storage.InMemory.EngineConfig.InMemorySizeRatio == 0 {
			spec.Mongod.Storage.InMemory.EngineConfig.InMemorySizeRatio = defaultInMemorySizeRatio
		}
	case v1alpha1.StorageEngineWiredTiger:
		if spec.Mongod.Storage.WiredTiger == nil {
			spec.Mongod.Storage.WiredTiger = &v1alpha1.MongodSpecWiredTiger{}
		}
		if spec.Mongod.Storage.WiredTiger.CollectionConfig == nil {
			spec.Mongod.Storage.WiredTiger.CollectionConfig = &v1alpha1.MongodSpecWiredTigerCollectionConfig{}
		}
		if spec.Mongod.Storage.WiredTiger.EngineConfig == nil {
			spec.Mongod.Storage.WiredTiger.EngineConfig = &v1alpha1.MongodSpecWiredTigerEngineConfig{}
		}
		if spec.Mongod.Storage.WiredTiger.EngineConfig.CacheSizeRatio == 0 {
			spec.Mongod.Storage.WiredTiger.EngineConfig.CacheSizeRatio = defaultWiredTigerCacheSizeRatio
		}
		if spec.Mongod.Storage.WiredTiger.IndexConfig == nil {
			spec.Mongod.Storage.WiredTiger.IndexConfig = &v1alpha1.MongodSpecWiredTigerIndexConfig{
				PrefixCompression: true,
			}
		}
	}

	if spec.Mongod.OperationProfiling == nil {
		spec.Mongod.OperationProfiling = &v1alpha1.MongodSpecOperationProfiling{
			Mode: defaultOperationProfilingMode,
		}
	}
	if len(spec.Replsets) == 0 {
		spec.Replsets = []*v1alpha1.ReplsetSpec{{
			Name: defaultReplsetName,
			Size: defaultMongodSize,
		}}
	} else {
		for _, replset := range spec.Replsets {
			if replset.Size == 0 {
				replset.Size = defaultMongodSize
			}
		}
	}
	if spec.RunUID == 0 && h.serverVersion.Platform != v1alpha1.PlatformOpenshift {
		spec.RunUID = defaultRunUID
	}
}

// newPSMDBStatefulSet returns a PSMDB stateful set
func (h *Handler) newPSMDBStatefulSet(m *v1alpha1.PerconaServerMongoDB, replset *v1alpha1.ReplsetSpec, clusterRole *v1alpha1.ClusterRole) (*appsv1.StatefulSet, error) {
	h.addPSMDBSpecDefaults(&m.Spec)

	limits, err := parseSpecResourceRequirements(replset.Limits)
	if err != nil {
		return nil, err
	}
	requests, err := parseSpecResourceRequirements(replset.Requests)
	if err != nil {
		return nil, err
	}
	resources := &corev1.ResourceRequirements{
		Limits:   limits,
		Requests: requests,
	}

	ls := labelsForPerconaServerMongoDB(m, replset)
	set := &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name + "-" + replset.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: m.Name + "-" + replset.Name,
			Replicas:    &replset.Size,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Affinity:      newPSMDBPodAffinity(ls),
					RestartPolicy: corev1.RestartPolicyAlways,
					InitContainers: []corev1.Container{
						h.newPSMDBInitContainer(m),
					},
					Containers: []corev1.Container{
						h.newPSMDBMongodContainer(m, replset, clusterRole, resources),
					},
					SecurityContext: &corev1.PodSecurityContext{
						FSGroup: h.getContainerRunUID(m),
					},
					Volumes: []corev1.Volume{
						{
							Name: mongodToolsVolName,
						},
						{
							Name: m.Spec.Secrets.Key,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: m.Spec.Secrets.Key,
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: newPSMDBMongodVolumeClaims(m, resources, mongodDataVolClaimName, replset.StorageClass),
		},
	}
	addOwnerRefToObject(set, asOwner(m))
	return set, nil
}

// newPSMDBService returns a core/v1 API Service
func newPSMDBService(m *v1alpha1.PerconaServerMongoDB, replset *v1alpha1.ReplsetSpec) *corev1.Service {
	ls := labelsForPerconaServerMongoDB(m, replset)
	service := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name + "-" + replset.Name,
			Namespace: m.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       mongodPortName,
					Port:       m.Spec.Mongod.Net.Port,
					TargetPort: intstr.FromInt(int(m.Spec.Mongod.Net.Port)),
				},
			},
			ClusterIP: "None",
			Selector:  ls,
		},
	}
	addOwnerRefToObject(service, asOwner(m))
	return service
}
