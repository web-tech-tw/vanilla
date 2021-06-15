// Vanilla
// (c)2021 SuperSonic (https://github.com/supersonictw)

package fs

type Interface interface {
	Validate() bool
	Refactor() Response
}
