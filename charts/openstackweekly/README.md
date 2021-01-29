# openstackweekly charts
Helm chart for openstackweeklys

### Installing the charts
From root directory of openstackweekly. Please edit the values.yaml inside charts before applying.
```
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install zufardhiyaulhaq/openstackweekly --name-template openstackweekly -f values.yaml
```

### Configuration

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| community | string | `"OpenStack Indonesia Community"` |  |
| cronSchedule | string | `"0 8 * * 0"` |  |
| github.branch | string | `"master"` |  |
| github.organization | string | `"zufardhiyaulhaq"` |  |
| github.repository | string | `"community-ops"` |  |
| github.repository_path | string | `"./manifest/openstack-community/"` |  |
| github.token | string | `"your_token"` |  |
| image.name | string | `"openstackweekly"` |  |
| image.repository | string | `"zufardhiyaulhaq/openstackweekly"` |  |
| image.tag | string | `"0.0.1"` |  |
| image_url | string | `"https://object-storage-ca-ymq-1.vexxhost.net/swift/v1/6e4619c416ff4bd19e1c087f27a43eea/www-images-prod/openstack-logo/OpenStack-Logo-Vertical.png"` |  |
| jobHistoryLimit | int | `1` |  |
| namespace | string | `"openstack-community"` |  |
| tags | string | `"weekly,openstack"` |  |

check & modify values.yaml for details
