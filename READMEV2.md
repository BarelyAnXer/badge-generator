| app         | port |
| ----------- | ---- |
| frontend    | 4173 |
| backend     | 8080 |
| kubebuilder | none |

 Note: swap out the values of github username to yours
### Frontend
1. docker build -t ghcr.io/barelyanxer/badge-generator-frontend:latest . --no-cache
2. docker tag badge-generator-frontend ghcr.io/barelyanxer/badge-generator-frontend (review)
3. docker push ghcr.io/barelyanxer/badge-generator-frontend
### Backend
1. docker build -t ghcr.io/barelyanxer/badge-generator-backend:latest .
2. docker push ghcr.io/barelyanxer/badge-generator-backend
### Kubebuilder
1. make build-installer IMG=ghcr.io/barelyanxer/my-kubebuilder-project:latest (for generating the install)
2. make docker-build docker-push IMG=ghcr.io/barelyanxer/my-kubebuilder-project:latest (for creating kubebuilder image and pushig it to the repo)
3. kind load docker-image ghcr.io/barelyanxer/my-kubebuilder-project:latest --name <your-kind-cluster-name> (for loading the docker image and putting it in cluster)
4. make deploy IMG=ghcr.io/barelyanxer/my-kubebuilder-project:latest (can only run in the source code but its to make the crd usable in cluster its an alternative to make deploy)
5. kubectl apply -f https://raw.githubusercontent.com/BarelyAnXer/badge-generator/main/dist/install.yaml (for installing / applying "deploy" controller in another way)

create a secret to pull the image
kubectl create secret docker-registry ghcr-secret \
  --docker-server=ghcr.io \
  --docker-username=your-username \
  --docker-password=your-secret \
  --docker-email=your-email \
  --namespace my-kubebuilder-project-system

create service for the frontend and backend to expose them

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
  namespace: my-kubebuilder-project-system
spec:
  selector:
    run: badge-frontend
  ports:
    - protocol: TCP
      port: 4173
      targetPort: 4173
      nodePort: 32001
  type: NodePort
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-service
  namespace: my-kubebuilder-project-system
spec:
  selector:
    run: badge-backend
  ports:
  - protocol: TCP
    port: 8080        
    targetPort: 8080  
  type: ClusterIP
```

and then create a pod for the frontend and backend follow similar template for backend pod 

1. Generate the YAML File for the Pod
```bash
k run badge-frontend --image=ghcr.io/barelyanxer/badge-generator-frontend:latest --dry-run=client -o yaml > badge-frontend.yaml
```

 2. Edit the YAML to Add the Secret

```yaml
spec:
  containers:
  - name: badge-frontend
    image: ghcr.io/barelyanxer/badge-generator-backend:latest
  imagePullSecrets:
  - name: ghcr-secret
```
3.  Apply the YAML

```bash
kubectl apply -f badge-frontend.yaml
```

then do the same thing with backend