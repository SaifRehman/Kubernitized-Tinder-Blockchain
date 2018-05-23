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


## Contributors
| [<img src="https://media.licdn.com/dms/image/C5603AQGqNuagFBJDQQ/profile-displayphoto-shrink_200_200/0?e=1532563200&v=beta&t=CCItySx0oYEV7tzTR6DiQ4ybh6dmwToQQy2_TCxU5rc" width="100px;"/><br /><sub><b>SaifRehman</b></sub>](https://www.linkedin.com/in/saif-ur-rehman/)<br />[💻]("Code") [📖]( "Documentation") [🤔](#ideas-keppel "Ideas, Planning, & Feedback") [⚠️]("Tests") | [<img src="https://media.licdn.com/dms/image/C5103AQHfEw0aX0CPJw/profile-displayphoto-shrink_800_800/0?e=1532563200&v=beta&t=SWLTO9FTwCBfdjd8pewRDJd8_EsT7nSv-ZY3rxC_oM8" width="100px;"/><br /><sub><b>Areej Essa</b></sub>](https://www.linkedin.com/in/areej-essa-b19336114/)<br />[💻]("Code") [🤔](#ideas-mappum "Ideas, Planning, & Feedback") [⚠️]("Tests") [🔌](#plugin-mappum "Plugin/utility libraries")
| :---: | :---: |

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
- [x] Creating Blockchain services
- [ ] Connecting to real database
- [ ] Creating cron jobs for matching 
- [ ] Creating frontend app
- [ ] Deploying to ICP (IBM Cloud Private)

## Reading images locally
```
$ eval $(minikube docker-env)
```
## Build and run Docker images
1. Creating user-service
```
$ cd user-service
$ docker build -t user-service .
$ docker run -i -t -p 30094:30094 user-service:latest 
```
2. Creating location-service
```
$ cd location-service
$ docker build -t location-service .
$ docker run -i -t -p 30095:30095 location-service:latest 
> Add google map api key in main.go
```
3. Creating blockchain-node1
```
$ cd blockchain-validator1
$ docker build -t blockchain-validator1 .
$ docker run -i -t -p 30090:30090 -p 30091:30091 blockchain-validator1:latest 
```
4. Creating blockchain-node2
```
$ cd blockchain-validator2
$ docker build -t blockchain-validator2 .
$ docker run -i -t -p 30092:30092 -p 30093:30093 blockchain-validator2:latest 
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
$ kubectl dahsboard
```
4. running service-deployment.yml file
```
$ kubectl create -f service-deployment.yml
```
5. get minikube ip
```
$ minikube ip
```

## Services
1. Blockchain node1
- [x] Tendermint RPC port 30090
- [x] Tendermint p2p port 30091
2. Blockchain node2
- [x] Tendermint RPC port 30092
- [x] Tendermint p2p port 30093
3. User Service
- [x] port 30094
4. Location Service
- [x] port 30095

## Apis
### User service 
* (GET) :30094/all
* (GET) :30094/user/{id}
* (POST) :30094/create
```JavaScript
   { "id": "1",
    "name": "Matt",
    "gender": "Male",
    "age": 23,
    "userlocation": {
        "Lat": 54.234,
        "Long": 55.234
    }
   }
```
* (DELETE) :30094/delete/{id}

### Location service
* (GET) :30095/location

### Blockchain Node1 
*  (GET) :30090/ # Tendermint RPC Port

### Blockchain Node2
* (GET) :30092/ # Tendermint RPC Port

## Authors
- [Saif ur Rehman](https://www.linkedin.com/in/saif-ur-rehman/)
- [Areej Essa](https://www.linkedin.com/in/areej-essa-b19336114/)
