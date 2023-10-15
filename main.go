
package networkdiagnosis

import (
	"fmt"
	"net"
)

// Pingテスト
func PerformPing(address string) error {
	conn, err := net.Dial("ip4:icmp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Ping successful to", address)
	return nil
}


// トラフィック監視

// 