
# Table: azure_compute_virtual_machine_linux_configuration_ssh_public_keys
SSHPublicKey contains information about SSH certificate public key and the path on the Linux VM where the public key is placed
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|virtual_machine_id|uuid|Unique ID of azure_compute_virtual_machines table (FK)|
|path|text|Specifies the full path on the created VM where ssh public key is stored If the file already exists, the specified key is appended to the file Example: /home/user/ssh/authorized_keys|
|key_data|text|SSH public key certificate used to authenticate with the VM through ssh The key needs to be at least 2048-bit and in ssh-rsa format <br><br> For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docsmicrosoftcom/en-us/azure/virtual-machines/linux/mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftocjson)|
