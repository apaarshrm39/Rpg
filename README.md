rgp

apaar.dev = group
v1alpha1  = version

# Steps

1) Install the code Generator binary 

```bash
go get k8s.io/code-generator
```

2) Save the path of the binary in a variable

```bash

execdir=~/go/pkg/mod/k8s.io/code-generator@v0.21.10

```

3) Run the code Generator command

```bash

"${execdir}"/generate-groups.sh all github.com/apaarshrm39/rgp/pkg/client github.com/apaarshrm39/rgp/pkg/apis apaar.dev:v1alpha1 --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.."  --go-header-file "${execdir}"/hack/boilerplate.go.txt

```

4) Install the controller gen binary

```bash 

go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1

```

5) Make manifest files for crd

binary=~/go/bin

$binary/controller-gen paths=github.com/apaarshrm39/rgp/pkg/apis/apaar.dev/v1alpha1 crd:trivialVersions=true output:crd:artifacts:config=manifests

```