	
**BOSH SoftLayer Pool Server Release**
-------------

BOSH SoftLayer pool server provides APIs to utilize virtual guests pooling on SoftLayer. This is a BOSH release for it.

**Releases and stemcells**
-------------
Except this release,  following releases are also required.

- [postgres](http://bosh.io/releases/github.com/cloudfoundry/postgres-release) 
- [bosh-softlayer-cpi](http://bosh.io/releases/github.com/cloudfoundry-incubator/bosh-softlayer-cpi-release) 

SoftLayer light stemcell is needed for deployment and can be found in [bosh.io](http://bosh.io/)

**Bootstrap on SoftLayer**
-------------
You can use bosh-init from BOSH community to bootstrap a pool server on SoftLayer.

- To install bosh-init, please refer to [install-bosh-init](http://bosh.io/docs/install-bosh-init.html) and its usage can be found in [using-bosh-init](http://bosh.io/docs/using-bosh-init.html).
> **Note:**
>  Please make sure the machine installed bosh-init can access SoftLayer private network and you can enable SoftLayer VPN if it is outside of SoftLayer data center. This is because it need to communicate with the target VM over SoftLayer private network to accomplish a successful deployment.

- Prepare a deployment manifest

You can find a deployment manifest example under docs named `vps-init-example.yml` which can deploy a virtual guest pooling server and please replace release, stemcell, resource and credential information accordingly.
> **Note:**
>  For releases and stemcells, please either use url like the example manifest does or download them to your local machine and specify its location.
>  
>  - bosh-softlayer-pool-server-release
>  - postgres-release
>  - bosh-softlayer-cpi-release
>  - SoftLayer light stemcell

Here is an example for key properties of jobs.
```
jobs:
- name: vps
  instances: 1

  templates:
  - {name: postgres, release: postgres}
  - {name: vps, release: bosh-softlayer-pool-server}

  resource_pool: vms

  networks:
  - name: default

  properties:
    databases: &20585760
      roles:
      - name: postgres
      password: postgres
      address: 127.0.0.1
      port: 5432
      databases:
      - name: bosh
    vps:
      host: 0.0.0.0
      port: 8889
      log_level: debug
      sql:
        db_username: postgres
        db_password: postgres
        db_host: 127.0.0.1
        db_port: 5432
        db_schema: bosh
        db_driver: postgres

```
- Kick-off deployment

```
bosh-init deploy <your-manifest.yml>
```

- After deployment completes, you can login the environment and take a check. Run `monit summary`, normally if everything works well you can get an output as below.
```
~# monit summary
The Monit daemon 5.2.5 uptime: 33m

Process 'postgres'                  running
Process 'vps'                       running
System 'system_localhost'           running
```
> **Note:**
> If any job is not running, run `monit restart` <job-name> to restart it. If this doesn't work out, you can check logs under /var/vcap/sys/log and do further investigation.

- To fully enable virtual guest pooling on SoftLayer, except deploying a pool server, you also need to make director to connect to it and enable pooling feature. Please refer to guide on [bosh-softlayer-cpi-release](https://github.com/cloudfoundry-incubator/bosh-softlayer-cpi-release).