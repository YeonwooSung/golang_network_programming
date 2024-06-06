package util

import (
	"log"
	"fmt"

	"github.com/google/gopacket/pcap"
)

// GetDeviceList - get list of devices
func GetDeviceList() {
	// Get device list - run as ADMINISTRATOR!
    devices, err := pcap.FindAllDevs()
    if err != nil {
        panic(err)
    }

	// Print device information
	for _, device := range devices {
		log.Println("Name: ", device.Name)
		log.Println("Description: ", device.Description)
		log.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			log.Println("- IP address: ", address.IP)
			log.Println("- Subnet mask: ", address.Netmask)
		}
	}
}

func SelectTargetEthDevice() string {
	var deviceName string

	// Let user choose device
	fmt.Println("Enter device name: ")
	fmt.Scanln(&deviceName)

	return deviceName
}
