package bincode
import(
	"fmt"
	"reflect"
	"encoding/binary"
)
func SerializeData(v reflect.Value) ([]byte, error) {
	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case reflect.Uint8:
		return []byte{uint8(v.Uint())}, nil
	case reflect.Int16:
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v.Int()))
		return b, nil
	case reflect.Uint16:
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v.Uint()))
		return b, nil
	case reflect.Int32:
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(v.Int()))
		return b, nil
	case reflect.Uint32:
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(v.Uint()))
		return b, nil
	case reflect.Int64:
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v.Int()))
		return b, nil
	case reflect.Uint64:
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, v.Uint())
		return b, nil
	case reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Uint8:
			// Handle []uint8 (byte slice) explicitly
			length := v.Len()
			b := make([]byte, 8+length) // 8 bytes for length prefix
			binary.LittleEndian.PutUint64(b[:8], uint64(length))
			for i := 0; i < length; i++ {
				b[8+i] = byte(v.Index(i).Uint())
			}
			return b, nil
		default:
			return nil, fmt.Errorf("unsupported type: %v, elem: %v", v.Kind(), v.Type().Elem().Kind())
		}
	case reflect.Array:
		switch v.Type().Elem().Kind() {
		case reflect.Uint8:
			b := make([]byte, 0, v.Len())
			for i := 0; i < v.Len(); i++ {
				b = append(b, byte(v.Index(i).Uint()))
			}
			return b, nil
		}
		return nil, fmt.Errorf("unsupported type: %v, elem: %v", v.Kind(), v.Type().Elem().Kind())
	case reflect.String:
		str := v.String()
		length := len(str)
		b := make([]byte, 8+length) // 8 bytes for length prefix
		binary.LittleEndian.PutUint64(b[:8], uint64(length))
		copy(b[8:], []byte(str))
		return b, nil
	case reflect.Ptr:
		if v.IsNil() {
			return []byte{0}, nil
		}
		d, err := SerializeData(v.Elem())
		if err != nil {
			return nil, err
		}
		b := make([]byte, 1+len(d))
		b[0] = 1
		copy(b[1:], d[:])
		return b, nil
	case reflect.Struct:
		data := make([]byte, 0, 1024)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			d, err := SerializeData(field)
			if err != nil {
				return nil, err
			}
			data = append(data, d...)
		}
		return data, nil
	case reflect.Map:
		keys := v.MapKeys()
		output := make([]byte, 0)

		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(len(keys)))
		output = append(output, b...)

		for _, key := range keys {
			keyBytes, err := SerializeData(key)
			if err != nil {
				return nil, err
			}
			valBytes, err := SerializeData(v.MapIndex(key))
			if err != nil {
				return nil, err
			}

			keyLenBytes := make([]byte, 8)
			binary.LittleEndian.PutUint64(keyLenBytes, uint64(len(keyBytes)))
			output = append(output, keyBytes...)
			output = append(output, valBytes...)
		}
		return output, nil
	}
	return nil, fmt.Errorf("unsupported type: %v", v.Kind())
}