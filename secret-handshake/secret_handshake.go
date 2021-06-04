package secret

func reverse(secretHandshake []string) []string {

	for i := 0; i < len(secretHandshake)/2; i++ {
		j := len(secretHandshake) - i - 1
		secretHandshake[i], secretHandshake[j] = secretHandshake[j], secretHandshake[i]
	}
	return secretHandshake
}

// Handshake returns a secret handshake based on a uint code
func Handshake(n uint) []string {
	secretHandshake := []string{}

	if n&uint(1) != 0 {
		secretHandshake = append(secretHandshake, "wink")
	}
	if n&uint(2) != 0 {
		secretHandshake = append(secretHandshake, "double blink")
	}
	if n&uint(4) != 0 {
		secretHandshake = append(secretHandshake, "close your eyes")
	}
	if n&uint(8) != 0 {
		secretHandshake = append(secretHandshake, "jump")
	}

	if n&uint(16) != 0 {
		secretHandshake = reverse(secretHandshake)
	}

	return secretHandshake
}
