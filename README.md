# helmsync
one stop shopping for helm chart build/test/deploy   

## Usage  

```
docker build -t helmsync .   
docker run -it v /path/to/my/helm/repo/:/tmp -e REPOPATH="/tmp" -e VERSIONNUMBER="1.2.1" -e ARTIFACTORYURL="https:///my.artifactory.com/repo" -e IGNOREINFO="TRUE" helmsync  
```

### Current functionality      
parse git diff output to determine which charts changed in a monorepo  
helm lint all updated charts  
deploy all updated charts once linting passes  

### Planned functionality  
deploy all charts to a kind cluster to test a successful install    
run inspec or similar tests against kind cluster post install  


#### todo

todo:


artifactory
- tests
- think about an abstraction layer

E2E testing
- make sure the git key is properly read

docs

- docker command for using key
- deploykey for github deploys
