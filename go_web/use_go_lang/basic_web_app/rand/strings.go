package rand

//this is an arbitrary number
const RememberTokenBytes = 32

//RememberToken is a helper function designed to generate remember tokens
//of a predetermined byte size.
func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}

//Bytes will help us generate n random bytes, or wil
//return an error if there was one. This uses the
func Bytes(n int) ([]byte, error) {
	b := make([]byte, error) 
		_, err := rand.Read(b)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
}

//return a string that is the base64 URL encoded version of that byte slice
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

