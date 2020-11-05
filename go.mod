module github.com/Walkbase/minio-resources-operator

go 1.13

replace github.com/robotinfra/minio-resources-operator => ./

require (
	github.com/go-ini/ini v1.62.0 // indirect
	github.com/hashicorp/vault/api v1.0.4
	github.com/minio/minio v0.0.0-20201105153408-bd77f29fc4a3
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/operator-framework/operator-sdk v0.19.4
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.6.3
)

replace github.com/operator-framework/operator-sdk => github.com/operator-framework/operator-sdk v0.19.4
