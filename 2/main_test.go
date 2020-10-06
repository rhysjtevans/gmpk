package main

import "testing"

func TestGetJsonObject(t *testing.T) {
    json := `{"compute":{"azEnvironment":"AzurePublicCloud","customData":"","isHostCompatibilityLayerVm":"false","location":"uksouth","name":"nfrx","offer":"UbuntuServer","osType":"Linux","placementGroupId":"","plan":{"name":"","product":"","publisher":""},"platformFaultDomain":"0","platformUpdateDomain":"0","provider":"Microsoft.Compute","publicKeys":[{"keyData":"","path":"/home/azureuser/.ssh/authorized_keys"}],"publisher":"Canonical","resourceGroupName":"soldat","resourceId":"/subscriptions/___/resourceGroups/soldat/providers/Microsoft.Compute/virtualMachines/nfr","securityProfile":{"secureBootEnabled":"false","virtualTpmEnabled":"false"},"sku":"18.04-LTS","storageProfile":{"dataDisks":[{"caching":"ReadWrite","createOption":"Attach","diskSizeGB":"256","image":{"uri":""},"lun":"0","managedDisk":{"id":"/subscriptions/___/resourceGroups/soldat/providers/Microsoft.Compute/disks/nfr_DataDisk_0","storageAccountType":"Standard_LRS"},"name":"nfr_DataDisk_0","vhd":{"uri":""},"writeAcceleratorEnabled":"false"}],"imageReference":{"id":"","offer":"UbuntuServer","publisher":"Canonical","sku":"18.04-LTS","version":"latest"},"osDisk":{"caching":"ReadWrite","createOption":"FromImage","diffDiskSettings":{"option":""},"diskSizeGB":"30","encryptionSettings":{"enabled":"false"},"image":{"uri":""},"managedDisk":{"id":"/subscriptions/_____/resourceGroups/SOLDAT/providers/Microsoft.Compute/disks/","storageAccountType":"StandardSSD_LRS"},"name":"","osType":"Linux","vhd":{"uri":""},"writeAcceleratorEnabled":"false"}},"subscriptionId":"","tags":"","tagsList":[],"version":"18.04.202009010","vmId":"","vmScaleSetName":"","vmSize":"Standard_DS1_v2","zone":""},"network":{"interface":[{"ipv4":{"ipAddress":[{"privateIpAddress":"10.0.0.5","publicIpAddress":"x.x.x.x"}],"subnet":[{"address":"10.0.0.0","prefix":"24"}]},"ipv6":{"ipAddress":[]},"macAddress":""}]}}`
    result := GetJsonObject(json,"compute.name") 
    
    if result != "nfrx" {
        t.Error("Expected Compute.name==nfrx but got", result)
    }
}
