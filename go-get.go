// +build !windows

package winapi

func init() {
	panic(`runtime.GOOS != "windows"`)
}
