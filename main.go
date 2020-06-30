package main
import (
	"fmt"
	"log"
	"crypto/rand"
	libvirt "github.com/libvirt/libvirt-go"
	//"github.com/libvirt/libvirt-go-xml"
  
)

func getConnection(ipAddress string,user string)(*libvirt.Connect){
	conn, err := libvirt.NewConnect("qemu+ssh://"+user+"@"+ipAddress+"/system")
	if err!=nil{
		log.Fatal(err)
	}
	return conn
}
func buildTestDomain(imageName string,conn libvirt.Connect) (*libvirt.Domain) {
	// xmlFile, err := os.Open("ubuntu.xml")
    // // if we os.Open returns an error then handle it
    // if err != nil {
    //     fmt.Println(err)
	// }
	// fmt.Println(xmlFile)
	b := make([]byte, 16)
   _, err := rand.Read(b)
   if err != nil {
    log.Fatal(err)
   }
   uuid := fmt.Sprintf("%x-%x-%x-%x-%x",b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
   fmt.Println(uuid)
	dom, err := conn.DomainDefineXML(`<domain type="kvm">
		<name>` + imageName  + `</name>
				<uuid>`+uuid+`</uuid>
				<memory unit='KiB'>1048576</memory>
				<currentMemory unit='KiB'>1048576</currentMemory>
				<vcpu placement='static'>1</vcpu>
				<resource>
					<partition>/machine</partition>
				</resource>
				<os>
					<type arch='x86_64' machine='pc-i440fx-bionic'>hvm</type>
					<boot dev='hd'/>
				</os>
				<features>
					<acpi/>
					<apic/>
					<vmport state='off'/>
				</features>
				<cpu mode='custom' match='exact' check='full'>
					<model fallback='forbid'>Broadwell-noTSX-IBRS</model>
					<feature policy='require' name='vme'/>
					<feature policy='require' name='f16c'/>
					<feature policy='require' name='rdrand'/>
					<feature policy='require' name='hypervisor'/>
					<feature policy='require' name='arat'/>
					<feature policy='require' name='xsaveopt'/>
					<feature policy='require' name='abm'/>
				</cpu>
				<clock offset='utc'>
					<timer name='rtc' tickpolicy='catchup'/>
					<timer name='pit' tickpolicy='delay'/>
					<timer name='hpet' present='no'/>
				</clock>
				<on_poweroff>destroy</on_poweroff>
				<on_reboot>restart</on_reboot>
				<on_crash>destroy</on_crash>
				<pm>
					<suspend-to-mem enabled='no'/>
					<suspend-to-disk enabled='no'/>
				</pm>
				<devices>
					<emulator>/usr/bin/kvm-spice</emulator>
					<disk type='file' device='cdrom'>
      <source file='/home/swamym/Downloads/ubuntu-18.04.4-desktop-amd64.iso'/>
      <target dev='hdc'/>
      <readonly/>
    </disk>
					<disk type='file' device='disk'>
					<driver name='qemu' type='qcow2'/>
					<source file='/var/lib/libvirt/images/generic.qcow2'/>
					<backingStore/>
					<target dev='hda' bus='ide'/>
					<alias name='ide0-0-0'/>
					<address type='drive' controller='0' bus='0' target='0' unit='0'/>
					</disk>
					<disk type='file' device='cdrom'>
					<target dev='hdb' bus='ide'/>
					<readonly/>
					<alias name='ide0-0-1'/>
					<address type='drive' controller='0' bus='0' target='0' unit='1'/>
					</disk>
					<controller type='usb' index='0' model='ich9-ehci1'>
					<alias name='usb'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x7'/>
					</controller>
					<controller type='usb' index='0' model='ich9-uhci1'>
					<alias name='usb'/>
					<master startport='0'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0' multifunction='on'/>
					</controller>
					<controller type='usb' index='0' model='ich9-uhci2'>
					<alias name='usb'/>
					<master startport='2'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x1'/>
					</controller>
					<controller type='usb' index='0' model='ich9-uhci3'>
					<alias name='usb'/>
					<master startport='4'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x2'/>
					</controller>
					<controller type='pci' index='0' model='pci-root'>
					<alias name='pci.0'/>
					</controller>
					<controller type='ide' index='0'>
					<alias name='ide'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x1'/>
					</controller>
					<controller type='virtio-serial' index='0'>
					<alias name='virtio-serial0'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
					</controller>
					<interface type='network'>
					<mac address='52:54:00:59:30:69'/>
					<source network='default' bridge='virbr0'/>
					<target dev='vnet1'/>
					<model type='rtl8139'/>
					<alias name='net0'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
					</interface>
					<serial type='pty'>
					<source path='/dev/pts/6'/>
					<target type='isa-serial' port='0'>
						<model name='isa-serial'/>
					</target>
					<alias name='serial0'/>
					</serial>
					<console type='pty' tty='/dev/pts/6'>
					<source path='/dev/pts/6'/>
					<target type='serial' port='0'/>
					<alias name='serial0'/>
					</console>
					<channel type='spicevmc'>
					<target type='virtio' name='com.redhat.spice.0' state='connected'/>
					<alias name='channel0'/>
					<address type='virtio-serial' controller='0' bus='0' port='1'/>
					</channel>
					<input type='mouse' bus='ps2'>
					<alias name='input0'/>
					</input>
					<input type='keyboard' bus='ps2'>
					<alias name='input1'/>
					</input>
					<graphics type='spice' port='5901' autoport='yes' listen='127.0.0.1'>
					<listen type='address' address='127.0.0.1'/>
					<image compression='off'/>
					</graphics>
					<sound model='ich6'>
					<alias name='sound0'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
					</sound>
					<video>
					<model type='qxl' ram='65536' vram='65536' vgamem='16384' heads='1' primary='yes'/>
					<alias name='video0'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
					</video>
					<redirdev bus='usb' type='spicevmc'>
					<alias name='redir0'/>
					<address type='usb' bus='0' port='1'/>
					</redirdev>
					<redirdev bus='usb' type='spicevmc'>
					<alias name='redir1'/>
					<address type='usb' bus='0' port='2'/>
					</redirdev>
					<memballoon model='virtio'>
					<alias name='balloon0'/>
					<address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
					</memballoon>
				</devices>
				<seclabel type='dynamic' model='apparmor' relabel='yes'>
					<label>libvirt-bf834055-8945-4e5d-858b-8bd283595aa5</label>
					<imagelabel>libvirt-bf834055-8945-4e5d-858b-8bd283595aa5</imagelabel>
				</seclabel>
				<seclabel type='dynamic' model='dac' relabel='yes'>
					<label>+64055:+134</label>
					<imagelabel>+64055:+134</imagelabel>
				</seclabel>
				</domain>`)
	if err != nil {
		panic(err)
	}
	return dom
}
func create(imageName string,conn libvirt.Connect){
	dom := buildTestDomain(imageName,conn)
	defer func() {
		dom.Free()
		if res, _ := conn.Close(); res != 0 {
			fmt.Println("Close() == %d, expected 0", res)
		}
	}()
	if err := dom.Create(); err != nil {
		fmt.Println(err)
		return
	}
	state, reason, err := dom.GetState()
	if err != nil {
		fmt.Println(err)
		return
	}
	if state != libvirt.DOMAIN_RUNNING {
		fmt.Println("Domain should be running")
		return
	}
	if libvirt.DomainRunningReason(reason) != libvirt.DOMAIN_RUNNING_BOOTED {
		fmt.Println("Domain reason should be booted")
		return
	}
}
func listRunningDomains(conn libvirt.Connect){
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%d running domains:\n", len(doms))

fmt.Printf("\n**********Domains**********\n")

for i, dom := range doms {
	name, err := dom.GetName()

    if err == nil {
        fmt.Printf("%d  %s  running\n",i+1, name)
	}

    dom.Free()
}
}
func listDomain(conn libvirt.Connect){
	doms, err := conn.ListDomains()
	if err != nil {
		log.Fatal(err)
		return
	}
	if doms == nil {
		log.Fatal("ListDefinedDomains shouldn't be nil")
		return
	}
	fmt.Println(doms)
}
func createImage(){
	
}
func main() {
	var imageName string
	fmt.Println("Enter domain name")
	fmt.Scanf("%s",&imageName)
	conn := getConnection("192.168.100.9","swamym")
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	//create(imageName,*conn)
	listDomain(*conn)
	listRunningDomains(*conn)
	//fmt.Println(ni)
}