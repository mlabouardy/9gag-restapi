node{
  def registryUrl= 'http://10.100.248.18:5000'
  def registryCredId = 'nexus'
  
  stage('Checkout'){
    checkout scm
  }

  stage('Build'){
    image = docker.build 'mlabouardy/9gag-api'
  }

  stage('Push'){
    docker.withRegistry(registryUrl, registryCredId){
      image.push 'latest'
    }
  }
}
