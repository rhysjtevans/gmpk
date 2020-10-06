package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/tidwall/gjson"
    "flag"
)

var jsonKey string

func getMetadata(){
    client := &http.Client{}
    req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance", nil)
    req.Header.Add("Metadata", "True")

    q := req.URL.Query()
    q.Add("format", "json")
    q.Add("api-version", "2020-06-01")
    req.URL.RawQuery = q.Encode()

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Errored when sending request to the server")
        fmt.Println("You are probably not running this on an azure virtual machine")
    }
    defer resp.Body.Close()
    
    respBody, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(resp.Status)
    fmt.Println(string(respBody))
    
}

func GetJsonObject(str string, jsonKey string) string {
    if jsonKey != "" {
        return gjson.Get(str, jsonKey).String()
    }else{
        return str
    }
}

func main() {
    flag.StringVar(&jsonKey, "key", "","A decimal notated index key of the json object you want to retrieve")
    flag.Parse()

    json := `{"compute":{"azEnvironment":"AzurePublicCloud","customData":"","isHostCompatibilityLayerVm":"false","location":"uksouth","name":"nfr","offer":"UbuntuServer","osType":"Linux","placementGroupId":"","plan":{"name":"","product":"","publisher":""},"platformFaultDomain":"0","platformUpdateDomain":"0","provider":"Microsoft.Compute","publicKeys":[{"keyData":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCEccaKMLKoqgEEi3ln/EfUiQyE\r\nBoPEP6ByIpVCBlAUlfa6rLzHNunpN5N+n9Uykrx7SA7wfz4yM9EEjlMwD13gjZsb\r\nqTnQxLscmpYh2lJ1nlfWJEpaK5v4g1qZJZdz6FQ7JyXetPAcMUmCmaJkPzQmnCRR\r\n8zvslxpia9w2bH9fXUdVzEhy1bk9GsA7H6EPJiFLZk9n/6+3a8VvYXsAjcuzGNzq\r\nRRwy+sRrSQlVfVNkgDHbs0gR+LF1u6oCVoPf0cMVZ/utt1Tm2EjIv5Ns55l8Gig8\r\nMyOsIlnKes4i3NEt26I2SxQtxtPtu/wYxkQrRkqUhbsSdcWkXqPlVUfCAQgac3IN\r\nTvBm8y8iosHUhjPxskFzOCKOEcu1+Tm+p/6L9sFJ+Je2W00RLSwM3W2cY7MKa7iE\r\n09PzLWRi/vMO6E02kv5LX9CWG0XtLI0iBb6ybmGgTIFEjkYrcuEf97eRzK2omygV\r\nr5vTkvcPlt75m0VPU0T3SDt0bZJCeCOVLo7KbCs= generated-by-azure\r\n","path":"/home/azureuser/.ssh/authorized_keys"}],"publisher":"Canonical","resourceGroupName":"soldat","resourceId":"/subscriptions/d72ddf83-3391-4f89-a2c9-0b69cf10712e/resourceGroups/soldat/providers/Microsoft.Compute/virtualMachines/nfr","securityProfile":{"secureBootEnabled":"false","virtualTpmEnabled":"false"},"sku":"18.04-LTS","storageProfile":{"dataDisks":[{"caching":"ReadWrite","createOption":"Attach","diskSizeGB":"256","image":{"uri":""},"lun":"0","managedDisk":{"id":"/subscriptions/d72ddf83-3391-4f89-a2c9-0b69cf10712e/resourceGroups/soldat/providers/Microsoft.Compute/disks/nfr_DataDisk_0","storageAccountType":"Standard_LRS"},"name":"nfr_DataDisk_0","vhd":{"uri":""},"writeAcceleratorEnabled":"false"}],"imageReference":{"id":"","offer":"UbuntuServer","publisher":"Canonical","sku":"18.04-LTS","version":"latest"},"osDisk":{"caching":"ReadWrite","createOption":"FromImage","diffDiskSettings":{"option":""},"diskSizeGB":"30","encryptionSettings":{"enabled":"false"},"image":{"uri":""},"managedDisk":{"id":"/subscriptions/d72ddf83-3391-4f89-a2c9-0b69cf10712e/resourceGroups/SOLDAT/providers/Microsoft.Compute/disks/nfr_OsDisk_1_fd72a22e09874687adc9f5337a5ca19a","storageAccountType":"StandardSSD_LRS"},"name":"nfr_OsDisk_1_fd72a22e09874687adc9f5337a5ca19a","osType":"Linux","vhd":{"uri":""},"writeAcceleratorEnabled":"false"}},"subscriptionId":"d72ddf83-3391-4f89-a2c9-0b69cf10712e","tags":"","tagsList":[],"version":"18.04.202009010","vmId":"131675be-bda9-40d4-8b46-01e637aeae2b","vmScaleSetName":"","vmSize":"Standard_DS1_v2","zone":""},"network":{"interface":[{"ipv4":{"ipAddress":[{"privateIpAddress":"10.0.0.5","publicIpAddress":"51.105.33.158"}],"subnet":[{"address":"10.0.0.0","prefix":"24"}]},"ipv6":{"ipAddress":[]},"macAddress":"0022483F7A18"}]}}`
    fmt.Println(GetJsonObject(json,jsonKey))
    
}
