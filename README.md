# kvm-remote-go
requirements
1. install kvm 
2. configure /etc/libvirt/libvirtd.conf for ssh connection
3. test the connection virsh qemu+ssh://username@ipaddress/system

golang requirements

1. clone the project
2. go get github.com/libvirt/libvirt-go
3. go get github.com/libvirt/libvirt-go-xml
4. /project folger: go run main.go

##
change the ip address
