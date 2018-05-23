#  Kubernitized Proof of Stake Blockchain on ICP
<h1 align="center">
  <br>
  <a href="https://github.com/SaifRehman/Kubernitized-Tinder-Blockchain"><img src="https://cdn-images-1.medium.com/max/2000/1*HyemvyVt7JI25k-_cTKMcg.png" alt="Lotion" width="IBM"></a>
  <br>
      Kubernitized Blockchain Tinder app on IBM Cloud Private
  <br>
  <br>
</h1>

<h4 align="center">Powered by Tendermint  ✨ Lotion ✨  Kubernetes and IBM Cloud Private</h4>

<p align="center">
  <a>
    <img src="https://img.shields.io/travis/keppel/lotion/master.svg"
         alt="Travis Build">
  </a>
</p>
<br>

## Pre Requisites
1. Signup to [IBM Cloud](http://ibm.biz/ioblockchain)
2. Install [Cloud Foundry CLI](https://docs.cloudfoundry.org/cf-cli/install-go-cli.html)
3. Install [Bluemix CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started)
4. Install [Kubernetes CLI](https://kubernetes.io/docs/user-guide/prereqs/)
5. Install Minikube [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
6. Install Docker [Docker](https://docs.docker.com/install/)

## Tasks
- [x] Creating user service
- [x] Creating location service
- [x] Creating Blockchain service
- [ ] Creating cron jobs for matching 
- [ ] Creating frontend app
- [ ] Deploying to ICP (IBM Cloud Private)

## Build Docker images
1. Creating user-service
```
$ cd user-service
$ docker build -t user-service .
```
2. Creating location-service
```
$ cd location-service
$ docker build -t location-service .
> Add google map api key in main.go
```
3. Creating blockchain-node1
```
$ cd blockchain-validator1
$ docker build -t blockchain-validator1 .
```
4. Creating blockchain-node2
```
$ cd blockchain-validator2
$ docker build -t blockchain-validator2 .
```

## Deploy to minikube
1. Start Minikube
```
$ minikube start
```
2. Get cluster info
```
$ kubectl cluster-info
```
3. minikube dahsboard
```
$ kubectl cluster-info
```